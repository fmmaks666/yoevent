<script setup></script>

<template>
  <div class="spinner"></div>
</template>

<style scoped>
.spinner {
  width: 64px;
  height: 64px;
  max-width: 64px;
  max-height: 64px;
  border-radius: 50%;
  border: solid 8px green;
  border-top: solid 8px transparent;
  animation: spin 1s ease-in-out infinite;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}
</style>
