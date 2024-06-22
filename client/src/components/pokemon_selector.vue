<script setup>
defineProps(['type', 'names'])
</script>

<template>
  <main>
    Pokemon Selector! Type: {{ type }}
    <select>
        <option v-for="name in names">{{ name }}</option>
    </select>
  </main>
</template>

<style scoped>
</style>
