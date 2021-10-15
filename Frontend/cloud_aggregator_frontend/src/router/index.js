import Vue from 'vue'
import Router from 'vue-router'
import TasksInfo from '../components/containers/TasksInfo'
import TaskForm from '../components/containers/TasksForm'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    // {
    //   path: '/',
    //   name: 'MainPage',
    //   component: MainPage
    // }
    {
      path: '/tasksInfo',
      name: 'tasksInfo',
      component: TasksInfo
    },
    {
      path: '/taskForm',
      name: 'taskForm',
      component: TaskForm
    }
  ]
})
