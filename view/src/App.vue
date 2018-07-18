<template>
	<div id="app">
		<createTodolist></createTodolist>
		<router-view></router-view>
		<todolist></todolist>
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