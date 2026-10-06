<script setup>
import Card from '../components/Card.vue'
import Spinner from '../components/Spinner.vue'
import DeleteButton from '../components/DeleteButton.vue'

defineProps({
  id: {
    type: Number,
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
  until: {
    type: String,
    required: true,
    default: '',
  },
  registered: {
    type: Number,
    required: true,
  },
  'max-registrations': {
    type: Number,
    required: true,
  },
  mode: {
    type: String,
    required: false,
    default: 'register',
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
defineEmits(['action', 'delete', 'details'])
</script>

<template>
  <!-- Should have three modes: register, info, delete -->
  <Card class="event-box">
    <div class="info">
      <h3 class="green">
        {{ title }}
        <span class="regs"> · {{ registered }}/{{ maxRegistrations }}</span>
        <span class="until">· {{ until }}</span>
      </h3>
      <p>{{ description }}</p>
    </div>
    <div class="usable">
      <template v-if="!loading">
        <button
          v-if="mode == 'register'"
          :disabled="disabled ? true : undefined"
          :data-variant="active ? 'primary' : 'secondary'"
          @click="$emit('action', id)"
        >
          Я буду тут
        </button>
        <button
          v-else-if="mode == 'info'"
          :disabled="disabled ? true : undefined"
          :data-variant="active ? 'primary' : 'secondary'"
          @click="$emit('action', id)"
        >
          Деталі
        </button>
        <!-- Fix this satan's button -->
        <button
          v-else-if="mode == 'delete'"
          data-variant="destructive"
          @click="$emit('delete', id)"
        >
          Скасувати
        </button>
      </template>
      <Spinner class="spinner" v-else />
    </div>
  </Card>
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

.until {
  display: inline-block;
  font-size: 1rem;
  font-weight: normal;
  color: darkgray;
}

.regs {
  display: inline-block;
  font-weight: bold;
  font-size: 16px;
  color: var(--color-bold);
}

.spinner {
  align-self: flex-end;
}

button {
  justify-self: end;
  align-self: start;
  width: 100%;
}

.delete-btn {
  width: 100%;
  height: 70%;
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
  text-align: left;
}

.usable {
  width: 30%;
  justify-self: stretch;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-content: center;
}

.event-box {
  display: flex;
  flex-direction: row;
  width: 100%;
  justify-self: stretch;
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
  }
}
</style>
