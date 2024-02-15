import { createApp } from 'vue';
// Import Pinia into your config file
import { createPinia } from 'pinia';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
// import { md3 } from 'vuetify/blueprints'
import App from './App.vue';
//
const vuetify = createVuetify({
  components,
  directives,
});
createApp(App)
// Add the line below to the file to instantiate it
  .use(createPinia())
  .use(vuetify)
  .mount('#app');
