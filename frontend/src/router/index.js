import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/Home.vue'
import CreatePost from '../views/CreatePost.vue'
import SinglePost from '../views/SinglePost.vue'
const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path:'/create-post',
    name:'createPost',
    component:CreatePost
  },
  {

    path:'/post/:id',
    name:'singlePost',
    component:SinglePost,
  }
]

const router = createRouter({
  history: createWebHistory("/"),
  routes
})
export default router