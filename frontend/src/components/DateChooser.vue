<script setup>
const model = defineModel()
</script>

<template>
  <div class="date-chooser">
    <input type="number" v-model="model.year" />
    <select v-model="model.month">
      <option :value="1">Січень</option>
      <option :value="2">Лютий</option>
      <option :value="3">Березень</option>
      <option :value="4">Квітень</option>
      <option :value="5">Травень</option>
      <option :value="6">Червень</option>
      <option :value="7">Липень</option>
      <option :value="8">Серпень</option>
      <option :value="9">Вересень</option>
      <option :value="10">Жовтень</option>
      <option :value="11">Листопад</option>
      <option :value="12">Грудень</option>
    </select>
  </div>
</template>

<style scoped>
.date-chooser {
  display: flex;
  gap: 8px;
  justify-content: center;
}
@media (min-width: 1024px) {
}
</style>
