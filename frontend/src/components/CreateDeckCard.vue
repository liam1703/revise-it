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
      <h3>Create new deck</h3>
      <div class="my-4">
        <label for="deckName" class="text-sm font-medium">Deck name:</label>
        <input v-model="deckName" id="deckName" type="text" class="mt-2 p-2 text-lg w-full rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500">
        <label for="deckDescription" class="text-sm font-medium">Deck description (Optional):</label> 
        <textarea v-model="deckDescription" name="deckDescription" id="deckDescription" cols="30" rows="8" class="mt-2 p-2 text-lg w-full rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500"></textarea>
      </div>
    </div>
  </TheModal>
</template>

<script setup>
import { ref } from 'vue';
import TheModal from './TheModal.vue';
const createDeckModal = ref(false);
</script>
