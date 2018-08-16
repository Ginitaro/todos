<template>
	<v-app>
		<router-view></router-view>
		<todolist></todolist>
	</v-app>
</template>

<script>
import todolist from './components/Todolist.vue'

export default {
	components: {
		todolist
	}
}
</script>