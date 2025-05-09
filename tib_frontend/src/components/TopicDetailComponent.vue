<template>
  <div class="topic-detail">
    <el-card v-if="topic">
      <h2>活动详情</h2>
      <el-descriptions :column="1" border>
        <el-descriptions-item label="标题">{{ topic.Title }}</el-descriptions-item>
        <el-descriptions-item label="地点">{{ topic.Location }}</el-descriptions-item>
        <el-descriptions-item label="时间">{{ topic.EventTime }}</el-descriptions-item>
        <el-descriptions-item label="简介">{{ topic.Description || '暂无简介' }}</el-descriptions-item>
      </el-descriptions>
    </el-card>
  </div>

  <div class="comment-section">
    <h3>评论区</h3>

    <!-- 新增评论 -->
    <el-form :model="newComment" @submit.prevent="handleAddComment">
      <el-form-item label="添加评论">
        <el-input
            type="textarea"
            v-model="newComment.content"
            placeholder="写下你的评论..."
            rows="3"
        />
      </el-form-item>
      <el-button type="primary" @click="handleAddComment">提交</el-button>
    </el-form>

    <el-divider />

    <!-- 评论列表 -->
    <div v-if="comments != null">
      <el-card v-for="comment in comments" :key="comment.comment_id" class="mb-2">
        <div class="comment-header">
          <strong>{{ comment.username + "   "}}</strong>
          <span class="comment-time">{{ formatTime(comment.created_at) }}</span>
        </div>
        <div class="comment-content">
          {{ comment.content }}
        </div>
        <div v-if="comment.user_id === this.user_id">
          <el-button type="danger" @click="handleDelComment(comment.comment_id)">删除评论</el-button>
        </div>
      </el-card>
    </div>
    <div v-else class="empty-text">暂无评论</div>
  </div>

</template>

<script>
import {getTopicById} from "@/services/topicService";
import { addComment, getComment, delComment} from "@/services/commentService";
export default {
  data() {
    return {
      topic: null,
      topic_id: parseInt(localStorage.getItem("topicId")),
      user_id : parseInt(localStorage.getItem("userId")),
      comments:[],
      newComment:{
        content:''
      },
    };
  },
  async created() {
    try {
      const res = await getTopicById(this.topic_id);
      console.log(res)

      this.topic = res.data.message;
      await this.loadComments()
    } catch (e) {
      this.$message.error("加载活动详情失败");
    }
  },
  methods: {
    async loadComments() {
      try {
        const commentRes = await getComment(this.topic_id);
        this.comments = commentRes.data.comments;
      } catch (e) {
        console.error('Failed to load topics:', e);
      }
    },
    async handleAddComment(){
      const content = this.newComment.content.trim();
      if (!content){
        alert("评论不能为空");
        return;
      }
      try{
        await addComment(this.user_id, this.topic_id,content);
        this.newComment.content = "";
        await this.loadComments()
      }catch (err){
        alert(err);
      }
    },
    async handleDelComment(comment_id){
      try{
        await delComment(comment_id);
        this.loadComments()
        alert("删除评论成功！");
      }catch (err){
        alert("删除评论失败！");
      }
    },
    formatTime(timeStr) {
      return new Date(timeStr).toLocaleString();
    },
  }
};
</script>

<style scoped>
.topic-detail {
  padding: 20px;
}
.comment-section {
  margin: 20px;
}
.empty-text {
  color: #888;
  padding: 10px;
}
</style>
