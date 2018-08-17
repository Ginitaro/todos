<template>
    <v-form ref="form" class="flex pt-4" @submit.prevent lazy-validation>
        <v-text-field
            name="title" 
            v-model="title"
            v-bind:rules="titleRules"
            v-on:keyup.enter="createTodoList($event, title)"
            append-icon="list"
            label="Add a category..."
            prepend-inner-icon="add"
        ></v-text-field>       
    </v-form>
</template>
    
<script>
export default {
    data: function() {
        return {
            title: '',
            titleRules: [
                v => !!v || 'Title is required',
            ],
        }
    },
    methods: {
        createTodoList: function(e, title) {
            if(this.$refs.form.validate()){
                this.$store.dispatch('addTodoListItem', title)
                this.$refs.form.reset()
                // Blur input after sending data
                e.srcElement.blur()
            }

        }
    }
}
</script>