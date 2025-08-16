import React from 'react'

type Props = {
  status: string
  totalProgress: number
  overallProgress: number
  server: 'ONLINE'|'OFFLINE'
  onStart: () => void
}

export default function StatusPanel({status,totalProgress,overallProgress,server,onStart}:Props){
  return (
    <section className="glass p-4">
      <div className="grid md:grid-cols-2 gap-4 items-center">
        <div className="space-y-2">
          <div className="text-sm opacity-80">Current: <span className="opacity-100">{status}</span></div>
          <div className="text-sm opacity-80">Total progress:</div>
          <div className="progress"><div style={{width: `${Math.round(totalProgress*100)}%`}}/></div>
          <div className="text-sm opacity-80">Overall:</div>
          <div className="progress"><div style={{width: `${Math.round(overallProgress*100)}%`}}/></div>
          <div className="text-sm">Server: <span className={server==='ONLINE'?'text-green-400':'text-red-400'}>{server}</span></div>
        </div>
        <div className="flex items-center justify-center">
          <button className="btn btn-primary text-2xl px-12 py-4 rounded-2xl" onClick={onStart}>START</button>
        </div>
      </div>
    </section>
  )
}
