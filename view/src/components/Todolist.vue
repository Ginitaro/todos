<template>
    <div class="panel todo-controls">
        <label class="panel-block task" v-for="todolist in todolistdata">
            <div class="item">
                <div class="item-name">
                    {{ todolist.Title }}
                </div>
                
                <createTodo v-bind:parent="todolist.ID"></createTodo>
                
                <button @click="removeTodoListItem(todolist.ID)" class="delete" aria-label="delete"></button>
                
            </div>
            <ol>
                <li v-for="todo in todolist.Todos">
                    {{ todo.Name }}
                    <input type="checkbox" v-on:change="toggleTodo($event, todolist.ID, todo.ID)" v-model="todo.Done">
                </li>
            </ol>
        </label>
    </div>
</template>
    
<script>
import createTodo from './createTodo.vue'
export default {
    components: {
        createTodo
    },
    methods: {
        removeTodoListItem: function(id) {
            this.$store.dispatch('removeTodoListItem', id)
        },
        toggleTodo: function(e, todolistId, todoId) {
            let checked = e.target.checked;
            this.$store.dispatch('toggleTodo', {todolistId, todoId, checked})
        }
    },
    computed: {
        todolistdata() {
            return this.$store.getters.getTodoList;
        }
    },
    created() {
        this.$store.dispatch('getTodoList');
    }
}
</script>