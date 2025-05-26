import axios from 'axios';

// Determine the base URL based on the environment
const API_BASE_URL = process.env.NODE_ENV === 'development'
  ? 'http://localhost:8080/api/v1' // Backend server for development
  : 'https://your-production-api-url.com/api/v1'; // TODO: Replace with your actual production API URL

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Optional: Add a request interceptor for things like adding auth tokens
apiClient.interceptors.request.use(
  (config) => {
    // const token = localStorage.getItem('userToken'); // Example: get token from localStorage
    // if (token) {
    //   config.headers.Authorization = `Bearer ${token}`;
    // }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Optional: Add a response interceptor for global error handling
apiClient.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error('API Error:', error.response || error.message || error);
    // You could also dispatch a global error notification here
    if (error.response && error.response.status === 401) {
      // Handle unauthorized errors, e.g., redirect to login
      // window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default apiClient; 