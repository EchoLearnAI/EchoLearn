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
        const fetchedUser = await getUserById(storedUserId);
        setUser(fetchedUser.data || fetchedUser); // Adjust based on your API response structure
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
    // userData could be the full user object from createUser or just { ID: ... }
    setUser(userData);
    localStorage.setItem('echolearn_userId', userData.ID);
    setUserError(null);
  }, []);

  const logout = useCallback(() => {
    setUser(null);
    localStorage.removeItem('echolearn_userId');
    // Potentially clear session context as well, or handle that in SessionProvider
    setUserError(null);
    // Maybe redirect to login page: navigate('/'); or handle in component
  }, []);

  // Function to refresh user data if needed
  const refreshUser = useCallback(async () => {
    if (user && user.ID) {
      setIsLoadingUser(true);
      try {
        const fetchedUser = await getUserById(user.ID);
        setUser(fetchedUser.data || fetchedUser);
        setUserError(null);
      } catch (error) {
        console.error('Failed to refresh user:', error);
        setUserError(error);
        // Potentially log out user if refresh fails critically
      } finally {
        setIsLoadingUser(false);
      }
    }
  }, [user]);

  return (
    <UserContext.Provider value={{ user, isLoadingUser, userError, login, logout, refreshUser, loadUserFromStorage }}>
      {children}
    </UserContext.Provider>
  );
};

export default UserContext; 