// main.js
import Vue from "vue";
import App from "./App.vue";
import router from "./router"; // Importa el enrutador desde el archivo router.js

Vue.config.productionTip = false;

new Vue({
  render: h => h(App),
  router, 
}).$mount("#app");
