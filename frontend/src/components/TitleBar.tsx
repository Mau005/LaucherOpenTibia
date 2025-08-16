import React from 'react'
const R = (window as any).runtime

export default function TitleBar(){
  return (
    <div className="titlebar drag flex items-center justify-between px-3 py-2">
      <div className="text-sm opacity-80 select-none">AinhoSoft</div>
      <div className="flex items-center gap-2 no-drag">
        <button className="tb-btn" onClick={()=>R.WindowMinimise()}>—</button>
        <button className="tb-btn" onClick={()=>R.WindowToggleMaximise()}>▢</button>
        <button className="tb-btn tb-close" onClick={()=>R.Quit()}>✕</button>
      </div>
    </div>
  )
}
