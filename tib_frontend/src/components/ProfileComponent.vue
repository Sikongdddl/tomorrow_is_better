<template>
  <div class="profile-container">
    <el-card>
      <h2>个人信息</h2>
      <div v-if="user.avatar_url" class="avatar-wrapper">
        <img :src="fullAvatarURL" alt="用户头像" class="avatar-img" />
      </div>
      <input type="file" @change="handleAvatarUpload" />
      <el-descriptions :column="1" border>
        <el-descriptions-item label="用户名">{{ user.username }}</el-descriptions-item>
        <el-descriptions-item label="违规次数">{{ user.violation_count }}</el-descriptions-item>
        <el-descriptions-item label="成功参与次数">{{ user.success_participation_cnt }}</el-descriptions-item>
        <el-descriptions-item label="注册时间">{{ user.created_at }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

<!--    <el-card style="margin-top: 20px;">-->
<!--      <h2>我参加的活动</h2>-->
<!--      <el-table :data="topics" style="width: 100%">-->
<!--        <el-table-column prop="title" label="活动标题" />-->
<!--        <el-table-column prop="location" label="地点" />-->
<!--        <el-table-column prop="event_time" label="时间" />-->
<!--      </el-table>-->
<!--    </el-card>-->
  </div>
</template>

<script>
import { getUserInfByID, uploadAvatar} from "@/services/userService";

export default {
  data() {
    return {
      user: {},
      userId: null,
      topics: []
    };
  },
  async created() {
    this.userId = parseInt(localStorage.getItem('userId'), 10);
    console.log(this.userId);
    const resUser = await getUserInfByID(this.userId);

    this.user = resUser.data;
    console.log(this.user.avatar_url);
    //
    // const resTopics = await axios.get(`/users/${userID}/topics`);
    // this.topics = resTopics.data.topics;
  },
  computed: {
    fullAvatarURL() {
      // 确保从后端返回的路径拼接为完整 URL
      return this.user.avatar_url
          ? `http://localhost:8080${this.user.avatar_url}`
          : '';
    }
  },
  methods:{
    async handleAvatarUpload(event) {
      const file = event.target.files[0];
      const userID = localStorage.getItem("userId");
      if (!file || !userID) return;

      try {
        const response = await uploadAvatar(userID, file);
        this.user.avatar_url = response.data.avatar_url;
        this.$message.success("头像上传成功！");
        const resUser = await getUserInfByID(this.userId);
        this.user = resUser.data;
      } catch (err) {
        this.$message.error("头像上传失败");
        console.error(err);
      }
    }
  }
}
</script>

<style scoped>
.profile-container {
  padding: 20px;
}

.avatar-wrapper {
  margin-bottom: 10px;
  text-align: center;
}

.avatar-img {
  width: 120px;
  height: 120px;
  border-radius: 50%;
  object-fit: cover;
  border: 2px solid #ccc;
}
</style>
