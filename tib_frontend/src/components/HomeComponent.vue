<template>
  <div class="home">
    <el-card class="login-card" v-if="isLogin">
      <h2>Login</h2>
      <el-form :model="form" ref="form" label-width="100px" @submit.prevent="login">
        <el-form-item label="Username" :rules="[{ required: true, message: 'Please input your username', trigger: 'blur' }]">
          <el-input v-model="form.username" placeholder="Username"></el-input>
        </el-form-item>

        <el-form-item label="Password" :rules="[{ required: true, message: 'Please input your password', trigger: 'blur' }]">
          <el-input type="password" v-model="form.password" placeholder="Password"></el-input>
        </el-form-item>

        <el-button type="primary" @click="login">Login</el-button>
      </el-form>
      <p>Don't have an account? <a href="#" @click="toggleForm">Register here</a></p>
    </el-card>

    <el-card class="register-card" v-else>
      <h2>Register</h2>
      <el-form :model="form" ref="form" label-width="100px" @submit.prevent="register">
        <el-form-item label="Username" :rules="[{ required: true, message: 'Please input your username', trigger: 'blur' }]">
          <el-input v-model="form.username" placeholder="Username"></el-input>
        </el-form-item>

        <el-form-item label="Password" :rules="[{ required: true, message: 'Please input your password', trigger: 'blur' }]">
          <el-input type="password" v-model="form.password" placeholder="Password"></el-input>
        </el-form-item>

        <el-button type="primary" @click="register">Register</el-button>
      </el-form>
      <p>Already have an account? <a href="#" @click="toggleForm">Login here</a></p>
    </el-card>
  </div>
</template>

<script>
import { login, register } from '@/services/userService'; // 导入封装的登录和注册请求

export default {
  data() {
    return {
      isLogin: true, // 控制当前显示登录还是注册
      form: {
        username: '',
        password: '',
        confirmPassword: '',
      },
    };
  },
  methods: {
    toggleForm() {
      this.isLogin = !this.isLogin;
      this.clearForm();
    },

    // 登录请求
    async login() {
      try {
        const response = await login(this.form.username, this.form.password);
        console.log('Login success:', response);
        localStorage.setItem('userId', response.data.user.id)
        this.$router.push('/topics');
        // 处理登录成功的逻辑，例如保存token或跳转页面
      } catch (error) {
        console.error('Login failed:', error);
        alert('Login failed: ' + (error.response?.data?.message || 'Unknown error'));
      }
    },

    // 注册请求
    async register() {
      try {
        const response = await register(this.form.username, this.form.password);
        console.log('Registration success:', response);
        // 处理注册成功的逻辑，例如跳转到登录页面
      } catch (error) {
        console.error('Registration failed:', error);
        alert('Registration failed: ' + (error.response?.data?.message || 'Unknown error'));
      }
    },

    // 清空表单数据
    clearForm() {
      this.form.username = '';
      this.form.password = '';
    }
  }
};
</script>

<style scoped>
.home {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh; /* 让页面占满整个视口高度 */
  margin: 0;
  background-color: #f4f4f4;
}

.el-card {
  width: 400px;
  padding: 20px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  margin: 0 auto;
}

a {
  color: #409eff;
}

a:hover {
  text-decoration: underline;
}
</style>
