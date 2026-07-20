<script setup>
defineProps({
  message: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="error-box">
    <h3>Помилка!</h3>
    <p>{{ message }}</p>
  </div>
</template>

<style scoped>
h3 {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: -0.4rem;
}

p {
  text-align: left;
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

.error-box {
  color: white;
  display: column;
  flex-direction: row;
  width: 100%;
  align-items: flex-start;
  justify-self: stretch;
  padding: 1rem;
  border: 0.1rem solid var(--color-border-error);
  border-radius: 1rem;
  background-color: var(--color-background-error);
}

/* When it's the first or last child ... */
.error-box + .error-box {
  border-top: 0rem;
}

@media (max-width: 1024px) {
  .error-box + .error-box {
    border-radius: 0;
  }

  .error-box:first-child {
    border-radius: 1rem 1rem 0 0;
  }

  .error-box:last-child {
    border-radius: 0 0 1rem 1rem;
  }

  .error-box:only-child {
    border-radius: 1rem;
  }
}
@media (min-width: 1024px) {
  p {
    text-align: left;
  }

  .greetings h1,
  .greetings h3 {
    text-align: left;
  }

  .error-box {
    color: white;
    display: flex;
    flex-direction: column;
    width: 100%;
    justify-self: stretch;
    align-items: flex-start;
    padding: 1rem;
    border: 0.2rem solid var(--color-border-error);
    border-radius: 1rem;
    background-color: var(--color-background-error);
    filter: drop-shadow(0px 0.2rem 0.1rem darkgray);
  }
}
</style>
