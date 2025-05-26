import apiClient from './apiClient';

/**
 * Starts a new game session.
 * @param {object} sessionData - Data to start the session.
 * @param {string} sessionData.userId - The ID of the user starting the session.
 * @param {string} sessionData.mode - The game mode (e.g., 'mistakes', 'timed', 'infinite', 'category_challenge').
 * @param {string} [sessionData.categoryName] - The category name, if mode is 'category_challenge'.
 * @param {number} [sessionData.maxMistakes] - Maximum mistakes allowed, if mode is 'mistakes'.
 * @param {number} [sessionData.totalQuestions] - Total questions for the session, if mode is 'timed' or 'category_challenge' (can also be used with 'mistakes').
 * @returns {Promise<object>} The newly created session object.
 */
export const startSession = async (sessionData) => {
  try {
    const response = await apiClient.post('/session/start', sessionData);
    return response.data;
  } catch (error) {
    console.error('Error starting session:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Submits an answer for the current question in a session.
 * @param {object} answerData - Data for submitting the answer.
 * @param {string} answerData.sessionId - The ID of the current session.
 * @param {string} answerData.questionId - The ID of the question being answered.
 * @param {string} answerData.selectedOptionId - The ID of the option selected by the user.
 * @returns {Promise<object>} Feedback for the submitted answer, including whether it was correct and the session status.
 */
export const submitAnswer = async (answerData) => {
  try {
    const response = await apiClient.post('/session/submit', answerData);
    return response.data;
  } catch (error) {
    console.error('Error submitting answer:', error.response?.data || error.message);
    throw error.response?.data || error;
  }
};

/**
 * Fetches the summary of a game session.
 * @param {string} sessionId - The ID of the session to get the summary for.
 * @returns {Promise<object>} The session summary, including score, mistakes, and answered questions.
 */
export const getSessionSummary = async (sessionId) => {
  try {
    const response = await apiClient.get(`/session/${sessionId}/summary`);
    return response.data;
  } catch (error) {
    console.error(`Error fetching session summary for ${sessionId}:`, error.response?.data || error.message);
    throw error.response?.data || error;
  }
}; 