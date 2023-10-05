<template>
  <v-container align="center">
    <h2>Register customer</h2>
    <v-form @submit.prevent="registerCustomer">
      <v-text-field v-model="newClient.name" label="Name" required></v-text-field>
      <v-text-field v-model="newClient.email" label="Email" required></v-text-field>
      <v-btn type="submit" color="primary">Submit</v-btn>
    </v-form>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      newClient: {
        name: '',
        email: ''
      }
    };
  },
  methods: {
    async registerCustomer() {
      try {
        await this.$api.post('/Customer/add', this.newClient);
        
        this.newClient.name = '';
        this.newClient.email = '';

        alert('Customer registered successfully');
      } catch (error) {
        console.error('Error:', error);
        alert('Error: ' + error);
      }
    }
  }
};
</script>

<style>
.v-container {
  width: 30%!important;
}
</style>
