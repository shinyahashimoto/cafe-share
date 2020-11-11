<template>
  <div class="hello">
    <!-- <span>{{ msg }}</span> -->

    <!-- ジャンボトロン -->
    <div class="jumbotron">
      <div class="container"></div>
    </div>

    <section>
      <div class="row mx-5">
        <div
          v-for="(shop, i) in shops"
          :key="i"
          class="col-sm-3 col-xs-6 p-0 text-left"
        >
          <button @click="toDetailPage(shop.id)">
            <div>
              <b-card>
                <div class="text-left">
                  <img src="../assets/img/スタバ横浜.jpg" class="img-fluid" />
                  <!-- <button :onclick="toDetailPage()">ページせんい</button> -->
                  <div class="mb-3">{{ shop.shopName }}</div>
                  {{ shop.Id }}
                  <div>
                    <span>営業時間：</span>
                    <span>{{ shop.openTime }}</span>
                    〜
                    <span>{{ shop.closeTime }}</span>
                  </div>
                  <!-- 住所 -->
                  <span class="text-muted">{{ shop.address }}</span>
                  <!-- 特徴 -->
                  <div>
                    <span v-for="(imageName, i) in shop.imageNames" :key="i">
                      <img
                        :src="require('../assets/img/' + imageName)"
                        class="img-fluid"
                        width="25"
                        height="25"
                        style="margin-right: 5px;"
                      />
                    </span>
                  </div>
                  <div class="text-right mt-3">
                    <span>
                      <img
                        src="../assets/img/ハート.jpg"
                        class="img-fluid"
                        width="25"
                        height="25"
                      />
                      <span>{{ shop.likeCount }}</span>
                    </span>
                    <span>
                      <img
                        src="../assets/img/星.jpeg"
                        class="img-fluid"
                        width="25"
                        height="25"
                      />
                      <span>{{ shop.favoriteCount }}</span>
                    </span>
                    <span>
                      <img
                        src="../assets/img/コメント.png"
                        class="img-fluid"
                        width="25"
                        height="25"
                      />
                      <span>{{ shop.commentCount }}</span>
                    </span>
                  </div>
                </div>
              </b-card>
            </div>
          </button>
        </div>
      </div>
    </section>

    <!-- ページング -->
    <pagination></pagination>
  </div>
</template>

<script>
import axios from "axios";
import Pagination from "./common/Pagenation";

export default {
  name: "Top",
  data() {
    return {
      msg: "Welcome to Your Vue.js App",
      loginId: 0,
      shop: {
        id: 0,
        shopName: "",
        address: "",
        shopUrl: "",
        shopImage: "",
        img_name: "",
        openTime: "",
        closeTime: "",
        imageNames: [],
        // likeCount: 0,
      },
      shops: [],
      length: 12,
    };
  },
  components: {
    Pagination,
  },
  destoryed() {
    console.log("破壊");
  },
  // beforeCreate: async function() {
  //   const shops = await this.$store.getters.getShopAll;
  //   console.log(shops);
  // },
  beforeCreate: async function() {
    // if (this.shops[0] !== null) {
    //   console.log("aru");
    //   console.log(this.shops).length;
    // } else {
    //   console.log("nai");
    // }
    this.loginId = this.$store.getters.loginInfo.id;

    let res = await axios.get("http://localhost:8000/shopAll");
    await res.data.filter((shop) => {
      shop.openTime = this.formatTime(shop.openTime);
      shop.closeTime = this.formatTime(shop.closeTime);
      // this.$store.commit("setShop", shop);
      this.shops.push(shop);
    });
    console.log(res.data);
    // this.shops = this.$store.getters.getShopAll;
  },
  methods: {
    tomypage() {
      this.$router.push({ name: "MyPage", params: { id: this.loginId } });
    },
    formatTime(time) {
      if (time.indexOf("午前") !== -1) {
        // console.log(str1.indexOf(str2))
        // console.log('存在します')
        const splitTime = time.split(" ");
        return splitTime[0];
      } else {
        var splitTime = time.split(" ");
        splitTime = splitTime[0].split(":");
        splitTime = parseInt(splitTime);
        splitTime = splitTime + 12;
        return splitTime + ":00";
      }
    },
    toDetailPage(shopId) {
      this.$router.push({ name: "ShowShop", params: { id: shopId } });
      // this.$router.push((name, "ShowShop"));
    },
    clickCallback: function(pageNum) {
      console.log(pageNum);
    },
  },
  computed: {
    sharedState() {
      return this.$store.state;
    },
    listStart() {
      return 12 * (this.$store.state.paginationNUmber - 1); // 表示したいデータの最初の値。0、20、40...
    },
  },
};
</script>

<!-- Add 'scoped' attribute to limit CSS to this component only -->
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
button {
  margin: 10px 0;
  padding: 10px;
}
.jumbotron {
  background: url(../assets/img/header.jpeg);
  background-size: cover;
  height: 300px;
  object-position: 0 100%;
}
</style>
