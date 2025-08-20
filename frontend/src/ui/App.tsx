import React, { useEffect, useState } from 'react'
import LauncherShell from '../components/LauncherShell'
import TitleBar from '../components/TitleBar'
import StatusBox from '../components/StatusBox'
import NewsBox from '../components/NewsBox'
import promo from '../assets/wallpaper.png'

declare global {
  interface Window {
    runtime: any
    go: any
  }
}

export default function App(){
  // estado real del updater
  const [status, setStatus]   = useState('Launcher started, have a good game!')
  const [totalP, setTotalP]   = useState(0)
  const [fileP, setFileP]     = useState(0)
  const [busy, setBusy]       = useState(false)
  const [installDir, setInstallDir] = useState<string>('')

  useEffect(() => {
    const off1 = window.runtime?.EventsOn?.('update:status', (s:string)=> setStatus(s))
    const off2 = window.runtime?.EventsOn?.('update:totalProgress', (p:number)=> setTotalP(Math.round((p||0)*100)))
    const off3 = window.runtime?.EventsOn?.('update:fileProgress',  (p:number)=> setFileP(Math.round((p||0)*100)))
    return () => { off1?.(); off2?.(); off3?.() }
  }, [])

  useEffect(() => {
    (async () => {
      try {
        const dir = await window.go.main.App.GetInstallDir()
        setInstallDir(dir)
      } catch (e) {
        console.error('GetInstallDir error:', e)
      }
    })()
  }, [])

  const handleGameStart = async () => {
    if (busy) return
    setBusy(true)
    setStatus('Checking...')
    setTotalP(0); setFileP(0)

    try {
      const manifestURL = '/client/manifest'
      if (!installDir) throw new Error('Resolviendo carpeta de instalación...')

      // 1) chequeo + actualización
      await window.go.main.App.UpdateFromManifest(manifestURL, installDir)

      // 2) lanzar el juego si todo OK
      setStatus('Launching game...')
      // ajusta al ejecutable real dentro de ./Client
      await window.go.main.App.StartGame()

      setStatus('Game launched!')
    } catch (e:any) {
      setStatus(`Error: ${e?.message || String(e)}`)
      setBusy(false) // reintentar
      return
    }

    setBusy(false)
  }

  return (
    <LauncherShell>
      {/* Barra personalizada (drag/no-drag) */}
      <TitleBar />

      {/* Header fino dentro del marco */}
      <div className="skin-header flex items-center justify-between">
        <div className="text-sm opacity-80">© AinhoSoft</div>
        <div className="text-sm opacity-80">Settings • WinMode</div>
      </div>

      {/* Cuerpo del launcher */}
      <div className="p-4 lg:p-6 grid grid-cols-1 lg:grid-cols-[1fr_360px] gap-6">
        {/* Promo / Showcase */}
        <section className="skin-panel overflow-hidden">
          <img src={promo} className="w-full object-cover" />
        </section>

        {/* Lateral derecho: Status + News (tu misma apariencia) */}
        <div className="flex flex-col gap-6">
          <StatusBox />
          <NewsBox />
        </div>

        {/* Barra inferior: mensaje + barras + START (misma apariencia) */}
        <div className="lg:col-span-2 flex flex-col lg:flex-row items-center gap-4 skin-panel p-4">
          <div className="flex-1 w-full">
            <div className="text-sm opacity-90">{status}</div>

            <div className="mt-2 text-xs opacity-80">Total</div>
            <div className="progress"><div style={{width:`${totalP}%`}}/></div>

            <div className="mt-2 text-xs opacity-80">Update</div>
            <div className="progress"><div style={{width:`${fileP}%`}}/></div>
          </div>

          <button
            className="start-btn disabled:opacity-60 disabled:cursor-not-allowed"
            onClick={handleGameStart}
            disabled={busy}
          >
            {busy ? 'PLEASE WAIT…' : 'GAME START'}
          </button>
        </div>
      </div>
    </LauncherShell>
  )
}
