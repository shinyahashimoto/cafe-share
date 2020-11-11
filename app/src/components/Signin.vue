<template>
  <div class="signin">
    <b-card class="px-5">
      <h2>ログイン画面</h2>
      <p>メールアドレス</p>
      <input type="text" class="form-control" v-model="email" />
      <p>パスワード</p>
      <input type="password" class="form-control" v-model="password" />

      <button @click="login" type="button" class="btn btn-success px-3">
        ログイン
      </button>
    </b-card>
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Signin",
  data: function() {
    return {
      email: "",
      password: "",
      msg: "",
    };
  },
  methods: {
    login: async function() {
      await axios
        .post("http://localhost:8000/login", {
          email: this.email,
          password: this.password,
        })
        .then((res) => {
          this.$store.dispatch("setLoginInfo", res.data);
          this.$router.push({ name: "Top" });
          // this.$router.push({
          //   path:
          //     "backuri" in this.$route.query &&
          //     this.$route.query.backuri.match(/^\//)
          //       ? this.$route.query.backuri
          //       : "/",
          // });

          this.msg = res.data;
          this.$router.push({ name: "Top" });
        })
        .catch((error) => {
          console.log(error);
          alert("異常が発生しました。");
        });
    },
  },
};
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.signin {
  margin-top: 20px;
  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center;
}
input {
  margin: 10px 0;
  padding: 10px;
}
button {
  margin: 10px 0;
  padding: 10px;
}
</style>
