<template class="mb-5">
  <div class="mb-5">
    <!-- <v-app> -->
    <v-app-bar id="page-header" app dark>
      <v-toolbar-title
        ><router-link class="text-white" to="/"
          >Cafe Share</router-link
        ></v-toolbar-title
      >
      <v-spacer></v-spacer>
      <router-link class="text-white" to="/showShop">
        ショップ詳細
      </router-link>

      <router-link v-if="loginId != 0" class="text-white" to="/registerShop"
        >ショップ登録</router-link
      >

      <router-link v-if="loginId == 0" class="text-white" to="/signup"
        >アカウント登録</router-link
      >
      <router-link
        v-if="loginId == 0"
        id="login"
        class="text-white"
        to="/signin"
        >ログイン</router-link
      >
      <button @click="logout">
        ログアウト
      </button>

      <router-link
        v-if="loginId != 0"
        id="login"
        class="text-white"
        v-bind:to="{ name: 'MyPage', params: { id: loginId } }"
        >Mypage</router-link
      >
    </v-app-bar>
    <!-- </v-app> -->
  </div>
</template>

<script>
export default {
  data() {
    return {
      loginId: 0,
    };
  },
  // Created() {
  //   this.loginId = this.$store.getters.loginInfo.id;
  // },
  methods: {
    logout() {
      this.$store.dispatch("removeLoginInfo");
      alert("ログアウトしました。");
      // もしトップにいたら、リダイレクト処理が必要なのかな？
      // this.$router.go({ name: "Top" });
      this.$router.push({ name: "Top" });
    },
  },
  computed: {
    changeLoginInfo() {
      return this.$store.getters.loginInfo.id;
    },
  },
  watch: {
    changeLoginInfo(val, old) {
      this.loginId = val;
      console.log("watch", val, old);
    },
  },
};
</script>
>

<style scoped>
#page-header {
  background-color: #8b4513;
}
#login {
  margin-left: 10px;
}
</style>
