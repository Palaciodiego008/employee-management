<template>
    <div class="py-4">
      <h1 class="text-2xl font-bold mb-4">Employee Hierarchy</h1>
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
        axios.get("http://localhost:8080/hierarchy")
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
  