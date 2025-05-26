import React, { useContext, useEffect, useState } from 'react';
import { useParams, useNavigate, Link } from 'react-router-dom';
import styled from 'styled-components';
import SessionContext from '../contexts/SessionContext';
import UserContext from '../contexts/UserContext';

const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: ${props => props.theme.spacings.large};
  width: 100%;
  max-width: 800px;
  margin: 0 auto;
`;

const Title = styled.h1`
  color: ${props => props.theme.colors.primary};
  margin-bottom: ${props => props.theme.spacings.large};
`;

const SummaryDetails = styled.div`
  background-color: ${props => props.theme.colors.light}80; // Slightly transparent light background
  padding: ${props => props.theme.spacings.large};
  border-radius: ${props => props.theme.radii.medium};
  box-shadow: ${props => props.theme.shadows.small};
  width: 100%;
  margin-bottom: ${props => props.theme.spacings.large};
  text-align: center;

  p {
    font-size: ${props => props.theme.fontSizes.medium};
    margin-bottom: ${props => props.theme.spacings.small};
    color: ${props => props.theme.colors.text};
  }
  strong {
    color: ${props => props.theme.colors.primary};
  }
`;

const QuestionsReview = styled.div`
  width: 100%;
  margin-top: ${props => props.theme.spacings.large};

  h3 {
    text-align: center;
    margin-bottom: ${props => props.theme.spacings.medium};
    color: ${props => props.theme.colors.secondary};
  }
`;

const QuestionItem = styled.div`
  background-color: ${props => props.theme.colors.white};
  padding: ${props => props.theme.spacings.medium};
  border: 1px solid ${props => props.theme.colors.border};
  border-left: 5px solid ${props => (props.isCorrect ? props.theme.colors.success : props.theme.colors.danger)};
  border-radius: ${props => props.theme.radii.small};
  margin-bottom: ${props => props.theme.spacings.medium};

  p {
    margin-bottom: ${props => props.theme.spacings.xsmall};
  }
`;

const QuestionText = styled.p`
  font-weight: bold;
  color: ${props => props.theme.colors.dark};
`;

const AnswerText = styled.p`
  font-style: italic;
  color: ${props => props.theme.colors.secondary};
`;

const ExplanationText = styled.p`
  font-size: ${props => props.theme.fontSizes.small};
  color: ${props => props.theme.colors.info};
  margin-top: ${props => props.theme.spacings.small};
`;

const LoadingMessage = styled.p`
  font-size: ${props => props.theme.fontSizes.large};
  color: ${props => props.theme.colors.secondary};
`;

const ErrorMessage = styled.p`
  color: ${props => props.theme.colors.danger};
`;

const StyledButton = styled(Link)`
  display: inline-block;
  margin-top: ${props => props.theme.spacings.large};
  padding: ${props => props.theme.spacings.medium} ${props => props.theme.spacings.large};
  background-color: ${props => props.theme.colors.primary};
  color: ${props => props.theme.colors.white};
  border: none;
  border-radius: ${props => props.theme.radii.medium};
  font-size: ${props => props.theme.fontSizes.medium};
  text-decoration: none;
  cursor: pointer;
  transition: background-color 0.2s ease-in-out;

  &:hover {
    background-color: ${props => props.theme.colors.primary}d0;
    color: ${props => props.theme.colors.white}; 
  }
`;

const SummaryPage = () => {
  const { sessionId } = useParams();
  const navigate = useNavigate();
  const { user } = useContext(UserContext);
  const { fetchSessionSummary, endCurrentSession, isLoadingSession, sessionError } = useContext(SessionContext);

  const [summaryData, setSummaryData] = useState(null);

  useEffect(() => {
    if (!user) {
      navigate('/create-user');
      return;
    }
    if (sessionId) {
      fetchSessionSummary(sessionId)
        .then(data => {
          if (data) {
            setSummaryData(data);
          } else {
            // Error already handled in context, but could set local error too
            setSummaryData(null); // Ensure no stale data
          }
        })
        .catch(err => { 
            // This catch is for safety, context should log it
            console.error("Direct catch in SummaryPage fetch: ", err);
            setSummaryData(null); 
        });
    }

    // Cleanup function to end session state when leaving summary page
    return () => {
      endCurrentSession();
    };
  }, [sessionId, fetchSessionSummary, endCurrentSession, user, navigate]);

  if (isLoadingSession && !summaryData) return <PageContainer><LoadingMessage>Loading summary...</LoadingMessage></PageContainer>;
  if (sessionError && !summaryData) return <PageContainer><ErrorMessage>Error: {sessionError.message || 'Could not load session summary.'}</ErrorMessage></PageContainer>;
  if (!summaryData) return <PageContainer><LoadingMessage>No summary data available.</LoadingMessage></PageContainer>;

  const { score, mistakes_made, mode, answered_questions_details } = summaryData;
  const totalAnswered = answered_questions_details?.length || 0;

  return (
    <PageContainer>
      <Title>Session Summary</Title>
      <SummaryDetails>
        <p><strong>Mode:</strong> {mode?.replace('_', ' ').split(' ').map(w => w.charAt(0).toUpperCase() + w.substring(1).toLowerCase()).join(' ')}</p>
        <p><strong>Final Score:</strong> {score}</p>
        <p><strong>Total Questions Answered:</strong> {totalAnswered}</p>
        {mode === 'mistakes' && <p><strong>Mistakes Made:</strong> {mistakes_made}</p>}
        {/* Add more details as available and relevant, e.g., time taken if tracked */}
      </SummaryDetails>

      {answered_questions_details && answered_questions_details.length > 0 && (
        <QuestionsReview>
          <h3>Questions Review</h3>
          {answered_questions_details.map((item, index) => (
            <QuestionItem key={index} isCorrect={item.is_correct}>
              <QuestionText>{index + 1}. {item.question_text}</QuestionText>
              <AnswerText>Your answer: {item.selected_option_text} {item.is_correct ? '(Correct)' : '(Incorrect)'}</AnswerText>
              {!item.is_correct && item.correct_option_text && (
                <AnswerText>Correct answer: {item.correct_option_text}</AnswerText>
              )}
              {item.explanation && <ExplanationText>Explanation: {item.explanation}</ExplanationText>}
            </QuestionItem>
          ))}
        </QuestionsReview>
      )}

      <StyledButton to='/'>Back to Home</StyledButton>
    </PageContainer>
  );
};

export default SummaryPage; 