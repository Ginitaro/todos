<template>
    <nav class="panel todo-controls">
        <div class="panel-heading">
            Todo App
        </div>
        <form class="panel-block" @submit.prevent>
            <div class="field is-grouped">
                <div class="control is-expanded is-medium">
                    <p class="control has-icons-left">
                        <input class="input is-medium" type="text" placeholder="Add a category..." name="title" v-model="title">
                        <span class="icon is-medium is-left">
                            <i class="fas fa-plus"></i>
                        </span>
                    </p>
                </div>
                <button class="button is-medium is-info" v-on:click="createTodoList(title)">Add</button>
            </div>
            </div>
        </form>
        <div class="panel-tabs">
            <a class="is-active">All</a>
            <a>Unfinished</a>
            <a>Finished</a>
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