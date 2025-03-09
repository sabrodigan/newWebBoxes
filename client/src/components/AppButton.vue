<template>
  <ion-button :id="id" :class="styleClass" @click="onButtonClick" :disabled="isDisabled">
    {{ title }}
    <ion-spinner :name="spinnerName" v-if="isDisabled"></ion-spinner>
  </ion-button>
</template>

<script setup>
import { ref, computed } from 'vue';
import { IonButton, IonSpinner } from '@ionic/vue';

defineProps({
  id: {
    type: String,
    default: ''
  },
  title: {
    type: String,
    default: ''
  },
  styleClass: {
    type: String,
    default: ''
  },
  spinnerName: {
    type: String,
    default: 'bubbles'
  },
  action: {
    type: Function,
    default: () => {}
  },
  dependentIdsToBeBlocked: {
    type: Array,
    default: () => []
  }
});

const isButtonDisabled = ref(false);
const isDisabled = computed(() => isButtonDisabled.value);

const onButtonClick = async () => {
  this.isButtonDisabled.value = true;
  try {
    await Promise.resolve(action(dependentIdsToBeBlocked));
  } catch (err) {
    console.error('Error running action:', err);
  } finally {
    isButtonDisabled.value = false;
  }
};


</script>

<style scoped></style>