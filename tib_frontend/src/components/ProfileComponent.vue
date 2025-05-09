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
        <el-descriptions-item label="成功参与次数">{{ user.success_participation }}</el-descriptions-item>
        <el-descriptions-item label="注册时间">{{ user.created_at }}</el-descriptions-item>
      </el-descriptions>
    </el-card>

    <el-card style="margin-top: 20px;">
      <h2>我提出的活动</h2>
      <el-table :data="createdTopics" style="width: 100%">
        <el-table-column prop="Title" label="活动标题" />
        <el-table-column prop="Location" label="地点" />
        <el-table-column prop="EventTime" label="时间" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button type="danger" @click="leave(scope.row.ID)">Leave</el-button>
          </template>
        </el-table-column>
        <el-button type="danger" @click="leave(data.ID)">Leave</el-button>
      </el-table>
    </el-card>

    <el-card style="margin-top: 20px;">
      <h2>我参加的活动</h2>
      <el-table :data="participatedTopics" style="width: 100%">
        <el-table-column prop="Title" label="活动标题" />
        <el-table-column prop="Location" label="地点" />
        <el-table-column prop="EventTime" label="时间" />
        <el-table-column label="操作">
          <template #default="scope">
            <el-button type="danger" @click="leave(scope.row.ID)">Leave</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
</template>

<script>
import { getUserInfByID, uploadAvatar, getCreatedTopics, getParticipatedTopics } from "@/services/userService";
import {leaveTopic} from "@/services/topicService"
export default {
  data() {
    return {
      user: {},
      userId: null,
      createdTopics: [],
      participatedTopics: [],
    };
  },
  async created() {
    this.userId = parseInt(localStorage.getItem('userId'), 10);
    console.log(this.userId);
    const resUser = await getUserInfByID(this.userId);
    this.user = resUser.data;
    await this.loadTopics();
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
    async loadTopics() {
      const resCreated = await getCreatedTopics(this.userId);
      this.createdTopics = resCreated.data;

      const resJoined = await getParticipatedTopics(this.userId);
      this.participatedTopics = resJoined.data;
    },

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
    },
    async leave(topicId){
      try{
        await leaveTopic(this.userId, topicId);
        alert('Leave successfully');
        this.loadTopics(this.userId);
      } catch (e){
        alert('Leave failed: '+ (e.response?.data?.error || 'Unknown error'));
      }
    },
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
