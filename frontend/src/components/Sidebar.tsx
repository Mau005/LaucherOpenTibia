import React from 'react'

const links = [
  {label: 'Website', href: '#'},
  {label: 'Forum', href: '#'},
  {label: 'Ranking', href: '#'},
  {label: 'Statistics', href: '#'},
  {label: 'News', href: '#'},
]

export default function Sidebar(){
  return (
    <aside className="glass p-4 flex flex-col gap-3">
      {links.map(l => (
        <button key={l.label} className="side-btn">
          <span className="font-semibold">{l.label}</span>
          <span>›</span>
        </button>
      ))}
      <div className="mt-auto flex items-center justify-end gap-3 opacity-80">
        <span>💬</span><span>🐦</span><span>🎮</span>
      </div>
    </aside>
  )
}
