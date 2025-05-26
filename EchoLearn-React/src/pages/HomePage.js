import React, { useContext } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import UserContext from '../contexts/UserContext';
import SessionContext from '../contexts/SessionContext';

const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: ${props => props.theme.spacings.large};
  text-align: center;
`;

const Title = styled.h1`
  color: ${props => props.theme.colors.primary};
  margin-bottom: ${props => props.theme.spacings.medium};
`;

const WelcomeMessage = styled.p`
  font-size: ${props => props.theme.fontSizes.large};
  margin-bottom: ${props => props.theme.spacings.large};
`;

const ButtonGroup = styled.div`
  display: flex;
  flex-direction: column; // Stack buttons vertically on mobile
  gap: ${props => props.theme.spacings.medium};
  margin-bottom: ${props => props.theme.spacings.large};

  ${props => props.theme.media.sm`
    flex-direction: row; // Buttons side-by-side on larger screens
  `}
`;

const StyledButton = styled.button`
  padding: ${props => props.theme.spacings.medium} ${props => props.theme.spacings.large};
  background-color: ${props => props.theme.colors.primary};
  color: ${props => props.theme.colors.white};
  border: none;
  border-radius: ${props => props.theme.radii.medium};
  font-size: ${props => props.theme.fontSizes.medium};
  cursor: pointer;
  transition: background-color 0.2s ease-in-out;
  min-width: 200px; // Ensure buttons have a decent width

  &:hover {
    background-color: ${props => props.theme.colors.primary}d0;
  }

  &:disabled {
    background-color: ${props => props.theme.colors.disabled};
    cursor: not-allowed;
  }
`;

const LogoutButton = styled(StyledButton)`
  background-color: ${props => props.theme.colors.secondary};
  &:hover {
    background-color: ${props => props.theme.colors.secondary}d0;
  }
`;

const AuthLink = styled(Link)`
  font-size: ${props => props.theme.fontSizes.medium};
  color: ${props => props.theme.colors.primary};
  text-decoration: underline;
  margin-top: ${props => props.theme.spacings.large};
`;

const HomePage = () => {
  const { user, logout, isLoadingUser } = useContext(UserContext);
  const { startNewSession, isLoadingSession } = useContext(SessionContext);
  const navigate = useNavigate();

  const handleStartSession = async (mode) => {
    let sessionParams = { mode };
    if (mode === 'mistakes') {
      sessionParams.maxMistakes = 5; // Example default
      sessionParams.totalQuestions = 10; // Example default
    } else if (mode === 'category_challenge') {
      // For now, we don't have category selection UI here. 
      // This would need to be expanded, or category chosen on a subsequent screen.
      // Defaulting to a placeholder category or a random one if backend supports
      sessionParams.categoryName = 'Verb Tenses'; // Placeholder
      sessionParams.totalQuestions = 10;
    } else if (mode === 'infinite') {
      // No specific params needed beyond mode for infinite
    }

    const newSession = await startNewSession(sessionParams);
    if (newSession) {
      navigate('/quiz');
    }
    // Errors are handled in SessionContext and could be displayed via a global notification system or here
  };

  if (isLoadingUser) {
    return <PageContainer><p>Loading user...</p></PageContainer>;
  }

  return (
    <PageContainer>
      <Title>Welcome to EchoLearn!</Title>
      {user ? (
        <>
          <WelcomeMessage>Hello, {user.name || 'User'}!</WelcomeMessage>
          <p>Choose a game mode to start learning:</p>
          <ButtonGroup>
            <StyledButton onClick={() => handleStartSession('mistakes')} disabled={isLoadingSession}>
              {isLoadingSession ? 'Starting...' : 'Mistakes Mode'}
            </StyledButton>
            <StyledButton onClick={() => handleStartSession('infinite')} disabled={isLoadingSession}>
              {isLoadingSession ? 'Starting...' : 'Infinite Mode'}
            </StyledButton>
            <StyledButton onClick={() => handleStartSession('category_challenge')} disabled={isLoadingSession}>
              {isLoadingSession ? 'Starting...' : 'Category Challenge'}
            </StyledButton>
          </ButtonGroup>
          <LogoutButton onClick={logout} disabled={isLoadingSession}>
            Logout
          </LogoutButton>
        </>
      ) : (
        <>
          <WelcomeMessage>Please log in or create an account to start.</WelcomeMessage>
          <ButtonGroup>
            <StyledButton as={Link} to="/create-user">
              Create Account
            </StyledButton>
            {/* <StyledButton as={Link} to="/login">Login</StyledButton> */}
            {/* We don't have a separate login page yet, assuming creation implies login for now */}
          </ButtonGroup>
          
        </>
      )}
    </PageContainer>
  );
};

export default HomePage; 