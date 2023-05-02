<template>
  <div class="search-gathering">
    <h1>搜索聚会</h1>
    <p>请填写以下信息，以搜索相关的桌面游戏聚会。</p>

    <n-form ref="form" :model="form" layout="vertical">
      <n-form-item label="地点">
        <n-input v-model="form.location" />
      </n-form-item>
      <n-form-item label="日期">
        <n-date-picker v-model="form.date" type="date" />
      </n-form-item>
      <n-form-item label="桌面游戏类型">
        <n-input v-model="form.gameType" />
      </n-form-item>
      <n-form-item>
        <n-button type="primary" @click="searchGatherings">搜索聚会</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script>
export default {
  name: 'SearchGathering',
  data() {
    return {
      form: {
        location: '',
        date: '',
        gameType: '',
      },
    };
  },
  methods: {
    async searchGatherings() {
      // Call API to search gatherings based on the form data
      console.log('Search criteria:', this.form);
      // Display search results (not implemented in this example)
    },
  },
};
</script>

<style scoped>
.search-gathering {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  width: 100%;
  max-width: 600px;
  margin: 0 auto;
  padding: 2rem;
  box-sizing: border-box;
}

h1 {
  margin-bottom: 1rem;
}

p {
  margin-bottom: 2rem;
}
</style>
