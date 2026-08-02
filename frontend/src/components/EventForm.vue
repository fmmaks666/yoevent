<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router' // IDEA: Move this logic to DataView
import { useMutation } from '@tanstack/vue-query'
import { useAppStore } from '../stores/app.js'
import { postVisitor } from '../api/api.js'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'
import DataOverview from '../components/DataOverview.vue'

// Cancelable in other words we're editing
const props = defineProps({
  cancelable: {
    type: Boolean,
    required: false,
    default: false,
  },
  event: {
    type: Object,
    required: false,
    default: undefined,
  },
})

const emit = defineEmits(['submit', 'cancel'])

const router = useRouter()
const route = useRoute()

const { authToken, visitorData, getData, setData } = useAppStore()
const { mutateAsync, isPending, isError, error } = useMutation({
  mutationFn: async (data) => {
    const res = await postVisitor(data)
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})

const data = ref({
  title: '',
  description: '',
  date: '',
  time: '',
  needs_registration: false,
  is_cancelled: false,
  is_private: false,
  is_onetime: false,
})

if (props.event) data.value = formatData({ ...props.event })

async function onSubmit() {
  emit('submit', normalizeData(data.value))
  // Make a request to fucking get the hash
  try {
  } catch (e) {
    console.error(e)
  }
}

function formatData(data) {
  if (data.is_onetime) {
    const formatted = { ...data }
    const date = new Date(data.date)
    // TODO: Fix time: it's in UTC
    formatted.date = formatted.date.slice(0, 16)
    return formatted
  }
  const date = new Date(data.time)
  const hour = date.getHours()
  const min = date.getMinutes()
  const time = `${String(hour).padStart(2, '0')}:${String(min).padStart(2, '0')}`
  const formatted = { ...data }
  formatted.time = time
  return formatted
}

function normalizeData(data) {
  if (data.is_onetime) {
    const [dateSeg, timeSeg] = data.date.split('T')
    const [year, month, day] = dateSeg.split('-').map(Number)
    const [hour, min] = timeSeg.split(':').map(Number)

    const date = new Date(year, month - 1, day, hour, min)
    const normalized = { ...data }
    normalized.time = null
    normalized.date = date.toISOString()
    return normalized
  }
  const [hour, min] = data.time.split(':')
  // WHAT A HACKY WAY OF EXISTING, Neh?
  console.log(hour, min)
  const time = new Date(1970, 0, 1, hour, min) // TODO: Switch to just HH:mm which might be tough in the future?
  const normalized = { ...data }
  normalized.date = null
  normalized.time = time.toISOString()
  return normalized
}
</script>

<template>
  <h1>{{ cancelable ? 'Редагувати' : 'Створити' }} подію</h1>
  <form @submit.prevent="onSubmit()">
    <label for="title">Інформація про події</label>
    <input type="text" v-model.trim="data.title" placeholder="Назва події" id="title" required />
    <input type="text" v-model.trim="data.description" placeholder="Опис" required />
    <div class="group">
      <input type="checkbox" v-model="data.is_onetime" name="onetime" id="onetime" />
      <label for="onetime">Одноразова подія</label>
    </div>
    <div v-if="cancelable" class="group">
      <input type="checkbox" v-model="data.is_private" name="private" id="private" />
      <label for="private">Схована подія</label>
    </div>

    <input v-if="data.is_onetime" type="datetime-local" v-model.trim="data.date" required />
    <template v-else>
      <input type="time" v-model.trim="data.time" required />
      <select v-model="data.weekday" required>
        <option :value="1">Понеділок</option>
        <option :value="2">Вівторок</option>
        <option :value="3">Середа</option>
        <option :value="4">Четвер</option>
        <option :value="5">П'ятниця</option>
        <option :value="6">Субота</option>
        <option :value="0">Неділя</option>
      </select>
    </template>
    <input type="submit" value="Прийняти" />
    <button data-variant="secondary" v-if="cancelable" class="cancel" @click="$emit('cancel')">
      Скасувати
    </button>
  </form>
  <div v-if="isError">
    <ErrorBox :message="error?.message" />
  </div>
  <div v-if="isPending" class="spinner">
    <Spinner />
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
  margin-bottom: 8px;
}

form {
  display: flex;
  flex-direction: column;
  gap: 8px;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
  margin-bottom: 2rem; /* TODO: SWITCH TO PXs */
}

form input,
form div,
select,
button {
  width: 25%;
}

.entry input {
  margin-top: 0.2rem;
  width: 100%;
}

.group {
  display: flex;
  gap: 8px;
}

.group input[type='checkbox'] {
  width: 28px;
  height: 28px;
  min-width: 28px;
  min-height: 28px;
}

.spinner {
  display: flex;
  flex-direction: column;
  width: 100%;
  align-items: center;
}

@media (max-width: 1024px) {
  form input,
  form div,
  select,
  button {
    width: 80%;
  }

  .group {
    align-self: center;
  }

  /*.group label {
    flex: 4;
    text-align: left;
  } */
  /* .group input {
    flex: 1;
    width: 2rem;
  } */
}
</style>
