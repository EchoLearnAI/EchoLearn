import apiClient from './apiClient';

export const startSession = async (userId, mode) => {
  try {
    const response = await apiClient.post('/session/start', { user_id: userId, mode });
    return response.data;
  } catch (error) {
    console.error('Error starting session:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

export const submitAnswer = async (sessionId, questionId, selectedOptionId) => {
  try {
    const response = await apiClient.post('/session/submit', {
      session_id: sessionId,
      question_id: questionId,
      selected_option_id: selectedOptionId,
    });
    return response.data; // Contains result, session_active, session_details, explanation
  } catch (error) {
    console.error('Error submitting answer:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

export const getSessionSummary = async (sessionId) => {
  try {
    const response = await apiClient.get(`/session/${sessionId}/summary`);
    return response.data;
  } catch (error) {
    console.error('Error fetching session summary:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
}; 