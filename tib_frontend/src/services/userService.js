import axios from 'axios'; // 引入axios实例

const baseURL = 'http://localhost:8080';

const userService = axios.create({
    baseURL,
    timeout: 10000,
})

// 登录请求
export const login = (username, password) => {
    return userService.post('/login', {
        username,
        password,
    }); // 返回响应数据，处理成功
};

// 注册请求
export const register = (username, password) => {
    return userService.post('/register', {
        username,
        password,
    }); // 返回响应数据，处理成功
};
