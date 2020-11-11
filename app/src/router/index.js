import Vue from "vue";
import Router from "vue-router";
import Top from "@/components/Top";
import Signup from "@/components/Signup";
import Signin from "../components/Signin";
import RegisterShop from "@/components/RegisterShop";
import ShowShop from "@/components/ShowShop";
import MyPage from "@/components/MyPage";
// import NProgress from "nprogress";

import store from "@/store/index.js";

Vue.use(Router);

const router = new Router({
  mode: "history",
  linkActiveClass: "active", // なんのために？
  routes: [
    {
      path: "*",
      redirect: "/login",
    },
    {
      path: "/",
      name: "Top",
      component: Top,
    },
    {
      path: "/signup",
      name: "Signup",
      component: Signup,
    },
    {
      path: "/signin",
      name: "Signin",
      component: Signin,
    },
    {
      path: "/registerShop",
      name: "RegisterShopNew",
      component: RegisterShop,
      meta: { requiresAuth: true },
    },
    {
      path: "/registerShop/:id",
      name: "RegisterShop",
      component: RegisterShop,
      meta: { requiresAuth: true },
    },
    {
      path: "/showShop/:id",
      name: "ShowShop",
      component: ShowShop,
    },
    {
      path: "/myPage/:id",
      name: "MyPage",
      component: MyPage,
    },
  ],
});

router.beforeEach((to, from, next) => {
  // const shops = this.$store.getters.getShopAll;

  let currentUser = store.getters.userName;
  let requiresAuth = to.matched.some((record) => record.meta.requiresAuth);
  if (requiresAuth && !currentUser) next("signin");
  else if (!requiresAuth && currentUser) next();
  else next();
});

export default router;
