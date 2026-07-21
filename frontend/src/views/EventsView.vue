<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getEvents, postVisit } from '../api/api.js'
import { formatDate } from '../utils/utils.js'
import { useAppStore } from '../stores/app.js'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const ev = ref(null)
const loadings = ref({})
const { authToken, visitorData, getData, setData } = useAppStore()
const router = useRouter()

const { data, isPending, isError, error } = useQuery({
  queryKey: ['events'],
  queryFn: async () => {
    const visitor = getData().visitor || ''
    const res = await getEvents(visitor)
    const json = await res.json()

    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  mutateAsync,
  data: dataEvent,
  isPending: isPendingEvent,
  isError: isErrorEvent,
  error: errorEvent,
} = useMutation({
  mutationFn: async (event) => {
    const visitor = getData().visitor || ''
    const res = await postVisit(event, visitor)
    const json = await res.json()

    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const client = useQueryClient()

async function onAction(id) {
  try {
    const oldData = getData()
    if (!oldData || Object.keys(oldData).length <= 0) {
      router.push({
        path: '/data',
        query: { event: id },
      })
      return
    }
    loadings.value[id] = true
    const res = await mutateAsync(id)
    client.invalidateQueries({ queryKey: ['events'] })
    loadings.value[id] = false
  } catch (e) {
    console.error(e)
    loadings.value[id] = false
  }
}

function isVisited(e) {
  if (e.is_onetime) {
    return e.last_visit === e.date
  }

  let days = 0
  const now = new Date()
  // ADD 3 hours to make the same with our time TODO: Play with the locales
  //now.setHours(now.getHours() + 3)
  const weekday = now.getDay()
  const targetT = new Date(e.time)
  const targetW = e.weekday
  const timeNow = now.getHours() * 60 * 60 + now.getMinutes() * 60
  const timeTarget = targetT.getHours() * 60 * 60 + targetT.getMinutes() * 60
  days = ((weekday - targetW + 6) % 7) + 1
  //return date.toString()
  if (weekday === targetW && timeNow >= timeTarget) {
    days = 0
  }
  const date = new Date()
  date.setDate(now.getDate() - days)
  date.setHours(targetT.getHours() + 3, targetT.getMinutes(), 0, 0) // THIS BS's in +00:00

  // MY SECOND GREATEST SIN

  return e.last_visit == date.toISOString().slice(0, 19) + '+03:00'
}
</script>

<template>
  <main>
    <h2>Події</h2>
    <ErrorBox v-if="isErrorEvent" :message="errorEvent?.message" class="event-error" />
    <div class="list">
      <Spinner v-if="isPending" class="spinner" />
      <ErrorBox v-else-if="isError" :message="error?.message" />
      <template v-else-if="data && data.length > 0">
        <Event
          @action="onAction"
          v-for="e in data"
          :key="e.event_id"
          :id="e.event_id"
          :title="e.title"
          :description="e.description"
          :time="formatDate(e)"
          :active="isVisited(e)"
          :loading="loadings[e.event_id] ? true : false"
          :disabled="isPendingEvent || isVisited(e)"
        />
      </template>
      <h3 v-else>Подій немає</h3>
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

.event-error {
  margin-bottom: 8px;
}

.error {
  width: 100%;
  margin-bottom: 16px;
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
