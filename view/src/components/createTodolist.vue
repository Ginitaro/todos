<template>
    <v-form class="flex" @submit.prevent>
        <v-text-field
            name="title" 
            v-model="title"
            v-on:keyup.enter="createTodoList(title)"
            append-icon="list"
            flat
            label="Add a category..."
            prepend-inner-icon="add"
            solo-inverted
        ></v-text-field>       
    </v-form>
</template>
    
<script>
export default {
    data: function() {
        return {
            title: '',
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