# build-sign-release.ps1
param(
  [string]$Version       = "1.0.0",
  [string]$BuildArch     = "windows/amd64",
  [string]$AppExeName    = "Ainho_Launcher.exe",
  [string]$PfxPath       = "C:\Users\mau\Documents\ainhosoft.pfx",
  [string]$PfxPassword   = "",
  [string]$CertSubject   = "",
  [string]$TimestampUrl  = "http://timestamp.digicert.com",
  [switch]$SkipSign,
  [switch]$SkipInstaller    # <--- NUEVO
)

$ErrorActionPreference = "Stop"
Set-StrictMode -Version Latest

function Require-File($path, $name) {
  if (-not (Test-Path $path)) { throw "No se encontró $name en $path" }
}

function Add-NsisToPath {
  $nsisRoot = 'C:\Program Files (x86)\NSIS'
  $nsisBin  = Join-Path $nsisRoot 'Bin'
  if (-not (($env:Path -split ';') -contains $nsisRoot)) { $env:Path += ";$nsisRoot" }
  if (-not (($env:Path -split ';') -contains $nsisBin )) { $env:Path += ";$nsisBin"  }
}

function Find-Installer([string]$root) {
  $candidates = @()

  $binDir = Join-Path $root 'build\bin'
  if (Test-Path $binDir) {
    $candidates += Get-ChildItem $binDir -Filter *.exe -Recurse -ErrorAction SilentlyContinue |
      Where-Object { $_.Name -match '(setup|installer|install)' -and $_.Name -notmatch 'Webview2|Edge' } |
      Sort-Object LastWriteTime -Descending
  }

  $wailsInst = Join-Path $root 'build\windows\installer'
  if (Test-Path $wailsInst) {
    $candidates += Get-ChildItem $wailsInst -Filter *.exe -Recurse -ErrorAction SilentlyContinue |
      Where-Object { $_.Name -match '(setup|installer|install)' -and $_.Name -notmatch 'Webview2|Edge' } |
      Sort-Object LastWriteTime -Descending
  }

  $packagingDir = Join-Path $root 'packaging'
  if (Test-Path $packagingDir) {
    $candidates += Get-ChildItem $packagingDir -Filter *.exe -Recurse -ErrorAction SilentlyContinue |
      Where-Object { $_.Name -match '(setup|installer|install)' -and $_.Name -notmatch 'Webview2|Edge' } |
      Sort-Object LastWriteTime -Descending
  }

  if (-not $candidates) {
    $candidates += Get-ChildItem $root -Filter *.exe -Recurse -ErrorAction SilentlyContinue |
      Where-Object { $_.Name -match '(setup|installer|install)' -and $_.Name -notmatch 'Webview2|Edge' } |
      Sort-Object LastWriteTime -Descending
  }

  if ($candidates) { return ($candidates | Select-Object -First 1 -ExpandProperty FullName) }
  return $null
}

# -------- Paths base --------
$Root   = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $Root
$BinDir = Join-Path $Root "build\bin"
$OutDir = $BinDir

# -------- Herramientas --------
if (-not $SkipInstaller) {
  Add-NsisToPath   # solo si vamos a crear instalador
}

# makensis (solo si hace falta)
$MakensisExe = $null
if (-not $SkipInstaller) {
  try { $MakensisExe = (Get-Command makensis.exe -ErrorAction SilentlyContinue).Source } catch {}
  if (-not $MakensisExe) { $MakensisExe = "C:\Program Files (x86)\NSIS\makensis.exe" }
  Require-File $MakensisExe "makensis.exe"
}

# signtool
$SigntoolExe = (Get-ChildItem "$env:ProgramFiles(x86)\Windows Kits\10\bin\*\x64\signtool.exe" -Recurse -ErrorAction SilentlyContinue |
  Sort-Object { $_.VersionInfo.FileVersionRaw } -Descending |
  Select-Object -First 1 -ExpandProperty FullName)
if (-not $SigntoolExe) { $SigntoolExe = "C:\Program Files (x86)\Windows Kits\10\bin\10.0.26100.0\x64\signtool.exe" }
Require-File $SigntoolExe "signtool.exe"

# wails
if (-not (Get-Command wails -ErrorAction SilentlyContinue)) {
  throw "No se encontró 'wails' en PATH. Instálalo: go install github.com/wailsapp/wails/v2/cmd/wails@latest"
}

# -------- Build Wails (.exe) --------
Write-Host "== Compilando Wails (solo binario) ==" -ForegroundColor Cyan
$env:APP_VERSION = $Version
wails build -clean -platform $BuildArch

if (-not (Test-Path $BinDir)) { throw "No existe $BinDir. ¿Falló la compilación?" }

# Resuelve EXE
$ExePath = Join-Path $BinDir $AppExeName
if (-not (Test-Path $ExePath)) {
  $ExePath = Get-ChildItem $BinDir -Filter *.exe -Recurse -ErrorAction SilentlyContinue |
    Where-Object { $_.Name -notmatch '(setup|installer|install)' } |
    Sort-Object Length -Descending |
    Select-Object -First 1 -ExpandProperty FullName
}
if (-not $ExePath) { throw "No encontré el ejecutable principal en $BinDir." }

# -------- Instalador (opcional) --------
$InstallerPath = $null
if ($SkipInstaller) {
  Write-Host "== Saltando creación de instalador NSIS por -SkipInstaller ==" -ForegroundColor Yellow
} else {
  Write-Host "== Generando script de instalador NSIS con Wails ==" -ForegroundColor Cyan
  wails build -nsis -platform $BuildArch

  # Compilar NSIS manual si hay .nsi
  $NsisCandidates = @()
  $pathsToCheck = @(
    (Join-Path $Root "build\windows\installer\project.nsi"),
    (Join-Path $Root "build\windows\installer\installer.nsi"),
    (Join-Path $Root "build\nsis\project.nsi"),
    (Join-Path $Root "packaging\installer.nsi")
  )
  foreach ($p in $pathsToCheck) { if (Test-Path $p) { $NsisCandidates += $p } }

  if ($NsisCandidates.Count -eq 0) {
    Write-Warning "No se encontró ningún .nsi (build\\windows\\installer\\*.nsi, build\\nsis\\project.nsi o packaging\\installer.nsi)."
  } else {
    # Define ARCH para wails_tools.nsh
    $archLower = $BuildArch.ToLowerInvariant()
    $archDefine = @()
    if ($archLower -match 'amd64') {
      $archDefine += "/DARG_WAILS_AMD64_BINARY=`"$ExePath`""
    } elseif ($archLower -match 'arm64') {
      $archDefine += "/DARG_WAILS_ARM64_BINARY=`"$ExePath`""
    } else {
      Write-Warning "BuildArch '$BuildArch' no reconocido para defines NSIS. Intento AMD64 por defecto."
      $archDefine += "/DARG_WAILS_AMD64_BINARY=`"$ExePath`""
    }

    foreach ($nsi in $NsisCandidates) {
      Write-Host "== Ejecutando makensis (/V4) ==" -ForegroundColor Cyan
      Write-Host "NSI: $nsi"
      & "$MakensisExe" /V4 @archDefine "$nsi"
      if ($LASTEXITCODE -ne 0) {
        throw "NSIS falló con código $LASTEXITCODE. Revisa el log anterior."
      }
    }
  }

  # Localiza instalador real (ignora WebView2)
  $InstallerPath = Find-Installer -root $Root
}

# -------- Salida --------
Write-Host "EXE: $ExePath"
if ($InstallerPath) { Write-Host "INSTALADOR: $InstallerPath" }

# -------- Firma --------
function Sign-File($path) {
  Write-Host "Firmando: $path" -ForegroundColor Yellow
  if ($PfxPath) {
    if (-not (Test-Path $PfxPath)) { throw "No existe PFX: $PfxPath" }
    if ($PfxPassword) {
      & "$SigntoolExe" sign /fd SHA256 /tr $TimestampUrl /td SHA256 /f "$PfxPath" /p "$PfxPassword" "$path"
    } else {
      & "$SigntoolExe" sign /fd SHA256 /tr $TimestampUrl /td SHA256 /f "$PfxPath" "$path"
    }
  } elseif ($CertSubject) {
    & "$SigntoolExe" sign /fd SHA256 /tr $TimestampUrl /td SHA256 /n "$CertSubject" "$path"
  } else {
    throw "No se especificó ni -PfxPath ni -CertSubject. Usa -SkipSign si no quieres firmar."
  }
}

if (-not $SkipSign) {
  Write-Host "== Firmando artefactos ==" -ForegroundColor Cyan
  Sign-File $ExePath
  if ($InstallerPath) { Sign-File $InstallerPath }
} else {
  Write-Warning "Saltando firma por -SkipSign"
}

# -------- Hashes --------
Write-Host "== Generando SHA256SUMS.txt ==" -ForegroundColor Cyan
$SumFile = Join-Path $OutDir "SHA256SUMS.txt"
Remove-Item $SumFile -ErrorAction SilentlyContinue | Out-Null
$items = @($ExePath)
if ($InstallerPath) { $items += $InstallerPath }
foreach ($i in $items) {
  $h = Get-FileHash $i -Algorithm SHA256
  "{0}  {1}" -f $h.Hash.ToLower(), (Split-Path $i -Leaf) | Out-File -FilePath $SumFile -Append -Encoding utf8
}
Write-Host "SHA256SUMS en: $SumFile"

# -------- Resumen --------
Write-Host "`n== RESUMEN ==" -ForegroundColor Green
Write-Host "Versión:      $Version"
Write-Host "Arquitectura: $BuildArch"
Write-Host "EXE:          $ExePath"
if ($InstallerPath) { Write-Host "Instalador:   $InstallerPath" }
Write-Host "Hashes:       $SumFile"
Write-Host "Firmado:      " -NoNewline
if (-not $SkipSign) { Write-Host "Sí" } else { Write-Host "No" }
