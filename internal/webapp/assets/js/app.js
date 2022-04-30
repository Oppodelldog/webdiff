const Home = {template: '<div><h6>Home</h6></div>'}
const Diff = {template: '<diff/>'}
const Download = {template: '<div><h6>Download</h6></div>'}

const routes = [
    {path: '/', component: Home},
    {path: '/diff', component: Diff},
    {path: '/download', component: Download},
]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHashHistory(),
    routes,
})

const app = Vue.createApp({});

app.component('diff', componentDiff);

app.use(router)
app.mount('#app')
