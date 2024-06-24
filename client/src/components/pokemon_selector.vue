<script setup>
import { computed, ref } from 'vue';

const props = defineProps(['names', 'pokedex', 'color'])
const selected = ref(props.names[0])
const src = ref()

function update_image_src(){
  props.pokedex.getPokemonByName(selected.value).then(result => {
    src.value = result.sprites.front_default
  })
}

update_image_src()
</script>

<template>
  <main>
    <div>
    <img :src>
    <select v-model="selected" @change="update_image_src">
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