<template>
    <v-container fluid grid-list-lg>
        <v-layout row>
            <v-flex xs12 sm4 offset-sm4>
                <v-card>
                    <v-toolbar color="blue darken-4" dark extended tabs class="pa-0">
                        
                        <createTodolist></createTodolist>
                        
                        <v-tabs
                        slot="extension"
                        centered
                        color="transparent"
                        slider-color="white"
                        v-model="active"
                        >
                            <v-tab href="#all">All</v-tab>
                            <v-tab href="#unfinished">Unfinished</v-tab>
                            <v-tab href="#finished">Finished</v-tab>
                        </v-tabs>
                        
                    </v-toolbar>
                    
                    <v-tabs-items v-model="active">
                        <v-tab-item id="all" ref="all">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in alltodolistdata" :key="todolist.ID" class="py-2">
                                    <div slot="header" class="todolist-title">
                                        <div>
                                            <v-btn outline fab small color="black">
                                                <v-icon>list</v-icon>
                                            </v-btn>
                                            <label>{{ todolist.Title }}</label>
                                        </div>
                                    </div>
                                    <v-card>
                                        <v-btn flat small color="error" v-on:click="openDialog(todolist.ID)">Remove</v-btn>    
                                        <createTodo v-bind:parent="todolist.ID"></createTodo>                 
                                        <v-list>
                                            
                                            <v-list-tile v-for="(todo, index) in todolist.Todos" :key="index" @click="toggleTodo(!todo.Done, todolist.ID, todo.ID)" ripple>
                                                <v-list-tile-action>
                                                    <v-checkbox @click="" v-model="todo.Done"></v-checkbox>
                                                </v-list-tile-action>                           
                                                <v-list-tile-content>
                                                    <v-list-tile-title class="todo-name" v-bind:class="{ done: todo.Done }">{{ todo.Name }}</v-list-tile-title>
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
                        </v-tab-item>
                        <v-tab-item id="unfinished">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in unifinishedtodolistdata" :key="todolist.ID" class="py-2">
                                    <div slot="header" class="todolist-title">
                                        <div>
                                            <v-btn outline fab small color="grey">
                                                <v-icon>list</v-icon>
                                            </v-btn>
                                            <label>{{ todolist.Title }}</label>
                                        </div>
                                    </div>
                                    <v-card>
                                        <v-btn flat small color="error" v-on:click="removeTodoListItem(todolist.ID)">Remove</v-btn>                          
                                        <createTodo v-bind:parent="todolist.ID"></createTodo>                 
                                        <v-list>
                                            
                                            <v-list-tile v-for="(todo, index) in todolist.Todos" :key="index" @click="toggleTodo(!todo.Done, todolist.ID, todo.ID)" ripple>
                                                <v-list-tile-action>
                                                    <v-checkbox @click="" v-model="todo.Done"></v-checkbox>
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
                        </v-tab-item>
                        <v-tab-item id="finished">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in finishedtodolistdata" :key="todolist.ID" class="py-2">
                                    <div slot="header" class="todolist-title">
                                        <div>
                                            <v-btn outline fab small color="grey">
                                                <v-icon>list</v-icon>
                                            </v-btn>
                                            <label>{{ todolist.Title }}</label>  
                                        </div>
                                    </div>
                                    <v-card>
                                        <v-btn flat small color="error" v-on:click="removeTodoListItem(todolist.ID)">Remove</v-btn>                               
                                        <createTodo v-bind:parent="todolist.ID"></createTodo>                 
                                        <v-list>
                                            
                                            <v-list-tile v-for="(todo, index) in todolist.Todos" :key="index" @click="toggleTodo(!todo.Done, todolist.ID, todo.ID)" ripple>
                                                <v-list-tile-action>
                                                    <v-checkbox @click="" v-model="todo.Done"></v-checkbox>
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
                        </v-tab-item>
                    </v-tabs-items>
                    
                </v-card>
            </v-flex>
        </v-layout>
        <v-dialog v-model="dialog" max-width="290">
            <v-card>
                <v-card-title class="headline">Are you sure you want to remove whole category?</v-card-title>

                <v-card-text>
                    Todos within it will be lost.
                </v-card-text>

                <v-card-actions>
                    <v-spacer></v-spacer>

                    <v-btn
                        color="blue darken-1"
                        flat="flat"
                        @click="dialog = false"
                    >No</v-btn>

                    <v-btn
                        color="blue darken-1"
                        flat="flat"
                        @click="removeTodoListItem(removeId)"
                    >Yes</v-btn>
                </v-card-actions>
                
            </v-card>
        </v-dialog>
    </v-container>
</template>
    
<script>
import createTodo from './createTodo.vue'
import createTodolist from './createTodolist.vue'

export default {
    data: function() {
        return {
            active: 'all',
            dialog: false,
            removeId: -1,
        }
    },
    components: {
        createTodo,
        createTodolist
    },
    methods: {
        removeTodoListItem: function(id) {
            this.$store.dispatch('removeTodoListItem', id)
            this.dialog = false
        },
        toggleTodo: function(status, todolistId, todoId) {
            this.$store.dispatch('toggleTodo', {todolistId, todoId, status})
        },
        removeTodo: function(todolistId, todoId) {
            this.$store.dispatch('removeTodo', {todolistId, todoId})
        },
        openDialog: function(todolistId) {
            this.dialog = true
            this.removeId = todolistId
        },
    },
    computed: {
        alltodolistdata() {
            return this.$store.getters.getTodoList
        },
        finishedtodolistdata() {
            return this.$store.getters.filteredByFinished
        },
        unifinishedtodolistdata() {
            return this.$store.getters.filteredByUnfinished
        },
    },
    mounted() {
        this.$store.dispatch('getTodoList');
    }
}
</script>