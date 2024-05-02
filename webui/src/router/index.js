import { createRouter, createWebHistory } from 'vue-router'
import RealEstateView from '@/views/Fe26Layout.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: RealEstateView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/FileExplorer.vue')
    },
    {
      path: '/notes',
      name: 'notes',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/Notes/Notes.vue')
    }
  ]
})

export default router
