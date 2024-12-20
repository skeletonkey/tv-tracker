<script setup>
import { ref, computed } from 'vue'
import Home from './pages/HomePage.vue'
import About from './pages/AboutPage.vue'
import NotFound from './pages/NotFoundPage.vue'
import Search from './pages/SearchPage.vue'

const routes = {
  '/': Home,
  '/about': About,
  '/search': Search
}

const currentPath = ref(window.location.hash)

window.addEventListener('hashchange', () => {
  currentPath.value = window.location.hash
})

const currentView = computed(() => {
  return routes[currentPath.value.slice(1) || '/'] || NotFound
})
</script>

<template>
  <a href="#/">Home</a> |
  <a href="#/search">Search</a> |
  <a href="#/about">About</a> |
  <a href="#/non-existent-path">Broken Link</a>
  <component :is="currentView" />
</template>