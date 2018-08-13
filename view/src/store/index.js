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
		removeTodoListItem(state, removedId) {
			var index = state.todolist.findIndex(todolistitem => todolistitem.ID === removedId);
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
		},
		addTodo({ commit }, params) {
			
			const formData = new URLSearchParams();
			formData.append('name', params.text);
			
			
			axios.post('api/'+params.id+'/create', formData)
			.then(xhr => {
				this.response = xhr.data;
				this.dispatch('getTodoList')
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err')
			});
		},
		toggleTodo({ commit }, params) {
			
			const formData = new URLSearchParams();
			formData.append('status', params.status);
			
			axios.patch('api/'+params.todolistId+'/todo/'+params.todoId+'/toggle', formData)
			.then(xhr => {
				this.response = xhr.data;
				console.log(this.response)
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err')
			});
		}
	},
	getters: {
		getTodoList: state => {
			return state.todolist
		},
		getTodoListId: state => {
			return state
		}
  	}
})
