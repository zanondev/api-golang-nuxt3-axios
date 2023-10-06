<template>
  <v-container align="center" class="d-flex">
    <div class="form-container" aling="center">
      <h2>Register customer</h2>
      <v-form @submit.prevent="registerCustomer">
        <v-text-field
          v-model="customerModel.name"
          label="Name"
          required
        ></v-text-field>
        <v-text-field
          v-model="customerModel.email"
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
          <tr
            class="text-center"
            v-for="customer in customers"
            :key="customer.id"
          >
            <td>{{ customer.id }}</td>
            <td>{{ customer.name }}</td>
            <td>{{ customer.email }}</td>
            <td>
              <v-btn color="secondary" @click="getCustomerToUpdate(customer)"
                >Edit</v-btn
              >
              <v-btn
                color="secondary"
                @click="deleteCustomer(customer)"
                class="ml-2"
                >Delete</v-btn
              >
            </td>
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
      customerModel: {
        name: "",
        email: "",
      },
      customers: [],
      isUpdating: false,
      selectedCustomerId: null,
    };
  },
  methods: {
    async registerCustomer() {
      try {
        await this.$api.post("/Customer/add", this.customerModel);

        this.customerModel.name = "";
        this.customerModel.email = "";

        alert("Customer registered successfully");
        await this.getCustomers();
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    async getCustomers() {
      try {
        const response = await this.$api.get("/Customers");
        this.customers = response.data;
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    getCustomerToUpdate(customer) {
      try {
        this.isUpdating = true;
        this.selectedCustomerId = customer.id;
        this.customerModel.name = customer.name;
        this.customerModel.email = customer.email;
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    async updateCustomer() {
      try {
        this.isUpdating = false;

        const updatedCustomer = {
          id: this.selectedCustomerId,
          name: this.customerModel.name,
          email: this.customerModel.email,
        };

        await this.$api.put("/Customer/update", updatedCustomer);

        this.customerModel.name = "";
        this.customerModel.email = "";

        alert("Customer updated successfully");
        await this.getCustomers();
      } catch (error) {
        console.error("Error:", error);
        alert("Error: " + error);
      }
    },
    async deleteCustomer(customer) {
      try {
        await this.$api.delete(`/Customer/remove/${customer.id}`);

        alert("Customer deleted successfully");
        await this.getCustomers();
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
