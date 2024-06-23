<script setup>
import { ref, computed } from 'vue';

const props = defineProps(['type', 'names', 'pokedex'])
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
    Pokemon Selector!
    <img :src>
    <select v-model="selected" @change="update_image_src">
      <option v-for="name in names">{{ name }}</option>
    </select>
  </main>
</template>

<style scoped></style>