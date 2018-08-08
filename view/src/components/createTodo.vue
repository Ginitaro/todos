<template>
    <form class="panel-block" @submit.prevent>
        <div class="field is-grouped">
            <div class="control is-expanded is-medium">
                <input class="input is-medium" type="text" placeholder="Type something here..." v-model="text">
            </div>
            <button class="button is-medium is-info" v-on:click="addTodo(parent, text)">Add</button>
        </div>
    </form>
</template>
    
<script>
export default {
    data: function() {
        return {
            text: ''
        }
    },
    props: ['parent'],
    methods: {
        addTodo: function(id, text) {
            this.$store.dispatch('addTodo', {text, id});
            this.text = '';
        }
    }
}
</script>