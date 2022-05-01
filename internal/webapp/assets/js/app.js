const Home = {template: '<home/>'}
const Diff = {template: '<diff/>'}
const Download = {template: '<download/>'}
const Browse = {template: '<browse/>'}
const ServerLog = {template: '<serverlog/>'}

const routes = [
    {path: '/', component: Home},
    {path: '/diff', component: Diff},
    {path: '/download', component: Download},
    {path: '/browse', component: Browse},
]

const router = VueRouter.createRouter({
    history: VueRouter.createWebHashHistory(),
    routes,
})

const app = Vue.createApp({});

app.component('diff', componentDiff);
app.component('download', componentDownload);
app.component('home', componentHome);
app.component('browse', componentBrowse);
app.component('serverlog', componentServerLog);

app.use(router)
app.mount('#app')

connectWebsocket('ws://' + document.location.host + "/ws")