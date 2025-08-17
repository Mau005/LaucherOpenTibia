import React, { useEffect, useMemo, useRef, useState } from 'react'

declare global { interface Window { go:any } }

type InfoResponse = {
  ServerInfo: { Uptime: string }
  Players: { Online: number; Max: number; Peak?: number }
}

function fmtHMS(totalSec: number) {
  const s = Math.max(0, Math.floor(totalSec))
  const hh = String(Math.floor(s / 3600)).padStart(2, '0')
  const mm = String(Math.floor((s % 3600) / 60)).padStart(2, '0')
  const ss = String(s % 60).padStart(2, '0')
  return `${hh}:${mm}:${ss}`
}

export default function StatusBox() {
  const [online, setOnline] = useState(false)
  const [players, setPlayers] = useState({ Online: 0, Max: 0 })
  const [uptime, setUptime] = useState(0) // segundos
  const timerRef = useRef<number | null>(null)

  // 1) Cargar /client/info una sola vez (vía Go)
  useEffect(() => {
    const url = '/client/info' // puedes cambiarlo luego en Go
    window.go.main.App.GetServerInfo(url)
      .then((info: InfoResponse) => {
        const upStr = (info?.ServerInfo?.Uptime || '').trim()
        const up = Number(upStr) || 0
        const isOnline = up > 0
        setOnline(isOnline)
        setPlayers({
          Online: info?.Players?.Online ?? 0,
          Max: info?.Players?.Max ?? 0,
        })
        setUptime(isOnline ? up : 0)
      })
      .catch(() => {
        setOnline(false)
        setPlayers({ Online: 0, Max: 0 })
        setUptime(0)
      })
  }, [])

  // 2) Si está ONLINE, contar uptime localmente; si está OFFLINE, quedarse en 0
  useEffect(() => {
    if (timerRef.current) {
      window.clearInterval(timerRef.current)
      timerRef.current = null
    }
    if (online) {
      timerRef.current = window.setInterval(() => {
        setUptime(prev => prev + 1)
      }, 1000)
    } else {
      setUptime(0)
    }
    return () => {
      if (timerRef.current) window.clearInterval(timerRef.current)
      timerRef.current = null
    }
  }, [online])

  const statusClass = useMemo(
    () => (online ? 'text-green-400' : 'text-red-400'),
    [online]
  )
  const statusLabel = online ? 'ONLINE' : 'OFFLINE'

  return (
    <section className="skin-panel p-4">
      <div className="card-title mb-2">SERVER STATUS</div>

      <div className="space-y-1 text-sm">
        <div className="flex justify-between">
          <span>Connect Server:</span>
          <span className={statusClass}>{statusLabel}</span>
        </div>

        <div className="flex justify-between">
          <span>Game Server:</span>
          <span className={statusClass}>{statusLabel}</span>
        </div>

        <div className="flex justify-between">
          <span>Server Time:</span>
          <span className={online ? 'text-green-400' : 'opacity-70'}>
            {fmtHMS(uptime)}
          </span>
        </div>

        <div className="flex justify-between">
          <span>Players Online:</span>
          <span className={online ? 'text-green-300' : 'opacity-70'}>
            {players.Online}{players.Max ? ` / ${players.Max}` : ''}
          </span>
        </div>
      </div>
    </section>
  )
}
