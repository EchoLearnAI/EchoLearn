import React from 'react';
import { UserProvider } from './context/UserContext';
import { SessionProvider } from './context/SessionContext';
import AppNavigator from './navigation/AppNavigator';
import { GestureHandlerRootView } from 'react-native-gesture-handler'; // Required for React Navigation
import { SafeAreaProvider } from 'react-native-safe-area-context';

const App = () => {
  return (
    <GestureHandlerRootView style={{ flex: 1 }}>
      <SafeAreaProvider>
        <UserProvider>
          <SessionProvider>
            <AppNavigator />
          </SessionProvider>
        </UserProvider>
      </SafeAreaProvider>
    </GestureHandlerRootView>
  );
};

export default App; 