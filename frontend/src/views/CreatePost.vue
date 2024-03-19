<template>
    <h1>Add Post</h1>
  <form @submit.prevent="addPost">
    <div class="from-group">
        <label>Project Title</label>
        <input type="text" v-model="title">
    </div>
    <div class="from-group">
        <label>Post Body</label>
        <textarea  cols="30" rows="10" v-model="body"></textarea>
    </div>
   
    <button class="btn btn-primary">Add Post</button>
  </form>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

let router = useRouter();
let title = ref('');
let body = ref('');

let addPost = async ()=>{
    await axios.post('http://localhost:3000/api/v1/post',{
        user_id:"65f3de0d76caeccb7518c2e6", // for now use defalut user
        title:title.value,
        content:body.value
    })
    router.push('/')
}
</script>

<style>
    h1 {
            color: darkred;
            margin-bottom: 50px;
        }
    form {
        width: 400px;
        margin: 0 auto 100px;
    }
    .from-group label {
        display: block;
        padding-bottom: 10px;
        text-align: left;
        font-size: 16px;
        font-weight: bold;
    }
    .from-group input,
    .from-group textarea {
        width: 100%;
        box-sizing: border-box;
        padding: 10px;
        resize: none;
        margin-bottom: 10px;
        outline: none;
        font-size: 16px;
    }
    button {
        padding: 10px;
        cursor: pointer;
    }
    .tag {
        display: inline-block;
        margin: 0 10px 20px;
        color: #444;
        background: #ddd;
        padding: 8px;
        border-radius: 20px;
        font-size: 14px;
    }
</style>