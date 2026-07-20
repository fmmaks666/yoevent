<script setup>
import { useRoute, useRouter } from 'vue-router'
import { useQRCode } from '@vueuse/integrations/useQRCode'
import { useQuery } from '@tanstack/vue-query'
import { getEvent } from '../api/api.js'
import AdminForm from '../components/AdminForm.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const router = useRouter()
const route = useRoute()
</script>

<template>
  <div class="admin-view">
    <AdminForm />
    <!--
    <Spinner v-if="isPending" class="spinner" />
    <Error v-else-if="isError || (data && data.error)" :message="error || data.error" />
    <template v-else>
      <h1>Відвідувачам "{{ data.title }}"</h1>
      <img :src="qr" alt="Позначити, що ви були на події" />
    </template>
    <button class="returnBtn" @click="router.back()">Повернутись</button>
    -->
  </div>
</template>

<style scoped>
.admin-view {
  text-align: center;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 2rem; /* TODO: Switch to pixels */
}

h1,
h3 {
  text-align: center;
  font-weight: bold;
}

img {
  width: 80vw;
  align-self: center;
}

.returnBtn {
  align-self: center;
  width: 80%;
}

@media (min-width: 1024px) {
  img {
    width: 30vw;
  }
  .returnBtn {
    width: 30%;
  }
}
</style>
