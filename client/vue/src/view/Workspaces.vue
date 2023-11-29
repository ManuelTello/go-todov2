<template>
    <input type="text" v-model="new_workspace" placeholder="New workspace ...">
    <input type="button" value="+" v-on:click="AddNewWorkspace">
    <div v-if="workspaces.lenght > 0 || workspaces != null" v-for="w in workspaces">
        <router-link :to="'/u/manuel/' + w.Name">{{ w.Name }}</router-link>
    </div>
    <div v-else>
        No workspaces to show
    </div>
</template>

<script>
export default {
    name: "workspaces",
    data() {
        return {
            new_workspace: "",
            workspaces: [
                {
                    Id: 0,
                    Name: "Golang todo"
                },
                {
                    Id: 1,
                    Name: "Things for nina"
                }
            ]
        }
    },
    methods: {
        AddNewWorkspace() {
            console.log("New workspace added: ", this.new_workspace)
        }
    },
    mounted() {
        console.log(this.$route)
    }
}
</script>