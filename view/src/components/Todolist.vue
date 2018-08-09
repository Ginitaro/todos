<template>
    <div class="panel todo-controls">
        <div class="panel-block task" v-for="todolist in todolistdata" v-bind:class="{active:todolist.isOpen}">
            <div class="item">
                <div class="item-name" v-on:click="todolist.isOpen = !todolist.isOpen" >
                    <span class="circle"><i class="fa fa-list"></i></span> 
                    <span class="todo-title">{{ todolist.Title }}</span>
                </div>
                
                <!-- <a href="#" @click="removeTodoListItem(todolist.ID)" class="delete" aria-label="delete"></a> -->
                
            </div>
            <div class="item-content">
                <createTodo v-bind:parent="todolist.ID"></createTodo>
                <ol>
                    <li v-for="todo in todolist.Todos">
                        {{ todo.Name }}
                        <input type="checkbox" v-on:change="toggleTodo($event, todolist.ID, todo.ID)" v-model="todo.Done">
                    </li>
                </ol>
            </div>
        </div>
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