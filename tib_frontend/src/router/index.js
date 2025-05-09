import { createRouter, createWebHistory } from 'vue-router';
import HomeComponent  from '../components/HomeComponent.vue'
import TopicComponent from "@/components/TopicComponent.vue";
import ProfileComponent from "@/components/ProfileComponent.vue";
import TopicDetailComponent from "@/components/TopicDetailComponent.vue";

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
    },
    {
        path:'/topic/:id',
        name: 'TopicDetail',
        component: TopicDetailComponent,
        props: true,
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;
