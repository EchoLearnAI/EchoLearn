import apiClient from './apiClient';

/**
 * Creates a new user.
 * @param {object} userData - The user data.
 * @param {string} userData.name - The user's name.
 * @param {string} userData.email - The user's email.
 * @returns {Promise<object>} The created user data.
 */
export const createUser = async (userData) => {
  try {
    const response = await apiClient.post('/users', userData);
    return response.data;
  } catch (error) {
    console.error('Error creating user:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Fetches a user by their ID.
 * @param {string} userId - The ID of the user to fetch.
 * @returns {Promise<object>} The user data.
 */
export const getUserById = async (userId) => {
  try {
    const response = await apiClient.get(`/users/${userId}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching user ${userId}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
}; 