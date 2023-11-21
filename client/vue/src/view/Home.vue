<template>
    <form @submit.prevent="HandleSubmit">
        <input type="submit" value="add">
        <textarea v-model="form.todo"></textarea>
    </form>
    <div v-if="tasks.length > 0 || tasks != null">
        <h1>tasks</h1>
        <div v-for="t in tasks">
            <h1>{{ t.task.todo }}</h1>
            <button v-on:click="TaskRemove(t.id)">x</button>
        </div>
    </div>
</template>

<script>
import axios from "axios"
const url = "http://localhost:8080/"

export default {
    name: "Home",
    data() {
        return {
            tasks: [],
            form: {
                todo: null
            },
        }
    },
    methods: {
        MapList(list) {
            let aux = list
            let newl = []

            while (aux != null) {
                newl.push(aux)
                aux = aux.next
            }

            return newl
        },
        async TaskFetch() {
            try {
                const request = await axios.get(url)
                if (request.status == 200) {
                    const data = await request.data
                    const lista = this.MapList(data)
                    this.tasks = lista
                }
            } catch (err) {

            }
        },
        async TaskRemove(taskid) {
            try {
                const request = await axios.delete(url + "remove", {
                    headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    },
                    data: {
                        "id": taskid
                    }
                })

                if (request.status == 200) {
                    this.tasks = this.tasks.filter(t => t.id != taskid)
                }

            } catch (err) {

            }
        },
        async TaskAdd() {
            try {
                const request = await axios({
                    method: "POST", url: url + "add", headers: {
                        "Content-Type": "application/x-www-form-urlencoded"
                    }, data: {
                        "todo": this.form.todo
                    }
                })
                return request
            } catch (err) {

            }
        },
        async HandleSubmit() {
            const addrequest = await this.TaskAdd()

            if (addrequest.status == 200) {
                const nt = {
                    id: this.tasks.length == 0 ? 1 : this.tasks[this.tasks.length - 1].id + 1,
                    task: {
                        todo: this.form.todo,
                    }
                }
                this.tasks.push(nt)
            }
        }
    },
    beforeMount() {
        this.TaskFetch()
    }
}
</script>