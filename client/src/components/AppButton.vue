<template>
  <ion-button :id="id" :class="style" @click="OnButtonClick" :disabled="isDisabled">
  </ion-button>
</template>

import { IonButton, IonSpinner } from '@ionic/vue';

<script>
  export default {
    name: 'AppButton',
    components: {
      IonButton,
      IonSpinner
    },
    props: {
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
      },
    computed: {
        isDisabled() {
          return true
      }
    },
    methods: {
      async onButtonClick() {
        await this.action(this.dependentIdsToBeBlocked);
      }
    },
    }
  }
</script>

<style scoped>

</style>
