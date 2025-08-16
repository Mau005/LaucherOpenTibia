import React from 'react'
import mage from '../assets/mage_overlay.png'

export default function LauncherShell({ children }: { children: React.ReactNode }) {
  return (
    <div className="relative h-screen">
      {/* contenedor centrado, sin fondos alrededor */}
      <div className="relative">
        {/* marco interno */}
        <div className="skin-panel relative">
          {children}
        </div>

        {/* MAGO sobresaliendo por fuera del panel */}
        <img
          src={mage}
          alt=""
          className="
            pointer-events-none
            absolute
            -left-2        /* sale por la izquierda */
            -top-5         /* súbelo un poco más si quieres: -top-28/-top-32 */
            h-[35vh]        /* alto grande para que se vea poderoso */
            z-50
            drop-shadow-[0_24px_48px_rgba(0,0,0,0.55)]
          "
        />
      </div>
    </div>
  )
}
