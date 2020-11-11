const auth = {
  state: {
    login: {
      id: 0,
      userName: "ゲスト",
      token: false,
      // expire: 0,
    },
  },
  getters: {
    userName: (state) => state.login.userName,
    loginInfo: (state) => state.login,
  },
  mutations: {
    SET_LOGIN_INFO: (state, login) => {
      state.login.id = login.id;
      state.login.userName = login.userName; // ユーザー名
      state.login.token = login.token; // ログイントークン
      // state.login.expire = Math.floor(1000 * login.expire); // APIからUNIXタイム(秒)で有効期限が返ってくるものとし、ミリ秒に変換しておく
    },
    REMOVE_LOGIN_INFO: (state) => {
      state.login.id = 0;
      state.login.userName = "";
      state.login.token = false;
    },
  },
  actions: {
    setLoginInfo({ commit }, login) {
      commit("SET_LOGIN_INFO", login);
    },
    removeLoginInfo({ commit }) {
      commit("REMOVE_LOGIN_INFO");
    },
  },
};

export default auth;
