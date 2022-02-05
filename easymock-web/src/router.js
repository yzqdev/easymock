import { createRouter, createWebHistory } from "vue-router";

import docs from "@/view/docs/index.vue";
import login from "@/view/Login.vue";
import group from "@/view/Group.vue";
import editor from "@/view/editor/index.vue";
import project from "@/view/Project.vue";
import profile from "@/view/Profile.vue";
import createProject from "@/view/new/index.vue";
import logOut from "@/components/log-out.vue";
import dashboard from "@/view/Dashboard.vue";
import detail from "@/view/project-detail/index.vue";
import layout from "@comp/Layout.vue";

const routes = [
  { path: "/login", component: login },
  { path: "/log-out", component: logOut },
  {
    path: "/",
    component: layout,
    children: [
      { path: "/", component: project },
      { path: "workbench", component: project },
      { path: "group/:id", component: project },
      { path: "group", component: group },
      { path: "docs", component: docs },
      { path: "changelog", component: docs },
      { path: "dashboard", component: dashboard },
      { path: "profile", component: profile },
      { path: "new", component: createProject },
      { path: "project/:id", component: detail },
      { path: "editor/:projectId", component: editor },
      { path: "editor/:projectId/:id", component: editor },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});
export default router;
