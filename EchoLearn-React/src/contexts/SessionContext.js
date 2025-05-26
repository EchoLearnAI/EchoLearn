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
  const [answerFeedback, setAnswerFeedback] = useState(null);

  const fetchQuestionForSession = useCallback(async (session) => {
    if (!session || !session.ID || !session.current_question_id) {
      if (session && session.mode === 'category_challenge' && session.category_name) {
        const questions = await getQuestionsByCategory(session.category_name);
        setCurrentQuestion(questions.data[0] || null); 
        return questions.data[0] || null;
      } else if (session && session.current_question_id === "") { 
        setCurrentQuestion(null);
        return null;
      } else if (session && session.current_question_id) { 
        setCurrentQuestion(null);
        return null;
      }
      setCurrentQuestion(null); 
      return null;
    }
    try {
      const questionResponse = await getQuestionById(session.current_question_id);
      // Assuming getQuestionById from service returns { data: questionObject } or questionObject directly
      // Let's assume service returns questionObject directly as per previous fixes pattern
      setCurrentQuestion(questionResponse || null); // if service returns {data: q}, then questionResponse.data
      return questionResponse || null;
    } catch (error) {
      console.error('[SessionContext] Error fetching question for session:', error);
      setSessionError(error);
      setCurrentQuestion(null);
      return null;
    }
  }, []);

  const startNewSession = useCallback(async (sessionParams) => {
    console.log('[SessionContext] Attempting to start new session with params:', sessionParams);
    if (!user || !user.id) {
      const noUserError = new Error('User not logged in or user ID is missing. Cannot start session.');
      console.error('[SessionContext] User check failed:', noUserError, 'User object:', user);
      setSessionError(noUserError);
      return null;
    }
    console.log('[SessionContext] User verified:', user.id);

    setIsLoadingSession(true);
    setSessionError(null);
    setAnswerFeedback(null);
    try {
      const sessionData = { user_id: user.id, ...sessionParams };
      console.log('[SessionContext] Calling apiStartSession with data:', sessionData);
      
      const newSessionObject = await apiStartSession(sessionData); 
      
      console.log('[SessionContext] Received from apiStartSession:', newSessionObject);

      if (!newSessionObject || !newSessionObject.id) {
        console.error('[SessionContext] Failed to create session or session object is invalid:', newSessionObject);
        setSessionError(new Error('Failed to initialize session from server.'));
        setCurrentSession(null);
        setCurrentQuestion(null);
        return null;
      }

      setCurrentSession(newSessionObject);
      console.log('[SessionContext] Session set in context:', newSessionObject);

      if (newSessionObject.current_question_id) {
        console.log('[SessionContext] Session has current_question_id, fetching question:', newSessionObject.current_question_id);
        await fetchQuestionForSession(newSessionObject);
      } else {
        console.log('[SessionContext] Session has no current_question_id, setting currentQuestion to null.');
        setCurrentQuestion(null); 
      }
      console.log('[SessionContext] startNewSession successful, returning session object.');
      return newSessionObject; // Return the actual session object
    } catch (error) {
      console.error('[SessionContext] Error during apiStartSession call or subsequent logic:', error);
      setSessionError(error);
      setCurrentSession(null);
      setCurrentQuestion(null);
      return null;
    } finally {
      setIsLoadingSession(false);
      console.log('[SessionContext] startNewSession finished.');
    }
  }, [user, fetchQuestionForSession]);

  const submitUserAnswer = useCallback(async (sessionId, questionId, selectedOptionId) => {
    setIsLoadingSession(true);
    setSessionError(null);
    try {
      // apiSubmitAnswer from sessionService.js returns response.data directly
      const feedback = await apiSubmitAnswer({ sessionId, questionId, selectedOptionId });
      setAnswerFeedback(feedback);
      
      const sessionUpdate = feedback?.session_update;
      setCurrentSession(prevSession => ({
        ...prevSession,
        ...(sessionUpdate || {}),
        score: sessionUpdate?.score ?? prevSession?.score,
        mistakes_made: sessionUpdate?.mistakes_made ?? prevSession?.mistakes_made,
        is_active: sessionUpdate?.is_active ?? prevSession?.is_active,
        current_question_id: sessionUpdate?.current_question_id ?? prevSession?.current_question_id,
      }));

      if (sessionUpdate?.is_active && sessionUpdate?.current_question_id) {
        await fetchQuestionForSession(sessionUpdate); // Pass the session_update part which contains the next question ID
      } else {
        setCurrentQuestion(null); 
      }
      return feedback;
    } catch (error) {
      console.error('[SessionContext] Error submitting answer:', error);
      setSessionError(error);
      return null;
    } finally {
      setIsLoadingSession(false);
    }
  }, [fetchQuestionForSession]);

  const fetchSessionSummary = useCallback(async (sessionId) => {
    setIsLoadingSession(true);
    setSessionError(null);
    try {
      // apiGetSessionSummary from sessionService.js returns response.data directly
      const summary = await apiGetSessionSummary(sessionId);
      if (currentSession && currentSession.id === sessionId) {
        setCurrentSession(prev => ({...prev, ...summary, is_active: false }));
      }
      setCurrentQuestion(null); 
      return summary;
    } catch (error) {
      console.error('[SessionContext] Error fetching session summary:', error);
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
        setAnswerFeedback, 
      }}
    >
      {children}
    </SessionContext.Provider>
  );
};

export default SessionContext; 