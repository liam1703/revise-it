<template>
  <div class="mb-6">
    <button @click.prevent="createDeckModal = true"
      class="block w-full p-6 max-h-sm h-72  border rounded-lg shadow text-center hover:bg-gray-700 bg-gray-800 border-gray-700">
      <h5 class="mb-12 text-2xl font-bold tracking-tight text-white ">Create new deck</h5>
      <p class="font-normal text-gray-400">
        Click here to create a new deck of flashcards.
      </p>
    </button>
  </div>
  <TheModal v-if="createDeckModal" >
    <div>
      yo yo yo
    </div>
  </TheModal>



</template>

<script setup>
import { ref } from 'vue';
import TheModal from './TheModal.vue';
const createDeckModal = ref(false);
</script>
