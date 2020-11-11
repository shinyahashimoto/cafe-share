import Vue from "vue";
import Vuex from "vuex";
import createPersistedState from "vuex-persistedstate";
import auth from "@/store/modules/auth";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
  },
  plugins: [
    createPersistedState({
      key: "example",
      paths: ["auth.login"],
      storage: window.sessionStorage,
    }),
  ],
  state: {
    shop: {
      id: 0,
      shopName: "",
      address: "",
      shopUrl: "",
      shopImage: "",
      img_name: "",
      openTime: "",
      closeTime: "",
      impression: "",
      selectedTags: [],
    },
    shops: [],
    paginationNUmber: 1,
  },
  getters: {
    getShopAll(state) {
      return state.shops;
    },
    // getShopById(state) {
    // function(_id) {
    // state.shops.filter((shop) => {
    //   if (shop.id === shopId) {
    //     return shop;
    //   }
    // });
    // }
    // },
  },

  mutations: {
    setShop(state, shopData) {
      state.shops.push(shopData);
    },
  },
  actions() {
    // setShop(commit);
  },
});
