<template>
    <v-form ref="form" class="flex" @submit.prevent lazy-validation>
        <v-text-field
            v-model="text"
            name="text"
            v-on:keyup.enter="addTodo($event, parent, text)"
            v-bind:rules="textRules"
            append-icon="list"
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
            text: '',
            textRules: [
                v => !!v || 'Text is required',
            ],
        }
    },
    props: ['parent'],
    methods: {
        addTodo: function(e, id, text) {
            if(this.$refs.form.validate()){
                this.$store.dispatch('addTodo', {text, id});
                this.text = '';
                this.$refs.form.reset()
                // Blur input after sending data
                e.srcElement.blur()
            }
        }
    }
}
</script>