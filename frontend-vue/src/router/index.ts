import { createRouter, createWebHistory } from "vue-router";
import LoginView from "../views/dev/graphql/LoginView.vue";
import SignupView from "../views/dev/graphql/SignupView.vue";
import SeenArticlesView from "../views/dev/graphql/SeenArticlesView.vue";
import UnseenArticlesView from "../views/dev/graphql/UnseenArticlesView.vue";
import SeenSentsView from "../views/dev/graphql/SeenSentsView.vue";
import UnseenSentsView from "../views/dev/graphql/UnseenSentsView.vue";
import ExamineCorrectView from "../views/dev/graphql/ExamineCorrectView.vue";
import ExamineIncorrectView from "../views/dev/graphql/ExamineIncorrectView.vue";
import ReviewDueView from "../views/dev/graphql/ReviewDueView.vue";
import ListComponents from "../views/dev/ListComponents.vue";
import ListSentences from "../views/dev/ListSentences.vue";
import SummaryView from "../views/SummaryView.vue";
import ArticlesView from "../views/ArticlesView.vue";
import LearnView from "../views/LearnView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "ListSentences",
      component: ListSentences,
      meta: { transition: "slide-left" },
    },
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
      props: true,
    },
    {
      path: "/articles/:id/summary",
      name: "SummaryView",
      component: SummaryView,
      meta: { transition: "slide-left" },
    },
    {
      path: "/list-components",
      name: "ListComponents",
      component: ListComponents,
    },
    {
      path: "/dev/graphql",
      name: "LoginView",
      component: LoginView,
    },
    {
      path: "/dev/graphql/signup",
      name: "SignupView",
      component: SignupView,
    },
    {
      path: "/dev/graphql/articles/seen",
      name: "SeenArticlesView",
      component: SeenArticlesView,
    },
    {
      path: "/dev/graphql/articles/unseen",
      name: "UnseenArticlesView",
      component: UnseenArticlesView,
    },
    {
      path: "/dev/graphql/sents/seen",
      name: "SeenSentsView",
      component: SeenSentsView,
    },
    {
      path: "/dev/graphql/sents/unseen",
      name: "UnseenSentsView",
      component: UnseenSentsView,
    },
    {
      path: "/dev/graphql/examine/correct",
      name: "ExamineCorrectView",
      component: ExamineCorrectView,
    },
    {
      path: "/dev/graphql/examine/incorrect",
      name: "ExamineIncorrectView",
      component: ExamineIncorrectView,
    },
    {
      path: "/dev/graphql/review/due",
      name: "ReviewDue",
      component: ReviewDueView,
    },
  ],
});

export default router;
