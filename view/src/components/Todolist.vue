<template>
    <v-container fluid grid-list-lg>
        <v-layout row>
            <v-flex xs12 sm6 offset-sm3>
                <v-card>
                    
                    <createTodolist></createTodolist>
                    
                    <v-expansion-panel focusable>
                        <v-expansion-panel-content v-for="todolist in todolistdata">
                            <div slot="header">
                                <v-btn outline fab small color="grey">
                                    <v-icon>list</v-icon>
                                </v-btn>
                                <label class="todo-title">{{ todolist.Title }}</label>
                                <v-btn color="error" flat small v-on:click="removeTodoListItem(todolist.ID)">Remove Category</v-btn>
                            </div>
                            <v-card>
                                                             
                                <createTodo v-bind:parent="todolist.ID"></createTodo>
                                                       
                                <v-card-text class="todo" v-for="todo in todolist.Todos">
                                    
                                    <v-list-tile>
                                        <v-list-tile-action>
                                            <v-checkbox v-on:change="toggleTodo($event, todolist.ID, todo.ID)" v-model="todo.Done"></v-checkbox>
                                        </v-list-tile-action>
                                        <v-list-tile-content @click="todo.Done = !todo.Done">
                                            <v-list-tile-title>{{ todo.Name }}</v-list-tile-title>
                                        </v-list-tile-content>
                                    </v-list-tile>
                                     
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
        toggleTodo: function(status, todolistId, todoId) {
            this.$store.dispatch('toggleTodo', {todolistId, todoId, status})
        },
        removeTodo: function(todolistId, todoId) {
            this.$store.dispatch('removeTodo', {todolistId, todoId})
        }
    },
    computed: {
        todolistdata() {
            
            let getTodoList = this.$store.getters.getTodoList
            
            if (getTodoList !== null){
                for (var i = getTodoList.length - 1; i >= 0; i--) {
                    this.$set(getTodoList[i], "isOpen", false)
                }
                return getTodoList
            } else {
                return null
            }

        }
    },
    created() {
        this.$store.dispatch('getTodoList');
    }
}
</script>