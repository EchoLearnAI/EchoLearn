import apiClient from './apiClient';

/**
 * Fetches a random question.
 * @returns {Promise<object>} A random question object.
 */
export const getRandomQuestion = async () => {
  try {
    const response = await apiClient.get('/questions/random');
    return response.data;
  } catch (error) {
    console.error('Error fetching random question:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Fetches questions by category name.
 * @param {string} categoryName - The name of the category.
 * @returns {Promise<Array<object>>} A list of questions in the specified category.
 */
export const getQuestionsByCategory = async (categoryName) => {
  try {
    const response = await apiClient.get(`/questions/category/${categoryName}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching questions for category ${categoryName}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Fetches a question by its ID.
 * @param {string} questionId - The ID of the question.
 * @returns {Promise<object>} The question object.
 */
export const getQuestionById = async (questionId) => {
  try {
    const response = await apiClient.get(`/questions/${questionId}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching question ${questionId}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Adds a new question.
 * @param {object} questionData - The question data to add.
 * @returns {Promise<object>} The added question data.
 */
export const addQuestion = async (questionData) => {
  try {
    const response = await apiClient.post('/questions', questionData);
    return response.data;
  } catch (error) {
    console.error('Error adding question:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
}; 