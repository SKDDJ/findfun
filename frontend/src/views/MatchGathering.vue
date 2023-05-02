<template>
  <div class="match-gathering">
    <h1>匹配聚会</h1>
    <p>请选择您喜欢的桌面游戏类型和地点，然后开始匹配聚会。</p>

    <n-form ref="form" :model="form" layout="vertical">
      <n-form-item label="桌面游戏类型">
        <n-input v-model="form.gameType" />
      </n-form-item>
      <n-form-item label="地点">
        <n-input v-model="form.location" />
      </n-form-item>
      <n-form-item>
        <n-button type="primary" @click="startMatching">开始匹配</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script>
export default {
  name: 'MatchGathering',
  data() {
    return {
      form: {
        gameType: '',
        location: '',
      },
    };
  },
  methods: {
    async startMatching() {
      // Call API to match gatherings based on the form data
      console.log('Matching criteria:', this.form);
      // Display matching results (not implemented in this example)
    },
  },
};
</script>

<style scoped>
.match-gathering {
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
