import React from 'react'

export default function NewsPanel(){
  const md = `## Scheduled Maintenance
Our servers will be under maintenance today.

- New outfits
- Bug fixes
`
  return (
    <section className="glass p-4">
      <h3 className="card-title mb-2">News</h3>
      <div className="h-56 overflow-auto pr-2 space-y-3">
        <article className="bg-white/5 p-3 rounded-xl">
          <div className="font-semibold mb-1">Scheduled Maintenance</div>
          <p className="text-sm opacity-80">Our servers will be under maintenance today. New outfits & bug fixes.</p>
        </article>
        <article className="bg-white/5 p-3 rounded-xl">
          <div className="font-semibold mb-1">Patch 1.0.1</div>
          <p className="text-sm opacity-80">NPCs fixed, portals repaired, looting changes & free premium week.</p>
        </article>
      </div>
    </section>
  )
}
