import React from 'react'

export default function StatusBox(){
  return (
    <section className="skin-panel p-4">
      <div className="card-title mb-2">SERVER STATUS</div>
      <div className="space-y-1 text-sm">
        <div className="flex justify-between"><span>Connect Server:</span><span className="text-red-400">OFFLINE</span></div>
        <div className="flex justify-between"><span>Game Server:</span><span className="text-red-400">OFFLINE</span></div>
        <div className="flex justify-between"><span>Server Time:</span><span className="text-green-400">14:55:32</span></div>
      </div>
    </section>
  )
}
