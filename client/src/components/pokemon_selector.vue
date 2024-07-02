<script setup>
import { computedAsync } from '@vueuse/core';

const props = defineProps(['names', 'pokedex', 'color', 'index'])

const selected = defineModel();


const src = computedAsync(async () => {
  let x = await props.pokedex.getPokemonByName(selected.value)
  console.log(selected.value)
  return x.sprites.front_default
})

selected.value = props.names[0]

</script>

<template>
  <main>
    <div>
    <img :src>
    <select v-model="selected">
      <option v-for="name in names">{{ name }}</option>
    </select>
    <span></span>
  </div>
  </main>
</template>

<style scoped>
  div {
    padding: 0.5em;
    background-color: v-bind(color);
    display: inline-flex;
    flex-direction: column;
    align-items: center;
    flex-basis: auto;
    border-radius: 40px;
    min-width: 16em;
  }
  img {
    background-color: white;
    border-radius: 10%;
  }
</style>