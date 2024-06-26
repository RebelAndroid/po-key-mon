import {defineStore} from 'pinia'
import { ref } from 'vue'

export const useSelectedStore = defineStore('selected', () => {
    const indices = ref([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0])

    return {indices}
})