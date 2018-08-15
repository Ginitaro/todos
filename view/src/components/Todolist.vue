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
                        v-model="alltodolistdata.doneTask"
                        >
                            <v-tab href="#all" v-on:click="stripExpanded(alltodolistdata)">All</v-tab>
                            <v-tab href="#unfinished" v-on:click="stripExpanded(unifinishedtodolistdata)">Unfinished</v-tab>
                            <v-tab href="#finished" v-on:click="stripExpanded(finishedtodolistdata)">Finished</v-tab>
                        </v-tabs>
                        
                    </v-toolbar>
                    
                    <v-tabs-items v-model="alltodolistdata.doneTask">
                        <v-tab-item id="all">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in alltodolistdata" :key="todolist.ID" class="py-2">
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
                                            
                                            <v-list-tile v-for="todo in todolist.Todos" :key="todo.ID">
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
                        </v-tab-item>
                        <v-tab-item id="unfinished">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in unifinishedtodolistdata" :key="todolist.ID" class="py-2">
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
                                            
                                            <v-list-tile v-for="todo in todolist.Todos" :key="todo.ID">
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
                        </v-tab-item>
                        <v-tab-item id="finished">
                            <v-expansion-panel expand>
                                <v-expansion-panel-content v-for="todolist in finishedtodolistdata" :key="todolist.ID" class="py-2">
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
                                            
                                            <v-list-tile v-for="todo in todolist.Todos" :key="todo.ID">
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
                        </v-tab-item>
                    </v-tabs-items>
                    
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
        stripExpanded: function (list) {
            for (var i = list.length - 1; i >= 0; i--) {
                list[i]["expanded"] = false
            }
        }
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