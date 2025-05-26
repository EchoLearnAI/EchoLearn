import React, { useState, useContext } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import styled from 'styled-components';
import UserContext from '../contexts/UserContext';
import { loginUser } from '../api/authService'; // We'll create this service and function

const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: ${({ theme }) => theme.spacings.large};
  background-color: ${({ theme }) => theme.colors.background};
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
  gap: ${({ theme }) => theme.spacings.medium};
  width: 100%;
  max-width: 400px;
  padding: ${({ theme }) => theme.spacings.large};
  background-color: ${({ theme }) => theme.colors.surface};
  border-radius: ${({ theme }) => theme.radii.medium};
  box-shadow: ${({ theme }) => theme.shadows.medium};
`;

const Title = styled.h1`
  color: ${({ theme }) => theme.colors.primary};
  text-align: center;
  margin-bottom: ${({ theme }) => theme.spacings.medium};
`;

const Input = styled.input`
  padding: ${({ theme }) => theme.spacings.small} ${({ theme }) => theme.spacings.medium};
  border: 1px solid ${({ theme }) => theme.colors.border};
  border-radius: ${({ theme }) => theme.radii.small};
  font-size: ${({ theme }) => theme.fontSizes.medium};
  &:focus {
    outline: none;
    border-color: ${({ theme }) => theme.colors.primary};
  }
`;

const Button = styled.button`
  padding: ${({ theme }) => theme.spacings.medium};
  background-color: ${({ theme }) => theme.colors.primary};
  color: ${({ theme }) => theme.colors.onPrimary};
  border: none;
  border-radius: ${({ theme }) => theme.radii.small};
  font-size: ${({ theme }) => theme.fontSizes.medium};
  cursor: pointer;
  transition: background-color 0.3s ease;

  &:hover {
    background-color: ${({ theme }) => theme.colors.primaryVariant};
  }

  &:disabled {
    background-color: ${({ theme }) => theme.colors.disabled};
    cursor: not-allowed;
  }
`;

const ErrorMessage = styled.p`
  color: ${({ theme }) => theme.colors.error};
  text-align: center;
  font-size: ${({ theme }) => theme.fontSizes.small};
`;

const LinkText = styled.p`
  text-align: center;
  margin-top: ${({ theme }) => theme.spacings.medium};
  font-size: ${({ theme }) => theme.fontSizes.small};
  color: ${({ theme }) => theme.colors.textSecondary};

  a {
    color: ${({ theme }) => theme.colors.primary};
    text-decoration: none;
    &:hover {
      text-decoration: underline;
    }
  }
`;

function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { login } = useContext(UserContext);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    if (!email.trim() || !password.trim()) {
      setError('Email and Password are required.');
      setIsLoading(false);
      return;
    }

    try {
      const user = await loginUser({ email, password });
      if (user && user.id) { // Assuming API returns user object directly
        login(user); // UserContext.login expects the user object
        navigate('/'); // Navigate to home or dashboard
      } else {
        setError('Login failed. Please check your credentials.');
      }
    } catch (err) {
      const errorMessage = err.response?.data?.message || err.response?.data?.details || err.message || 'An unexpected error occurred.';
      setError(`Login failed: ${errorMessage}`);
      console.error('Login error:', err.response || err);
    }
    setIsLoading(false);
  };

  return (
    <PageContainer>
      <Form onSubmit={handleSubmit}>
        <Title>Login to EchoLearn</Title>
        {error && <ErrorMessage>{error}</ErrorMessage>}
        <Input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          disabled={isLoading}
        />
        <Input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          disabled={isLoading}
        />
        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Logging In...' : 'Login'}
        </Button>
        <LinkText>
          Don't have an account? <Link to="/create-user">Create one</Link>
        </LinkText>
      </Form>
    </PageContainer>
  );
}

export default LoginPage; 