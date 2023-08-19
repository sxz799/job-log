import axios from 'axios'


const request = axios.create({
    baseURL: import.meta.env.VITE_BASE_PATH,
    timeout: 3000
})

export default request
