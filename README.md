
<p align="center">
    <img src="./docs/logo.png" alt="GoWVP Logo" width="550"/>
</p>

<p align="center">
    <a href="https://github.com/gowvp/gb28181/releases"><img src="https://img.shields.io/github/v/release/ixugo/goweb?include_prereleases" alt="Version"/></a>
    <a href="https://github.com/ixugo/goweb/blob/master/LICENSE.txt"><img src="https://img.shields.io/dub/l/vibe-d.svg" alt="License"/></a>
</p>

# 开箱即用的 GB/T28181 协议视频平台

go wvp 是 Go 语言实现的开源 GB28181 解决方案，基于GB28181-2022标准实现的网络视频平台，支持 rtmp/rtsp，客户端支持网页版本和安卓 App。支持rtsp/rtmp等视频流转发到国标平台，支持rtsp/rtmp等推流转发到国标平台。


## 在线演示平台

+ [在线演示平台 :)](http://gowvp.golang.space:15123/)


|![](./docs/demo/home.webp)|![](./docs/demo/play.webp)|
|-|-|



## 应用场景：
+ 支持浏览器无插件播放摄像头视频。
+ 支持国标设备(摄像机、平台、NVR等)设备接入
+ 支持非国标(rtsp, rtmp，直播设备等等)设备接入，充分利旧。
+ 支持国标级联。多平台级联。跨网视频预览。
+ 支持跨网网闸平台互联。


## 开源库

感谢 @panjjo 大佬的开源库 [panjjo/gosip](https://github.com/panjjo/gosip)，GoWVP 的 sip 信令基于此库，出于底层封装需要，并非直接 go mod 依赖该项目，而是源代码放到了 pkg 包中。

流媒体服务基于@夏楚 [ZLMediaKit](https://github.com/ZLMediaKit/ZLMediaKit)

播放器使用@dexter [jessibuca](https://github.com/langhuihui/jessibuca/tree/v3)

项目框架基于 @ixugo [goweb](https://github.com/ixugo/goweb)

Java 语言 WVP @648540858 [wvp-GB28181-pro](https://github.com/648540858/wvp-GB28181-pro)

## GoWVP, GB/T28181 交流群
<h1>点个 Star 再扫码进群呀</h1>
<img src="./wechat.jpg" alt="wechat" width="200"/>

## QA

> 怎么没有前端资源? 如何加载网页呢?

前端资源打包后放到项目根目录，重命名为 `www` 即可正常加载。

> 有没有代码相关的学习资料?

[GB/T28181 开源日记[1]：从 0 到实现 GB28181 协议的完整实践](https://juejin.cn/post/7456722441395568651)

[GB/T28181 开源日记[2]：搭建服务端，解决跨域，接口联调](https://juejin.cn/post/7456796962120417314)

[GB/T28181 开源日记[3]：使用 React 组件构建监控数据面板](https://juejin.cn/post/7457228085826764834)

[GB/T28181 开源日记[4]：使用 ESlint 辅助开发](https://juejin.cn/post/7461539078111789108)

[GB/T28181 开源日记[5]：使用 react-hook-form 完成表单](https://juejin.cn/post/7461899974198181922)

[GB/T28181 开源日记[6]：React 快速接入 jessibuca.js 播放器](https://juejin.cn/post/7462229773982351410)

[GB/T28181 开源日记[7]：实现 RTMP 鉴权与播放](https://juejin.cn/post/7463504223177261119)

[GB/T28181 开源日记[8]：国标开发速知速会](https://juejin.cn/post/7468626309699338294)

> 有没有使用资料?

**RTMP**

[RTMP 推拉流规则](https://juejin.cn/post/7463124448540934194)

[如何使用 OBS RTMP 推流到 GB/T28181平台](https://juejin.cn/post/7463350947100786739)

[海康摄像机 RTMP 推流到开源 GB/T28181 平台](https://juejin.cn/post/7468191617020313652)

[大华摄像机 RTMP 推流到开源 GB/T28181 平台](https://juejin.cn/spost/7468194672773021731)

**GB/T28181**

[GB28181 七种注册姿势](https://juejin.cn/post/7465274924899532838)

## 文档

GoWVP [在线接口文档](apifox.com/apidoc/shared-7b67c918-5f72-4f64-b71d-0593d7427b93)

ZLM使用文档 [github.com/ZLMediaKit/ZLMediaKit](https://github.com/ZLMediaKit/ZLMediaKit)


## 快速开始

即将发布安装包 和 docker 版本。

如果你是 Go 语言开发者并熟悉 docker，可以提前下载源代码，本地编程运行。

**前置条件**

+ Golang
+ Docker & Docker Compose
+ Make

**操作流程**

+ 1. 克隆本项目
+ 2. 修改 configs/config.toml 中 `WebHookIP` 为你的局域网 IP
+ 3. 执行 `make build/linux && docker compose up -d`
+ 4. 自动创建了 zlm.conf 文件夹，获取 config.ini 的 api 秘钥，填写到 `configs/config.toml` 的 `Secret`
+ 5. 执行 `docker compose restart`
+ 6. 浏览器访问 `http://localhost:15123`



##  如何参与开发?

1. fork 本项目
2. 编辑器 run/debug 设置配置输出目录为项目根目录
3. 修改，提交 PR，说明修改内容

## 功能特性
- [x] 集成 web 界面
- [x] 接入设备
  - [x] 视频预览
  - [x] 无限制接入路数，能接入多少设备只取决于你的服务器性能
  - [ ] 云台控制，控制设备转向，拉近，拉远
  - [ ] 预置位查询，使用与设置
  - [ ] 查询 NVR/IPC 上的录像与播放，支持指定时间播放与下载
  - [x] 无人观看自动断流，节省流量
  - [x] 视频设备信息同步
  - [ ] 离在线监控
  - [x] 支持直接输出RTSP、RTMP、HTTP-FLV、Websocket-FLV、HLS多种协议流地址
  - [ ] 支持通过一个流地址直接观看摄像头，无需登录以及调用任何接口
  - [x] 支持 UDP 和 TCP 两种国标信令传输模式
  - [ ] 支持 UDP 和 TCP 被动,TCP 主动 三种国标流传输模式
  - [x] 支持检索,通道筛选
  - [ ] 支持通道子目录查询
  - [ ] 支持过滤音频，防止杂音影响观看
  - [x] 支持国标网络校时
  - [x] 支持播放 H264 和 H265
  - [ ] 报警信息处理，支持向前端推送报警信息
  - [ ] 语音对讲
  - [ ] 支持业务分组和行政区划树自定义展示以及级联推送
  - [ ] 支持订阅与通知方法
    - [ ] 移动位置订阅
    - [ ] 移动位置通知处理
    - [ ] 报警事件订阅
    - [ ] 报警事件通知处理
    - [ ] 设备目录订阅
    - [ ] 设备目录通知处理
  -  [ ] 移动位置查询和显示
  - [x] 支持手动添加设备和给设备设置单独的密码
-  [ ] 支持平台对接接入
- [x] 支持自动配置ZLM媒体服务, 减少因配置问题所出现的问题
- [x] 支持启用 udp 多端口模式, 提高 udp 模式下媒体传输性能
- [x] 支持局域网/互联网/特殊网络环境部署
- [x] 支持 gowvp 与 zlm 分开部署，提升平台并发能力
- [ ] 支持拉流RTSP/RTMP，分发为各种流格式，或者推送到其他国标平台
- [ ] 支持推流RTSP/RTMP，分发为各种流格式，或者推送到其他国标平台
- [x] 支持推流鉴权
- [x] 支持接口鉴权
- [ ] 云端录像，推流/代理/国标视频均可以录制在云端服务器，支持预览和下载
- [ ] 支持跨域请求，支持前后端分离部署
- [x] 支持 PostgreSQL 数据库
- [ ] 支持录制计划, 根据设定的时间对通道进行录制. 暂不支持将录制的内容转发到国标上级



## 授权协议

本项目自有代码使用宽松的MIT协议，在保留版权信息的情况下可以自由应用于各自商用、非商业的项目。

但是本项目也零碎的使用了一些其他的开源代码，在商用的情况下请自行替代或剔除； 由于使用本项目而产生的商业纠纷或侵权行为一概与本项目及开发者无关，请自行承担法律风险。

在使用本项目代码时，也应该在授权协议中同时表明本项目依赖的第三方库的协议