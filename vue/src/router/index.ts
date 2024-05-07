/*
 * @Author: Callay 2415993100@qq.com
 * @Date: 2024-05-06 22:25:37
 * @LastEditors: Callay 2415993100@qq.com
 * @LastEditTime: 2024-05-07 19:19:41
 * @FilePath: \checkIn\vue\src\router\index.ts
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect:'/login'
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue')
    },
    {
      path:'/index',
      name:'index',
      component:()=>import('@/views/index/IndexView.vue'),
      children:[
        {
          path:'home',
          name:'home',
          component:()=>import('@/views/index/HomeView.vue')
        },
        {
          path:'history',
          name:'history',
          component:()=>import('@/views/index/HistoryView.vue')
        },
        {
          path:'organization',
          name:'organization',
          component:()=>import('@/views/index/OrganizationView.vue')
        },
        {
          path:'userInfo',
          name:'userInfo',
          component:()=>import('@/views/index/UserInfoView.vue')
        }
      ]
    }
  ]
})

export default router
