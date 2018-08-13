// CORE
import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import router from './router';
import store from './store';

// CSS
import '../assets/app.css'

Vue.use(Vuetify)

new Vue({
	router,
	store,
    el: '#app',
    render: h => h(App)
});

