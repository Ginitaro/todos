<template>
    <div class="panel todo-controls">
        <label class="panel-block task" v-for="todolist in todolistdata">
            <div class="item">
                <div class="item-check">
                    <input type="checkbox" v-bind:checked="todolist.Done" v-bind:name="todolist.ID">
                </div>
                <div class="item-name">
                    {{ todolist.Title }}
                </div>
                <button @click="removeTodo(todolist.ID)" class="delete" aria-label="delete"></button>
            </div>
            <ol>
                <li v-for="todo in todolist.Todos">
                    {{ todo.Name }}
                    <input type="checkbox" v-bind:checked="todo.Done" v-bind:name="todo.ID">
                </li>
            </ol>
        </label>
    </div>
</template>
    
<script>
export default {
    methods: {
        removeTodo: function(id) {
            this.$store.dispatch('removeTodoListItem', id)
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