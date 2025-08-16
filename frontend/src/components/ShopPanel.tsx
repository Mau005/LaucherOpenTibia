import React from 'react'

const items = [
  {name: 'Starter Pack', price: '$2.99'},
  {name: 'Premium 30d', price: '$6.99'},
  {name: 'Outfit Bundle', price: '$3.49'},
  {name: 'XP Boost (7d)', price: '$1.99'},
]

export default function ShopPanel(){
  return (
    <section className="glass p-4">
      <h3 className="card-title mb-2">Shop</h3>
      <ul className="divide-y divide-white/10">
        {items.map(i => (
          <li key={i.name} className="py-2 flex items-center gap-3">
            <div className="w-8 h-8 rounded bg-white/10" />
            <div className="flex-1">
              <div className="font-medium">{i.name}</div>
              <div className="text-xs opacity-70">Instant delivery</div>
            </div>
            <div className="mr-3 opacity-80">{i.price}</div>
            <button className="btn btn-primary">Buy</button>
          </li>
        ))}
      </ul>
    </section>
  )
}
