<template>
    <nav class="panel todo-controls">
        <div class="panel-heading">
            Todo List App
        </div>
        <form class="panel-block" @submit.prevent>
            <div class="field is-grouped">
                <div class="control is-expanded is-medium">
                        <input class="input is-medium" type="text" placeholder="Type something here..." name="title" v-model="title">
                </div>
                <button class="button is-medium is-info" v-on:click="createTodoList(title)">Add</button>
            </div>
            </div>
        </form>
        <div class="panel-tabs">
            <a class="is-active">All</a>
            <a>Complete</a>
            <a>Done</a>
        </div>
    </nav>
</template>
    
<script>
export default {
    data: function() {
        return {
            title: ''
        }
    },
    methods: {
        createTodoList: function(title) {
            this.$store.dispatch('addTodoListItem', title);
            this.title = '';
        }
    }
}
</script>