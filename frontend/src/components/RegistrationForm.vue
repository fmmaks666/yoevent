<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router' // IDEA: Move this logic to DataView
import { useMutation } from '@tanstack/vue-query'
import { useAppStore } from '../stores/app.js'
import { postVisitor } from '../api/api.js'
import { toLocalTime, toUTC } from '../utils/utils.js'
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
  until: '',
  max_registrations: 0,
  is_private: false,
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
  // The date's in UTC therefore we convert it to LOCAL TIME
  if (data.is_onetime) {
    const formatted = { ...data }
    const date = new Date(data.date)
    /*const [hour, min] = toLocalTime(date)
    date.setHours(hour)
    date.setMinutes(min)
    // TODO: Fix time: it's in UTC
    console.log(data.date, hour, min, date, date.toISOString())
    // PUT the properly formatted string here!!
    // yyyy-mm-ddThh:mm
  */
    formatted.date = toLocalTime(data.date) // formatted.date.slice(0, 16)
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
  // FUCK JS
  // We get this shoot in LOCAL TIME, therefore we have to convert it to UTC
  /* const [dateSeg, timeSeg] = data.date.split('T')
    const [year, month, day] = dateSeg.split('-').map(Number)
    const [hour, min] = timeSeg.split(':').map(Number)

    // Convert damn hours and mins to Euriope/Kyiv

    const date = new Date(year, month - 1, day, hour, min) */
  const normalized = { ...data }
  normalized.time = null
  // TODO: Basically we need to replace this fucked up line
  //normalized.date = date.toISOString()
  normalized.date = toUTC(data.date)
  return normalized
}
</script>

<template>
  <h1>{{ cancelable ? 'Редагувати' : 'Створити' }} реєстрацію</h1>
  <form @submit.prevent="onSubmit()">
    <label for="title">Інформація про реєстрацію</label>
    <input
      type="text"
      v-model.trim="data.title"
      placeholder="Назва реєстрації"
      id="title"
      required
    />
    <input type="text" v-model.trim="data.description" placeholder="Опис" required />
    <div v-if="cancelable" class="group">
      <input type="checkbox" v-model="data.is_private" name="private" id="private" />
      <label for="private">Схована реєстрація</label>
    </div>
    <input type="datetime-local" v-model.trim="data.until" required />
    <div class="group" name="max-registrations" id="max-registrations">
      <input
        type="number"
        v-model="data.max_registrations"
        name="max-registrations"
        id="max-registrations"
      />
      <label for="max-registrations">Вільних місць</label>
    </div>
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
