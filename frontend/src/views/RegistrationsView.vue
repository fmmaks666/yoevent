<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getRegistrations, postEventRegistration, deleteEventRegistration } from '../api/api.js'
import { formatDateRegistration, sortEvents } from '../utils/utils.js'
import { useAppStore } from '../stores/app.js'
import Registration from '../components/Registration.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const ev = ref(null)
const loadings = ref({})
const { authToken, visitorData, getData, setData } = useAppStore()
const router = useRouter()

const { data, isPending, isError, error } = useQuery({
  queryKey: ['registrations'],
  queryFn: async () => {
    const visitor = getData().visitor || ''
    const res = await getRegistrations(visitor)
    const json = await res.json()

    if (json && json.error) {
      throw new Error(json.error)
    }
    return sortEvents(json)
  },
})

const {
  mutateAsync,
  data: dataReg,
  isPending: isPendingReg,
  isError: isErrorReg,
  error: errorReg,
} = useMutation({
  mutationFn: async (reg) => {
    const visitor = getData().visitor || ''
    const res = await postEventRegistration(reg, visitor)
    const json = await res.json()

    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const {
  mutateAsync: deleteReg,
  data: dataDelete,
  isPending: isPendingDelete,
  isError: isErrorDelete,
  error: errorDelete,
} = useMutation({
  mutationFn: async (reg) => {
    const visitor = getData().visitor || ''
    const res = await deleteEventRegistration(reg, visitor)

    if (!res.ok) {
      const json = res.json()
      if (json && json.error) {
        throw new Error(json.error)
      }
    }
    return {}
  },
})
const client = useQueryClient()

async function onAction(id) {
  try {
    const oldData = getData()
    if (!oldData || Object.keys(oldData).length <= 0) {
      router.push({
        path: '/data',
        query: { registration: id },
      })
      return
    }
    loadings.value[id] = true
    const res = await mutateAsync(id)
    client.invalidateQueries({ queryKey: ['registrations'] })
    loadings.value[id] = false
  } catch (e) {
    console.error(e)
    loadings.value[id] = false
  }
}

async function onDelete(id) {
  try {
    // Would be stupid to request data lol
    loadings.value[id] = true
    const res = await deleteReg(id)
    client.invalidateQueries({ queryKey: ['registrations'] })
    loadings.value[id] = false
  } catch (e) {
    console.error(e)
    loadings.value[id] = false
  }
}
</script>

<template>
  <main>
    <h2>Реєстарація</h2>
    <h3>На події у вільному просторі "YO!"</h3>
    <ErrorBox v-if="isErrorReg" :message="errorReg?.message" class="event-error" />
    <div class="list">
      <Spinner v-if="isPending" class="spinner" />
      <ErrorBox v-else-if="isError" :message="error?.message" />
      <template v-else-if="data && data.length > 0">
        <Registration
          @action="onAction"
          @delete="onDelete"
          v-for="r in data"
          :key="r.registration_id"
          :id="r.registration_id"
          :title="r.title"
          :description="r.description"
          :until="formatDateRegistration(r)"
          :registered="r.registered"
          :max-registrations="r.max_registrations"
          :mode="r.is_registered ? 'delete' : 'register'"
          :loading="loadings[r.id] ? true : false"
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
  margin-bottom: 8px;
}
h3 {
  margin-bottom: 16px;
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
  h3 {
    margin-bottom: 64px;
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
