import axios from 'axios';

// IMPORTANT: Configure this URL based on your setup!
// See README.md for details (Android Emulator, iOS Simulator, Physical Device)
const API_BASE_URL = 'http://localhost:8080/api/v1'; // Default for iOS Sim / if backend is on same machine
// const API_BASE_URL = 'http://10.0.2.2:8080/api/v1'; // Common for Android Emulator

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

apiClient.interceptors.response.use(
  response => response,
  error => {
    // Log error or handle specific error codes globally if needed
    console.error('API Error:', error.response || error.message);
    // You might want to navigate to an error screen or show a global message
    return Promise.reject(error);
  },
);

export default apiClient; 