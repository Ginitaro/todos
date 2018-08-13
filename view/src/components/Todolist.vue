<template>
    <v-container fluid grid-list-lg>
        <v-layout row>
            <v-flex xs12 sm4 offset-sm4>
                <v-card>
                    <v-toolbar color="blue" dark extended tabs class="pa-0">
                        
                        <createTodolist></createTodolist>
                        
                        <v-tabs
                        slot="extension"
                        centered
                        color="transparent"
                        slider-color="white"
                        >
                            <v-tab>All</v-tab>
                            <v-tab>Unfinished</v-tab>
                            <v-tab>Finished</v-tab>
                        </v-tabs>
                        
                    </v-toolbar>
                    
                    <v-expansion-panel>
                        <v-expansion-panel-content v-for="todolist in todolistdata" class="py-2">
                            <div slot="header" class="todolist-title" @click="todolist.expanded = !todolist.expanded">
                                <div>
                                    <v-btn outline small v-if="todolist.expanded" color="red" v-on:click="removeTodoListItem(todolist.ID)">
                                        <v-icon>clear</v-icon>Remove
                                    </v-btn>
                                    <v-btn outline fab small v-else color="grey">
                                        <v-icon>list</v-icon>
                                    </v-btn>
                                    <label>{{ todolist.Title }}</label>     
                                </div>
                            </div>
                            <v-card>                              
                                <createTodo v-bind:parent="todolist.ID"></createTodo>                 
                                <v-list>
                                    
                                    <v-list-tile v-for="todo in todolist.Todos">
                                        <v-list-tile-action>
                                            <v-checkbox v-on:change="toggleTodo($event, todolist.ID, todo.ID)" v-model="todo.Done"></v-checkbox>
                                        </v-list-tile-action>                           
                                        <v-list-tile-content>
                                            <v-list-tile-title>{{ todo.Name }}</v-list-tile-title>
                                        </v-list-tile-content>
                                        <v-list-tile-action>
                                            <v-btn flat icon color="error" v-on:click="removeTodo(todolist.ID, todo.ID)">
                                                <v-icon>clear</v-icon>
                                        </v-btn>
                                        </v-list-tile-action>
                                    </v-list-tile>
                                     
                                </v-list>
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
        },
    },
    computed: {
        todolistdata(e) {
            
            let getTodoList = this.$store.getters.getTodoList
            
            if (getTodoList !== null){
                for (var i = getTodoList.length - 1; i >= 0; i--) {
                    this.$set(getTodoList[i], "expanded", false)
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