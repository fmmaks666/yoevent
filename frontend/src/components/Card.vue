<script setup></script>

<template>
  <div class="card">
    <slot></slot>
  </div>
</template>

<style scoped>
.card {
  text-align: center;
  width: 100%;
  padding: 1rem;
  border: 0.1rem solid var(--color-border);
  border-radius: 16px;
  background-color: var(--color-background-soft);
}

.card + .card {
  border-top: 0;
}

@media (max-width: 1024px) {
  .card + .card {
    border-radius: 0;
  }

  .card:first-child {
    border-radius: 16px 16px 0 0;
  }

  .card:last-child {
    border-radius: 0 0 16px 16px;
  }

  .card:only-child {
    border-radius: 16px;
  }
}

@media (min-width: 1024px) {
  .card {
    width: 40%;
    margin: 0 auto; /* HACK */
    justify-self: stretch;
    padding: 1rem;
    border: 0.2rem solid var(--color-border);
    border-color: white;
    border-radius: 16px;
    background-color: var(--color-background-soft);
    filter: drop-shadow(0px 0.2rem 0.1rem darkgray);
  }
}
</style>
