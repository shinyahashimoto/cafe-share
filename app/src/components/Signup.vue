<template>
  <div class="signup">
    <b-card class="px-5">
      <h3>アカウント作成</h3>
      <ValidationObserver name="a" v-slot="{ registerUser }">
        <p>ユーザー名</p>
        <ValidationProvider name="名前" rules="required">
          <div slot-scope="ProviderProps">
            <input type="text" class="form-control" v-model="userName" />
            <p class="error">{{ ProviderProps.errors[0] }}</p>
          </div>
        </ValidationProvider>
        <p>メールアドレス</p>
        <ValidationProvider rules="required" v-slot="{ errors }">
          <input type="text" class="form-control" v-model="email" />
          <p class="text-danger">{{ errors[0] }}</p>
        </ValidationProvider>
        <p>パスワード</p>
        <ValidationProvider
          rules="required|confirmed:confirmation"
          v-slot="{ errors }"
        >
          <input type="password" class="form-control" v-model="password" />
          <p class="text-danger">{{ errors[0] }}</p>
        </ValidationProvider>
        <p>パスワード確認用</p>
        <ValidationProvider vid="confirmation" v-slot="{ errors }">
          <input type="password" class="form-control" v-model="confirmation" />
          <p class="text-danger">{{ errors[0] }}</p>
        </ValidationProvider>
        <button
          @click.prevent="register(registerUser)"
          type="button"
          class="btn btn-success px-3"
        >
          登録
        </button>
      </ValidationObserver>
    </b-card>
    <div>{{ passwordMessage }}</div>
    <h1>{{ msg }}</h1>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Signup",
  data() {
    return {
      userName: "",
      email: "",
      password: "",
      confirmation: "",
      passwordMessage: "",
      msg: "",
    };
  },
  methods: {
    register: async function() {
      if (this.password != this.confirmation) {
        this.passwordMessage = "パスワードと確認用パスワードが異なります";
        return;
      }

      let res = await axios.post("http://localhost:8000/users", {
        userName: this.userName,
        email: this.email,
        password: this.password,
      });
      this.msg = res.data;
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
.signup {
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
