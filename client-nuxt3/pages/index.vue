<template>
  <div>
    <h1>Clientes</h1>
    <v-form @submit.prevent="cadastrarCliente">
      <v-text-field v-model="novoCliente.name" label="Nome" required></v-text-field>
      <v-text-field v-model="novoCliente.email" label="Email" required></v-text-field>
      <v-btn type="submit" color="primary">Cadastrar</v-btn>
    </v-form>
  </div>
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
        // Realize a solicitação POST para a sua API com os dados do novo cliente
        await this.$api.post('/Customer/add', this.novoCliente);
        
        // Limpe os campos do formulário após o cadastro bem-sucedido
        this.novoCliente.name = '';
        this.novoCliente.email = '';

        // Você também pode adicionar uma mensagem de sucesso aqui usando o Vuetify's snackbar
        this.$snackbar.show({
          text: 'Cliente cadastrado com sucesso',
          color: 'success'
        });
      } catch (error) {
        console.error('Erro ao cadastrar cliente:', error);
        // Lide com erros de solicitação, por exemplo, exibindo uma mensagem de erro usando o Vuetify's snackbar
        this.$snackbar.show({
          text: 'Erro ao cadastrar cliente',
          color: 'error'
        });
      }
    }
  }
};
</script>

<style>

</style>