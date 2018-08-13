<template>
    <v-text-field
    v-model="text"
    v-on:keyup.enter="addTodo(parent, text)"
    append-icon="list"
    flat
    label="Add todo..."
    prepend-inner-icon="add"
    class="create-todo"
    ></v-text-field>
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