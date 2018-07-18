const routes = [
	{ path: '/', todolist },
	{ path: '/create', component: createTodolist },
]

const router = new VueRouter({
	routes, // short for `routes: routes`
})
