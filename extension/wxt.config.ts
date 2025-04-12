import { defineConfig } from 'wxt';

// See https://wxt.dev/api/config.html
export default defineConfig({
  modules: ['@wxt-dev/module-react'],
  manifestVersion: 3,
  manifest: {
    name: "ytb2mpv - Open YouTube Videos in mpv",
    short_name: "ytb2mpv",
    description: "Easily open your YouTube videos in mpv or download them just one click.",
    host_permissions: [
      "http://localhost:53918/*",
    ],
    author: {
      email: "michio.haiyaku@gmail.com"
    },
    homepage_url: "https://github.com/michioxd/ytb2mpv",
    web_accessible_resources: [
      {
        resources: [
          "*"
        ],
        matches: [
          "https://*.youtube.com/*"
        ]
      }
    ]
  },
  vite: () => ({
    build: {
      cssCodeSplit: false,
      minify: "terser",
      terserOptions: {
        parse: {
          html5_comments: false,
        },
        format: {
          comments: false
        }
      }
    }
  })
});
