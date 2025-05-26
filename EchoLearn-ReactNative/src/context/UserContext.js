import React, { createContext, useState, useEffect } from 'react';
import AsyncStorage from '@react-native-async-storage/async-storage';
import * as userService from '../api/userService';

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const loadUser = async () => {
      setIsLoading(true);
      try {
        const storedUserId = await AsyncStorage.getItem('userId');
        if (storedUserId) {
          const fetchedUser = await userService.getUserById(storedUserId);
          setUser(fetchedUser);
        } else {
          setUser(null); // Explicitly set to null if no stored user
        }
      } catch (error) {
        console.error('Failed to load user from storage or API', error);
        setUser(null); // Ensure user is null on error
        await AsyncStorage.removeItem('userId'); // Clear inconsistent stored ID
      }
      setIsLoading(false);
    };
    loadUser();
  }, []);

  const login = async (createdUser) => {
    setUser(createdUser);
    await AsyncStorage.setItem('userId', createdUser.id);
  };

  const logout = async () => {
    setUser(null);
    await AsyncStorage.removeItem('userId');
    // Potentially clear session context as well
  };

  return (
    <UserContext.Provider value={{ user, setUser, login, logout, isLoadingUser: isLoading }}>
      {children}
    </UserContext.Provider>
  );
}; 