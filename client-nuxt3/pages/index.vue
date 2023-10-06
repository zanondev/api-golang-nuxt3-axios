<template>
  <v-container align="center">
    <div class="form-container" aling="center">
      <h2>Register customer</h2>
      <v-form @submit.prevent="registerCustomer">
        <v-text-field
          v-model="newClient.name"
          label="Name"
          required
        ></v-text-field>
        <v-text-field
          v-model="newClient.email"
          label="Email"
          required
        ></v-text-field>
        <v-btn type="submit" color="primary">Submit</v-btn>
        <v-btn color="primary" @click="getCustomers" style="margin-left: 10px"
          >Get</v-btn
        >
      </v-form>
    </div>
    <div class="table-container">
      <v-table>
        <thead>
          <tr>
            <th class="text-center">ID</th>
            <th class="text-center">Name</th>
            <th class="text-center">Email</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="customer in customers" :key="customer.id">
            <td>{{ customer.id }}</td>
            <td>{{ customer.name }}</td>
            <td>{{ customer.email }}</td>
          </tr>
        </tbody>
      </v-table>
    </div>
  </v-container>
</template>

<script>
export default {
  data() {
    return {
      newClient: {
        name: "",
        email: "",
      },
      customers: [],
    };
  },
  methods: {
    async registerCustomer() {
      try {
        const response = await this.$api.post("/Customer/add", this.newClient);

        this.newClient.name = "";
        this.newClient.email = "";

        alert("Customer registered successfully");
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    async getCustomers() {
      try {
        const response = await this.$api.get("/Customers");
        // Atualize a lista de clientes com os dados obtidos
        this.customers = response.data;
        alert("Customers retrieved successfully");
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
  },
};
</script>

<style>
.form-container {
  width: 30% !important;
}

.table-container {
  width: 80% !important;
}
</style>
