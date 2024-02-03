<template>
  <navbar />
  <h1 class="text-4xl text-center my-12">My Decks</h1>
  <div class="flex w-screen justify-center">
  <div class="sm:gap-4 w-full lg:w-2/3 xl:w-7/12 sm:grid lg:grid-cols-3 sm:grid-cols-2">
    <create-deck-card />
    <template v-for="deck in decks" v-bind:key="deck">
      <deck-card  :deck="deck"/>
    </template>
  </div>
</div>
</template>

<script setup>
import Navbar from '../components/Navbar.vue';
import CreateDeckCard from '../components/CreateDeckCard.vue';
// createdeckcard is a bad name?
import DeckCard from '../components/DeckCard.vue';

const decks = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
</script>
