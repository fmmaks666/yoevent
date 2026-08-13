<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router' // IDEA: Move this logic to DataView
import { useMutation } from '@tanstack/vue-query'
import { useAppStore } from '../stores/app.js'
import { postVisitor } from '../api/api.js'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const router = useRouter()
const route = useRoute()

const emit = defineEmits(['submit'])

const hash = ref('')

//const oldData = getData()
//if (oldData && Object.keys(oldData).length > 0) data.value = oldData

function onSubmit() {
  // Make a request to fucking get the hash
  emit('submit', hash.value)
  hash.value = ''
}
</script>

<template>
  <form @submit.prevent="onSubmit()">
    <input type="text" v-model="hash" placeholder="ID відвідувача" required />

    <input type="submit" value="Додати" />
    <!-- <input type="submit" value="Нова людина" /> -->
  </form>
  <!-- <div v-if="isError">
    <Error :message="error" />
  </div>
  <div v-if="isPending" class="spinner">
    <Spinner />
  </div>
  -->
</template>

<style scoped>
form {
  display: flex;
  flex-direction: row;
  gap: 0.5rem;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  margin-bottom: 2rem; /* TODO: SWITCH TO PXs */
}

form input,
form div {
  width: 25%;
}

form {
  width: 350px;
}

input[type='text'] {
  width: 70%;
}

input[type='submit'] {
  width: 30%;
}

.group {
  display: flex;
}

.spinner {
  display: flex;
  flex-direction: column;
  width: 100%;
  align-items: center;
}

@media (max-width: 1024px) {
  form input,
  form div {
    width: 80%;
  }

  .group {
    align-self: center;
  }

  .group label {
    flex: 4;
    text-align: left;
  }
  .group input {
    /* flex: 1; */
    width: 2rem;
  }
}
</style>
