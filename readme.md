# AinhoSoft Client Launcher

## 📌 Description

This project is a cross-platform launcher for an OpenTibia-style server, developed in **Golang + Wails** with a React + TailwindCSS frontend.
It displays server status, news, updates, and launches the official client.

---

## 🚀 Requirements

- Golang 1.20+
- KrayAcc: https://github.com/Mau005/KrayAccOpenTibia (manifest service)
- Node.js 15+
- Wails CLI
```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

---

## ▶️ Run in development mode

### Windows (PowerShell)
```powershell
Push-Location frontend; npm run build; Pop-Location; wails dev
```

### Windows (CMD)
```cmd
cd frontend && npm run build && cd .. && wails dev
```

###Linux/Mac
```bash
cd frontend && npm run build && cd .. && wails dev
```

---

## 📦 Compilation

```bash
wails build
```

The binary will be generated in the `build/bin` folder.

---

## ⚙️ External Configuration

The launcher can take endpoints and configurations from a Go file that exposes functions such as:

```go
func (a *App) GetServerInfo(infoURL string) (*InfoResponse, error)
func (a *App) GetNews(newsURL string) ([]NewsItem, error)
```

This allows you to unify and modify IP addresses or endpoints from an **external configurator**.

---

## 🌐 Endpoints used
- `/client/info` → Server information (Uptime, Players, Status).
- `/get_news_short` → Short news and announcements.
- `/client/manifest` → Manifest dependency.

---

## 👨‍💻 Credits

Project developed by **Mau005**
Contributions and improvements are welcome.

---

## 📜 License

MIT License © 2025 AinhoSoft