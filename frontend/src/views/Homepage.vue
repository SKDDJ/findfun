<template>
  <div class="homepage">
    <h1>Findfun桌游聚会平台</h1>
    <p>欢迎来到Findfun桌游聚会平台！在这里，你可以发布、搜索、匹配和评价桌游聚会。</p>

    <div class="buttons">
      <n-button type="primary" @click="goToPublishGathering">发布聚会</n-button>
      <n-button type="primary" @click="goToSearchGathering">搜索聚会</n-button>
      <n-button type="primary" @click="goToMatchGathering">匹配聚会</n-button>
      <n-button type="primary" @click="goToEvaluateSystem">评价系统</n-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Homepage',
  methods: {
    goToPublishGathering() {
      this.$router.push('/publish-gathering');
    },
    goToSearchGathering() {
      this.$router.push('/search-gathering');
    },
    goToMatchGathering() {
      this.$router.push('/match-gathering');
    },
    goToEvaluateSystem() {
      this.$router.push('/evaluate-system');
    },
  },
};
</script>

<style scoped>
.homepage {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  text-align: center;
}

.buttons {
  display: flex;
  gap: 20px;
  margin-top: 20px;
}
</style>
