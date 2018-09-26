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

if ('serviceWorker' in navigator) {
	window.addEventListener('load', () => {
		navigator.serviceWorker.register('/service-worker.js').then(registration => {
			console.log('SW registered: ', registration);
		})
		.catch(registrationError => {
			console.log('SW registration failed: ', registrationError);
		});
	});
}
