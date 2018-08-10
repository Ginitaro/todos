<template>
	<div id="app">
		<v-app>
			<router-view></router-view>
			<todolist></todolist>
		</v-app>
	</div>
</template>

<script>
import todolist from './components/Todolist.vue'
import createTodolist from './components/createTodolist.vue'

export default {
	components: {
		todolist,
		createTodolist
	}
}
</script>