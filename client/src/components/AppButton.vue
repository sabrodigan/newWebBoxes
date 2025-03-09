<template>
  <ion-button :id="id" :class="styleClass" @click="onButtonClick" :disabled="isDisabled">
    {{ title }}
    <ion-spinner :name="spinnerName" v-if="isDisabled"></ion-spinner>
  </ion-button>
</template>

<script lang="js">
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
        Required: true
      },
      dependentIdsToBeDisabled: {
        type: Array,
        default: []
      }
    },
    data() {
      return {
        isButtonDisabled: false,
      };
    },
    methods: {
      async onButtonClick() {
        this.isButtonDisabled = true;
        setTimeout(async () => {
          await this.action(this.dependentIdsToBeDisabled)
          this.isButtonDisabled = false;
        }, 6000);
      }
    },
    computed: {
      isDisabled() {
        return this.isButtonDisabled
        }

    }}
</script>


<style scoped></style>