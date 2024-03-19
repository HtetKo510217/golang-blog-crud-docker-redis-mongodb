<template>
    <div class="home">
      <h1>Home</h1>
      <PostList :posts="posts"></PostList>
    </div>
  </template>
  
  <script setup>
  import PostList from '../components/PostList.vue'
  import getPosts from '../composables/getPosts'
  import { onMounted } from 'vue'
  const {posts,load} = getPosts()
  
  onMounted(() => {
    load()
  })
 
  </script>

  