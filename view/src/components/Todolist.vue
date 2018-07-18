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
                
                
                <form class="panel-block" @submit.prevent>
                    <div class="field is-grouped">
                        <div class="control is-expanded is-medium" v-bind:class="{ active: is-loading }">
                            <input class="input is-medium" type="text" placeholder="Type something here..." v-model="todoName">
                        </div>
                        <button class="button is-medium is-info" v-on:click="addTodo(todolist.ID, todoName)">Add</button>
                    </div>
                </form>
                
                <button @click="removeTodoListItem(todolist.ID)" class="delete" aria-label="delete"></button>
                
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
        removeTodoListItem: function(id) {
            this.$store.dispatch('removeTodoListItem', id)
        },
        addTodo: function(id, text) {
            this.$store.dispatch('addTodo', {text, id});
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