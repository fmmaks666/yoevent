<script setup>
import { useRouter } from 'vue-router'
import { useQuery } from '@tanstack/vue-query'
import { getEvents } from '../api/api.js'
import Event from '../components/Event.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'

const router = useRouter()

const { data, isPending, isError, error } = useQuery({
  queryKey: ['events'],
  queryFn: async () => {
    const res = await getEvents('')
    return res.json()
  },
})

function onAction(id) {
  router.push(`/qr/${id}`)
}
</script>

<template>
  <main>
    <h2>Показати QR</h2>
    <div class="list">
      <Spinner v-if="isPending" class="spinner" />
      <ErrorBox v-else-if="isError || (data && data.error)" :message="error || data.error" />
      <template v-else-if="data && data.length > 0">
        <Event
          @action="onAction"
          v-for="e in data"
          :id="e.event_id"
          :title="e.title"
          :description="e.description"
          :time="e.date"
          button-text="Показати QR"
          :active="true"
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

  /*
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: stretch;
    align-self: stretch;
  } */
}
</style>
