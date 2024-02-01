<template>
  <div class="bg-slate-200 h-screen w-screen">
    <navbar />
    <div class="w-screen">
      <h1 class="text-center text-6xl mt-16 font-semibold text-indigo-600">Revision backed by science.</h1>
      <h4 class="text-center text-xl mt-8 text-black">Create questions and learn using active recall.</h4>
      <div class="text-center w-screen mt-6">
        <button @click.prevent="goToLogin" class="text-center border border-solid bg-purple-600 text-white p-3 hover:rounded-lg">Begin learning</button>
      </div>
      <div class="bg-gray-300 text-center rounded-lg w-10/12 mx-auto mt-24 p-8">
        <h5 class="text-indigo-600 font-bold text-xl">What do we do?</h5>
      </div>
    </div>
  </div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue';
function goToLogin() {
  window.location.href = '/login';
}
</script>
