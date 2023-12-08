<template>
    <div>
      <h1 class="text-2xl font-bold mb-4">Update Manager</h1>
      <form @submit.prevent="updateManager" class="max-w-md">
        <div class="mb-4">
          <label for="employeeId" class="block text-gray-700 text-sm font-bold mb-2">Employee ID:</label>
          <input v-model="employeeId" type="number" id="employeeId" name="employeeId" class="form-input w-full" required />
        </div>
        <div class="mb-4">
          <label for="newManagerId" class="block text-gray-700 text-sm font-bold mb-2">New Manager ID:</label>
          <input v-model="newManagerId" type="number" id="newManagerId" name="newManagerId" class="form-input w-full" required />
        </div>
        <button type="submit" class="bg-green-500 text-white font-bold py-2 px-4 rounded">
          Update Manager
        </button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        employeeId: null,
        newManagerId: null,
      };
    },
    methods: {
      updateManager() {
        // Send a request to your API to update the manager
        axios.put(`http://localhost:8080/api/employees/${this.employeeId}/managers/${this.newManagerId}`)
          .then(response => {
            console.log("Manager updated successfully:", response.data);
            // Optionally, you can redirect or show a success message here
          })
          .catch(error => {
            console.error("Error updating manager:", error);
            // Handle the error (e.g., show an error message)
          });
      },
    },
  };
  </script>
  