import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/layout'

// in development-env not use lazy-loading, because lazy-loading too many pages will cause webpack hot update too slow. so only in production use lazy-loading;
// detail: https://panjiachen.github.io/vue-element-admin-site/#/lazy-loading

Vue.use(Router)

/**
* hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
* alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
*                                if not set alwaysShow, only more than one route under the children
*                                it will becomes nested mode, otherwise not show the root menu
* redirect: noredirect           if `redirect:noredirect` will no redirect in the breadcrumb
* name:'router-name'             the name is used by <keep-alive> (must set!!!)
* meta : {
    title: 'title'               the name show in subMenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    breadcrumb: false            if false, the item will hidden in breadcrumb(default is true)
  }
**/
export const constantRouterMap = [
  { path: '/login', component: () => import('@/views/login/index'), hidden: true },
  { path: '/404', component: () => import('@/views/404'), hidden: true },
  {
    path: '/',
    name: 'Home',
    component: Layout,
    redirect: '/',
    hidden: true,
    children: [{
      path: '',
      component: () => import('@/views/home/index')
    }]
  },
  {
    path: '/about',
    name: 'About',
    component: Layout,
    redirect: '/',
    hidden: true,
    children: [{
      path: '',
      component: () => import('@/views/About')
    }]
  },
  {
    path: '/articles',
    name: 'Articles',
    component: Layout,
    hidden: true,
    redirect: '/article/detail',
    children: [{
      path: 'detail/:id',
      component: () => import('@/views/article/view/index')
    },
    {
      path: 'new',
      component: () => import('@/views/article/new/index')
    },
    {
      path: 'edit/:id',
      component: () => import('@/views/article/edit/index')
    }]
  },
  {
    path: '/resources',
    name: 'Resources',
    component: Layout,
    hidden: true,
    redirect: '/',
    children: [{
      path: '',
      component: () => import('@/views/resource/index')
    }]
  },
  {
    path: '/users',
    name: 'Users',
    component: Layout,
    hidden: true,
    redirect: '/',
    children: [{
      path: 'register',
      component: () => import('@/views/user/registration/index')
    }]
  },
  {
    path: '/challenges',
    name: 'Challenges',
    component: Layout,
    hidden: true,
    redirect: '/challenges/home',
    children: [{
      path: 'home',
      component: () => import('@/views/challenge/home/index')
    },
    {
      path: 'edit/:id',
      component: () => import('@/views/challenge/edit/index')
    },
    {
      path: 'new',
      component: () => import('@/views/challenge/new/index')
    },
    {
      path: 'groups/:id',
      component: () => import('@/views/challenge/group/index')
    }]
  },
  {
    path: '/policies',
    name: 'Policies',
    component: Layout,
    hidden: true,
    redirect: '/policies/home',
    children: [{
      path: 'home',
      component: () => import('@/views/policy/home/index')
    }]
  },
  { path: '/home', redirect: '/', hidden: true },
  { path: '*', redirect: '/', hidden: true }
]

export default new Router({
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
