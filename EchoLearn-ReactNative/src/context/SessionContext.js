import React, { createContext, useState, useContext } from 'react';
import * as sessionService from '../api/sessionService';
import { UserContext } from './UserContext';

export const SessionContext = createContext();

export const SessionProvider = ({ children }) => {
  const { user } = useContext(UserContext);
  const [currentSession, setCurrentSession] = useState(null);
  const [currentQuestion, setCurrentQuestion] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  const startNewSession = async (mode) => {
    if (!user) {
      setError('User not logged in to start a session.');
      return null;
    }
    setIsLoading(true);
    setError(null);
    try {
      const session = await sessionService.startSession(user.id, mode);
      setCurrentSession(session);
      return session;
    } catch (err) {
      console.error('Failed to start session', err);
      setError(err.message || 'Failed to start session');
      setCurrentSession(null);
      return null;
    } finally {
      setIsLoading(false);
    }
  };

  const submitUserAnswer = async (questionId, selectedOptionId) => {
    if (!currentSession) {
      setError('No active session to submit answer to.');
      return null;
    }
    setIsLoading(true);
    setError(null);
    try {
      const result = await sessionService.submitAnswer(
        currentSession.id,
        questionId,
        selectedOptionId,
      );
      setCurrentSession(result.session_details); // Update session state from response
      return result; // Contains feedback and updated session
    } catch (err) {
      console.error('Failed to submit answer', err);
      setError(err.message || 'Failed to submit answer');
      return null;
    } finally {
      setIsLoading(false);
    }
  };

  const fetchSessionSummary = async (sessionIdToFetch) => {
    const id = sessionIdToFetch || currentSession?.id;
    if (!id) {
      setError('No session ID provided to fetch summary.');
      return null;
    }
    setIsLoading(true);
    setError(null);
    try {
      const summary = await sessionService.getSessionSummary(id);
      // If this is the current session being summarized, update it
      if (currentSession && currentSession.id === id) {
        setCurrentSession(summary.session); 
      }
      return summary;
    } catch (err) {
      console.error('Failed to fetch session summary', err);
      setError(err.message || 'Failed to fetch session summary');
      return null;
    } finally {
      setIsLoading(false);
    }
  };
  
  const endCurrentSession = () => {
    // This might involve calling summary or just clearing local state
    // For infinite mode, calling summary implicitly ends it on backend
    if (currentSession && currentSession.mode === 'infinite' && currentSession.is_active) {
        fetchSessionSummary(currentSession.id); // This will mark it as ended on backend
    }
    setCurrentSession(null);
    setCurrentQuestion(null);
  };


  return (
    <SessionContext.Provider
      value={{
        currentSession,
        setCurrentSession, // For direct manipulation if needed, e.g., after summary
        currentQuestion,
        setCurrentQuestion, // To load questions into the quiz screen
        startNewSession,
        submitUserAnswer,
        fetchSessionSummary,
        endCurrentSession,
        isLoadingSession: isLoading,
        sessionError: error,
      }}>
      {children}
    </SessionContext.Provider>
  );
}; 