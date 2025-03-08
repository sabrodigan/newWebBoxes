<template>
  <ion-button :id="id" :class="style" @click="OnButtonClick" :disabled="isDisabled">
    {{ title }}
    <ion-spinner :name="spinnerName" v-if="isDisabled"></ion-spinner>
  </ion-button>
</template>

<script>
  import { IonButton, IonSpinner } from '@ionic/vue';

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
      }
    },
    data() {
      return {
        isButtonDisabled: false,
      }
    },
    computed: {
      isDisabled() {
        return this.isButtonDisabled
      }
    },
    methods: {
      async onButtonClick() {
        this.isButtonDisabled = true;
        await this.action(this.dependentIdsToBeBlocked);
        this.isButtonDisabled = false;
      }
    }
  }
</script>

<style scoped>

</style>