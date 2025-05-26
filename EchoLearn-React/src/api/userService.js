import apiClient from './apiClient';

/**
 * Creates a new user.
 * @param {object} userData - Object containing name, email, and password.
 * @param {string} userData.name - The user's name.
 * @param {string} userData.email - The user's email.
 * @param {string} userData.password - The user's password.
 * @returns {Promise<object>} The created user object from the API (nested under 'data').
 */
export const createUser = async (userData) => {
  try {
    const response = await apiClient.post('/users', userData);
    return response.data.data;
  } catch (error) {
    console.error('Error creating user:', error.response ? error.response.data : error.message);
    throw error;
  }
};

/**
 * Gets a user by their ID.
 * @param {string} userId - The ID of the user to retrieve.
 * @returns {Promise<object>} The user object from the API (nested under 'data').
 */
export const getUserById = async (userId) => {
  try {
    const response = await apiClient.get(`/users/${userId}`);
    return response.data.data;
  } catch (error) {
    console.error('Error getting user by ID:', error.response ? error.response.data : error.message);
    throw error;
  }
}; 