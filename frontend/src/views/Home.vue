<template>
  <div class="bg-slate-200 h-screen w-screen">
    <navbar />
    <h1 class="mb-12">Homepage</h1>
    <router-link to="/login" class="m-4 border border-solid p-3 rounded bg-red-300"><button>Login</button></router-link>
  </div>
</template>

<script setup>
import navbar from '../components/navbar.vue';
</script>
