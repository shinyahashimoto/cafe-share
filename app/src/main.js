import Vue from "vue";
import App from "./App.vue";
import store from "./store";
import router from "./router";
import vuetify from "./plugins/vuetify";
import "./validate/validate.js";
import VueCtkDateTimePicker from "vue-ctk-date-time-picker";
import "vue-ctk-date-time-picker/dist/vue-ctk-date-time-picker.css";

Vue.component("VueCtkDateTimePicker", VueCtkDateTimePicker);

Vue.config.productionTip = false;

// vue-bootstarp
import BootstrapVue from "bootstrap-vue";
import "bootstrap/dist/css/bootstrap.css";
import "bootstrap-vue/dist/bootstrap-vue.css";
Vue.use(BootstrapVue);

import * as GoogleMaps from "vue2-google-maps";
Vue.use(GoogleMaps, {
  load: {
    key: "AIzaSyCiWBKPMaAeah1OvcZxoSzMehqXD8JjHv4",
    libraries: "places",
  },
});

new Vue({
  vuetify,
  store,
  router,
  render: (h) => h(App),
}).$mount("#app");
