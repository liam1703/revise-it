<template>
  <div class="bg-slate-200 h-screen w-screen">
    <navbar />
    <div>
      <h1 class="text-center text-6xl mt-16 font-semibold">Revision backed by science.</h1>
      <h4 class="text-center text-xl mt-8">Create questions and learn using active recall.</h4>
      <div class="text-center w-screen mt-6">
        <button class="text-center border border-solid bg-purple-600 text-white p-3 hover:rounded-lg">Begin learning</button>
      </div>

    </div>
  </div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue';
</script>
