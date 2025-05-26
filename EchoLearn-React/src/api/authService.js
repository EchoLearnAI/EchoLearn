import apiClient from './apiClient';

/**
 * Logs in a user.
 * @param {object} credentials - User credentials.
 * @param {string} credentials.email - The user's email.
 * @param {string} credentials.password - The user's password.
 * @returns {Promise<object>} The user object from the API (nested under 'data').
 */
export const loginUser = async (credentials) => {
  try {
    const response = await apiClient.post('/auth/login', credentials);
    return response.data.data; // Assuming backend returns { data: userObject }
  } catch (error) {
    console.error('Error logging in user:', error.response ? error.response.data : error.message);
    throw error;
  }
}; 