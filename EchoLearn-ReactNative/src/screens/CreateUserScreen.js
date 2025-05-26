import React, { useState, useContext } from 'react';
import { View, Text, TextInput, Button, StyleSheet, Alert, ActivityIndicator } from 'react-native';
import { UserContext } from '../context/UserContext';
import * as userService from '../api/userService';

const CreateUserScreen = ({ navigation }) => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { login } = useContext(UserContext);

  const handleCreateUser = async () => {
    if (!name.trim()) {
      Alert.alert('Validation Error', 'Name cannot be empty.');
      return;
    }
    setIsLoading(true);
    try {
      const userData = { name, email: email.trim() || undefined }; // Send email only if provided
      const createdUser = await userService.createUser(userData);
      login(createdUser); // Update context and AsyncStorage
      // Navigation to HomeScreen is handled by AppNavigator due to user state change
    } catch (error) {
      Alert.alert('Creation Failed', error.message || 'Could not create user.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Welcome to EchoLearn!</Text>
      <Text style={styles.subtitle}>Let's get you started.</Text>
      <TextInput
        style={styles.input}
        placeholder="Enter your name (required)"
        value={name}
        onChangeText={setName}
      />
      <TextInput
        style={styles.input}
        placeholder="Enter your email (optional)"
        value={email}
        onChangeText={setEmail}
        keyboardType="email-address"
        autoCapitalize="none"
      />
      {isLoading ? (
        <ActivityIndicator size="large" color="#007bff" />
      ) : (
        <Button title="Create Profile & Start Learning" onPress={handleCreateUser} />
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
    backgroundColor: '#fff',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  subtitle: {
    fontSize: 16,
    color: 'gray',
    marginBottom: 30,
    textAlign: 'center',
  },
  input: {
    width: '100%',
    height: 50,
    borderColor: 'gray',
    borderWidth: 1,
    borderRadius: 8,
    marginBottom: 15,
    paddingHorizontal: 15,
    fontSize: 16,
  },
});

export default CreateUserScreen; 