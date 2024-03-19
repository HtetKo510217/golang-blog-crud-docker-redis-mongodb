<template>
    <div class="container">
    <div class="row">
      <div class="col-md-12">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">{{ post.title }}</h5>
            <p class="card-text">{{ post.content }}</p>
          </div>
        </div>
      </div>
      
      <div class="col-md-4 mb-4 mt-3" v-for="comment in comments" :key="comment._id">
          <div class="custom-comment-design">
              <h5>{{ comment.content }}</h5>
          </div>
      </div>

      <div class="col-md-12 mt-3">
        <form @submit.prevent="addComment">
          <div class="from-group">
              <label>Comment</label>
              <textarea  cols="30" rows="10" v-model="comment"></textarea>
          </div>
          <button class="btn btn-primary">Add Comment</button>
        </form>
      </div>
    </div>
  </div>
  
  </template>
  
  <script setup>
  import { useRoute } from 'vue-router'
  import axios from 'axios';
  import { ref, onMounted } from 'vue';

  const route = useRoute()
  const id = route.params.id
  const post = ref({})
  const comment = ref('')
  const comments = ref({})
  const getSinglePost = async () => {
    const res = await axios.get(`http://localhost:3000/api/v1/post/${id}`)
    post.value = res.data
  }

  const getSingleComment = async () => {
    const res = await axios.get(`http://localhost:3000/api/v1/post/${id}/comment`)
    comments.value = res.data
  }

  let addComment = async () => {
    await axios.post(`http://localhost:3000/api/v1/post/${id}/comment`, {
      post_id: id,
      content: comment.value
    })
    comment.value = ''
    getSingleComment()
  }

  onMounted(() => {
    getSinglePost()
    getSingleComment()
  })




  </script>
  
  <style>
  .custom-comment-design {
    background-color: #f5f5f5;
    padding: 10px;
    border-radius: 4px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
  </style>