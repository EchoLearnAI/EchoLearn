import React, { useContext } from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import { UserContext } from '../context/UserContext';

import CreateUserScreen from '../screens/CreateUserScreen';
import HomeScreen from '../screens/HomeScreen';
import QuizScreen from '../screens/QuizScreen';
import SummaryScreen from '../screens/SummaryScreen';
import { View, ActivityIndicator, StyleSheet } from 'react-native';

const Stack = createStackNavigator();

const AppNavigator = () => {
  const { user, isLoadingUser } = useContext(UserContext);

  if (isLoadingUser) {
    return (
      <View style={styles.loaderContainer}>
        <ActivityIndicator size="large" color="#0000ff" />
      </View>
    );
  }

  return (
    <NavigationContainer>
      <Stack.Navigator initialRouteName={user ? "Home" : "CreateUser"}>
        {user ? (
          <>
            <Stack.Screen name="Home" component={HomeScreen} options={{ title: 'EchoLearn' }} />
            <Stack.Screen name="Quiz" component={QuizScreen} options={{ title: 'Quiz' }} />
            <Stack.Screen name="Summary" component={SummaryScreen} options={{ title: 'Quiz Summary' }} />
          </>
        ) : (
          <Stack.Screen name="CreateUser" component={CreateUserScreen} options={{ title: 'Welcome' }} />
        )}
      </Stack.Navigator>
    </NavigationContainer>
  );
};

const styles = StyleSheet.create({
  loaderContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
});

export default AppNavigator; 