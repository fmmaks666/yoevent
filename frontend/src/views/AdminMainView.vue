<script setup>
import { watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getEventsAdmin, checkAuth, postEvent, downloadVisits } from '../api/api.js'
import { formatDate } from '../utils/utils.js'
import { useAppStore } from '../stores/app.js'
import AdminView from './AdminView.vue'
import AdminForm from '../components/AdminForm.vue'
import EventForm from '../components/EventForm.vue'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const { authToken, setAuthToken } = useAppStore()

const router = useRouter()
const route = useRoute()
const client = useQueryClient()

const {
  mutateAsync,
  isPending: isPendingDownload,
  isError: isErrorDownload,
  error: errorDownload,
} = useMutation({
  mutationFn: async (eventId) => {
    const res = await downloadVisits(authToken, eventId)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  data: events,
  isPending: isPendingEvents,
  isError: isErrorEvents,
  error: errorEvents,
} = useQuery({
  queryKey: ['admin', 'events'],
  queryFn: async () => {
    const res = await getEventsAdmin(authToken)
    const json = await res.json()

    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  mutateAsync: mutateEvent,
  isPending: isPendingEvent,
  isError: isErrorEvent,
  error: errorEvent,
} = useMutation({
  mutationFn: async (data) => {
    const res = await postEvent(authToken, data)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

async function onAction(id) {
  try {
    /*
    const res = await mutateAsync(id)
    const url = window.URL.createObjectURL(new Blob([res.csv]))
    const a = document.createElement('a')
    a.href = url
    a.download = `${id}.csv`
    a.click()
    window.URL.revokeObjectURL(url)
    */
    router.push(`/admin/event/${id}`)
  } catch (err) {}
}

async function createEvent(data) {
  await mutateEvent(data)
  client.invalidateQueries({ queryKey: ['events'] })
  client.invalidateQueries({ queryKey: ['admin', 'events'] })
}
</script>

<template>
  <AdminView>
    <h2>Панель адміністратора</h2>

    <EventForm @submit="createEvent" />
    <div class="creation">
      <Spinner v-if="isPendingEvent" class="spinner" />
      <ErrorBox v-else-if="isErrorEvent" :message="errorEvent?.message" />
    </div>
    <div class="list">
      <Spinner v-if="isPendingEvents" class="spinner" />
      <ErrorBox v-else-if="isErrorEvents" :message="errorEvents?.message" />
      <template v-else-if="events && events.length > 0">
        <Event
          @action="onAction"
          v-for="e in events"
          :id="e.event_id"
          :title="e.title"
          :description="e.description"
          :time="formatDate(e)"
          buttonText="Деталі"
        />
      </template>
      <h3 v-else>Подій немає. Створіть їх?</h3>
    </div>
    <!-- <Error v-else-if="isError || (data && data.error)" :message="error || data.error" />
    <template v-else>
      <h1>Відвідувачам "{{ data.title }}"</h1>
      <img :src="qr" alt="Позначити, що ви були на події" />
    </template>
    <button class="returnBtn" @click="router.back()">Повернутись</button>
    -->
  </AdminView>
</template>

<style scoped>
.h1,
h3 {
  text-align: center;
  font-weight: bold;
}
h2 {
  text-align: center;
  margin-bottom: 1rem;
}

.creation {
  margin-bottom: 16px;
}

.spinner {
  margin: 0 auto;
}

.list {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

@media (min-width: 1024px) {
  h2 {
    margin-bottom: 8px;
  }
  .list {
    gap: 0.3rem;
  }
}
</style>
