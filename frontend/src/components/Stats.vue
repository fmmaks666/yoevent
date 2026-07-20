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
</script>

<template>
  <Card class="stats">
    <p>
      Кількість проведених подій: <span class="focus">{{ data.occurences }}</span> <br />
      Загальна кількість учасників: <span class="focus">{{ data.organizers + data.members }}</span>
      <br />
      Статус: Правовласники:
      <span class="focus"
        ><!-- {{ data.organizers }} -->
        —</span
      >
      <br />
      Статус: Носії обов'язків: <span class="focus">{{ data.members }}</span> <br />
      Стать: Чоловіки': <span class="focus">{{ data.men }}</span> <br />
      Стать: Жінки: <span class="focus">{{ data.women }}</span> <br />
      Стать: Інше: <span class="focus">—</span> <br />
      Кількість жінок до 30 років: <span class="focus">{{ data.women_below_30 }}</span> <br />
      Кількість чоловіків до 30 років: <span class="focus">{{ data.men_below_30 }}</span> <br />
      Кількість жінок з інвалідністю: <span class="focus">{{ data.disabled_women }}</span> <br />
      Кількість чоловіків з інвалідністю: <span class="focus">{{ data.disabled_men }}</span> <br />
      Кількість жінок ВПО: <span class="focus">{{ data.nonlocal_women }}</span> <br />
      Кількість чоловіків ВПО: <span class="focus">{{ data.nonlocal_men }}</span> <br />
      Вік до 15 років: <span class="focus">{{ data.age_less_than_15 }}</span> <br />
      Вік від 15 до 29 років: <span class="focus">{{ data.age_less_than_29 }}</span> <br />
      Вік від 30 до 35 років: <span class="focus">{{ data.age_more_than_30 }}</span> <br />
      Вік 35+ років: <span class="focus">{{ data.age_more_than_35 }}</span> <br />
      Вік: не визначено: <span class="focus">—</span> <br />
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

.stats {
  text-align: start;
  /* width: 100%; */
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
