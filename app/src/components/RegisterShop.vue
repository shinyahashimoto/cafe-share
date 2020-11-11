<template>
  <div id="register-shop">
    <div class="container">
      <div class="row pt-5">
        <div class="mx-auto">
          <b-card id="form-container" class="px-5">
            <h3 class="font-weight-bold">店舗登録</h3>
            <ValidationObserver name="a" v-slot="{ registerUser }">
              <p class="text-left">店名</p>
              <ValidationProvider name="shopName" rules="required">
                <div slot-scope="ProviderProps">
                  <input
                    type="text"
                    class="form-control"
                    v-model="shop.shopName"
                  />
                  <p class="error">{{ ProviderProps.errors[0] }}</p>
                </div>
              </ValidationProvider>
              <p class="text-left">住所</p>
              <ValidationProvider name="address" rules="required">
                <div slot-scope="ProviderProps">
                  <input
                    type="text"
                    class="form-control"
                    v-model="shop.address"
                  />
                  <p class="error">{{ ProviderProps.errors[0] }}</p>
                </div>
              </ValidationProvider>
              <p class="text-left">店舗URL</p>
              <ValidationProvider
                name="shopUrl"
                rules="required"
                v-slot="{ errors }"
              >
                <input
                  type="text"
                  class="form-control"
                  v-model="shop.shopUrl"
                />
                <p class="text-danger">{{ errors[0] }}</p>
              </ValidationProvider>

              <p class="text-left">店舗画像</p>
              <div class="contents">
                <label v-show="!shopImage" class="input-item__label">
                  画像を選択
                  <input type="file" @change="onFileChange" name="uploadFile" />
                </label>
                <div class="preview-item">
                  <img
                    v-show="shopImage"
                    class="preview-item-file"
                    :src="shopImage"
                    alt
                  />
                  <div
                    v-show="shopImage"
                    class="preview-item-btn"
                    @click="remove"
                  >
                    <p class="preview-item-name">{{ img_name }}</p>
                    <button class="btn btn-success">画像を削除する</button>
                  </div>
                </div>
              </div>
              <p class="text-left">営業時間</p>
              <div class="btn-toolbar row">
                <div class="btn-group col-5">
                  <VueCtkDateTimePicker
                    v-model="shop.openTime"
                    :minute-interval="30"
                    :label="'開始時間'"
                    :format="'hh:mm a'"
                    :formatted="'hh:mm a'"
                    :overlay="true"
                    :only-time="true"
                    :color="'red'"
                    :button-color="'red'"
                  >
                    <button
                      type="button"
                      class="btn btn-outline-secondary time-btn b-100"
                    >
                      Label
                    </button>
                  </VueCtkDateTimePicker>
                </div>
                <span class="d-flex align-items-center col-2">〜</span>
                <div class="btn-group float-right col-5">
                  <VueCtkDateTimePicker
                    v-model="shop.closeTime"
                    :minute-interval="30"
                    :label="'終了時間'"
                    :format="'hh:mm a'"
                    :formatted="'hh:mm a'"
                    :overlay="true"
                    :only-time="true"
                  >
                    <button
                      type="button"
                      class="btn btn-outline-secondary time-btn"
                    >
                      Label
                    </button>
                  </VueCtkDateTimePicker>
                </div>
              </div>

              <p class="text-left">定休日</p>
              <input type="text" class="form-control" v-model="shop.shopUrl" />

              <p class="text-left">タグ</p>

              <section>
                <span v-for="(tag, i) in tags" :key="i">
                  <input
                    :id="'tag' + i"
                    type="checkbox"
                    :value="tag.id"
                    v-model="shop.selectedTags"
                    data-toggle="buttons"
                    class="tag-input"
                  />
                  <label :for="'tag' + i" class="mb-3">{{ tag.name }}</label>
                </span>
              </section>

              {{ selectedTags }}

              <p class="text-left">使用してみての感想</p>
              <textarea
                id="impression-container"
                v-model="shop.impression"
                placeholder="add multiple lines"
              ></textarea>
              <div></div>

              <button
                @click.prevent="register(registerUser)"
                type="button"
                class="btn btn-success px-3"
              >
                登録
              </button>
            </ValidationObserver>
            {{ msg }}
          </b-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Signup",
  components: {},
  data() {
    return {
      shop: {
        id: 0,
        shopName: "店舗",
        address: "住所",
        shopUrl: "店舗URL",
        shopImage: "",
        uploadImage: {},
        img_name: "",
        openTime: "04:00 午前",
        closeTime: "05:00 午後",
        impression: "感想",
        selectedTags: [],
      },
      tags: [],
      msg: "",
    };
  },
  beforeCreate: async function() {
    // console.log(this.tags);
    const shopId = parseInt(this.$route.params.id, 10);
    // if(shopId)
    if (!isNaN(shopId)) {
      let res = await axios.get("http://localhost:8000/shop", {
        params: {
          id: shopId,
        },
      });
      this.shop = res.data;
      console.log(res.data);
    }
    // this.shop.selectedTags = [];
    let res = await axios.get("http://localhost:8000/shopA");
    this.tags = res.data;
    console.log(this.shop.id);
  },
  methods: {
    register: async function() {
      console.log(this.shop.uploadImage);

      var param = {
        id: this.shop.id,
        shopName: this.shop.shopName,
        address: this.shop.address,
        shopUrl: this.shop.shopUrl,
        shopImage: this.shop.shopImage,
        openTime: this.shop.openTime,
        closeTime: this.shop.closeTime,
        selectedTags: this.shop.selectedTags,
        impression: this.shop.impression,
        headers: {
          // multipartで送信
          "content-type": "multipart/form-data",
        },
      };
      console.log("パラメーター");
      console.log(param);

      let res = await axios
        .post("http://localhost:8000/shop", param)
        .catch((error) => {
          console.log(error.response.data.Detail);
        });
      alert("登録完了");
      this.$router.push({ name: "Top" });
      console.log(res.data);
    },
    onFileChange: async function(e) {
      const files = e.target.files || e.dataTransfer.files;
      // this.uploadImage = files[0];

      const reader = new FileReader();
      reader.readAsDataURL(files[0]);

      this.createImage(files[0]);
      this.shop.img_name = files[0].name;
      console.log(typeof this.shop.shopImage);
    },
    // アップロードした画像を表示
    createImage(file) {
      const reader = new FileReader();
      reader.onload = (e) => {
        this.shop.shopImage = e.target.result;
      };

      // reader.readAsDataURL(file);
      console.log(reader.readAsDataURL(file));
    },
    remove() {
      this.shop.shopImage = false;
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
#register-shop {
  background-color: #f5f5f5;
}
.time-btn {
  width: 140px;
}

label {
  background-color: #f5f5f5;
  border: none;
  outline: none;
  border-radius: 40px;
}
label:hover {
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.6);
}

input:checked + label {
  background-color: #8b4513;
  color: white;
}
.tag-input {
  display: none;
}
#form-container {
  width: 500px;
}
#impression-container {
  border: 2px solid #0a0; /* 枠線 */
  border-radius: 0.67em; /* 角丸 */
  padding: 0.5em; /* 内側の余白量 */
  background-color: snow; /* 背景色 */
  width: 23em; /* 横幅 */
  height: 120px; /* 高さ */
  font-size: 1em; /* 文字サイズ */
  line-height: 1.2; /* 行の高さ */
}
.form-control {
  margin-top: 0;
  margin-bottom: 10px;
}
p {
  margin: 0;
}
</style>
