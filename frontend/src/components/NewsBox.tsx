import React from 'react'

export default function NewsBox(){
  return (
    <section className="skin-panel p-4">
      <div className="card-title mb-2">NOTICE & EVENTS</div>
      <div className="h-44 overflow-auto pr-2 space-y-3 text-sm">
        <article className="bg-white/5 rounded-lg p-2">
          <div className="font-semibold">Scheduled Maintenance</div>
          <div className="opacity-80">Servers under maintenance today.</div>
        </article>
        <article className="bg-white/5 rounded-lg p-2">
          <div className="font-semibold">Patch 1.0.1</div>
          <div className="opacity-80">Fixes for NPCs, portals, loot routing…</div>
        </article>
      </div>
    </section>
  )
}
