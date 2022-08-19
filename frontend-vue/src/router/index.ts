import { createRouter, createWebHistory } from "vue-router";
import SummaryView from "../views/SummaryView.vue";
import ArticlesView from "../views/ArticlesView.vue";
import LearnView from "../views/LearnView.vue";
import ReviewView from "../views/ReviewView.vue";
import OnboardingView from "../views/OnboardingView.vue";
import LoginView from "../views/LoginView.vue";
import SignupView from "../views/SignupView.vue";
import { checkLoggedIn } from "../helpers/checkLoggedIn";
import { useLoginStore } from "../stores/loginStore";

async function GoToOnboardingIfNoToken(): Promise<{ path: string }> {
  const loginStore = useLoginStore()
  const isLoggedIn: boolean = await checkLoggedIn();
  if (!isLoggedIn) {
    return { path: "/" };
  } else {
    loginStore.piniaLogIn();
  }
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "OnboardingView",
      component: OnboardingView,
    },
    {
      path: "/login",
      name: "LoginView",
      component: LoginView,
    },
    {
      path: "/test-signup",
      name: "SignupView",
      component: SignupView,
    },
    {
      path: "/articles",
      name: "ArticlesView",
      component: ArticlesView,
      meta: { transition: "slide-left" },
      beforeEnter: [GoToOnboardingIfNoToken],
    },
    {
      path: "/articles/:id/learn",
      name: "LearnView",
      component: LearnView,
      meta: { transition: "slide-left" },
      beforeEnter: [GoToOnboardingIfNoToken],
    },
    {
      path: "/articles/:id/summary",
      name: "SummaryView",
      component: SummaryView,
      meta: { transition: "slide-left" },
      beforeEnter: [GoToOnboardingIfNoToken],
    },
    {
      path: "/review/:days?",
      name: "ReviewView",
      component: ReviewView,
      meta: { transition: "slide-left" },
      beforeEnter: [GoToOnboardingIfNoToken],
    },
  ],
});

export default router;
