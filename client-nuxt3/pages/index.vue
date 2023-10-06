<template>
  <v-container align="center" class="d-flex">
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
        <v-btn type="submit" color="secondary" v-if="!isUpdating">Submit</v-btn>
        <v-btn
          @click="updateCustomer"
          color="secondary"
          class="ml-2"
          v-if="isUpdating"
          >Edit</v-btn
        >
        <v-btn color="secondary" @click="getCustomers" class="ml-2">Get</v-btn>
      </v-form>
    </div>
    <div class="table-container">
      <v-table>
        <thead>
          <tr>
            <th class="text-center">ID</th>
            <th class="text-center">Name</th>
            <th class="text-center">Email</th>
            <th class="text-center">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="customer in customers" :key="customer.id">
            <td>{{ customer.id }}</td>
            <td>{{ customer.name }}</td>
            <td>{{ customer.email }}</td>
            <v-btn color="secondary" @click="getCustomerToUpdate(customer)"
              >Edit</v-btn
            >
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
      isUpdating: false,
    };
  },
  methods: {
    async registerCustomer() {
      try {
        await this.$api.post("/Customer/add", this.newClient);

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
    getCustomerToUpdate(customer) {
      try {
        this.isUpdating = true;

        
        this.newClient.name = customer.name;
        this.newClient.email = customer.email;
        // alert("Customers retrieved successfully");
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    async updateCustomer() {
      try {
        this.isUpdating = false;
        // alert("Customers retrieved successfully");
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
  width: 20em;
  margin-left: 0;
  padding-right: 20px;
}

.table-container {
  width: 50em;
  max-height: 30em;
  overflow-y: auto;
}
</style>
