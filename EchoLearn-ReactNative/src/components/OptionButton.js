import React from 'react';
import { TouchableOpacity, Text, StyleSheet } from 'react-native';

const OptionButton = ({ option, onPress, disabled, style }) => {
  return (
    <TouchableOpacity
      style={[styles.button, style]} // Allow custom styles to be passed
      onPress={() => onPress(option.id)}
      disabled={disabled}>
      <Text style={styles.buttonText}>{option.label ? `${option.label}. ${option.text}` : option.text}</Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    backgroundColor: '#DDDDDD',
    padding: 15,
    borderRadius: 8,
    marginVertical: 8,
    alignItems: 'flex-start',
    width: '100%',
  },
  buttonText: {
    fontSize: 16,
    color: '#333',
  },
});

export default OptionButton; 