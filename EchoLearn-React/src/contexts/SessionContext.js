import React, { createContext, useState, useCallback, useContext } from 'react';
import {
  startSession as apiStartSession,
  submitAnswer as apiSubmitAnswer,
  getSessionSummary as apiGetSessionSummary,
} from '../api/sessionService';
import { getRandomQuestion, getQuestionsByCategory, getQuestionById } from '../api/questionService';
import UserContext from './UserContext';

const SessionContext = createContext();

export const SessionProvider = ({ children }) => {
  const { user } = useContext(UserContext);
  const [currentSession, setCurrentSession] = useState(null);
  const [currentQuestion, setCurrentQuestion] = useState(null);
  const [isLoadingSession, setIsLoadingSession] = useState(false);
  const [sessionError, setSessionError] = useState(null);
  const [answerFeedback, setAnswerFeedback] = useState(null); // To store feedback from submitAnswer

  const fetchQuestionForSession = useCallback(async (session) => {
    if (!session || !session.ID || !session.current_question_id) {
      // If no current_question_id, it might be the start of a session where the first question needs to be fetched based on mode
      // Or the session might have ended. For now, let's assume a question ID is present if session is active.
      if (session && session.mode === 'category_challenge' && session.category_name) {
        // Example: Fetch a question for category challenge. This logic might need refinement.
        const questions = await getQuestionsByCategory(session.category_name);
        // Select a question that hasn't been answered. This needs more complex logic based on session.answered_questions_json
        setCurrentQuestion(questions.data[0] || null); // Placeholder
        return questions.data[0] || null;
      } else if (session && session.current_question_id === "") { // Indicates session ended
        setCurrentQuestion(null);
        return null;
      } else if (session && session.current_question_id) { //This path should not be hit ideally if the backend sends the question_id as empty string when it's over
         //This is a safeguard
        setCurrentQuestion(null);
        return null;
      }
      // Fallback or if session doesn't dictate the next question (e.g. random for infinite)
      // const question = await getRandomQuestion(); // This might be too simplistic
      // setCurrentQuestion(question.data || question);
      // return question.data || question;
      setCurrentQuestion(null); // Default to null if no specific logic matches
      return null;
    }
    try {
      const question = await getQuestionById(session.current_question_id);
      setCurrentQuestion(question.data || question);
      return question.data || question;
    } catch (error) {
      console.error('Error fetching question for session:', error);
      setSessionError(error);
      setCurrentQuestion(null);
      return null;
    }
  }, []);

  const startNewSession = useCallback(async (sessionParams) => {
    if (!user || !user.ID) {
      setSessionError(new Error('User not logged in. Cannot start session.'));
      return null;
    }
    setIsLoadingSession(true);
    setSessionError(null);
    setAnswerFeedback(null);
    try {
      const sessionData = { userId: user.ID, ...sessionParams };
      const newSession = await apiStartSession(sessionData);
      setCurrentSession(newSession.data || newSession);
      if (newSession.data?.current_question_id || newSession?.current_question_id) {
        await fetchQuestionForSession(newSession.data || newSession);
      } else {
        setCurrentQuestion(null); // No initial question, or session starts differently
      }
      return newSession.data || newSession;
    } catch (error) {
      console.error('Error starting new session:', error);
      setSessionError(error);
      setCurrentSession(null);
      setCurrentQuestion(null);
      return null;
    } finally {
      setIsLoadingSession(false);
    }
  }, [user, fetchQuestionForSession]);

  const submitUserAnswer = useCallback(async (sessionId, questionId, selectedOptionId) => {
    setIsLoadingSession(true);
    setSessionError(null);
    try {
      const feedback = await apiSubmitAnswer({ sessionId, questionId, selectedOptionId });
      setAnswerFeedback(feedback.data || feedback);
      // Update session state based on feedback
      setCurrentSession(prevSession => ({
        ...prevSession,
        ...(feedback.data?.session_update || feedback?.session_update || {}),
        score: feedback.data?.session_update?.score ?? prevSession?.score,
        mistakes_made: feedback.data?.session_update?.mistakes_made ?? prevSession?.mistakes_made,
        is_active: feedback.data?.session_update?.is_active ?? prevSession?.is_active,
        current_question_id: feedback.data?.session_update?.current_question_id ?? prevSession?.current_question_id,
      }));

      if (feedback.data?.session_update?.is_active && (feedback.data?.session_update?.current_question_id || feedback?.session_update?.current_question_id)) {
        await fetchQuestionForSession(feedback.data.session_update);
      } else {
        setCurrentQuestion(null); // Session ended or no next question
      }
      return feedback.data || feedback;
    } catch (error) {
      console.error('Error submitting answer:', error);
      setSessionError(error);
      // Potentially keep answerFeedback as null or set an error state for it
      return null;
    } finally {
      setIsLoadingSession(false);
    }
  }, [fetchQuestionForSession]);

  const fetchSessionSummary = useCallback(async (sessionId) => {
    setIsLoadingSession(true);
    setSessionError(null);
    try {
      const summary = await apiGetSessionSummary(sessionId);
      // The summary might not directly affect currentSession or currentQuestion state here
      // It's more for display on a summary page.
      // However, you might want to update parts of the session or ensure it's marked as inactive.
      if (currentSession && currentSession.ID === sessionId) {
        setCurrentSession(prev => ({...prev, ...summary.data, is_active: false }));
      }
      setCurrentQuestion(null); // No active question when viewing summary
      return summary.data || summary;
    } catch (error) {
      console.error('Error fetching session summary:', error);
      setSessionError(error);
      return null;
    } finally {
      setIsLoadingSession(false);
    }
  }, [currentSession]);

  const endCurrentSession = useCallback(() => {
    setCurrentSession(null);
    setCurrentQuestion(null);
    setAnswerFeedback(null);
    setSessionError(null);
    // No API call to explicitly end session on backend, but frontend state is cleared.
    // Backend session becomes inactive after summary or based on game rules.
  }, []);

  return (
    <SessionContext.Provider
      value={{
        currentSession,
        currentQuestion,
        isLoadingSession,
        sessionError,
        answerFeedback,
        startNewSession,
        submitUserAnswer,
        fetchSessionSummary,
        endCurrentSession,
        setAnswerFeedback, // Expose to allow clearing feedback manually if needed
      }}
    >
      {children}
    </SessionContext.Provider>
  );
};

export default SessionContext; 