<script setup>
import { ref } from 'vue'
import { formatDate } from '../utils/utils.js'
import Card from './Card.vue'

defineProps({
  data: {
    type: Object,
    required: true,
  },
})

defineEmits(['edit'])
</script>

<template>
  <Card>
    <p>
      <span class="focus">{{ data.title }}</span>
      <br />
      Опис: <span class="focus description">{{ data.description }}</span> <br />
      <span class="focus">{{ data.is_onetime ? 'Одноразова подія' : 'Регулярна подія' }}</span>
      <br />
      <span class="focus">{{ formatDate(data) }}</span>
      <br />
      <!-- Номер телефону: <span class="focus">{{ data.phone_number }}</span> <br />
      Стать: <span class="focus">{{ data.sex == 1 ? 'Чоловіча' : 'Жіноча' }}</span> <br />
      Соціальний статус:
      <span class="focus">{{ data.is_local ? 'Місцевий мешканець' : 'ВПО' }}</span> <br />
      <span class="focus">{{ data.is_disabled ? 'Є інвалідність' : 'Немає інвалідностей' }}</span>
      <br />
      Згода на обробку особистих даних:
      <span class="focus">{{ data.agreed_to_privacy ? 'Так' : 'Ні' }}</span>
    </p> -->
      <button @click="$emit('edit')">Редагувати</button>
    </p>
  </Card>
</template>

<style scoped>
h3 {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: -0.4rem;
}

button {
  margin-top: 8px;
  width: 80%;
}
.focus {
  font-weight: 600;
}

.description {
  text-align: start;
}

@media (min-width: 680px) {
  button {
    justify-self: end;
    align-self: end;
    width: 60%;
  }
}

.info {
  width: 70%;
}

@media (min-width: 1024px) {
  button {
    width: 30%;
  }
}
</style>
