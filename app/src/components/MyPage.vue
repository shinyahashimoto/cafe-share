<template>
  <div>
    <section>
      <div class="row mx-5">
        <div v-for="(shop, i) in shops" :key="i" class="col-3 p-0 text-left">
          <button @click="toDetailPage(shop.id)">
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
                <span class="text-muted">{{ shop.address }}</span>
                <div>
                  <img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  />
                  <img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  />
                  <img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  /><img
                    src="../assets/img/電源.jpg"
                    class="img-fluid"
                    width="25"
                    height="25"
                  />
                </div>
                <div class="text-right mt-3">
                  <span>
                    <img
                      src="../assets/img/ハート.jpg"
                      class="img-fluid"
                      width="25"
                      height="25"
                    />
                    <span>0</span>
                  </span>
                  <span>
                    <img
                      src="../assets/img/星.jpeg"
                      class="img-fluid"
                      width="25"
                      height="25"
                    />
                    <span>0</span>
                  </span>
                  <span>
                    <img
                      src="../assets/img/コメント.png"
                      class="img-fluid"
                      width="25"
                      height="25"
                    />
                  </span>
                </div>
              </div>
            </b-card>
          </button>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "MyPage",
  data() {
    return {
      shop: {
        id: 0,
        shopName: "",
        address: "",
        shopUrl: "",
        shopImage: "",
        img_name: "",
        openTime: "",
        closeTime: "",
        selectedTags: [],
      },
      shops: [],
    };
  },
  beforeCreate: async function() {},
  created: async function() {
    var loginId = this.$store.getters.loginInfo.id;
    let res = await axios.get("http://localhost:8000/myPage", {
      params: {
        userId: loginId,
      },
    });

    console.log(res);
    console.log(res.data);
    await res.data.filter((shop) => {
      shop.openTime = this.formatTime(shop.openTime);
      shop.closeTime = this.formatTime(shop.closeTime);
      this.shops.push(shop);
    });
  },
  methods: {
    formatTime(time) {
      if (time.indexOf("午前") !== -1) {
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
