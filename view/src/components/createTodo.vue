<template>
    <v-form class="flex" @submit.prevent>
        <v-text-field
        v-model="text"
        v-on:keyup.enter="addTodo(parent, text)"
        append-icon="list"
        flat
        label="Add todo..."
        prepend-inner-icon="add"
        class="create-todo px-5 pt-3 mt-0"
        ></v-text-field>
    </v-form>
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