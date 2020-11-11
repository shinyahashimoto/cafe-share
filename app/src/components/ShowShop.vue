<template>
  <div class="register-shop">
    <b-card class="px-5">
      <!-- <div id="map">
        <GmapMap
          :center="center"
          :zoom="zoom"
          style="width: 100%; height: 100%;"
        >
          <GmapMarker
            v-for="(m, id) in marker_items"
            :position="m.position"
            :title="m.title"
            :key="id"
            :icon="m.icon"
            :clickable="true"
            :draggable="true"
          >
          </GmapMarker>
        </GmapMap>
      </div> -->
      <div class="container-fluid">
        <div class="row">
          <div class="col-7 text-left">
            <img src="../assets/img/スタバ横浜.jpg" class="img-fluid" />

            <span
              v-if="loginId !== 0"
              class="heart"
              v-on:click="likeBtn()"
              v-bind:class="{ active: isActive01 }"
            ></span>

            <span class="circle">
              <img
                v-if="loginId !== 0"
                v-on:click="favoriteBtn()"
                :src="require('../assets/img/' + favoriteImg + '.png')"
                width="30px"
              />
            </span>
            <i class="far fa-thumbs-up like-btn"></i>

            <!-- ショップ名 -->
            <div class="h3 font-weight-bold mt-5">
              {{ shop.shopName }}
              <div v-if="shop.userId === loginId">
                <!-- <button @click="toUpdatePage()"></button> -->
                <router-link
                  :to="{ name: 'RegisterShop', params: { id: shop.id } }"
                  >編集</router-link
                >
              </div>
              <div v-if="shop.userId === loginId">
                <button @click="deleteShop()">削除</button>
              </div>
            </div>
            <div>
              <span>
                <span class="heart active"></span>
                <span id="like-count" class="ml-2">
                  {{ shop.likeCount }}
                </span>
                <img
                  :src="require('../assets/img/星アイコン6.png')"
                  width="30px"
                />
                <span id="favorite-count">
                  {{ shop.favoriteCount }}
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
              </span>
            </div>
            <div>
              <div>投稿者の感想</div>
              {{ shop.impression }}
            </div>
            <div v-if="shop.comments.length">
              <div v-for="(target, i) in shop.comments" :key="i">
                <div>
                  {{ target.userName }}さんが{{ target.createdAt }}日にコメント
                </div>
                <div>
                  {{ target.content }}
                </div>
              </div>
            </div>
            <div>
              <div>コメント</div>
              <div v-if="loginId !== 0">
                <textarea
                  class="comment-container"
                  v-model="comment"
                  placeholder="コメント"
                ></textarea>
                <button @click="postComment()">投稿</button>
              </div>
              <div v-else>ログインしたらコメントができます。</div>
            </div>
            <div>
              <div class="font-weight-bold">基本情報</div>
              <div id="shop-info" class="row">
                <div class="col-3 info-border">店名</div>
                <div class="col-9">
                  {{ shop.shopName }}
                </div>
                <div class="col-3">住所</div>
                <div class="col-9">
                  {{ shop.address }}
                </div>
                <div class="col-3">営業時間</div>
                <div class="col-9">
                  {{ shop.openTime }}~ {{ shop.closeTime }}
                </div>
                <div class="col-3">URL</div>
                <div class="col-9">
                  {{ shop.shopUrl }}
                </div>
                <div class="col-3">定休日</div>
                <div class="col-9">ー</div>
                <div class="col-3">特徴</div>
                <div class="col-9">
                  <section>
                    <span v-for="(tag, i) in shop.selectedTagNames" :key="i">
                      <label :for="'tag' + i" class="mb-3">
                        {{ tag }}
                      </label>
                    </span>
                  </section>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      {{ shop.selectedTags }}

      {{ shop.id }}
    </b-card>

    <span>ShopId : </span>
    {{ $route.params.id }}

    <img
      v-show="shop.shopImage"
      class="preview-item-file"
      :src="shop.shopImage"
      alt
    />
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Signup",
  components: {},
  data() {
    return {
      loginId: 0,
      shop: {
        id: "",
        shopName: "",
        address: "",
        shopUrl: "",
        shopImage: "",
        openTime: "",
        closeTime: "",
        impression: "",
        userId: "",
        selectedTags: [],
        selectedTagNames: [],
        comments: [],
        like: {
          id: 0,
          isDeleted: false,
        },
        favorite: {
          id: 0,
          isDeleted: false,
        },
        likeCount: 0,
        favoriteCount: 0,
      },
      isActive01: false,
      isActive02: false,

      comment: "",
      // center: { lat: 35.698304, lng: 139.766325 },
      // zoom: 18,
      // mapStyle: {
      //   disableDefaultUI: true, // 表示のオプションを指定できます。
      //   styles: [
      //     // カスタマイズで使用したスタイルなどはここに。
      //   ],
      // },
      // marker_items: [
      //   { position: { lat: 35.443674, lng: 139.637964 }, title: "title" },
      // ],
      favoriteImg: null,
    };
  },
  mounted() {
    let recaptchaScript = document.createElement("script");
    recaptchaScript.setAttribute(
      "src",
      "https://kit.fontawesome.com/03d83596d9.js"
    );
    document.head.appendChild(recaptchaScript);
  },
  watch: {
    $route() {
      console.log("routeが変わった");
    },
  },
  computed: {
    ShopData() {
      console.log("findbyid");
      const shopId = parseInt(this.$route.params.id, 10);
      const data = this.$store.getters.getShopById(shopId);
      return data;
    },
  },
  created: async function() {
    const shopId = parseInt(this.$route.params.id, 10);
    // const data = this.$store.getters.getShopById;
    // ショップ詳細画面でリロード対策
    this.loginId = this.$store.getters.loginInfo.id;
    let res = await axios.get("http://localhost:8000/shop", {
      params: {
        id: shopId,
        userId: this.loginId,
      },
    });

    // いいねの存在
    if (res.data.like.id === 0) {
      res.data.like.isDeleted = false;
    } else if (!res.data.like.isDeleted) {
      this.isActive01 = true;
    }
    this.shop = res.data;

    console.log("favorite");
    console.log(res.data.favorite.id);

    //お気に入りの存在
    if (!res.data.favorite.isDeleted) {
      this.isActive02 = true;
    }

    this.changeFavoriteImg();

    this.shop.openTime = this.formatTime(this.shop.openTime);
    this.shop.closeTime = this.formatTime(this.shop.closeTime);
    this.shop.comments.forEach((comment) => {
      comment.createdAt = comment.createdAt.slice(0, 10);
    });

    console.log("favorite_isdeleted");
    console.log(this.shop.favorite.isDeleted);

    // 応急処置
    // this.shop.favorite.id = 0;
    // this.shop.favorite.isDeleted = false;
    // this.shop.favoriteCount = 0;
  },
  methods: {
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
    likeBtn: async function() {
      if (this.isActive01) {
        this.shop.likeCount--;
      } else {
        this.shop.likeCount++;
      }
      this.isActive01 = !this.isActive01;
      var loginInfo = this.$store.getters.loginInfo;
      var param = {
        id: this.shop.like.id,
        shopId: this.shop.id,
        userId: loginInfo.id,
        isDeleted: this.shop.like.isDeleted,
      };
      console.log(param);
      let res = await axios.post("http://localhost:8000/like", param);
      this.shop.like.id = res.data.id;
      this.shop.like.isDeleted = res.data.isDeleted;
      console.log(res);
    },
    favoriteBtn: async function() {
      if (this.loginId === 0) {
        // return文言を追加
        return;
      }

      var param = {
        id: this.shop.favorite.id,
        shopId: this.shop.id,
        userId: this.loginId,
        isDeleted: this.shop.favorite.isDeleted,
      };
      console.log(param);
      let res = await axios.post("http://localhost:8000/favorite", param);
      this.shop.favorite.id = res.data.id;
      this.shop.favorite.isDeleted = res.data.isDeleted;
      if (this.isActive02) {
        this.shop.favoriteCount--;
      } else {
        this.shop.favoriteCount++;
      }
      this.isActive02 = !this.isActive02;
      this.changeFavoriteImg();
    },
    changeFavoriteImg() {
      if (this.isActive02) {
        this.favoriteImg = "星アイコン6";
      } else {
        this.favoriteImg = "星アイコン8";
      }
    },
    postComment: async function() {
      var loginInfo = this.$store.getters.loginInfo;
      var param = {
        shopId: this.shop.id,
        userId: loginInfo.id,
        content: this.comment,
      };
      await axios
        .post("http://localhost:8000/comment", param)
        .then((res) => {
          var comment = {
            content: res.data.content,
            createdAt: res.data.createdAt.slice(0, 10),
          };
          this.shop.comments.push(comment);
          this.comment = "";
        })
        .catch((error) => {
          console.log(error.response.data.Detail);
        });
    },
    deleteShop: async function() {
      await axios
        .delete("http://localhost:8000/shop", {
          params: {
            shopId: this.shop.id,
          },
        })
        .then(this.$router.push({ name: "Top" }))
        .catch((err) => {
          console.log(err);
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
label > input {
  display: none;
}

label {
  padding: 0 1rem;
  border: solid 1px #888;
}

.heart {
  display: inline-block;
  color: #333; /*線の色*/
  position: relative;
  margin-top: 6px; /*要調整*/
  margin-left: 5px; /*要調整*/
  width: 9px;
  height: 9px;
  border-left: solid 1px currentColor;
  border-bottom: solid 1px currentColor;
  transform: rotate(-45deg) scale(2); /*scale要調整*/
  background: white; /*背景色*/
}
.heart:before {
  content: "";
  position: absolute;
  top: -5px;
  left: -1px;
  width: 10px;
  height: 5px;
  border-radius: 5px 5px 0 0;
  border-top: solid 1px currentColor;
  border-left: solid 1px currentColor;
  border-right: solid 1px currentColor;
  background: white; /*背景色*/
}
.heart:after {
  content: "";
  position: absolute;
  top: -1px;
  left: 8px;
  width: 5px;
  height: 10px;
  border-radius: 0 5px 5px 0;
  border-top: solid 1px currentColor;
  border-right: solid 1px currentColor;
  border-bottom: solid 1px currentColor;
  background: white; /*背景色*/
}
.active,
.active:after,
.active::before {
  background: red;
}

#footer {
  height: 1000px;
}

#shop-info > div {
  border-bottom: solid 2px #8b4513;
}

.comment-container {
  border: 2px solid #0a0; /* 枠線 */
  border-radius: 0.67em; /* 角丸 */
  padding: 0.5em; /* 内側の余白量 */
  background-color: snow; /* 背景色 */
  width: 45em; /* 横幅 */
  height: 120px; /* 高さ */
  font-size: 1em; /* 文字サイズ */
  line-height: 1.2; /* 行の高さ */
}
.circle {
  width: 50px; /*幅*/
  height: 50px; /*高さ*/
  border-radius: 50%; /*角丸*/
  background-color: red;
}
</style>
