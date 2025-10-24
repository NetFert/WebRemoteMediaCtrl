# WebRemoteMediaCtrl

---
## Language

- [English](README.md)
- [中文](README_zh.md)


### 🧭 **项目简介**
> **WebRemoteMediaCtrl** 是一个基于 **Go (Gin)** 和 **RobotGo** 的网页媒体控制工具。
> 可通过任意设备的浏览器远程控制电脑媒体播放（播放、暂停、切歌、音量、静音等）。
> 支持生成 **二维码**，方便手机快速访问控制界面。

---

## ⏬ 下载使用

**下载:**
- [GitHub Releases](https://github.com/NetFert/WebRemoteMediaCtrl)
- [Gitee Releases](https://gitee.com/NetFert/WebRemoteMediaCtrl)

---

![Preview](record.gif)

---


## 🚀 功能简介

* 🎶 播放控制：播放、暂停、停止、上一曲、下一曲
* 🔊 音量控制：增大、减小、静音、取消静音
* 📱 扫码访问：生成二维码，在手机浏览器中直接控制
* ⚙️ 模式选择：支持 release / debug 模式
* 🌗 支持深色模式 / 浅色模式自动适配

---

## 🧩 项目结构

```
├── main.go              # 后端服务（Gin + robotgo）
├── static/
│   ├── index.html       # 前端控制页面
│   └── qrcode.min.js    # 二维码生成脚本
```

---

## 🧰 环境依赖

* Go 1.18+
* [robotgo](https://github.com/go-vgo/robotgo)
* [Gin Web 框架](https://github.com/gin-gonic/gin)

安装依赖：

```bash
go get github.com/gin-gonic/gin
go get github.com/go-vgo/robotgo
```

---

## ⚙️ 启动参数

程序支持通过命令行参数设置端口号和模式：

| 参数   | 默认值       | 说明                          |
| ---- | --------- | --------------------------- |
| `-p` | `9608`    | Web 服务监听端口                  |
| `-m` | `release` | 运行模式，可选 `debug` 或 `release` |

示例：

```bash
go run main.go -p 9608 -m release
```

---

## 💻 运行步骤

1. **编译或运行程序**

   ```bash
   go run main.go
   ```

   或编译：

   ```bash
   go build -o media-controller
   ./media-controller
   ```

2. **查看控制页面**
   启动后，在浏览器中访问：

   ```
   http://localhost:9608
   ```

3. **手机控制**

    * 点击网页底部的「🤳」按钮生成二维码
    * 使用手机扫码打开网页（需在同一局域网内）
    * 点击按钮即可控制电脑音乐播放

---

## 📡 接口说明

前端通过以下接口与后端通信：

```
GET /media?operation=<操作名称>
```

支持的操作：

| 操作参数          | 功能   |
| ------------- | ---- |
| `play`        | 播放   |
| `pause`       | 暂停   |
| `stop`        | 停止   |
| `previous`    | 上一曲  |
| `next`        | 下一曲  |
| `volume_up`   | 音量增加 |
| `volume_down` | 音量减小 |
| `mute`        | 静音   |
| `un_mute`     | 取消静音 |

响应格式：

```json
{
  "ok": true,
  "msg": "Success"
}
```

---

## 🧠 代码说明

### 后端（main.go）

* 使用 `flag` 获取运行参数（端口、模式）
* 使用 `gin` 提供 `/media` 路由
* 通过 `robotgo.KeyTap` 模拟系统多媒体键，实现播放控制
* 返回操作结果 JSON 响应

### 前端（index.html）

* 简洁的网页控制界面
* 使用 `fetch` 调用 `/media?operation=...`
* 提供 toast 提示反馈
* 使用 `qrcode.min.js` 动态生成当前网页二维码

---

## 📱 页面预览

```
⏮ ⏸ ▶ ⏹ ⏭
🔇 🔊 🔉- 🔊+
🤳  ← 点击生成二维码
```

---

## ⚠️ 注意事项

* 请确保电脑允许被 `robotgo` 控制（部分系统需开启辅助权限）
* 手机与电脑必须在同一局域网内

---

## 🧑‍💻 使用建议

* 可结合系统音乐播放器（Spotify、网易云音乐、QQ 音乐等）使用
* 可将程序放在开机启动项，实现手机随时远程控制电脑音乐

---