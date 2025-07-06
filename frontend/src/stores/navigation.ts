import { ref } from 'vue'
import { defineStore } from 'pinia'

export const useNavigationStore = defineStore('navigation', () => {
  const section = ref('')

  function navigate(sec: string) {
    section.value = sec
  }

  return { section, navigate }
})
