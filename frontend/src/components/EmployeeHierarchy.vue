<template>
  <div class="py-4">
    <h1 class="text-2xl font-bold mb-4">Employee Hierarchy</h1>
    <div class="mb-4">
      <router-link to="/add-employee" class="bg-blue-500 text-white py-2 px-4 rounded">
        Add Employee
      </router-link>
      <router-link to="/update-manager" class="bg-green-500 text-white py-2 px-4 ml-4 rounded">
        Update Manager
      </router-link>
    </div>
    <ul class="pl-4">
      <li v-for="employee in hierarchy" :key="employee.employee.id" class="mb-2">
        <span class="font-semibold">{{ employee.employee.name }}</span>
        <employee-hierarchy :hierarchy="employee.subordinates" v-if="employee.subordinates" />
      </li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      hierarchy: [],
    };
  },
  mounted() {
    this.getHierarchy();
  },
  methods: {
    getHierarchy() {
      axios.get("http://localhost:8080/api/employees/hierarchy")
        .then(response => {
          this.hierarchy = response.data;
        })
        .catch(error => {
          console.error("Error fetching hierarchy:", error);
        });
    },
  },
};
</script>
