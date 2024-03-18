<template>
  <div class="px-4 flex flex-wrap gap-2">
    <div v-for="c in categories" :key="c" class="px-6 py-0.5 border border-neutral-800 rounded-full">{{ c }}</div>
    <div class="px-6 py-0.5 border border-neutral-800 bg-green-500 text-white rounded-full">Add</div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
const categories = ref(["Automotive", "Books", "Electronics", "Home", "Movies", "Music", "Outdoors", "Sports", "Tools", "Toys"]);
</script>