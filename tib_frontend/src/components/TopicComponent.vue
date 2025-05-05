<template>
  <div class="topics-container">
    <el-card class="topic-card" v-for="topic in topics" :key="topic.ID">
      <h3>{{ topic.Title }}</h3>
      <p>Event Time: {{ new Date(topic.EventTime).toLocaleString() }}</p>
      <p>Location: {{ topic.Location }}</p>
      <p v-for = "participant in topic.Participants" :key="participant.ID">
        UserID: {{participant.UserID}}
        Role: {{participant.Role}}
      </p>
      <el-button type="success" @click="join(topic.ID)">Join</el-button>
    </el-card>

    <el-divider></el-divider>

    <el-card class="create-topic">
      <h2>Create New Topic</h2>
      <el-form :model="newTopic" label-width="100px">
        <el-form-item label="Title">
          <el-input v-model="newTopic.title" />
        </el-form-item>
        <el-form-item label="Event Time">
          <el-date-picker v-model="newTopic.event_time" type="datetime" />
        </el-form-item>
        <el-form-item label="Location">
          <el-input v-model="newTopic.location" />
        </el-form-item>
        <el-button type="primary" @click="create">Create</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import { fetchTopics, createTopic, joinTopic } from '@/services/topicService';

export default {
  data() {
    return {
      topics: [],
      newTopic: {
        title: '',
        event_time: '',
        location: ''
      },
      userId: null // 从 localStorage 获取用户 ID
    };
  },
  created() {
    this.userId = parseInt(localStorage.getItem('userId'), 10);
    this.loadTopics();
  },
  methods: {
    async loadTopics() {
      try {
        const res = await fetchTopics();
        this.topics = res.data.topics;
        console.log(this.topics[0]);
      } catch (e) {
        console.error('Failed to load topics:', e);
      }
    },
    async create() {
      try {
        await createTopic({
          ...this.newTopic,
          creator_id: this.userId
        });
        this.loadTopics();
      } catch (e) {
        alert('Create failed: ' + (e.response?.data?.error || 'Unknown error'));
      }
    },
    async join(topicId) {
      try {
        console.log(topicId)
        console.log(this.userId)
        await joinTopic(this.userId,topicId);
        alert('Joined successfully');
        this.loadTopics();
      } catch (e) {
        alert('Join failed: ' + (e.response?.data?.error || 'Unknown error'));
      }
    }
  }
};
</script>

<style scoped>
.topics-container {
  padding: 20px;
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 20px;
}

.topic-card {
  padding: 20px;
}

.create-topic {
  grid-column: 1 / -1;
  margin-top: 40px;
}
</style>
