import apiClient from './apiClient';

export const getRandomQuestion = async () => {
  try {
    const response = await apiClient.get('/questions/random');
    return response.data;
  } catch (error) {
    console.error('Error fetching random question:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

export const getQuestionsByCategory = async (categoryName) => {
  try {
    const response = await apiClient.get(`/questions/category/${categoryName}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching questions for category ${categoryName}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

export const getQuestionById = async (questionId) => {
  try {
    const response = await apiClient.get(`/questions/${questionId}`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching question ${questionId}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

// AddQuestion is an admin function, typically not used directly in the mobile app by regular users.
// export const addQuestion = async (questionData) => {
//   try {
//     const response = await apiClient.post('/questions', questionData);
//     return response.data;
//   } catch (error) {
//     console.error('Error adding question:', error.response?.data || error.message);
//     throw error.response?.data || error;
//   }
// }; 