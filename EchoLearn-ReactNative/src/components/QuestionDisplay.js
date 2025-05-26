import React from 'react';
import { View, Text, StyleSheet } from 'react-native';

const QuestionDisplay = ({ question }) => {
  if (!question) {
    return <Text style={styles.loadingText}>Loading question...</Text>;
  }

  return (
    <View style={styles.container}>
      <Text style={styles.questionText}>{question.text}</Text>
      {question.grammar_rule && (
        <View style={styles.grammarRuleContainer}>
          <Text style={styles.grammarTitle}>{question.grammar_rule.title}</Text>
          {/* <Text style={styles.grammarDescription}>{question.grammar_rule.description}</Text> */}
          {/* Examples could be shown on feedback or summary */}
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    marginBottom: 20,
    padding: 15,
    backgroundColor: '#f9f9f9',
    borderRadius: 8,
  },
  questionText: {
    fontSize: 18,
    fontWeight: 'bold',
    marginBottom: 10,
    color: '#333',
  },
  grammarRuleContainer: {
    marginTop: 10,
    paddingTop: 10,
    borderTopWidth: 1,
    borderTopColor: '#eee',
  },
  grammarTitle: {
    fontSize: 16,
    fontWeight: '500',
    color: '#555',
  },
  grammarDescription: {
    fontSize: 14,
    color: '#666',
    marginTop: 5,
  },
  loadingText: {
    fontSize: 18,
    textAlign: 'center',
    marginTop: 20,
  },
});

export default QuestionDisplay; 