<template>
  <div class="h-screen w-screen bg-slate-100">
    <div class="w-full  pt-24">
      <h1 class="text-4xl text-center">Revise-it</h1>
      <div class="bg-white shadow-xl w-96 p-4 rounded-xl mx-auto mt-16">
        <h3 class="mb-3">Log into your Revise-it account.</h3>
        <label for="email" class="text-sm font-medium">Email:</label>
        <input v-model="email" id="email" type="text" class="mt-2 mb-4 p-2 text-lg w-full rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
        <label for="password" class="text-sm font-medium">Password:</label>
        <input v-model="password" id="password" type="password" class="mt-2 p-2 text-lg w-full rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
        <div class="mt-12 text-center">
          <button class="border border-solid p-3 rounded background-base">Login</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const email = ref('');
const password = ref('');
</script>
