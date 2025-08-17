import React, { useEffect, useState } from 'react'

// íconos opcionales según IconID (ajusta rutas si quieres)
import iconCommunity from '../assets/newsicon_community_small.png'
import iconDev       from '../assets/newsicon_development_small.png'
import iconSupport   from '../assets/newsicon_support_small.png'
import iconTech      from '../assets/newsicon_technical_small.png'

declare global {
  interface Window { go:any }
}

type NewsItem = {
  IconID: number
  Description: string
  CreatedAt: string
}

const ICONS: Record<number, string> = {
  1: iconCommunity,
  2: iconDev,
  3: iconSupport,
  4: iconTech,
}

export default function NewsBox(){
  const [items, setItems]   = useState<NewsItem[]>([])
  const [loading, setLoad]  = useState(true)
  const [error, setError]   = useState<string | null>(null)

  useEffect(() => {
    const url = '/get_news_short' // <-- AJUSTA AQUÍ
    window.go.main.App.GetNews(url)
      .then((list: NewsItem[]) => { setItems(list); setError(null) })
      .catch((e: any) => setError(e?.message || String(e)))
      .finally(() => setLoad(false))
  }, [])

  return (
    <section className="skin-panel p-4">
      <div className="card-title mb-2">NOTICE &amp; EVENTS</div>

      {loading && <div className="text-sm opacity-80">Loading news…</div>}
      {error && !loading && <div className="text-sm text-red-400">Error: {error}</div>}

      {!loading && !error && (
        <div className="h-44 overflow-auto pr-2 space-y-2 text-sm">
          {items.length === 0 && <div className="opacity-70">No news available.</div>}

          {items.map((n, i) => {
            const date = new Date(n.CreatedAt)
            return (
              <article
                key={i}
                className="bg-white/5 rounded-lg px-3 py-2 w-full flex items-start gap-3"
              >

                <img
                  src={ICONS[n.IconID] || ICONS[1]}
                  className="w-8 h-8 object-contain flex-shrink-0"
                  alt=""
                />

                <div className="flex-1 min-w-0">
                  <div className="flex items-center justify-between gap-3">
                    <span className="font-semibold truncate">Update</span>
                    <span className="text-xs opacity-70 whitespace-nowrap">
                      {date.toLocaleDateString()}
                    </span>
                  </div>

                  <div
                    className={
                      // usa UNA de las dos líneas siguientes:
                      // "opacity-80 mt-0.5 whitespace-pre-wrap line-clamp-2"  // <- si tienes plugin line-clamp
                      "opacity-80 mt-0.5 whitespace-pre-wrap break-words"     // <- sin plugin, deja flujo natural
                    }
                  >
                    {n.Description}
                  </div>
                </div>
              </article>
            )
          })}
        </div>
      )}
    </section>
  )
}
