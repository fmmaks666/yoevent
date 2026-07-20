<script setup>
import { ref } from 'vue'
import { useAppStore } from '../stores/app.js'
import DataForm from '../components/DataForm.vue'
import DataOverview from '../components/DataOverview.vue'

const editing = ref(false)

const { authToken, visitorData, getData, setData } = useAppStore()

const oldData = getData()
const isCancelable = ref(false)
if (oldData && Object.keys(oldData).length > 0) {
  isCancelable.value = true
}

function onSubmitted() {
  isCancelable.value = true
  editing.value = false
}
</script>

<template>
  <div class="data-view">
    <DataForm @submit="onSubmitted()" @cancel="editing = false" v-if="editing || !isCancelable" />
    <DataOverview @edit="editing = !editing" :data="getData()" v-else />
  </div>
</template>

<style scoped>
.data-view {
  text-align: center;
  min-height: 100vh;
}
@media (min-width: 1024px) {
}
</style>
