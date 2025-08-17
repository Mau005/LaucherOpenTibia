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

## Compile

