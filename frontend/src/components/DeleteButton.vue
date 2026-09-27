<script setup>
import { ref } from 'vue'

defineProps({
  label: {
    type: String,
    required: false,
    default: 'Видалити',
  },
  labelYes: {
    type: String,
    required: false,
    default: 'Видалити!',
  },
})
const emit = defineEmits(['delete'])

const showConsent = ref(false)
</script>

<template>
  <div class="delete">
    <template v-if="showConsent">
      <button data-variant="secondary" @click="showConsent = !showConsent">Скасувати</button>
      <button data-variant="destructive" @click="emit('delete'); showConsent = !showConsent">{{ labelYes }}</button>
    </template>

    <button v-else data-variant="destructive" @click="showConsent = !showConsent">
      {{ label }}
    </button>
  </div>
</template>

<style scoped>
.delete {
  display: flex;
  flex-direction: row;
}
</style>
