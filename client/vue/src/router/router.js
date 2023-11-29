import { createRouter, createWebHistory } from "vue-router"

const routes = [
    {
        path: "/",
        name: "home",
        component: () => import("../view/Home.vue"),
    },
    {
        path: "/u/manuel/workspaces",
        name: "workspaces",
        component: () => import("../view/Workspaces.vue")
    },
    {
        path: "/u/manuel/:workspacename",
        name: "workspace",
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

export default router