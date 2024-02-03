<template>
  <navbar />
  <h1 class="text-4xl text-center my-12">My Decks</h1>
  <div class="flex w-screen justify-center">
  <div class="lg:gap-8 w-full lg:w-3/4 xl:w-7/12 grid grid-cols-3 grid-rows-3">
    <create-deck-card />
    <template v-for="deck in decks" v-bind:key="deck">
      <deck-card  :deck="deck"/>
    </template>
  </div>
</div>s
</template>

<script setup>
import Navbar from '../components/Navbar.vue';
import CreateDeckCard from '../components/CreateDeckCard.vue';
// createdeckcard is a bad name?
import DeckCard from '../components/DeckCard.vue';

const decks = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
</script>
