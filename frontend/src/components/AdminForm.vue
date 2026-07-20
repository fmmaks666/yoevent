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

const emit = defineEmits(['login'])

const { authToken, setAuthToken, visitorData, getData, setData } = useAppStore()

const data = ref({
  username: 'admin',
  password: '',
})

//const oldData = getData()
//if (oldData && Object.keys(oldData).length > 0) data.value = oldData

function onSubmit() {
  // Make a request to fucking get the hash
  setAuthToken(data.value.username, data.value.password)
  router.push('/admin/dashboard')
  try {
  } catch (error) {
    console.error(error)
  }
}
</script>

<template>
  <h1>Увійдіть, щоб отримати доступ</h1>
  <form @submit.prevent="onSubmit()">
    <label for="password">Пароль адміністратора</label>
    <input type="password" v-model="data.password" placeholder="Пароль" id="password" />

    <input type="submit" value="Увійти" />
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
  flex-direction: column;
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

.entry input {
  margin-top: 0.2rem;
  width: 100%;
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
