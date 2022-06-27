import axios from "axios";

const token = localStorage.getItem('token');
const RLPP = localStorage.getItem('RLPP');
const id = localStorage.getItem("id");

const axiosInstance = axios.create({
    baseURL: 'https://hicoder.herokuapp.com/api',
    headers: {
        authorization: `barer ${token}`,
        id: id,
        RLPP: RLPP,
    }
  })

axiosInstance.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default axiosInstance;