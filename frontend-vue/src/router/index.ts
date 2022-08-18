import { createRouter, createWebHistory } from "vue-router";
import SummaryView from "../views/SummaryView.vue";
import ArticlesView from "../views/ArticlesView.vue";
import LearnView from "../views/LearnView.vue";
import ReviewView from "../views/ReviewView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/articles",
      name: "ArticlesView",
      component: ArticlesView,
      meta: { transition: "slide-left" },
    },
    {
      path: "/articles/:id/learn",
      name: "LearnView",
      component: LearnView,
      meta: { transition: "slide-left" },
    },
    {
      path: "/articles/:id/summary",
      name: "SummaryView",
      component: SummaryView,
      meta: { transition: "slide-left" },
    },
    {
      path: "/review/:days?",
      name: "ReviewView",
      component: ReviewView,
      meta: { transition: "slide-left" },
    },
  ],
});

export default router;
