<script setup>
import { watch } from 'vue'
import { RouterLink } from 'vue-router'
import { useRoute, useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getEvents, checkAuth, downloadVisits } from '../api/api.js'
import { useAppStore } from '../stores/app.js'
import AdminForm from '../components/AdminForm.vue'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const { authToken, setAuthToken } = useAppStore()

const router = useRouter()
const route = useRoute()
const client = useQueryClient()

const { data, isSuccess, isPending, isError, error } = useQuery({
  queryKey: ['check'],
  queryFn: async () => {
    if (!authToken) {
      return false
    }
    const res = await checkAuth(authToken)
    if (!res.ok) {
      throw new Error('Wrong password')
    }
    return res.ok
  },
})

client.refetchQueries({ queryKey: ['check'] })
if (!data) {
  router.push('/admin/login')
}

if (!authToken || authToken === btoa(':')) {
  router.push('/admin/login')
}

// TODO: Make sure that the redirect works
watch(isError, (failed) => {
  if (failed) {
    console.log(error?.message)
    setAuthToken('', '')
    router.push('/admin/login')
  }
})
</script>

<template>
  <nav>
    <RouterLink to="/admin/dashboard">Панель адміністратора</RouterLink>
    <RouterLink to="/admin/visitors">Відвідувачі</RouterLink>
  </nav>

  <main>
    <div class="spinner">
      <Spinner v-if="isPending" class="spinner" />
    </div>
    <div>
      <ErrorBox v-if="isError && !data" :message="error?.message" />
    </div>
    <slot v-if="data && !isError"></slot>
    <!-- <template v-else>
      <h1>Відвідувачам "{{ data.title }}"</h1>
      <img :src="qr" alt="Позначити, що ви були на події" />
    </template>
    <button class="returnBtn" @click="router.back()">Повернутись</button>
    -->
  </main>
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
h2 {
  text-align: center;
  /* margin-bottom: 1rem; */
}

.list {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.spinner {
  margin: 0 auto;
}

nav {
  width: 100%;
  text-align: center;
  margin: 0 4px;
  margin-top: 4px;
  margin-bottom: 8px;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

nav {
  font-size: 1.1rem;
  display: flex;
  flex-direction: row;
  text-align: center;

  align-items: center;
  justify-content: center;
  padding: 0 0;
}

main {
  min-height: 100vh;
}

@media (min-width: 1024px) {
  /* h2 {
    margin-bottom: 5rem;
  } */
  .list {
    gap: 0.3rem;
  }
}
</style>
