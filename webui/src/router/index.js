import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user.js'

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: '/',
            redirect: (to) => {
                return { path: '/files/' }
            }
        },
        {
            path: '/files',
            redirect: (to) => {
                return { path: '/files/' }
            }
        },
        {
            path: '/files/:path(.*)',
            name: 'app',
            meta: {
                // requiresAuth: true
            },
            component: () => import('@/views/MainAppView.vue')
        },
        {
            path: '/login',
            name: 'login',
            meta: {
                hideFromAuth: true
            },
            component: () => import('@/views/LoginView.vue')
        }
    ]
})

// this checks for metadata in the router and redirects to login page if the user is not logged in
// the same happens if the user is logged in he is redirected to the entry away from the login page
// this relies on the user store
// based on: https://stackoverflow.com/questions/52653337/vuejs-redirect-from-login-register-to-home-if-already-loggedin-redirect-from
router.beforeEach((to, from, next) => {
    const user = useUserStore()

    const navigate = function (to, next) {
        if (to.matched.some((record) => record.meta.requiresAuth)) {
            if (!user.isLoggedIn) {
                next({ name: 'login' })
            } else {
                next() // go to wherever I'm going
            }
        } else if (to.matched.some((record) => record.meta.hideFromAuth)) {
            if (user.isLoggedIn) {
                next({ name: 'home' }) // hide logged in users from hitting the login page
            } else {
                next()
            }
        } else {
            next() // does not require auth, make sure to always call next()!
        }
    }

    if (user.isFirstLogin) {
        user.setFirtLoginFalse()
        const p = user.checkState()
        p.then(() => {
            navigate(to, next)
        })
    } else {
        navigate(to, next)
    }
})

export default router
