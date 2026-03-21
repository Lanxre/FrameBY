import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
	app: {
		head: {
			title: "app",
			charset: "utf-8",
			viewport: "width=device-width, initial-scale=1",
		},
	},
	devtools: { enabled: true },

	modules: ["@pinia/nuxt", "@vueuse/nuxt", "@nuxtjs/color-mode", "@nuxt/icon"],

	routeRules: {
	  "/api/_nuxt_icon/**": { proxy: { to: "" } },
    "/api/**": { 
      proxy: "http://localhost:8080/api/**",
    },
  },

	colorMode: {
		preference: "light",
		fallback: "light",
		classSuffix: "",
		storageKey: "theme",
	},

	icon: {
    clientBundle: {
      scan: true,
      includeCustomCollections: true,
    },
    serverBundle: "local",
  },

	vite: {
		plugins: [tailwindcss()],
	},

	css: ["~/assets/css/global.css"],

	postcss: {
		plugins: {
			autoprefixer: {},
		},
	},

	sourcemap: {
		client: "hidden",
	},

	runtimeConfig: {
		public: {
			apiBase: import.meta.env.NUXT_PUBLIC_API_BASE || "http://localhost:8080",
		},
	},

	compatibilityDate: "latest",
});
