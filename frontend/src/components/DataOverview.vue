<script setup>
import { ref } from 'vue'
import { formatBirthdate } from '../utils/utils.js'
defineProps({
  data: {
    type: Object,
    required: true,
  },
})

defineEmits(['edit'])
</script>

<template>
  <h1>Ваші дані</h1>
  <div class="data-overview">
    <p>
      <span class="focus">{{ data.last_name }} {{ data.first_name }} {{ data.patronymic }}</span>
      <br />
      Дата народження: <span class="focus">{{ formatBirthdate(data.birthdate) }}</span> <br />
      Номер телефону: <span class="focus">{{ data.phone_number }}</span> <br />
      Стать: <span class="focus">{{ data.sex == 1 ? 'Чоловіча' : 'Жіноча' }}</span> <br />
      Соціальний статус:
      <span class="focus">{{ data.is_local ? 'Місцевий мешканець' : 'ВПО' }}</span> <br />
      <span class="focus">{{ data.is_disabled ? 'Є інвалідність' : 'Немає інвалідностей' }}</span>
      <br />
      Згода на обробку особистих даних:
      <span class="focus">{{ data.agreed_to_privacy ? 'Так' : 'Ні' }}</span>
    </p>
    <button @click="$emit('edit')">Редагувати</button>
  </div>
</template>

<style scoped>
.data-overview {
  text-align: center;
  width: 100%;
  padding: 1rem;
  border: 0.1rem solid var(--color-border);
  border-radius: 1rem;
  background-color: var(--color-background-soft);
}

h3 {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: -0.4rem;
}

.dataoverview {
  font-size: 1.5em;
}

button {
  margin-top: 8px;
  width: 80%;
}
.focus {
  font-weight: 600;
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
  .data-overview {
    width: 40%;
    margin: 0 auto; /* HACK */
    justify-self: stretch;
    padding: 1rem;
    border: 0.2rem solid var(--color-border);
    border-color: white;
    border-radius: 1rem;
    background-color: var(--color-background-soft);
    filter: drop-shadow(0px 0.2rem 0.1rem darkgray);
  }
}
</style>
