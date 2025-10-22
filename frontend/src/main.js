import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import 'aos/dist/aos.css'
import AOS from 'aos'
AOS.init({ duration: 1000, once: true })

createApp(App).use(router).use(store).mount('#app')
