# WebRemoteMediaCtrl

---
## Language

- [English](README.md)
- [中文](README_zh.md)


### 🧭 **Project Overview**

> **WebRemoteMediaCtrl** is a lightweight web-based media control tool built with **Go (Gin)** and **RobotGo**.
> It allows you to remotely control your computer’s media playback (play, pause, skip, volume, mute, etc.) from any device’s browser.
> Supports **QR code** generation for quick access via mobile devices.

---

## 🚀 Features

* 🎶 **Playback Control:** Play, pause, stop, previous, next
* 🔊 **Volume Control:** Increase, decrease, mute, unmute
* 📱 **QR Access:** Generate a QR code to control from your phone
* ⚙️ **Modes:** Supports `release` and `debug`
* 🌗 **Auto Theme:** Adapts to light or dark mode automatically

---

## 🧩 Project Structure

```
├── main.go              # Backend service (Gin + RobotGo)
├── static/
│   ├── index.html       # Frontend control page
│   └── qrcode.min.js    # QR code generation script
```

---

## 🧰 Dependencies

* Go 1.18+
* [RobotGo](https://github.com/go-vgo/robotgo)
* [Gin Web Framework](https://github.com/gin-gonic/gin)

Install dependencies:

```bash
go get github.com/gin-gonic/gin
go get github.com/go-vgo/robotgo
```

---

## ⚙️ Launch Parameters

The program supports command-line flags for port and mode configuration:

| Flag | Default   | Description                    |
| ---- | --------- | ------------------------------ |
| `-p` | `9608`    | Web server listening port      |
| `-m` | `release` | Run mode: `debug` or `release` |

Example:

```bash
go run main.go -p 9608 -m release
```

---

## 💻 How to Run

1. **Run or Build the Program**

   ```bash
   go run main.go
   ```

   or build:

   ```bash
   go build -o media-controller
   ./media-controller
   ```

2. **Open the Control Page**
   Once started, visit:

   ```
   http://localhost:9608
   ```

3. **Control from Mobile**

    * Click the “🤳” button to generate a QR code
    * Scan the QR code with your phone (must be on the same local network)
    * Tap buttons on the web page to control your PC media playback

---

## 📡 API Reference

Frontend interacts with the backend via the following endpoint:

```
GET /media?operation=<operation_name>
```

Supported operations:

| Operation     | Description     |
| ------------- | --------------- |
| `play`        | Play            |
| `pause`       | Pause           |
| `stop`        | Stop            |
| `previous`    | Previous track  |
| `next`        | Next track      |
| `volume_up`   | Increase volume |
| `volume_down` | Decrease volume |
| `mute`        | Mute            |
| `un_mute`     | Unmute          |

Response format:

```json
{
  "ok": true,
  "msg": "Success"
}
```

---

## 🧠 Code Overview

### Backend (`main.go`)

* Uses `flag` to parse command-line arguments (port, mode)
* Provides `/media` route via `gin`
* Uses `robotgo.KeyTap` to simulate system multimedia keys
* Returns JSON responses for operation results

### Frontend (`index.html`)

* Simple, responsive control interface
* Uses `fetch` to call `/media?operation=...`
* Displays feedback via a toast message
* Uses `qrcode.min.js` to dynamically generate QR codes

---

## 📱 UI Preview

```
⏮ ⏸ ▶ ⏹ ⏭
🔇 🔊 🔉- 🔊+
🤳  ← Tap to generate QR code
```

---

## ⚠️ Notes

* Ensure your computer allows `RobotGo` to control system inputs (some OSs require accessibility permissions)
* Both your phone and computer must be connected to the **same local network (LAN)**

---

## 🧑‍💻 Tips & Suggestions

* Works well with most music players (Spotify, Netease Cloud Music, QQ Music, etc.)
* You can add this program to system startup to control your PC music anytime from your phone