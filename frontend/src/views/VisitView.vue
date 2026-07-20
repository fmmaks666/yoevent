<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useQuery } from '@tanstack/vue-query'
import { postVisit } from '../api/api.js'
import { useAppStore } from '../stores/app.js'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const router = useRouter()
const route = useRoute()

const { authToken, visitorData, getData, setData } = useAppStore()

// TODO: Optimize getData calls
// TODO: That is to say, you need some sanity checks here
if (getData() && Object.keys(getData()).length === 0) {
  console.error('REDIRECT MY ASS')
  router.push({
    path: '/data',
    query: { event: route.params.id },
  })
}

const { data, isPending, isError, error } = useQuery({
  queryKey: ['visit'],
  queryFn: async () => {
    const visitor = getData().visitor || ''
    const event = route.params.id || ''
    const res = await postVisit(event, visitor)
    return res.json()
  },
})
</script>

<template>
  <main>
    <div class="list">
      <Spinner v-if="isPending" class="spinner" />
      <ErrorBox v-else-if="isError || (data && data.error)" :message="error || data.error" />
      <template v-if="!isPending && !isError && !data.error">
        <h2>
          Дякуємо, <br />
          що відвідали <wbr /> "{{ data.event.title }}"
        </h2>
      </template>
    </div>
  </main>
</template>

<style scoped>
h2,
h3 {
  text-align: center;
  margin-bottom: 1rem;
}

.list {
  display: flex;
  flex-direction: column;
  align-items: center;
}

main {
  min-height: 100vh;
}

@media (min-width: 1024px) {
  h2 {
    margin-bottom: 5rem;
  }

  .list {
    gap: 0.3rem;
  }

  /* main {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: stretch;
    align-self: stretch;
  } */
}
</style>
