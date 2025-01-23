// Code generated by gowebx, DO AVOID EDIT.
package media

import "github.com/ixugo/goweb/pkg/orm"

const (
	StatusPushing = "PUSHING" // 推流中状态
	StatusStopped = "STOPPED" // 停止
)

const (
	RTMPIDPrefix = "r" // rtmp ID 前缀，取 rtmp 首字母
)

// StreamPush domain model
type StreamPush struct {
	ID             string    `gorm:"primaryKey" json:"id"`
	CreatedAt      orm.Time  `gorm:"column:created_at;notNull;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`                  // 创建时间
	UpdatedAt      orm.Time  `gorm:"column:updated_at;notNull;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`                  // 更新时间
	Name           string    `gorm:"column:name;notNull;default:'';comment:推流名称" json:"name"`                                             // 推流名称
	PushedAt       *orm.Time `gorm:"column:pushed_at;notNull;default:'1970-01-01 00:00:00';comment:最后一次推流时间" json:"pushed_at"`            // 最后一次推流时间
	StoppedAt      *orm.Time `gorm:"column:stopped_at;notNull;default:'1970-01-01 00:00:00';comment:最后一次停止时间" json:"stopped_at"`          // 最后一次停止时间
	App            string    `gorm:"column:app;notNull;default:'';uniqueIndex:idx_stream_pushs_app_stream;comment:应用名" json:"app"`        // 应用名
	Stream         string    `gorm:"column:stream;notNull;default:'';uniqueIndex:idx_stream_pushs_app_stream;comment:流 ID" json:"stream"` // 流 ID
	MediaServerID  string    `gorm:"column:media_server_id;notNull;default:'';comment:媒体服务器 ID" json:"media_server_id"`                   // 媒体服务器 ID
	ServerID       string    `gorm:"column:server_id;notNull;default:'';comment:服务器 ID" json:"server_id"`                                 // 服务器 ID
	Status         string    `gorm:"column:status;notNull;default:'';comment:推流状态(PUSHING)" json:"status"`                                // 推流状态(PUSHING)
	IsAuthDisabled bool      `gorm:"column:is_auth_disabled;notNull;default:false;comment:是否启用推流鉴权" json:"is_auth_disabled"`              // 是否启用推流鉴权

	// 自定义拉流鉴权参数，IsAuthDisabled=false 时生效
	Session string `gorm:"column:session;notNull;default:'';comment:session" json:"-"`
}

// TableName database table name
func (*StreamPush) TableName() string {
	return "stream_pushs"
}
