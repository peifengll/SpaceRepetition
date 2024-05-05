import {createRouter, createWebHistory, RouteRecordRaw} from 'vue-router'
import AppLayout from '@/layout/AppLayout.vue'
// import productRoutes from './modules/product'
import Register from '../pages/Register.vue'

const routes: RouteRecordRaw[] = [
    {
        path: '/',
        component: AppLayout,
        children: [
            {
                path: '/',
                name: 'come',
                component: () => import('@/pages/Home.vue'),
                meta: {
                    title: "首页"
                }
            },
            {
                path:'/editor',
                name:'editor',
                component:()=>import('@/pages/Editor.vue'),
                props: (route) => ({ id: route.query.id }),
                meta:{
                    title:"编辑页"
                }
            },
            {
                path:'/review',
                name:'review',
                component:()=>import('@/pages/Review.vue')
            },
            {
                path:'/announcement',
                name:'announcement',
                component:()=>import('@/pages/Announcement.vue')
            },
            {
                path:'/exportinfos',
                name:'exportinfos',
                component:()=>import('@/pages/ExportInfos.vue')
            },
            {
                path:'/reviewtatistics',
                name:'ReviewStatistics',
                component:()=>import('@/pages/ReviewStatistics.vue')
            },

           

        ]
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('@/pages/Login.vue')
    },
    {
        path:'/register',
        name:'register',
        component:Register,
    },
    {
        path:'/home',
        name:'home',
        component:()=>import('@/pages/Home.vue')
    },

    {
        path:'/test',
        name:'test',
        component:()=>import('@/pages/TestNew.vue')
    },


]

const router = createRouter({
    // history: createWebHashHistory(), // hash 路由模式
    history: createWebHistory(), // history 路由模式
    routes // 路由规则
})

export default router

