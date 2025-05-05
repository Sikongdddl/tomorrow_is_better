import axios from 'axios';

// 创建一个axios实例
const instance = axios.create({
    baseURL: 'http://localhost:8080', // 替换为你的后端API根路径
    timeout: 10000, // 设置请求超时
    headers: {
        'Content-Type': 'application/json',
    },
});
export default instance;
