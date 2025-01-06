import axios from 'axios'


const request = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 3000,
})

// 请求拦截器
request.interceptors.request.use(
    (config) => {
        // 在发送请求之前可以进行一些操作，比如添加请求头
        // config.headers.Authorization = `Bearer ${token}`;
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// 响应拦截器
request.interceptors.response.use(
    (response) => {
        // 在这里可以对响应数据进行处理
        return response.data;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default request
