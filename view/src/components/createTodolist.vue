<template>
    <nav class="panel todo-controls">
        <div class="panel-heading">
            Todo List App
        </div>
        <form class="panel-block" @submit.prevent>
            <div class="field is-grouped">
                <div class="control is-expanded is-medium"
                    v-bind:class="{ active: is-loading }">
                        <input class="input is-medium" type="text" placeholder="Type something here..." name="title" ref="title">
                </div>
                <button class="button is-medium is-info" v-on:click="createTodoList()">Add</button>
            </div>
        </form>
        <div class="panel-tabs">
            <a class="is-active">All</a>
            <a>Complete</a>
            <a>Done</a>
        </div>
        <div class="wrapper">
            <div class="row">
                {{ response }}
            </div>
        </div>
    </nav>
</template>
    
<script>
export default {
    data: function() {
        return {
            response: null
        }
    },
    methods: {
        createTodoList: function() {
            this.$store.dispatch('addTodoListItem', this.$refs.title.value);
        }
    }
}
</script>