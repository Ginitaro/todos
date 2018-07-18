<template>
  <div class="full-width center-content">
    <hello-component name="haha" />
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