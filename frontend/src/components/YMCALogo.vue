<script setup></script>

<template>
  <picture>
    <!-- <source srcset="/YMCA_NS_KHUST_WHITE.png" media="(prefers-color-scheme: dark)" /> -->
    <source srcset="/YMCA.png" media="(prefers-color-scheme: dark)" />

    <img src="/YMCA_NS_KHUST_BLACK.png" />
  </picture>
</template>

<style scoped>
img {
  width: 270px;
}
/* rems? */
@media (max-width: 512px) {
  img {
    width: 135px;
  }
}
</style>
