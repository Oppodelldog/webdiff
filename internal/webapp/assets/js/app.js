const Home = {template: '<home/>'}
const Diff = {template: '<diff/>'}
const Download = {template: '<download/>'}

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
app.component('download', componentDownload);
app.component('home', componentHome);

app.use(router)
app.mount('#app')
