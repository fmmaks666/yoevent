<script setup>
import { RouterLink, RouterView } from 'vue-router'
import { VueQueryDevtools } from '@tanstack/vue-query-devtools'
import DataView from './views/DataView.vue'
import YMCALogo from './components/YMCALogo.vue'

const author = import.meta.env.VITE_AUTHOR
const contact = import.meta.env.VITE_CONTACT_INFO
</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/"><YMCALogo /></RouterLink>
        <RouterLink to="/">Події</RouterLink>
        <RouterLink to="/data">Мої дані</RouterLink>
        <RouterLink to="/visits">Відвідане</RouterLink>
      </nav>
    </div>
  </header>

  <RouterView />
  <footer>
    <RouterLink to="/admin/dashboard"><u>Для адміністраторів</u></RouterLink>
    <br />
    YMCA Нижнє Селище-Хуст. Сайт розробляє {{ author }}. Проблеми, Ідеї, Побажання: {{ contact }} у
    Telegram
    <!-- ·
    <RouterLink to="/tools">Для організаторів</RouterLink> -->
  </footer>

  <VueQueryDevtools />
</template>

<style scoped>
header {
  line-height: 1.5;
  max-height: 100vh;
}

nav {
  width: 100%;
  text-align: center;
  margin: 0 4px;
  margin-top: 32px;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

nav {
  font-size: 1.1rem;
  display: flex;
  flex-direction: row;
  text-align: center;

  align-items: center;
  justify-content: center;
  padding: 0 0;
}

footer a {
  color: black;
}

footer {
  border-radius: 1rem 1rem 0 0;
  margin: 0;
  margin-top: 1.5rem;
  background-color: lightgray;
  color: black;
  padding: 1rem;
  text-align: center;
  text-justify: center;
}

@media (min-width: 1024px) {
  /* header {
    display: flex;
    place-items: center;
    flex-direction: row;
  } */

  /* header .wrapper {
    width: 100%;
    direction: row;
    display: flex;
    align-items: center;
    flex-wrap: wrap;
  } */

  nav {
    margin-top: 1rem;
    padding: 1rem 0;
    font-size: 1.3rem;
  }

  footer {
    border-radius: 1rem;
  }
}
</style>
