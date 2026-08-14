<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useMutation } from '@tanstack/vue-query'
import { useAppStore } from '../stores/app.js'
import { postVisitor, updateVisitor } from '../api/api.js'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'
import DataForm from '../components/DataForm.vue'
import DataOverview from '../components/DataOverview.vue'
import { convertData, defaultData } from '../utils/utils'

const router = useRouter()
const route = useRoute()

const editing = ref(false)

const { authToken, visitorData, getData, setData } = useAppStore()
const { mutateAsync, isPending, isError, error } = useMutation({
  mutationFn: async ({ data, update = false }) => {
    const res = await (update && data.hash ? updateVisitor(data) : postVisitor(data))
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const data = ref(defaultData())
const oldData = getData()
const isCancelable = ref(false)
if (oldData && Object.keys(oldData).length > 0) {
  isCancelable.value = true
  data.value = convertData(oldData)
}

async function onSubmit(d) {
  // Make a request to fucking get the hash

  try {
    // TODO: INVALIDATE BS
    data.value.hash = getData().visitor // WHY
    d.hash = getData().visitor
    const resData = await mutateAsync({ data: d, update: isCancelable.value })
    if (isError.value) return
    data.value.visitor = resData['visitor']
    setData(convertData(resData))

    isCancelable.value = true
    editing.value = false

    if (route.query.event) {
      router.push(`/v/${route.query.event}`)
    }
  } catch (e) {
    console.error(e)
    throw e
  }
}
</script>

<template>
  <div class="data-view">
    <DataForm
      :data="isCancelable ? getData() : undefined"
      :cancelable="isCancelable"
      @submit="onSubmit"
      @cancel="editing = false"
      v-if="editing || !isCancelable"
    />
    <DataOverview @edit="editing = !editing" :data="getData()" v-else />

    <div v-if="isError">
      <ErrorBox :message="error?.message" />
    </div>
    <div v-if="isPending" class="spinner">
      <Spinner />
    </div>
  </div>
</template>

<style scoped>
.data-view {
  text-align: center;
  min-height: 100vh;
}
.spinner {
  display: flex;
  flex-direction: column;
  width: 100%;
  align-items: center;
}
@media (min-width: 1024px) {
}
</style>
