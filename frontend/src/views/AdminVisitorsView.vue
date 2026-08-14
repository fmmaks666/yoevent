<script setup>
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useQuery, useMutation, useQueryClient } from '@tanstack/vue-query'
import { getVisitors, checkAuth, postVisitor } from '../api/api.js'
import { defaultData } from '../utils/utils.js'
import { useAppStore } from '../stores/app.js'
import AdminView from './AdminView.vue'
import AdminForm from '../components/AdminForm.vue'
import EventForm from '../components/EventForm.vue'
import EventOverview from '../components/EventOverview.vue'
import VisitorOverview from '../components/VisitorOverview.vue'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'
import Card from '../components/Card.vue'
import DateChooser from '../components/DateChooser.vue'
import DataForm from '../components/DataForm.vue'

const { authToken, setAuthToken, getData } = useAppStore()

const router = useRouter()
const route = useRoute()
const client = useQueryClient()

const { data, isPending, isError, error } = useQuery({
  queryKey: ['visitors'],
  queryFn: async () => {
    const res = await getVisitors(authToken)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const formKey = ref(0)
const visitorData = ref(defaultData())
const {
  mutateAsync,
  isPending: isPendingVisitor,
  isError: isErrorVisitor,
  error: errorVisitor,
} = useMutation({
  mutationFn: async (data) => {
    const res = await postVisitor(data)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})
async function onSubmit(d) {
  try {
    // TODO: INVALIDATE BS
    const resData = await mutateAsync(d)
    if (isError.value) return
    client.invalidateQueries({ queryKey: ['visitors'] })

    visitorData.value = defaultData()
    formKey.value++

    const hash = resData.visitor
    await window.navigator.clipboard.writeText(hash)
  } catch (e) {
    console.error(e)
  }
}
</script>

<template>
  <AdminView>
    <h2>Відвідувачі</h2>
    <DataForm
      :key="formKey"
      @submit="onSubmit"
      :data="visitorData"
      :cancelable="false"
      :show-greeting="false"
    />
    <div v-if="isErrorVisitor" class="error">
      <ErrorBox :message="errorVisitor?.message" />
    </div>
    <div v-if="isPendingVisitor" class="spinner">
      <Spinner />
    </div>
    <Spinner v-if="isPending" />
    <ErrorBox v-else-if="isError" :message="error?.message" />
    <div v-else-if="data && data.length > 0" class="list">
      <VisitorOverview v-for="v in data" :key="v.hash" :data="v" />
    </div>
    <h3 v-else>Немає людинок</h3>
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

.error {
  margin-bottom: 8px;
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
