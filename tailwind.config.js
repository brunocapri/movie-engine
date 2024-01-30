/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'internal/templates/**/*.templ',
  ],
  darkMode: 'class',
  theme: {
    extend: {
      fontFamily:{
        'sans': ['IBM Plex Mono', 'sans-serif'], // This sets it as the default sans-serif font
        'serif': ['IBM Plex Mono', 'serif'], // This sets it as the default serif font
        'mono': ['IBM Plex Mono', 'monospace'] // This sets it as the default monospace font
      }
    }
  },
  plugins: [],
  corePlugins: {
    preflight: true,
  }
}

