// CORE
import Vue from 'vue'
import App from './App.vue'
import router from './router';
import store from './store';

// CSS
import '../assets/app.css'

new Vue({
    el: '#app',
    render: h => h(App)
});
