<script setup>
import { ref, toRef } from 'vue'
import { useRoute, useRouter } from 'vue-router' // IDEA: Move this logic to DataView
import { useMutation } from '@tanstack/vue-query'
import { useAppStore } from '../stores/app.js'
import { postVisitor, updateVisitor } from '../api/api.js'
import ErrorBox from '../components/ErrorBox.vue'
import Spinner from '../components/Spinner.vue'
import DataOverview from '../components/DataOverview.vue'
import { convertData, defaultData } from '../utils/utils.js'

const props = defineProps({
  data: {
    //type: Object, nullable
    required: true,
  },
  cancelable: {
    type: Boolean,
    required: true,
  },
  showGreeting: {
    type: Boolean,
    required: false,
    default: true,
  },
})
const emit = defineEmits(['submit', 'cancel'])

const router = useRouter()
const route = useRoute()

/* const { mutateAsync, isPending, isError, error } = useMutation({
  mutationFn: async ({ data, update = false }) => {
    const res = await (update ? updateVisitor(data) : postVisitor(data))
    const json = await res.json()
    if (json && json.error) {
      throw new Error(json.error)
    }
    return json
  },
})*/

// TODO: SEPARATION OF DAMN CONCERNS, I WILL NEED TO REUSE THIS

const data = ref(props.data !== undefined ? convertData({ ...props.data }) : defaultData())
const isCancelable = props.cancelable
const showGreeting = props.showGreeting

async function onSubmit() {
  emit('submit', data.value)
}
</script>

<template>
  <template v-if="showGreeting">
    <h1 v-if="data.first_name">Привіт, {{ data.first_name }}</h1>
    <h1 v-else>Ваші дані</h1>
  </template>
  <form @submit.prevent="onSubmit()">
    <label for="name">ПІБ</label>
    <input type="text" v-model.trim="data.last_name" placeholder="Прізвище" id="name" required />
    <input type="text" v-model.trim="data.first_name" placeholder="Ім'я" required />
    <input type="text" v-model.trim="data.patronymic" placeholder="По батькові" required />

    <label for="birthdate">Дата народження</label>
    <input type="date" v-model="data.birthdate" id="birthdate" required />

    <label for="phone-num">Номер телефону</label>
    <input
      type="tel"
      v-model.trim="data.phone_number"
      placeholder="0685551010"
      id="phone-num"
      required
    />
    <input
      type="text"
      v-model.trim="data.residence"
      placeholder="Місто проживання"
      id="residence"
      required
    />
    <div class="group">
      <div class="group">
        <input type="radio" v-model.number="data.sex" name="sex" id="female" :value="2" />
        <label for="female">Жінка</label>
      </div>
      <div class="group">
        <input type="radio" v-model.number="data.sex" name="sex" id="male" :value="1" />
        <label for="male">Чоловік</label>
      </div>
    </div>

    <!-- <div class="group">
      <input type="checkbox" v-model="data.is_local" name="local" id="local"/>
      <label for="local">Місцевий мешканець (Інакше ВПО)</label>
    </div> -->
    <div class="group">
      <input type="checkbox" v-model="data.has_moved" name="local" id="moved" />
      <label for="moved">ВПО</label>
    </div>

    <div class="group">
      <input type="checkbox" v-model="data.is_disabled" name="disability" id="disability" />
      <label for="disability">Наявність інвалідності</label>
    </div>
    <div class="group">
      <input type="checkbox" v-model="data.agreed_to_privacy" name="privacy" id="privacy" />
      <label for="privacy"
        >Надаю добровільну згоду на обробку моїх персональних даних, необхідних для участі в заході
        та звітності проєкту, а також на фото- та відеозйомку з подальшим використанням отриманих
        матеріалів організацією відповідно до Закону України «Про захист персональних даних» №
        2297-VI від 01.06.2010 (зі змінами).</label
      >
    </div>

    <input type="submit" value="Прийняти" />
    <button data-variant="secondary" v-if="isCancelable" class="cancel" @click="$emit('cancel')">
      Скасувати
    </button>

    <!-- <input type="submit" value="Нова людина" /> -->
  </form>
</template>

<style scoped>
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
  align-items: flex-start;
}

.group input[type='checkbox'],
.group input[type='radio'] {
  width: 28px;
  height: 28px;
  min-width: 28px;
  min-height: 28px;
}

.group label {
  text-align: left;
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
