import Vue from 'vue'
import Router from 'vue-router'
import Layout from '@/views/layout'

Vue.use(Router)
export const routes = [
  { path: '/home', redirect: '/', hidden: true },
  {
    path: '/',
    name: 'Home',
    icon: 'el-icon-s-home',
    component: Layout,
    redirect: '/',
    children: [{
      path: '',
      component: () => import('@/views/article/home/index')
    }]
  },
  {
    path: '/articles',
    name: 'Blog',
    icon: 'el-icon-collection',
    component: Layout,
    redirect: '/',
    children: [
      {
        name: 'Latest Posts',
        path: '',
        component: () => import('@/views/article/home/index')
      },
      {
        name: 'New Post',
        path: 'new',
        component: () => import('@/views/article/new/index')
      },
      {
        hidden: true,
        path: 'detail/:id',
        component: () => import('@/views/article/view/index')
      },
      {
        hidden: true,
        path: 'edit/:id',
        component: () => import('@/views/article/edit/index')
      }]
  },
  {
    path: '/resources',
    name: 'Resources',
    icon: 'el-icon-help',
    component: Layout,
    redirect: '/',
    children: [{
      path: '',
      component: () => import('@/views/resource/index')
    }]
  },
  {
    path: '/challenges',
    name: 'Challenges',
    icon: 'el-icon-coffee-cup',
    component: Layout,
    redirect: '/challenges/home',
    children: [{
      path: 'home',
      name: 'List Challenges',
      component: () => import('@/views/challenge/home/index')
    },
    {
      path: 'new',
      name: 'New Challenge',
      roles: ['group-admin'],
      component: () => import('@/views/challenge/new/index')
    },
    {
      path: 'admin',
      name: 'Manage Challenges',
      roles: ['group-admin'],
      component: () => import('@/views/challenge/admin/index')
    },
    {
      path: 'edit/:id',
      hidden: true,
      component: () => import('@/views/challenge/edit/index')
    },
    {
      path: 'groups/:id',
      hidden: true,
      component: () => import('@/views/challenge/group/index')
    }]
  },
  {
    path: '/policies',
    name: 'Policies',
    icon: 'el-icon-setting',
    roles: ['group-admin'],
    component: Layout,
    redirect: '/policies/home',
    children: [{
      path: 'home',
      name: 'Manage Policies',
      component: () => import('@/views/policy/home/index')
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
    },
    {
      path: 'reset-password/:token',
      component: () => import('@/views/user/reset-password/index')
    },
    {
      path: 'forgot-password',
      component: () => import('@/views/user/forgot-password/index')
    }]
  },
  {
    path: '/about',
    name: 'About',
    icon: 'el-icon-info',
    component: Layout,
    redirect: '/',
    children: [{
      path: '',
      component: () => import('@/views/About')
    }]
  },
  { path: '/login', component: () => import('@/views/login/index'), hidden: true },
  { path: '/404', component: () => import('@/views/404'), hidden: true },
  { path: '*', redirect: '/', hidden: true },
  { path: '/play', component: () => import('@/views/play/index'), hidden: true }
]

export default new Router({
  mode: 'history',
  scrollBehavior: () => ({ y: 0 }),
  routes: routes
})
