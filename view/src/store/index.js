import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
import VueAxios from 'vue-axios'

// For self-signed certificate
// const instance = axios.create({
//   baseURL: 'https://localhost:3000/api/',
// });

Vue.use(Vuex, VueAxios, axios)

export default new Vuex.Store({
	state: {
		todolist: []
	},
	mutations: {
		setTodoList(state, todolist) {
			
			if (todolist !== null){
                for (var i = todolist.length - 1; i >= 0; i--) {
                    var done = true
                    if ("Todos" in todolist[i]){
                        if (todolist[i].Todos !== null) {
                            for (var j = todolist[i].Todos.length - 1; j >= 0; j--) {
                                if (todolist[i].Todos[j].Done == false)
                                    done = false
                            }
                        } else
                            done = false
                    } 
                    todolist[i]['taskDone'] = done
                }                            
            }
			
			state.todolist = todolist;	
			
		},
		removeTodoListItem(state, removedId) {
			var index = state.todolist.findIndex(todolistitem => todolistitem.ID === removedId);
            state.todolist.splice(index, 1);
		},
		removeTodo(state, params) {
			var todolistitem_index = state.todolist.findIndex(todolistitem => todolistitem.ID === params.todolistId);
			var todo_index = state.todolist[todolistitem_index].Todos.findIndex(todo => todo.ID === params.todoId);
            state.todolist[todolistitem_index].Todos.splice(todo_index, 1);
		},
		addTodoListItem(state, params) {
			if (state.todolist == null){
				state.todolist = []
			}
			state.todolist.push(params)
		},
		addTodo(state, params) {
			var todolistitem_index = state.todolist.findIndex(todolistitem => todolistitem.ID === params.id);
			if (state.todolist[todolistitem_index].Todos == null){
				state.todolist[todolistitem_index].Todos = []
			}
			state.todolist[todolistitem_index].Todos.push(params.newTodo)
		},
		toggleTodo(state, params) {
			this.dispatch('getTodoList')	
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
				commit('addTodoListItem', this.response)
				this.dispatch('getTodoList')
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err', xhr)
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
				//this.dispatch('getTodoList')
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
				Vue.set(params, "newTodo", this.response)
				//commit('addTodo', params)
				this.dispatch('getTodoList')
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err', xhr)
			});
		},
		toggleTodo({ commit }, params) {
			
			const formData = new URLSearchParams();
			formData.append('status', params.status);
			
			axios.patch('api/'+params.todolistId+'/todo/'+params.todoId+'/toggle', formData)
			.then(xhr => {
				this.response = xhr.data;
				commit('toggleTodo', params)
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err', err)
			});
		},
		removeTodo({ commit }, params) {
			
			const formData = new URLSearchParams();

			formData.append('todolist_id', params.todolistId);
			formData.append('todo_id', params.todoId);
			
			axios.post('api/'+params.todolistId+'/todo/'+params.todoId+'/remove', formData)
			.then(xhr => {
				this.response = xhr.data;
				console.log(this.response)
				commit('removeTodo', params);
			})
			.catch(xhr => {
				this.response = xhr.status;
				console.error('err')
			});
		},
	},
	getters: {
		getTodoListId: state => {
			return state
		},
		getTodoList: state => {
			return state.todolist
		},
		filteredByFinished: state => {
			return state.todolist.filter(todolist => todolist.taskDone)
		},
		filteredByUnfinished: state => {
			return state.todolist.filter(todolist => !todolist.taskDone)
		}
  	}
})
