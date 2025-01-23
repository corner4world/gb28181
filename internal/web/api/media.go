// Code generated by gowebx, DO AVOID EDIT.
package api

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/gowvp/gb28181/internal/conf"
	"github.com/gowvp/gb28181/internal/core/media"
	"github.com/gowvp/gb28181/internal/core/sms"
	"github.com/ixugo/goweb/pkg/hook"
	"github.com/ixugo/goweb/pkg/web"
)

type MediaAPI struct {
	mediaCore media.Core
	smsCore   sms.Core
	conf      *conf.Bootstrap
}

// NewMediaAPI 媒体相关接口
func NewMediaAPI(mc media.Core, sc sms.Core, conf *conf.Bootstrap) MediaAPI {
	return MediaAPI{mediaCore: mc, smsCore: sc, conf: conf}
}

func registerMediaAPI(g gin.IRouter, api MediaAPI, handler ...gin.HandlerFunc) {
	{
		group := g.Group("/stream_pushs", handler...)
		group.GET("", web.WarpH(api.findStreamPush))
		group.GET("/:id", web.WarpH(api.getStreamPush))
		group.PUT("/:id", web.WarpH(api.editStreamPush))
		group.POST("", web.WarpH(api.addStreamPush))
		group.DELETE("/:id", web.WarpH(api.delStreamPush))
	}
}

// >>> streamPush >>>>>>>>>>>>>>>>>>>>

func (a MediaAPI) findStreamPush(c *gin.Context, in *media.FindStreamPushInput) (*web.PageOutput, error) {
	items, total, err := a.mediaCore.FindStreamPush(c.Request.Context(), in)
	if err != nil {
		return nil, err
	}

	cacheFn := hook.UseCache(func(s string) (*sms.MediaServer, error) {
		v, err := a.smsCore.GetMediaServer(c.Request.Context(), s)
		if err != nil {
			slog.Error("GetMediaServer", "err", err)
		}
		return v, err
	})

	out := make([]*media.FindStreamPushOutputItem, len(items))
	for i, item := range items {
		rtmpAddrs := []string{"unknow"}
		mediaID := item.MediaServerID
		if mediaID == "" {
			mediaID = sms.DefaultMediaServerID
		}
		if svr, _, _ := cacheFn(mediaID); svr != nil {
			addr := fmt.Sprintf("rtmp://%s:%d/%s/%s", web.GetHost(c.Request), svr.Ports.RTMP, item.App, item.Stream)
			if !item.IsAuthDisabled {
				addr += fmt.Sprintf("?sign=%s", hook.MD5(a.conf.Server.RTMPSecret))
			}
			rtmpAddrs[0] = addr
		}

		out[i] = &media.FindStreamPushOutputItem{
			StreamPush: *item,
			PushAddrs:  rtmpAddrs,
		}
	}
	return &web.PageOutput{Items: out, Total: total}, err
}

func (a MediaAPI) getStreamPush(c *gin.Context, _ *struct{}) (*media.StreamPush, error) {
	streamPushID := c.Param("id")
	return a.mediaCore.GetStreamPush(c.Request.Context(), streamPushID)
}

func (a MediaAPI) editStreamPush(c *gin.Context, in *media.EditStreamPushInput) (*media.StreamPush, error) {
	streamPushID := c.Param("id")
	return a.mediaCore.EditStreamPush(c.Request.Context(), in, streamPushID)
}

func (a MediaAPI) addStreamPush(c *gin.Context, in *media.AddStreamPushInput) (*media.StreamPush, error) {
	return a.mediaCore.AddStreamPush(c.Request.Context(), in)
}

func (a MediaAPI) delStreamPush(c *gin.Context, _ *struct{}) (*media.StreamPush, error) {
	streamPushID := c.Param("id")
	return a.mediaCore.DelStreamPush(c.Request.Context(), streamPushID)
}
