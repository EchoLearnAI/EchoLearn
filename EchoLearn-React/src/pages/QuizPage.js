import React, { useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import styled, { css } from 'styled-components';
import SessionContext from '../contexts/SessionContext';
import UserContext from '../contexts/UserContext'; // To check if user is logged in

const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: ${props => props.theme.spacings.medium};
  width: 100%;
  max-width: 800px; // Max width for quiz content
  margin: 0 auto; // Center the quiz content
`;

const LoadingMessage = styled.p`
  font-size: ${props => props.theme.fontSizes.large};
  color: ${props => props.theme.colors.secondary};
  margin-top: ${props => props.theme.spacings.xlarge};
`;

const ErrorMessage = styled.p`
  color: ${props => props.theme.colors.danger};
  font-size: ${props => props.theme.fontSizes.medium};
  margin-top: ${props => props.theme.spacings.medium};
  background-color: ${props => props.theme.colors.danger}20; // Light red background
  padding: ${props => props.theme.spacings.medium};
  border-radius: ${props => props.theme.radii.small};
`;

const QuestionCard = styled.div`
  background-color: ${props => props.theme.colors.white};
  padding: ${props => props.theme.spacings.large};
  border-radius: ${props => props.theme.radii.medium};
  box-shadow: ${props => props.theme.shadows.medium};
  width: 100%;
  margin-bottom: ${props => props.theme.spacings.large};
`;

const QuestionText = styled.h2`
  font-size: ${props => props.theme.fontSizes.xlarge};
  color: ${props => props.theme.colors.text};
  margin-bottom: ${props => props.theme.spacings.small};
  line-height: 1.4;
`;

const GrammarRuleTitle = styled.p`
  font-size: ${props => props.theme.fontSizes.small};
  color: ${props => props.theme.colors.secondary};
  margin-bottom: ${props => props.theme.spacings.medium};
  font-style: italic;
`;

const OptionsList = styled.div`
  display: flex;
  flex-direction: column;
  gap: ${props => props.theme.spacings.small};
`;

const OptionButton = styled.button`
  padding: ${props => props.theme.spacings.medium};
  background-color: ${props => props.theme.colors.light};
  color: ${props => props.theme.colors.text};
  border: 1px solid ${props => props.theme.colors.border};
  border-radius: ${props => props.theme.radii.small};
  font-size: ${props => props.theme.fontSizes.medium};
  text-align: left;
  cursor: pointer;
  transition: background-color 0.2s ease-in-out, border-color 0.2s ease-in-out;

  &:hover {
    background-color: ${props => props.theme.colors.border};
  }

  &:disabled {
    cursor: not-allowed;
    opacity: 0.7;
  }

  ${props => 
    props.isCorrect &&
    css`
      background-color: ${props.theme.colors.success}30; // Light green
      border-color: ${props.theme.colors.success};
      color: ${props.theme.colors.success};
    `}

  ${props =>
    props.isIncorrect &&
    css`
      background-color: ${props.theme.colors.danger}30; // Light red
      border-color: ${props.theme.colors.danger};
      color: ${props.theme.colors.danger};
    `}
`;

const FeedbackArea = styled.div`
  margin-top: ${props => props.theme.spacings.medium};
  padding: ${props => props.theme.spacings.medium};
  border-radius: ${props => props.theme.radii.small};
  text-align: center;

  ${props => 
    props.isCorrect &&
    css`
      background-color: ${props.theme.colors.success}20;
      color: ${props.theme.colors.success};
    `}

  ${props =>
    !props.isCorrect && 
    props.message && // Only show if there is a message (i.e., answer submitted)
    css`
      background-color: ${props.theme.colors.danger}20;
      color: ${props.theme.colors.danger};
    `}
`;

const ExplanationText = styled.p`
    font-size: ${props => props.theme.fontSizes.small};
    margin-top: ${props => props.theme.spacings.xsmall};
`;

const SessionInfo = styled.div`
  display: flex;
  justify-content: space-between;
  width: 100%;
  padding: ${props => props.theme.spacings.small} 0;
  margin-bottom: ${props => props.theme.spacings.medium};
  font-size: ${props => props.theme.fontSizes.small};
  color: ${props => props.theme.colors.secondary};
`;

const NextButton = styled.button`
  margin-top: ${props => props.theme.spacings.medium};
  padding: ${props => props.theme.spacings.small} ${props => props.theme.spacings.large};
  background-color: ${props => props.theme.colors.primary};
  color: ${props => props.theme.colors.white};
  border-radius: ${props => props.theme.radii.small};
  font-size: ${props => props.theme.fontSizes.medium};
  cursor: pointer;

   &:disabled {
    background-color: ${props => props.theme.colors.disabled};
  }
`;

const FinishQuizButton = styled(NextButton)`
    background-color: ${props => props.theme.colors.secondary};
    &:hover {
        background-color: ${props => props.theme.colors.secondary}d0;
    }
`;

const QuizPage = () => {
  const { user } = useContext(UserContext);
  const {
    currentSession,
    currentQuestion,
    isLoadingSession,
    sessionError,
    answerFeedback,
    submitUserAnswer,
    fetchSessionSummary,
    setAnswerFeedback, // To clear feedback for next question
  } = useContext(SessionContext);
  const navigate = useNavigate();

  const [selectedOptionId, setSelectedOptionId] = useState(null);
  const [isAnswerSubmitted, setIsAnswerSubmitted] = useState(false);

  useEffect(() => {
    // If no user, redirect to create user page (or login if that existed)
    if (!user) {
      navigate('/create-user');
      return;
    }
    // If no active session or no question, navigate to home to start a new one.
    // This handles cases where user directly navigates to /quiz or session ends unexpectedly.
    if (!isLoadingSession && !currentSession) {
      navigate('/');
    }
    // If session is active but no question and not loading, might also be an issue or end of quiz
    // This will be handled by answerFeedback and session.is_active logic primarily

  }, [user, currentSession, isLoadingSession, navigate]);

  useEffect(() => {
    // When a new question loads, reset selection and submission state
    setSelectedOptionId(null);
    setIsAnswerSubmitted(false);
    // Clear previous feedback when new question loads, unless feedback is for current submission
    if (!isAnswerSubmitted) {
        setAnswerFeedback(null); 
    }
  }, [currentQuestion, setAnswerFeedback, isAnswerSubmitted]);

  useEffect(() => {
    // Check if session has ended based on feedback or session state
    if (answerFeedback && !answerFeedback.is_active) {
        // If session ended after an answer
        navigate(`/summary/${currentSession.ID}`);
    } else if (currentSession && !currentSession.is_active && currentSession.ID) {
        // If session is marked inactive (e.g. fetched summary elsewhere)
        navigate(`/summary/${currentSession.ID}`);
    }
  }, [answerFeedback, currentSession, navigate]);


  const handleOptionSelect = (optionId) => {
    if (isAnswerSubmitted) return; // Don't allow changing answer after submission
    setSelectedOptionId(optionId);
  };

  const handleSubmitAnswer = async () => {
    if (!selectedOptionId || !currentQuestion || !currentSession) return;

    setIsAnswerSubmitted(true);
    await submitUserAnswer(currentSession.ID, currentQuestion.ID, selectedOptionId);
    // Navigation to summary or next question is handled by useEffect listening to answerFeedback and currentSession
  };

  const handleNextQuestion = () => {
    // This function is mostly for UI if we need an explicit next button *after* feedback is shown
    // The core logic of fetching next question is in SessionContext after submitUserAnswer.
    // Here, we just clear the feedback and submission state for the UI.
    setSelectedOptionId(null);
    setIsAnswerSubmitted(false);
    setAnswerFeedback(null);
    // The SessionContext should have already updated currentQuestion if there is one
    if (!currentQuestion && currentSession && currentSession.is_active) {
        // This case should ideally not be hit if context manages question flow properly
        console.warn("Next button clicked but no current question available.")
    }
  };
  
  const handleFinishQuiz = async () => {
    if (currentSession && currentSession.ID) {
      // Fetch summary which also marks session as inactive on backend via GetSessionSummary logic
      await fetchSessionSummary(currentSession.ID);
      navigate(`/summary/${currentSession.ID}`);
    }
  };

  if (isLoadingSession && !currentQuestion && !answerFeedback) return <PageContainer><LoadingMessage>Loading Quiz...</LoadingMessage></PageContainer>
  if (sessionError) return <PageContainer><ErrorMessage>Error: {sessionError.message || 'Could not load quiz data.'}</ErrorMessage></PageContainer>
  if (!currentSession) return <PageContainer><LoadingMessage>No active session. Redirecting...</LoadingMessage></PageContainer>; // Should be caught by useEffect
  // If currentSession exists, but no currentQuestion and feedback is shown, it means quiz ended.
  if (currentSession && !currentQuestion && answerFeedback && !answerFeedback.is_active){
      // This state is handled by useEffect navigating to summary.
      // If somehow stuck, show loading or message.
      return <PageContainer><LoadingMessage>Quiz ended, preparing summary...</LoadingMessage></PageContainer>
  }
  if (!currentQuestion && currentSession && currentSession.is_active && !isLoadingSession) {
    // This means session is active, but no question loaded (e.g. after all questions in a finite mode)
    // Backend should have set current_question_id to empty string. Session context handles this.
    // Or, if it's an infinite mode quiz that was just started and the first question hasn't loaded yet.
    if(currentSession.mode === 'infinite' && !answerFeedback) {
        // This might be a brief state for infinite mode starting. Let it load.
        return <PageContainer><LoadingMessage>Loading first question...</LoadingMessage></PageContainer>
    }
    // For other modes, if no question and session active, it usually means end of quiz.
    // This should ideally be handled by navigation from submitAnswer feedback.
    return <PageContainer><LoadingMessage>No more questions. Finishing quiz...</LoadingMessage></PageContainer>
  }
  if (!currentQuestion) {
      return <PageContainer><LoadingMessage>Loading question...</LoadingMessage></PageContainer>
  }


  const { text, options, grammar_rule } = currentQuestion;
  const feedbackMessage = answerFeedback ? (answerFeedback.is_correct ? 'Correct!' : 'Incorrect.') : '';
  const explanation = answerFeedback?.explanation;

  return (
    <PageContainer>
      <SessionInfo>
        <span>Score: {currentSession.score}</span>
        {currentSession.mode === 'mistakes' && <span>Mistakes: {currentSession.mistakes_made}/{currentSession.max_mistakes}</span>}
        {currentSession.mode !== 'infinite' && <span>Question: {currentSession.answered_questions_json?.length + (isAnswerSubmitted || !currentSession.is_active ? 0 : 1) }/{currentSession.total_questions}</span>}
      </SessionInfo>

      <QuestionCard>
        {grammar_rule && grammar_rule.title && <GrammarRuleTitle>{grammar_rule.title}</GrammarRuleTitle>}
        <QuestionText>{text}</QuestionText>
        <OptionsList>
          {options.map((opt) => (
            <OptionButton
              key={opt.ID}
              onClick={() => handleOptionSelect(opt.ID)}
              disabled={isAnswerSubmitted || isLoadingSession}
              isCorrect={isAnswerSubmitted && opt.is_correct}
              isIncorrect={isAnswerSubmitted && selectedOptionId === opt.ID && !opt.is_correct}
            >
              {opt.text}
            </OptionButton>
          ))}
        </OptionsList>
      </QuestionCard>

      {!isAnswerSubmitted && (
        <NextButton onClick={handleSubmitAnswer} disabled={!selectedOptionId || isLoadingSession}>
          {isLoadingSession ? 'Submitting...' : 'Submit Answer'}
        </NextButton>
      )}

      {isAnswerSubmitted && answerFeedback && (
        <FeedbackArea isCorrect={answerFeedback.is_correct} message={feedbackMessage}>
          <strong>{feedbackMessage}</strong>
          {explanation && <ExplanationText>{explanation}</ExplanationText>}
          {currentSession.is_active && currentQuestion && ( // Show Next if session active and there was a question (even if next is null)
            <NextButton onClick={handleNextQuestion} disabled={isLoadingSession} style={{marginTop: '10px'}}>
              {isLoadingSession ? 'Loading...' : 'Next Question'}
            </NextButton>
          )}
        </FeedbackArea>
      )}
      
      {currentSession.mode === 'infinite' && (
        <FinishQuizButton onClick={handleFinishQuiz} disabled={isLoadingSession} style={{marginTop: '20px'}}>
            {isLoadingSession ? 'Finishing...' : 'Finish Quiz (Infinite Mode)'}
        </FinishQuizButton>
      )}

      {isLoadingSession && <LoadingMessage>Processing...</LoadingMessage>}
    </PageContainer>
  );
};

export default QuizPage; 