import Vue from 'vue'
import Router from 'vue-router'
Vue.use(Router)

let router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: () => import('@/components/index/Index'),
      meta: {
        guest: true,
        layout: 'Default'
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('@/components/Login'),
      meta: {
          guest: true,
          layout: 'Auth'
      }
    },
  ]
})
router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
      if (localStorage.getItem('token') == null) {
          next({
              path: '/login',
              params: { nextUrl: to.fullPath }
          })
      } else {
          let user = JSON.parse(localStorage.getItem('user'))
          if(to.matched.some(record => record.meta.is_admin)) {
              if(user.type == 1){
                  next()
              }
              else{
                  next({ name: 'ViewInvoic'})
              }
          }else {
              next()
          }
      }
  } else if(to.matched.some(record => record.meta.guest)) {
      if(localStorage.getItem('token') == null){
          next()
      }
      else{
          next()
      }
  }else {
      next()
  }
})
export default router
