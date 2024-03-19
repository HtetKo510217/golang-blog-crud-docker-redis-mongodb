<template>
   <div class="container">
    <div class="row">
      <div class="col-md-4 mb-4" v-for="post in posts" :key="post.id">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{{ post.title }}</h5>
            <p class="card-text">{{ post.content }}</p>
            <router-link :to="{name:'singlePost',params:{id:post._id}}" class="btn btn-primary">Read More</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
 </template>
 
 <script setup>
 import { defineProps } from 'vue'
 const props = defineProps({
     posts: Array
 })
 </script>
 
 <style>
 
 </style>