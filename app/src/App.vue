<template>
  <div id="app">
    <Header />
    <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
    <router-view />
    <Footer />
  </div>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";
import Header from "./components/common/Header";
import Footer from "./components/common/Footer";

export default {
  name: "App",
  components: {
    // HelloWorld,
    Header,
    Footer,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  /* background-image: url("./assets/coffee1.webp"); */
}
</style>
