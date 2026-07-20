<script setup>
import Spinner from '../components/Spinner.vue'

defineProps({
  id: {
    type: String,
    required: true,
  },
  title: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
  time: {
    type: String,
    required: true,
    default: '',
  },
  'button-text': {
    type: String,
    default: 'Я тут',
  },
  'is-onetime': {
    type: Boolean,
    default: true,
  },
  active: {
    type: Boolean,
    default: false,
  },
  loading: {
    type: Boolean,
    default: false,
  },
  disabled: {
    type: Boolean,
    default: false,
  },
})
defineEmits(['action'])
</script>

<template>
  <div class="event-box">
    <div class="info">
      <h3 class="green">
        {{ title }} <span class="time">· {{ time }}</span>
      </h3>
      <p>{{ description }}</p>
    </div>
    <div class="usable">
      <button
        v-if="!loading"
        :disabled="disabled ? true : undefined"
        :data-variant="active ? 'primary' : 'secondary'"
        @click="$emit('action', id)"
      >
        {{ buttonText }}
      </button>
      <Spinner v-else />
    </div>
  </div>
</template>

<style scoped>
button {
  text-align: center;
}
h3 {
  font-size: 1.2rem;
  font-weight: bold;
  margin-bottom: -0.4rem;
}

.time {
  display: inline-block;
  font-size: 1rem;
  font-weight: normal;
  color: darkgray;
}

button {
  justify-self: end;
  align-self: start;
  width: 100%;
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

.usable {
  width: 30%;
  justify-self: stretch;
  display: flex;
  justify-content: flex-end;
}

.event-box {
  display: flex;
  flex-direction: row;
  width: 100%;
  justify-self: stretch;
  padding: 16px;
  border: 0.1rem solid var(--color-border);
  border-radius: 16px;
  background-color: var(--color-background-soft);
}

/* When it's the first or last child ... */
.event-box + .event-box {
  border-top: 0rem;
}

@media (max-width: 1024px) {
  .event-box + .event-box {
    border-radius: 0;
  }

  .event-box:first-child {
    border-radius: 16px 16px 0 0;
  }

  .event-box:last-child {
    border-radius: 0 0 16px 16px;
  }

  .event-box:only-child {
    border-radius: 16px;
  }
}
@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }

  .event-box {
    display: flex;
    flex-direction: row;
    width: 100%;
    justify-self: stretch;
    padding: 16px;
    border: 0.2rem solid var(--color-border);
    border-color: white;
    border-radius: 16px;
    background-color: var(--color-background-soft);
    filter: drop-shadow(0px 0.2rem 0.1rem darkgray);
  }
}
</style>
