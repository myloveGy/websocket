import Vue from 'vue'
import HomeComponent from './components/Home.vue'
import './index.css'
new Vue({
  el: "#app",
  template: `<home-component/>`,
  components: {
    HomeComponent,
  }
})


