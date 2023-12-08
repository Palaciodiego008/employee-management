<template>
    <div>
      <h1 class="text-2xl font-bold mb-4">Add Employee</h1>
      <form @submit.prevent="addEmployee" class="max-w-md">
        <div class="mb-4">
          <label for="name" class="block text-gray-700 text-sm font-bold mb-2">Name:</label>
          <input v-model="newEmployee.name" type="text" id="name" name="name" class="form-input w-full" required />
        </div>
        <div class="mb-4">
          <label for="managerId" class="block text-gray-700 text-sm font-bold mb-2">Manager ID:</label>
          <input v-model="newEmployee.managerId" type="number" id="managerId" name="managerId" class="form-input w-full" />
        </div>
        <div class="mb-4">
          <label for="version" class="block text-gray-700 text-sm font-bold mb-2">Version:</label>
          <input v-model="newEmployee.version" type="number" id="version" name="version" class="form-input w-full" required />
        </div>
        <button type="submit" class="bg-blue-500 text-white font-bold py-2 px-4 rounded">
          Add Employee
        </button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        newEmployee: {
          name: "",
          managerId: 0,
          version: 1,
        },
      };
    },
    methods: {
      addEmployee() {
        // Send a request to your API to add a new employee
        axios.post("http://localhost:8080/employees", this.newEmployee)
          .then(response => {
            console.log("Employee added successfully:", response.data);
          })
          .catch(error => {
            console.error("Error adding employee:", error);
          });
      },
    },
  };
  </script>
  