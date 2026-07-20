<script setup>
import { ref, onMounted } from 'vue'
import { useQuery } from '@tanstack/vue-query'
import { getVisits } from '../api/api.js'
import { formatVisitDate } from '../utils/utils.js'
import { useAppStore } from '../stores/app.js'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const ev = ref(null)
const { authToken, visitorData, getData, setData } = useAppStore()

const { data, isPending, isError, error } = useQuery({
  queryKey: ['visits'],
  queryFn: async () => {
    const visitor = getData().visitor || ''
    const res = await getVisits(visitor)
    return res.json()
  },
})

function onAction(id) {
  console.log(id)
}
</script>

<template>
  <main>
    <h2>Мої відвідування</h2>
    <div class="list">
      <Spinner v-if="isPending" class="spinner" />
      <ErrorBox v-else-if="isError" :message="error?.message" />

      <template v-else-if="data && data.length > 0">
        <Event
          @action="onAction"
          v-for="e in data"
          :id="e.event.event_id"
          :title="e.event.title"
          :description="e.event.description"
          :time="formatVisitDate(e)"
          :active="true"
          :disabled="true"
        />
      </template>
      <h3 v-else>Тут нічого. Відвідайте щось?</h3>
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
