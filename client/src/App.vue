<script setup>
import { Pokedex } from "pokeapi-js-wrapper";
import pokemon_selector from "./components/pokemon_selector.vue"
let fire_type_names = ["volcarona", "charizard", "emboar", "moltres", "reshiram"]
const pokedex = new Pokedex();
</script>

<template>
  <main>
    Hello from VUE!
    <pokemon_selector type="fire" :names = "fire_type_names" :pokedex></pokemon_selector>
  </main>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
