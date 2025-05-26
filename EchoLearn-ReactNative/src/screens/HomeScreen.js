import React, { useContext } from 'react';
import { View, Text, Button, StyleSheet, Alert } from 'react-native';
import { UserContext } from '../context/UserContext';
import { SessionContext } from '../context/SessionContext';

const HomeScreen = ({ navigation }) => {
  const { user, logout } = useContext(UserContext);
  const { startNewSession, isLoadingSession, sessionError } = useContext(SessionContext);

  const handleStartQuiz = async (mode) => {
    const session = await startNewSession(mode);
    if (session) {
      navigation.navigate('Quiz', { gameMode: mode });
    } else {
      Alert.alert('Error', sessionError || 'Could not start quiz session.');
    }
  };

  const handleLogout = () => {
    logout();
    // Navigation to CreateUserScreen will be handled by AppNavigator
  };

  return (
    <View style={styles.container}>
      <Text style={styles.welcomeText}>Welcome, {user?.name || 'User'}!</Text>
      <Text style={styles.subtitle}>Choose a game mode to start learning:</Text>
      
      <View style={styles.buttonContainer}>
        <Button 
          title="Survival Mode" 
          onPress={() => handleStartQuiz('survival')} 
          disabled={isLoadingSession} 
        />
      </View>
      <View style={styles.buttonContainer}>
        <Button 
          title="5-Topic Mode" 
          onPress={() => handleStartQuiz('five_topic')} 
          disabled={isLoadingSession} 
        />
      </View>
      <View style={styles.buttonContainer}>
        <Button 
          title="Infinite Mode" 
          onPress={() => handleStartQuiz('infinite')} 
          disabled={isLoadingSession} 
        />
      </View>

      {isLoadingSession && <Text style={styles.loadingText}>Starting session...</Text>}
      {sessionError && <Text style={styles.errorText}>{sessionError}</Text>}

      <View style={[styles.buttonContainer, styles.logoutButton]}>
        <Button title="Logout" color="red" onPress={handleLogout} />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
    backgroundColor: '#f5f5f5',
  },
  welcomeText: {
    fontSize: 22,
    fontWeight: 'bold',
    marginBottom: 10,
    textAlign: 'center',
  },
  subtitle: {
    fontSize: 16,
    color: 'gray',
    marginBottom: 30,
    textAlign: 'center',
  },
  buttonContainer: {
    width: '80%',
    marginVertical: 10,
  },
  loadingText: {
    marginTop: 10,
    color: 'blue',
  },
  errorText: {
    marginTop: 10,
    color: 'red',
  },
  logoutButton: {
    marginTop: 40,
  }
});

export default HomeScreen; 