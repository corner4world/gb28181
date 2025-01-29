package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/gowvp/gb28181/internal/conf"
	"github.com/gowvp/gb28181/internal/core/media"
	"github.com/gowvp/gb28181/internal/core/media/store/mediadb"
	"github.com/gowvp/gb28181/internal/core/uniqueid"
	"github.com/gowvp/gb28181/internal/core/uniqueid/store/uniqueiddb"
	"github.com/gowvp/gb28181/internal/core/version"
	"github.com/gowvp/gb28181/internal/core/version/store/versiondb"
	"github.com/gowvp/gb28181/pkg/gbs"
	"github.com/ixugo/goweb/pkg/orm"
	"github.com/ixugo/goweb/pkg/web"
	"gorm.io/gorm"
)

var (
	ProviderVersionSet = wire.NewSet(NewVersion)
	ProviderSet        = wire.NewSet(
		wire.Struct(new(Usecase), "*"),
		NewHTTPHandler,
		NewVersionAPI,
		NewSMSCore, NewSmsAPI,
		NewWebHookAPI,
		NewUniqueID,
		NewMediaCore, NewMediaAPI,
		gbs.NewServer,
		NewGb28181API,
	)
)

type Usecase struct {
	Conf       *conf.Bootstrap
	DB         *gorm.DB
	Version    VersionAPI
	SMSAPI     SmsAPI
	WebHookAPI WebHookAPI
	UniqueID   uniqueid.Core
	MediaAPI   MediaAPI
	GB28181API Gb28181API

	SipServer *gbs.Server
}

// NewHTTPHandler 生成Gin框架路由内容
func NewHTTPHandler(uc *Usecase) http.Handler {
	cfg := uc.Conf.Server
	// 检查是否设置了 JWT 密钥，如果未设置，则生成一个长度为 32 的随机字符串作为密钥
	if cfg.HTTP.JwtSecret == "" {
		cfg.HTTP.JwtSecret = orm.GenerateRandomString(32) // 生成一个长度为 32 的随机字符串作为密钥
	}
	// 如果不处于调试模式，将 Gin 设置为发布模式
	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode) // 将 Gin 设置为发布模式
	}
	g := gin.New() // 创建一个新的 Gin 实例
	// 处理未找到路由的情况，返回 JSON 格式的 404 错误信息
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, "来到了无人的荒漠") // 返回 JSON 格式的 404 错误信息
	})
	// 如果启用了 Pprof，设置 Pprof 监控
	if cfg.HTTP.PProf.Enabled {
		web.SetupPProf(g, &cfg.HTTP.PProf.AccessIps) // 设置 Pprof 监控
	}

	setupRouter(g, uc) // 设置路由处理函数

	return g // 返回配置好的 Gin 实例作为 http.Handler
}

// NewVersion ...
func NewVersion(db *gorm.DB) version.Core {
	vdb := versiondb.NewDB(db)
	core := version.NewCore(vdb)
	isOK := core.IsAutoMigrate(dbVersion)
	vdb.AutoMigrate(isOK)
	if isOK {
		slog.Info("更新数据库表结构")
		if err := core.RecordVersion(dbVersion, dbRemark); err != nil {
			slog.Error("RecordVersion", "err", err)
		}
	}
	orm.EnabledAutoMigrate = isOK
	return core
}

// NewUniqueID 唯一 id 生成器
func NewUniqueID(db *gorm.DB) uniqueid.Core {
	return uniqueid.NewCore(uniqueiddb.NewDB(db).AutoMigrate(orm.EnabledAutoMigrate), 5)
}

func NewMediaCore(db *gorm.DB, uni uniqueid.Core) media.Core {
	return media.NewCore(mediadb.NewDB(db).AutoMigrate(orm.EnabledAutoMigrate), uni)
}
