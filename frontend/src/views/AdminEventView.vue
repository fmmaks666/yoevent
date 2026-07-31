<script setup>
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getEvent, checkAuth, downloadVisits, getEventStats, updateEvent } from '../api/api.js'
import { useAppStore } from '../stores/app.js'
import AdminView from './AdminView.vue'
import AdminForm from '../components/AdminForm.vue'
import EventForm from '../components/EventForm.vue'
import EventOverview from '../components/EventOverview.vue'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'
import Card from '../components/Card.vue'
import Stats from '../components/Stats.vue'
import DateChooser from '../components/DateChooser.vue'

const { authToken, setAuthToken, getData } = useAppStore()

const router = useRouter()
const route = useRoute()
const client = useQueryClient()

const isEditing = ref(false)
const date = new Date()
const selectedDate = ref({ year: date.getFullYear(), month: date.getMonth() + 1 })

const {
  mutateAsync,
  isPending: isPendingDownload,
  isError: isErrorDownload,
  error: errorDownload,
} = useMutation({
  mutationFn: async ({ eventId, date, all = false }) => {
    const res = await downloadVisits(authToken, eventId, date, all)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  data: stats,
  mutateAsync: fetchStats,
  isPending: isPendingStats,
  isError: isErrorStats,
  error: errorStats,
} = useMutation({
  mutationFn: async ({ eventId, date, all = false }) => {
    const res = await getEventStats(authToken, eventId, date, all)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  data: eventData,
  isPending: isPendingEvent,
  isError: isErrorEvent,
  error: errorEvent,
} = useQuery({
  queryKey: ['event'],
  queryFn: async () => {
    const visitor = getData().visitor || ''
    const eventId = route.params.id || ''
    const res = await getEvent(eventId)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  mutateAsync: update,
  isPending: isPendingUpdate,
  isError: isErrorUpdate,
  error: errorUpdate,
} = useMutation({
  mutationFn: async (data) => {
    const res = await updateEvent(authToken, data)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

async function onStats() {
  try {
    const id = eventData.value.event_id
    const res = await fetchStats({
      eventId: id,
      date: selectedDate.value,
      all: eventData.value.is_onetime,
    })
  } catch (err) {
    console.error(err)
  }
}

async function onDownload() {
  try {
    const id = eventData.value.event_id
    const res = await mutateAsync({
      eventId: id,
      date: selectedDate.value,
      all: eventData.value.is_onetime,
    })
    const url = window.URL.createObjectURL(new Blob([res.csv]))
    const a = document.createElement('a')
    a.href = url
    a.download = `${id}.csv`
    a.click()
    window.URL.revokeObjectURL(url)
  } catch (err) {}
}

async function onSubmitted(data) {
  const newEvent = { ...data }
  newEvent.event_id = eventData.value.event_id
  await update(newEvent)
  client.invalidateQueries({ queryKey: ['event'] })
  isEditing.value = false
}
</script>

<template>
  <AdminView>
    <h2>Панель адміністратора</h2>
    <Spinner v-if="isPendingEvent" class="spinner" />
    <ErrorBox v-else-if="isErrorEvent" :message="errorEvent?.message" />
    <template v-else>
      <EventForm
        v-if="isEditing"
        :cancelable="true"
        :event="eventData"
        @submit="onSubmitted"
        @cancel="isEditing = false"
      />
      <!-- TODO: This dickhead has a problem of throwing up when it's an invalid date, handle it
        somehow, brother -->
      <EventOverview v-else :data="eventData" @edit="isEditing = !isEditing" />
      <ErrorBox v-if="isErrorUpdate && isEditing" :message="errorUpdate?.message" />
      <Spinner v-else-if="isPendingUpdate" />
    </template>
    <!-- <Error v-else-if="isError || (data && data.error)" :message="error || data.error" />
    <template v-else>
      <h1>Відвідувачам "{{ data.title }}"</h1>
      <img :src="qr" alt="Позначити, що ви були на події" />
    </template>
    <button class="returnBtn" @click="router.back()">Повернутись</button>
    -->
    <template v-if="eventData && !eventData.is_onetime">
      <h3>Оберіть проміжок часу</h3>
      <DateChooser class="chooser" v-model="selectedDate" />
    </template>
    <div class="tools">
      <button @click="onStats">Показати статистику</button>
      <button data-variant="secondary" @click="onDownload">Список відвідувачів</button>
    </div>
    <Stats v-if="!isPendingStats && !isErrorStats && stats" :data="stats" />
  </AdminView>
</template>

<style scoped>
.h1,
h3 {
  text-align: center;
  font-weight: bold;
  margin: 16px auto;
}
h2 {
  text-align: center;
  margin-bottom: 1rem;
}

.list {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.tools {
  display: flex;
  gap: 8px;
  justify-content: center;
  margin-top: 8px;
  margin-bottom: 8px;
}

.spinner {
  margin: 0 auto;
}

.chooser {
  margin-bottom: 16px;
}

@media (min-width: 1024px) {
  h2 {
    margin-bottom: 5rem;
  }
  .list {
    gap: 0.3rem;
  }
}
</style>
