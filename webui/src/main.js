import { createApp } from 'vue';
// Import Pinia into your config file
import { createPinia } from 'pinia';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import { aliases, fa } from 'vuetify/iconsets/fa'
import App from './App.vue';

import 'primeflex/primeflex.css'
import "primeicons/primeicons.css";
import 'primevue/resources/themes/aura-light-green/theme.css'
import PrimeVue from 'primevue/config';



// import '@fortawesome/fontawesome-free/css/all.css' // Ensure your project is capable of handling css files
// const vuetify = createVuetify({
//   components,
//   directives,
//
//   icons: {
//     defaultSet: 'fa',
//     aliases,
//     sets: {
//       fa,
//     },
//   },
// });


createApp(App)
// Add the line below to the file to instantiate it
  .use(createPinia())
  // .use(vuetify)
  .use(PrimeVue)
  .mount('#app');
