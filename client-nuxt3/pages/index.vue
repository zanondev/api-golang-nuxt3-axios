<template>
  <v-container align="center">
    <h2>Register customer</h2>
    <v-form @submit.prevent="cadastrarCliente">
      <v-text-field v-model="novoCliente.name" label="Name" required></v-text-field>
      <v-text-field v-model="novoCliente.email" label="Email" required></v-text-field>
      <v-btn type="submit" color="primary">Submit</v-btn>
    </v-form>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      novoCliente: {
        name: '',
        email: ''
      }
    };
  },
  methods: {
    async cadastrarCliente() {
      try {
        await this.$api.post('/Customer/add', this.novoCliente);
        
        this.novoCliente.name = '';
        this.novoCliente.email = '';

        alert('Customer registered successfully');
      } catch (error) {
        console.error('Error:', error);
        alert('Error: ', error);
      }
    }
  }
};
</script>

<style>
.v-container{
  width: 30%!important;
}

</style>