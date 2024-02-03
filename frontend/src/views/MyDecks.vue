<template>
  <navbar />
  <h1 class="text-4xl text-center mt-12">My Decks</h1>
  <div class="w-full grid gap-4 grid-cols-3 grid-rows-3">
    <create-deck-card />
    <template v-for="deck in decks" v-bind:key="deck">
      <deck-card  :deck="deck"/>
    </template>
  </div>


</template>

<script setup>
import Navbar from '../components/Navbar.vue';
import CreateDeckCard from '../components/CreateDeckCard.vue';
// createdeckcard is a bad name?
import DeckCard from '../components/DeckCard.vue';

const decks = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
</script>
