<template>
    <form @submit.prevent>
        
         <v-text-field 
            v-model="text" 
            label="Add todo..."
            required
            single-line
            hide-details
            v-on:keyup.enter="addTodo(parent, text)"
            ></v-text-field>
        
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