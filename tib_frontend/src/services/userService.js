import axios from 'axios'; // 引入axios实例

const baseURL = 'http://localhost:8080';

const userService = axios.create({
    baseURL,
    timeout: 10000,
})

// 登录请求
export const login = (username, password) => {
    return userService.post('/login', {
        username: username,
        password: password,
    }); // 返回响应数据，处理成功
};

// 注册请求
export const register = (username, password) => {
    return userService.post('/register', {
        username: username,
        password: password,
    }); // 返回响应数据，处理成功
};

export const getUserInfByID = (userID) => {
    return userService.post('/getUserInfoByID', {
        user_id: userID
    });
}

// userID 2 username
export const getNameById = (userID) => {
    return userService.post('/getUserNameByID', {
        user_id: userID
    });
}

//upload avatar
export const uploadAvatar = (userID, file) => {
    const formData = new FormData();
    formData.append('user_id', userID);
    formData.append('avatar', file);

    return userService.post('/uploadAvatar', formData, {
        headers: {
            'Content-Type': 'multipart/form-data',
        },
    });
}