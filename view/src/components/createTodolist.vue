<template>
    <v-form @submit.prevent>
        
        <v-toolbar color="blue" dark>
            <div class="xs12">
                <v-flex>
                    <v-text-field 
                    name="title" 
                    v-model="title" 
                    label="Add a category..."
                    single-line
                    hide-details
                    v-on:keyup.enter="createTodoList(title)"
                    >
                    </v-text-field>
                    <v-btn icon v-on:click="createTodoList(title)"><v-icon>add</v-icon></v-btn>
                </v-flex>
                <div>
                    asdasdasdsahdstg
                </div>
            </div>
        </v-toolbar>
        
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