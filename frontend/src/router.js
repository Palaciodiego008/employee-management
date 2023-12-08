import { createRouter, createWebHistory } from "vue-router";


const routes = [
  {
    path: "/",
    name: "employeeHierarchy",
    component: () => import("./components/EmployeeHierarchy.vue"),
  },
  {
    path: "/add-employee",
    name: "addEmployee",
    component: () => import("./components/AddEmployee.vue"),
  },
  {
    path: "/update-manager",
    name: "updateManager",
    component: () => import("./components/UpdateManager.vue"),
  },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
  });

export default router;
