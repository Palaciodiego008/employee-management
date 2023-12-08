// router.js
import Vue from "vue";
import VueRouter from "vue-router";
import EmployeeHierarchy from "@/components/EmployeeHierarchy.vue";
import AddEmployee from "@/components/AddEmployee.vue";
import UpdateManager from "@/components/UpdateManager.vue";

Vue.use(VueRouter);

const routes = [
  { path: "/", component: EmployeeHierarchy },
  { path: "/add-employee", component: AddEmployee },
  { path: "/update-manager", component: UpdateManager },
];

const router = new VueRouter({
  routes,
});

export default router;
