// CORE
import Vue from 'vue'
import App from './App.vue'
import router from './router';
import store from './store';

// CSS
import 'bulma/css/bulma.css'
import '../assets/app.css'

new Vue({
	router,
	store,
    el: '#app',
    render: h => h(App)
});
