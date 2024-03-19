<template>
    <header class="header">
      <h1 class="title ">Blog</h1>
      <div class="routerLink">
        <router-link to="/">Home</router-link>
        <router-link :to="{name:'createPost'}">Create</router-link>
      </div>
    </header>
  </template>
  
  <script>
  
  </script>
  
  <style>
  
  .header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 15px 50px;
      box-shadow: 1px 1px 4px rgba(0,0,0,.5);
      margin-bottom: 20px;
  }
  .title {
    margin-bottom: 0;
  }
  .header .routerLink a {
    font-size: 18px;
    color: black;
    padding-left: 30px;
    text-decoration: none;
  }
  
  </style>