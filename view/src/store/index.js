import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(Vuex, VueAxios, axios)

export default new Vuex.Store({
	state: {
		todolist: []
	},
	mutations: {
		setTodoList(state, todolist) {
			state.todolist = todolist;
		},
		removeTodoListItem(state, todolistitem) {
			var index = state.todolist.findIndex(todo => todo.ID === todolistitem);
            state.todolist.splice(index, 1);
		}
	},
	actions: {
		getTodoList({ commit }) {
			axios
				.get('api/get_todolist')
				.then(response => {
					commit('setTodoList', response.data)
				})
		},
		addTodoListItem({ commit }, text) {
			axios.post('api/create', {
				title: text
			})
			.then(xhr => {
				this.response = xhr.data;
				this.dispatch('getTodoList')
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err')
			});
		},
		removeTodoListItem({ commit }, todolistitem) {
			console.log('Remove List-item with ID', todolistitem);

			const params = new URLSearchParams();
			params.append('todolist_id', todolistitem);

			axios.post('api/remove', params)
			.then(xhr => {
				commit('removeTodoListItem', todolistitem);
				this.response = xhr.data;
				// this.dispatch('getTodoList')
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('Failed ->', xhr)
			});
		}
	},
	getters: {
		getTodoList: state => {
			return state.todolist
		}
  	}
})
