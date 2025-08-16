/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html','./src/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        primary: '#d4af37',
        dark: '#0b0b0c',
        panel: 'rgba(20,20,25,0.75)'
      },
      boxShadow: {
        glow: '0 0 30px rgba(212,175,55,0.5)'
      },
      backdropBlur: {
        xs: '2px',
      }
    },
  },
  plugins: [],
}
