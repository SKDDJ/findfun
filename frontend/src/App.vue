<template>
  <div id="app">
    <n-config-provider>
      <router-view></router-view>
    </n-config-provider>
  </div>
</template>

<script>
export default {
  name: 'App',
};
</script>

<style>
/* Add any global styles for your application here */
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
