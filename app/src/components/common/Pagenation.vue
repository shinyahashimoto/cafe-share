<template>
  <paginate
    v-model="page"
    :page-count="paginationNumber"
    :page-range="12"
    :margin-pages="0"
    :click-handler="clickCallback"
    :prev-text="'Prev'"
    :next-text="'Next'"
    :container-class="'pagination'"
    :page-class="'page-item'"
  >
  </paginate>
</template>

<script>
import Vue from "vue";
import Paginate from "vuejs-paginate";
Vue.component("paginate", Paginate);

export default {
  data() {
    return {
      length: 12, // 1ページに表示するListの数
      page: this.$store.state.paginationNUmber, // storeからpage番号を取得
    };
  },
  methods: {
    clickCallback(pageNum) {
      this.page = pageNum; // pageNumはpagerの何番目をclickしたかを取得
      this.$store.state.paginationNUmber = pageNum; // 何番目をclickしたかをstoreの値にset
    },
  },
  computed: {
    paginationNumber() {
      return this.$store.state.hoge.length / 12; // pagerの数を設定するために全データの数を12で割っている
    },
  },
};
</script>
