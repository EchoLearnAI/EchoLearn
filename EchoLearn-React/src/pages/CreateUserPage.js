import React, { useState, useContext } from 'react';
import { useNavigate } from 'react-router-dom';
import styled from 'styled-components';
import UserContext from '../contexts/UserContext';
import { createUser as apiCreateUser } from '../api/userService';

const PageContainer = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  min-height: 80vh; // Take up most of the viewport height
  padding: ${props => props.theme.spacings.medium};
`;

const FormContainer = styled.form`
  display: flex;
  flex-direction: column;
  gap: ${props => props.theme.spacings.medium};
  padding: ${props => props.theme.spacings.large};
  background-color: ${props => props.theme.colors.light};
  border-radius: ${props => props.theme.radii.medium};
  box-shadow: ${props => props.theme.shadows.medium};
  width: 100%;
  max-width: 400px; // Max width for the form
`;

const Title = styled.h1`
  text-align: center;
  color: ${props => props.theme.colors.primary};
  margin-bottom: ${props => props.theme.spacings.large};
`;

const Input = styled.input`
  padding: ${props => props.theme.spacings.small} ${props => props.theme.spacings.medium};
  border: 1px solid ${props => props.theme.colors.border};
  border-radius: ${props => props.theme.radii.small};
  font-size: ${props => props.theme.fontSizes.medium};

  &:focus {
    outline: none;
    border-color: ${props => props.theme.colors.primary};
    box-shadow: 0 0 0 2px ${props => props.theme.colors.primary}30; // Light shadow on focus
  }
`;

const Button = styled.button`
  padding: ${props => props.theme.spacings.medium};
  background-color: ${props => props.theme.colors.primary};
  color: ${props => props.theme.colors.white};
  border: none;
  border-radius: ${props => props.theme.radii.small};
  font-size: ${props => props.theme.fontSizes.medium};
  cursor: pointer;
  transition: background-color 0.2s ease-in-out;

  &:hover {
    background-color: ${props => props.theme.colors.primary}d0; // Darken on hover
  }

  &:disabled {
    background-color: ${props => props.theme.colors.disabled};
    cursor: not-allowed;
  }
`;

const ErrorMessage = styled.p`
  color: ${props => props.theme.colors.danger};
  font-size: ${props => props.theme.fontSizes.small};
  text-align: center;
  margin-top: ${props => props.theme.spacings.small};
`;

const CreateUserPage = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const { login } = useContext(UserContext);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);
    setIsLoading(true);

    if (!name.trim() || !email.trim()) {
      setError('Name and Email are required.');
      setIsLoading(false);
      return;
    }
    // Basic email validation
    if (!/\S+@\S+\.\S+/.test(email)) {
        setError('Please enter a valid email address.');
        setIsLoading(false);
        return;
    }

    try {
      const newUser = await apiCreateUser({ name, email });
      login(newUser.data || newUser); // Assuming API returns user object in {data: ...} or directly
      navigate('/'); // Navigate to home page after successful creation/login
    } catch (err) {
      console.error('Create user error:', err);
      setError(err.message || 'Failed to create user. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <PageContainer>
      <FormContainer onSubmit={handleSubmit}>
        <Title>Create Your Account</Title>
        <Input
          type="text"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          disabled={isLoading}
          required
        />
        <Input
          type="email"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          disabled={isLoading}
          required
        />
        <Button type="submit" disabled={isLoading}>
          {isLoading ? 'Creating...' : 'Create Account'}
        </Button>
        {error && <ErrorMessage>{error}</ErrorMessage>}
      </FormContainer>
    </PageContainer>
  );
};

export default CreateUserPage; 