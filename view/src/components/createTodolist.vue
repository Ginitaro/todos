<template>
    <v-form @submit.prevent>
                
        <v-toolbar color="blue" dark extended>
            <v-text-field
            name="title" 
            v-model="title"
            v-on:keyup.enter="createTodoList(title)"
            append-icon="list"
            class="mx-3"
            flat
            label="Add a category..."
            prepend-inner-icon="add"
            solo-inverted
            ></v-text-field>
            <v-toolbar-title slot="extension" class="white--text toolbar-options">
                <v-chip label><v-icon left>list</v-icon>All</v-chip>
                <v-chip label><v-icon left>close</v-icon>Unfinished</v-chip>
                <v-chip label><v-icon left>playlist_add_check</v-icon>Finished</v-chip>
            </v-toolbar-title>
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