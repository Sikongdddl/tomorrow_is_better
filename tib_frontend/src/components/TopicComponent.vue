<template>
  <div class="header">
    <h1>下次一定！</h1>
    <h4>内测中，非常需要大家的修改建议。敬请批评指正：sikongdddl@gmail.com</h4>
    <el-button type="primary" @click="goToProfile" style="float: right;">个人主页</el-button>
  </div>
  <div class="topics-container">
    <el-card class="topic-card" v-for="topic in topics" :key="topic.ID">
      <h3>{{ topic.Title }}</h3>
      <p>Event Time: {{ new Date(topic.EventTime).toLocaleString() }}</p>
      <p>Location: {{ topic.Location }}</p>
      <p>
        Participant List:
      </p>
      <p v-for = "participant in topic.Participants" :key="participant.ID">
        {{participant.Role}}:   {{userMap[participant.UserID]}}
      </p>
      <el-button type="primary" @click="goToTopicDetail(topic.ID)">查看详情</el-button>
      <el-button type="success" @click="join(topic.ID)">Join</el-button>
      <el-button type="danger" @click="leave(topic.ID)">Leave</el-button>
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
import { fetchTopics, createTopic, joinTopic, leaveTopic } from '@/services/topicService';
import { getNameById} from "@/services/userService";

export default {
  data() {
    return {
      topics: [],
      newTopic: {
        title: '',
        event_time: '',
        location: ''
      },
      userId: null, // 从 localStorage 获取用户 ID
      userMap: {}
    };
  },
  created() {
    this.userId = parseInt(localStorage.getItem('userId'), 10);
    this.loadTopics();
  },
  methods: {
    goToProfile() {
      this.$router.push('/profile');
    },
    goToTopicDetail(id){
      this.$router.push({name: "TopicDetail", params: {id} });
    },
    async loadTopics() {
      try {
        const res = await fetchTopics();
        this.topics = res.data.topics;

        const userIDs = new Set();
        this.topics.forEach(topic => {

          topic.Participants.forEach(p => userIDs.add(p.UserID));
        });

        await Promise.all([...userIDs].map(id => this.getName(id)));
        console.log(this.userMap);
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
    },
    async leave(topicId){
      try{
        await leaveTopic(this.userId, topicId);
        alert('Leave successfully');
        this.loadTopics();
      } catch (e){
        alert('Leave failed: '+ (e.response?.data?.error || 'Unknown error'));
      }
    },
    async getName(userID) {
      if (this.userMap[userID]) return;
      const res = await getNameById(userID);
      this.userMap[userID] = res.data.username;
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

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}


.create-topic {
  grid-column: 1 / -1;
  margin-top: 40px;
}
</style>
