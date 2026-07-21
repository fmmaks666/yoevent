<script setup>
import { ref } from 'vue'
import { formatDate, formatBirthdate } from '../utils/utils.js'
import Card from './Card.vue'

defineProps({
  data: {
    type: Object,
    required: true,
  },
})

const showMore = ref(false)
</script>

<template>
  <Card class="visitor-overview">
    <p>
      <span class="focus">{{ data.last_name }} {{ data.first_name }} {{ data.patronymic }} </span>
    </p>
    <p v-if="showMore">
      Вік: <span class="focus">{{ formatBirthdate(data.birthdate) }}</span> <br />
      Номер телефону: <span class="focus">{{ data.phone_number }}</span> <br />
      Стать: <span class="focus">{{ data.sex == 1 ? 'Чоловіча' : 'Жіноча' }}</span> <br />
      Соціальний статус:
      <span class="focus">{{ data.is_local ? 'Місцевий мешканець' : 'ВПО' }}</span> <br />
      <span class="focus">{{ data.is_disabled ? 'Є інвалідність' : 'Немає інвалідностей' }}</span>
      <br />
      Згода на обробку особистих даних:
      <span class="focus">{{ data.agreed_to_privacy ? 'Так' : 'Ні' }}</span>
    </p>
    <div class="actions">
      <button @click="showMore = !showMore" data-variant="secondary">
        {{ showMore ? 'Менше' : 'Більше' }}
      </button>
      <!--<button @click="">Скопіювати ID</button>-->
    </div>
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

.visitor-overview {
  text-align: start;
  display: flex;
  flex-direction: column;
  width: 100%;
}

.actions {
  width: 100%;
  display: flex;
  gap: 4px;
  justify-content: end;
}

@media (min-width: 680px) {
  button {
    justify-self: end;
    align-self: end;
    width: 30%;
  }
}

@media (min-width: 1024px) {
  button {
    width: 20%;
  }
}
</style>
