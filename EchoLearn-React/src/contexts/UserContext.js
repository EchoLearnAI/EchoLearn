import React, { createContext, useState, useEffect, useCallback } from 'react';
import { getUserById } from '../api/userService'; // Assuming userService is in src/api

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [isLoadingUser, setIsLoadingUser] = useState(true);
  const [userError, setUserError] = useState(null);

  const loadUserFromStorage = useCallback(async () => {
    setIsLoadingUser(true);
    setUserError(null);
    try {
      const storedUserId = localStorage.getItem('echolearn_userId');
      if (storedUserId) {
        const fetchedUserResponse = await getUserById(storedUserId);
        // API returns { data: userObject }, so we use fetchedUserResponse.data
        if (fetchedUserResponse && fetchedUserResponse.data) {
          setUser(fetchedUserResponse.data);
        } else {
          setUser(null); // Or handle as an error if data is expected
          console.warn('User data not found in response after fetching by ID');
          localStorage.removeItem('echolearn_userId');
        }
      } else {
        setUser(null);
      }
    } catch (error) {
      console.error('Failed to load user from storage or API:', error);
      setUserError(error);
      setUser(null);
      localStorage.removeItem('echolearn_userId'); // Clear invalid/stale ID
    } finally {
      setIsLoadingUser(false);
    }
  }, []);

  useEffect(() => {
    loadUserFromStorage();
  }, [loadUserFromStorage]);

  const login = useCallback((userData) => {
    // userData is expected to be the user object { id: ..., name: ... }
    setUser(userData);
    if (userData && userData.id) {
      localStorage.setItem('echolearn_userId', userData.id);
    } else {
      console.warn('Login called with userData missing an id:', userData);
    }
    setUserError(null);
  }, []);

  const logout = useCallback(() => {
    setUser(null);
    localStorage.removeItem('echolearn_userId');
    setUserError(null);
  }, []);

  const refreshUser = useCallback(async () => {
    if (user && user.id) { // Use lowercase 'id'
      setIsLoadingUser(true);
      try {
        const fetchedUserResponse = await getUserById(user.id); // Use lowercase 'id'
        // API returns { data: userObject }, so we use fetchedUserResponse.data
        if (fetchedUserResponse && fetchedUserResponse.data) {
          setUser(fetchedUserResponse.data);
        } else {
          console.warn('User data not found in response during refreshUser');
          // Optionally logout or set error if user becomes invalid
          setUser(null); 
          localStorage.removeItem('echolearn_userId');
        }
        setUserError(null);
      } catch (error) {
        console.error('Failed to refresh user:', error);
        setUserError(error);
      } finally {
        setIsLoadingUser(false);
      }
    } else if (user) { // User exists but user.id is missing for some reason
        console.warn("RefreshUser called when user object is missing an id:", user);
    }
  }, [user]);

  return (
    <UserContext.Provider value={{ user, isLoadingUser, userError, login, logout, refreshUser, loadUserFromStorage }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext; 