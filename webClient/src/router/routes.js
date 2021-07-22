const routes = [
  {path: '/', component: () => import(`../app/components/home.vue`)},
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('../app/components/error404.vue')
  }
]


export default routes
