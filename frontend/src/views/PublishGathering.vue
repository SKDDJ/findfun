<template>
  <div class="publish-gathering">
    <h1>发布聚会</h1>
    <p>请填写以下信息，以发布桌面游戏聚会。</p>

    <n-form ref="form" :model="form" :rules="rules" layout="vertical">
      <n-form-item label="活动名称" prop="title">
        <n-input v-model="form.title" />
      </n-form-item>
      <n-form-item label="地点" prop="location">
        <n-input v-model="form.location" />
      </n-form-item>
      <n-form-item label="日期" prop="date">
        <n-date-picker v-model="form.date" type="date" />
      </n-form-item>
      <n-form-item label="时间" prop="time">
        <n-time-picker v-model="form.time" format="HH:mm" />
      </n-form-item>
      <n-form-item label="桌面游戏类型" prop="gameType">
        <n-input v-model="form.gameType" />
      </n-form-item>
      <n-form-item label="最大玩家数量" prop="maxPlayers">
        <n-input-number v-model="form.maxPlayers" />
      </n-form-item>
      <n-form-item>
        <n-button type="primary" @click="submitForm">发布聚会</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script>
export default {
  name: 'PublishGathering',
  data() {
    return {
      form: {
        title: '',
        location: '',
        date: '',
        time: '',
        gameType: '',
        maxPlayers: '',
      },
      rules: {
        title: [{ required: true, message: '请输入活动名称', trigger: 'blur' }],
        location: [{ required: true, message: '请输入地点', trigger: 'blur' }],
        date: [{ required: true, message: '请选择日期', trigger: 'blur' }],
        time: [{ required: true, message: '请选择时间', trigger: 'blur' }],
        gameType: [{ required: true, message: '请输入桌面游戏类型', trigger: 'blur' }],
        maxPlayers: [{ required: true, message: '请输入最大玩家数量', trigger: 'blur' }],
      },
    };
  },
  methods: {
    async submitForm() {
      try {
        await this.$refs.form.validate();
        // Call API to submit the form data
        console.log('Form data:', this.form);
        // After successful submission, navigate to the homepage
        this.$router.push('/');
      } catch (error) {
        console.error('Form validation failed:', error);
      }
    },
  },
};
</script>

<style scoped>
.publish-gathering {
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
