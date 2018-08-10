<template>
    <v-container fluid grid-list-lg>
        <v-layout row>
            <v-flex xs12 sm6 offset-sm3>
                <v-card>
                    
                    <createTodolist></createTodolist>
                    
                    <v-expansion-panel>
                        <v-expansion-panel-content v-for="todolist in todolistdata">
                            <div slot="header">
                                <span class="circle"><i class="fa fa-list"></i></span> 
                                <label class="todo-title">{{ todolist.Title }}</label>
                            </div>
                            <v-card>
                                <v-btn color="error" v-on:click="removeTodoListItem(todolist.ID)">Remove Category</v-btn>
                                <createTodo v-bind:parent="todolist.ID"></createTodo>
                                <v-card-text v-for="todo in todolist.Todos">
                                    {{ todo.Name }}
                                     <input type="checkbox" v-on:change="toggleTodo($event, todolist.ID, todo.ID)" v-model="todo.Done">
                                </v-card-text>
                            </v-card>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                    
                </v-card>
            </v-flex>
        </v-layout>
    </v-container>
</template>
    
<script>
import createTodo from './createTodo.vue'
import createTodolist from './createTodolist.vue'

export default {    
    components: {
        createTodo,
        createTodolist
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
            let getTodoList = this.$store.getters.getTodoList
            for (var i = getTodoList.length - 1; i >= 0; i--) {
                this.$set(getTodoList[i], "isOpen", false)
            }
            return getTodoList
        }
    },
    created() {
        this.$store.dispatch('getTodoList');
    }
}
</script>