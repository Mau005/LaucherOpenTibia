package main

import (
	"context"
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
)

/* =========================
   Tipos
========================= */

type FileEntry struct {
	Path   string `json:"path"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
	URL    string `json:"url,omitempty"`
}
type NewsShort struct {
	CreatedAt   time.Time
	IconID      uint8  `json:"IconID"`
	Description string `json:"Description"`
}
type NewsResponse struct {
	NewsShort []NewsShort `json:"NewsShort"`
}
type Manifest struct {
	App     string      `json:"app"`
	Version string      `json:"version"`
	BaseURL string      `json:"base_url"`
	Files   []FileEntry `json:"files"`
}

type ServerInfo struct {
	Uptime     string `json:"Uptime"`
	IP         string `json:"IP"`
	ServerName string `json:"ServerName"`
	Port       string `json:"Port"`
	Location   string `json:"Location"`
	URL        string `json:"URL"`
	Server     string `json:"Server"`
	Version    string `json:"Version"`
	Client     string `json:"Client"`
}
type Players struct {
	Online int `json:"Online"`
	Max    int `json:"Max"`
	Peak   int `json:"Peak"`
}

type InfoResponse struct {
	Version    string     `json:"Version"`
	ServerInfo ServerInfo `json:"ServerInfo"`
	Players    Players    `json:"Players"`
}

type Config struct {
	IPConnect  string `json:"IPConnect"`
	ClientPath string `json:"ClientPath"`
}

/* =========================
   Helpers de archivos/hashes
========================= */

func sha256File(path string) (int64, string, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, "", err
	}
	defer f.Close()
	h := sha256.New()
	n, err := io.Copy(h, f)
	if err != nil {
		return 0, "", err
	}
	return n, hex.EncodeToString(h.Sum(nil)), nil
}

func ensureDir(p string) error { return os.MkdirAll(filepath.Dir(p), 0755) }

// Reemplaza tu atomicReplace por esta versión con reintentos (Windows-friendly)
func atomicReplace(tmpPath, finalPath string) error {
	const attempts = 10
	for i := 0; i < attempts; i++ {
		// Si existe el destino, intenta borrarlo (Windows no sobrescribe con Rename)
		if _, err := os.Stat(finalPath); err == nil {
			if err := os.Remove(finalPath); err != nil {
				time.Sleep(200 * time.Millisecond)
				continue
			}
		}
		if err := os.Rename(tmpPath, finalPath); err != nil {
			time.Sleep(200 * time.Millisecond)
			continue
		}
		return nil
	}
	return fmt.Errorf("failed to replace %s after %d attempts", finalPath, attempts)
}

/* =========================
   Assets frontend
========================= */

//go:embed all:frontend/dist
var assets embed.FS
var cfg Config

/* =========================
   App
========================= */

type App struct {
	ctx context.Context
}

func NewApp() *App { return &App{} }

func (a *App) startup(ctx context.Context) { a.ctx = ctx }

func (a *App) emit(event string, data any) {
	if a.ctx != nil {
		wruntime.EventsEmit(a.ctx, event, data)
	}
}

func (a *App) safeCtx() context.Context {
	if a != nil && a.ctx != nil {
		return a.ctx
	}
	return context.Background()
}

/*
	=========================
	  Rutas locales (./Client)

=========================
*/
func userDataRoot() (string, error) {
	switch runtime.GOOS {
	case "windows":
		// %LOCALAPPDATA%\AinhoSoft\AinhoLauncher
		if base, ok := os.LookupEnv("LOCALAPPDATA"); ok {
			return filepath.Join(base, "AinhoSoft", "AinhoLauncher"), nil
		}
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "AppData", "Local", "AinhoSoft", "AinhoLauncher"), nil
	case "darwin":
		// ~/Library/Application Support/AinhoLauncher
		home, _ := os.UserHomeDir()
		return filepath.Join(home, "Library", "Application Support", "AinhoLauncher"), nil
	default:
		// Linux/otros: ~/.local/share/ainho-launcher
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".local", "share", "ainho-launcher"), nil
	}
}

func (a *App) GetInstallDir() (string, error) {
	root, err := userDataRoot()
	if err != nil {
		return "", err
	}
	clientDir := filepath.Join(root, "client")
	if err := os.MkdirAll(clientDir, 0o755); err != nil {
		return "", err
	}
	return clientDir, nil
}
func (a *App) GetGameExecutable() (string, error) {
	dir, err := a.GetInstallDir()
	if err != nil {
		return "", err
	}

	var exeName string
	switch runtime.GOOS {
	case "windows":
		exeName = fmt.Sprintf("bin/%s.exe", cfg.ClientPath)
	case "darwin":
		exeName = fmt.Sprintf("Ainho.app/Contents/MacOS/%s", cfg.ClientPath)
	default: // Linux
		exeName = cfg.ClientPath
	}

	return filepath.Join(dir, exeName), nil
}

// func (a *App) GetInstallDir() (string, error) {
// 	exe, err := os.Executable()
// 	if err != nil {
// 		return "", err
// 	}
// 	baseDir := filepath.Dir(exe)
// 	clientDir := filepath.Join(baseDir, "Client")
// 	if _, err := os.Stat(clientDir); os.IsNotExist(err) {
// 		if err := os.MkdirAll(clientDir, 0755); err != nil {
// 			return "", err
// 		}
// 	}
// 	return clientDir, nil
// }

/* =========================
   HTTP utils
========================= */

func (a *App) httpGET(url string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(a.safeCtx(), "GET", fmt.Sprintf("%s%s", cfg.IPConnect, url), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *App) httpGETNotSecuence(url string) (*http.Response, error) {
	req, _ := http.NewRequestWithContext(a.safeCtx(), "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/* =========================
   Info del servidor
========================= */

func (a *App) GetServerInfo(infoURL string) (*InfoResponse, error) {
	resp, err := a.httpGET(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http %d", resp.StatusCode)
	}
	var info InfoResponse
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}

/* =========================
   Updater
========================= */

// puedes ignorar archivos con esta función (opcional)
func skippable(path string) bool {
	return strings.HasPrefix(filepath.Base(path), ".")
}

func buildFileURL(baseURL, path, override string) string {
	if strings.HasPrefix(override, "http") {
		baseURL = override
	}
	baseURL = strings.TrimRight(baseURL, "/")
	path = strings.TrimLeft(path, "/")
	return baseURL + "/" + path
}

func (a *App) downloadWithProgress(url, dest string, totalBytes int64, doneSoFar *int64) error {
	resp, err := a.httpGETNotSecuence(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("download %s -> http %d", url, resp.StatusCode)
	}

	tmp := dest + ".part"
	if err := ensureDir(dest); err != nil {
		return err
	}
	out, err := os.Create(tmp)
	if err != nil {
		return err
	}

	buf := make([]byte, 64*1024)
	var fileDone int64
	cl := resp.ContentLength

	for {
		n, rerr := resp.Body.Read(buf)
		if n > 0 {
			if _, werr := out.Write(buf[:n]); werr != nil {
				out.Close()
				return werr
			}
			fileDone += int64(n)
			if doneSoFar != nil {
				*doneSoFar += int64(n)
			}
			if cl > 0 {
				a.emit("update:fileProgress", float64(fileDone)/float64(cl))
			}
			if totalBytes > 0 && doneSoFar != nil {
				a.emit("update:totalProgress", float64(*doneSoFar)/float64(totalBytes))
			}
		}
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			out.Close()
			return rerr
		}
	}

	if err := out.Sync(); err != nil {
		out.Close()
		return err
	}
	if err := out.Close(); err != nil {
		return err
	}

	return atomicReplace(tmp, dest)
}

// ==== método expuesto a frontend ====
func (a *App) GetNews(apiURL string) ([]NewsShort, error) {
	resp, err := a.httpGET(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http %d", resp.StatusCode)
	}
	var payload NewsResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}
	return payload.NewsShort, nil
}

func (a *App) UpdateFromManifest(manifestURL, installDir string) error {
	resp, err := a.httpGET(manifestURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("manifest http %d", resp.StatusCode)
	}
	var man Manifest
	if err := json.NewDecoder(resp.Body).Decode(&man); err != nil {
		return err
	}
	fmt.Println("Manifest OK. base_url:", man.BaseURL)

	overrideBase := ""

	var bytesToDownload int64
	toProcess := make([]FileEntry, 0, len(man.Files))

	for _, f := range man.Files {
		if skippable(f.Path) {
			fmt.Println("Skip:", f.Path)
			continue
		}
		local := filepath.Join(installDir, filepath.FromSlash(f.Path))
		need := false
		if st, err := os.Stat(local); err != nil || st.Size() != f.Size {
			need = true
		} else {
			_, sum, err := sha256File(local)
			if err != nil || !strings.EqualFold(sum, f.SHA256) {
				need = true
			}
		}
		if need {
			bytesToDownload += f.Size
			toProcess = append(toProcess, f)
		}
	}

	if len(toProcess) == 0 {
		a.emit("update:status", "Up to date")
		a.emit("update:totalProgress", 1.0)
		fmt.Println("No files to download.")
		return nil
	}

	var doneSoFar int64
	for _, f := range toProcess {
		a.emit("update:status", "Downloading "+f.Path)

		url := f.URL
		if url == "" {
			url = buildFileURL(man.BaseURL, f.Path, overrideBase)
		}
		dest := filepath.Join(installDir, filepath.FromSlash(f.Path))

		fmt.Println("Downloading:", url, "->", dest)

		if err := a.downloadWithProgress(url, dest, bytesToDownload, &doneSoFar); err != nil {
			return err
		}

		_, sum, err := sha256File(dest)
		if err != nil || !strings.EqualFold(sum, f.SHA256) {
			return fmt.Errorf("checksum mismatch for %s", f.Path)
		}
	}

	a.emit("update:status", "Ready")
	a.emit("update:totalProgress", 1.0)
	fmt.Println("Update completed in:", installDir)
	return nil
}

func (a *App) CheckUpdates() ([]string, error) {
	steps := []string{"Checking updates...", "Downloading client.bin", "Verifying files...", "Ready to start"}
	for _, s := range steps {
		a.emit("update:status", s)
		time.Sleep(600 * time.Millisecond)
	}
	for i := 0; i <= 100; i++ {
		a.emit("update:totalProgress", float64(i)/100.0)
		time.Sleep(20 * time.Millisecond)
	}
	return steps, nil
}

func (a *App) StartGame() (string, error) {
	var cmd *exec.Cmd
	path, err := a.GetGameExecutable()
	if err != nil {
		log.Println(err)
		return "error execute client", err
	}
	switch runtime.GOOS {
	case "windows", "linux", "darwin":
		cmd = exec.Command(path)
	default:
		return "", fmt.Errorf("OS not supported")
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	return "OK", nil
}

func (a *App) OpenLink(url string) error {
	if a.ctx == nil {
		return fmt.Errorf("app not initialised yet")
	}
	wruntime.BrowserOpenURL(a.ctx, url)
	return nil
}

/* =========================
   Wails
========================= */

func main() {

	cont, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(cont, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	// cfg.IPConnect = "https://ainho.ddns.net"
	// cfg.ClientPath = "Ainho"
	app := NewApp()

	err = wails.Run(&options.App{
		Title:            "Ainho Launcher",
		Width:            1200,
		Height:           800,
		Frameless:        true,
		DisableResize:    false,
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		Mac: &mac.Options{
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			TitleBar: &mac.TitleBar{
				FullSizeContent:            true,
				HideTitle:                  true,
				TitlebarAppearsTransparent: true,
			},
		},
		AssetServer: &assetserver.Options{Assets: assets},
		OnStartup:   app.startup,
		Bind:        []interface{}{app},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
