import { createRouter, createWebHistory } from 'vue-router';
import HomeComponent  from '../components/HomeComponent.vue'
import TopicComponent from "@/components/TopicComponent.vue";
import ProfileComponent from "@/components/ProfileComponent.vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: HomeComponent, // 设置首页路由
    },
    {
        path:'/topics',
        name:'Topic page',
        component: TopicComponent,
    },
    {
        path:'/profile',
        component: ProfileComponent
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;
