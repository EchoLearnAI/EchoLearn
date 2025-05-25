import React, { useEffect, useState, useContext } from 'react';
import { View, Text, StyleSheet, ActivityIndicator, Button, ScrollView, FlatList } from 'react-native';
import { SessionContext } from '../context/SessionContext';

const SummaryScreen = ({ route, navigation }) => {
  const { sessionId } = route.params;
  const { fetchSessionSummary, isLoadingSession, sessionError, endCurrentSession } = useContext(SessionContext);
  const [summaryData, setSummaryData] = useState(null);

  useEffect(() => {
    const loadSummary = async () => {
      const data = await fetchSessionSummary(sessionId);
      if (data) {
        setSummaryData(data);
      }
    };
    loadSummary();
    
    // Ensure local session state is cleared if this was the current session
    // and it has been fetched and displayed.
    return () => {
        // Check if the session being displayed was the one active in context
        // This logic might need refinement based on how currentSession is cleared in SessionContext
        endCurrentSession(); 
    };
  }, [sessionId, fetchSessionSummary, endCurrentSession]);

  if (isLoadingSession || !summaryData) {
    return (
      <View style={styles.centered}>
        <ActivityIndicator size="large" color="#007bff" />
        <Text>Loading summary...</Text>
      </View>
    );
  }

  if (sessionError && !summaryData) {
    return (
      <View style={styles.centered}>
        <Text style={styles.errorText}>Error loading summary: {sessionError}</Text>
        <Button title="Go Home" onPress={() => navigation.replace('Home')} />
      </View>
    );
  }

  const { session, answered_questions } = summaryData;

  const renderAnswerDetail = ({ item }) => (
    <View style={[styles.answerItem, !item.is_correct && styles.incorrectAnswerBackground]}>
      <Text style={styles.questionTextSmall}>{item.question_text}</Text>
      <Text>Your answer was: <Text style={item.is_correct ? styles.correctText : styles.incorrectText}>{item.is_correct ? 'Correct' : 'Incorrect'}</Text></Text>
      {!item.is_correct && item.explanation && (
        <Text style={styles.explanationSmall}>Explanation: {item.explanation}</Text>
      )}
    </View>
  );

  return (
    <ScrollView contentContainerStyle={styles.container}>
      <Text style={styles.title}>Quiz Summary</Text>
      <View style={styles.summaryBox}>
        <Text style={styles.summaryText}>Mode: {session.mode}</Text>
        <Text style={styles.summaryText}>Final Score: {session.score}</Text>
        <Text style={styles.summaryText}>Mistakes Made: {session.mistakes_made}</Text>
        {session.mode === 'five_topic' && <Text style={styles.summaryText}>Questions Answered: {session.current_question_index} / {session.total_questions}</Text>}
      </View>

      <Text style={styles.breakdownTitle}>Answer Breakdown:</Text>
      {answered_questions && answered_questions.length > 0 ? (
        <FlatList
          data={answered_questions}
          renderItem={renderAnswerDetail}
          keyExtractor={(item, index) => (item.question_id ? item.question_id.toString() : 'q_') + index.toString()} // Ensure key is a string and unique
          style={styles.list}
        />
      ) : (
        <Text>No answers recorded for this session.</Text>
      )}

      <View style={styles.buttonContainer}>
        <Button title="Play Again / Go Home" onPress={() => navigation.replace('Home')} />
      </View>
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    padding: 20,
    backgroundColor: '#f5f5f5',
  },
  centered: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    textAlign: 'center',
    marginBottom: 20,
  },
  summaryBox: {
    backgroundColor: '#fff',
    padding: 20,
    borderRadius: 8,
    marginBottom: 20,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.23,
    shadowRadius: 2.62,
    elevation: 4,
  },
  summaryText: {
    fontSize: 18,
    marginBottom: 8,
  },
  breakdownTitle: {
    fontSize: 20,
    fontWeight: 'bold',
    marginBottom: 10,
  },
  list: {
    // Adjust max height as needed, or remove if full scroll is desired
  },
  answerItem: {
    backgroundColor: '#fff',
    padding: 15,
    borderRadius: 5,
    marginBottom: 10,
    borderWidth: 1,
    borderColor: '#eee',
  },
  incorrectAnswerBackground: {
    backgroundColor: '#ffebee', // Light red for incorrect answers
  },
  questionTextSmall: {
    fontSize: 15,
    fontWeight: '500',
    marginBottom: 5,
  },
  correctText: {
    color: 'green',
    fontWeight: 'bold',
  },
  incorrectText: {
    color: 'red',
    fontWeight: 'bold',
  },
  explanationSmall: {
    fontSize: 13,
    color: '#555',
    marginTop: 5,
  },
  errorText: {
    color: 'red',
    fontSize: 16,
    textAlign: 'center',
  },
  buttonContainer: {
      marginTop: 30,
  }
});

export default SummaryScreen; 