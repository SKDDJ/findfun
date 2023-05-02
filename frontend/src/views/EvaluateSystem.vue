<template>
  <div class="evaluate-system">
    <h1>评价系统</h1>
    <p>请填写以下信息，以提交您对桌面游戏聚会的评价。</p>

    <n-form ref="form" :model="form" layout="vertical">
      <n-form-item label="姓名">
        <n-input v-model="form.name" />
      </n-form-item>
      <n-form-item label="电子邮件">
        <n-input v-model="form.email" />
      </n-form-item>
      <n-form-item label="电话号码">
        <n-input v-model="form.phoneNumber" />
      </n-form-item>
      <n-form-item label="聚会评价">
        <n-input type="textarea" v-model="form.evaluation" />
      </n-form-item>
      <n-form-item label="评分 (1-5 星)">
        <n-rate v-model="form.rating" />
      </n-form-item>
      <n-form-item>
        <n-button type="primary" @click="submitEvaluation">提交评价</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script>
export default {
  name: 'EvaluateSystem',
  data() {
    return {
      form: {
        name: '',
        email: '',
        phoneNumber: '',
        evaluation: '',
        rating: 0,
      },
    };
  },
  methods: {
    async submitEvaluation() {
      // Call API to submit the evaluation based on the form data
      console.log('Evaluation data:', this.form);
      // Display a success message or handle errors (not implemented in this example)
    },
  },
};
</script>

<style scoped>
.evaluate-system {
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
