import React, { useState, useEffect, useContext } from 'react';
import { View, Text, StyleSheet, ActivityIndicator, Alert, Button, ScrollView } from 'react-native';
import { SessionContext } from '../context/SessionContext';
import * as questionService from '../api/questionService';
import QuestionDisplay from '../components/QuestionDisplay';
import OptionButton from '../components/OptionButton';

const QuizScreen = ({ route, navigation }) => {
  const { gameMode } = route.params;
  const {
    currentSession,
    submitUserAnswer,
    isLoadingSession,
    sessionError,
    endCurrentSession,
    fetchSessionSummary
  } = useContext(SessionContext);

  const [question, setQuestion] = useState(null);
  const [isLoadingQuestion, setIsLoadingQuestion] = useState(false);
  const [feedback, setFeedback] = useState({ message: '', isCorrect: false, correctOptionId: null, explanation: '' });
  const [answered, setAnswered] = useState(false);

  const fetchNextQuestion = async () => {
    setIsLoadingQuestion(true);
    setFeedback({ message: '', isCorrect: false, correctOptionId: null, explanation: '' });
    setAnswered(false);
    try {
      // For MVP, always fetch a random question. 
      // 5-Topic mode might need specific category logic if not handled by backend session state.
      const nextQuestion = await questionService.getRandomQuestion();
      setQuestion(nextQuestion);
    } catch (error) {
      Alert.alert('Error', 'Could not fetch next question.');
      console.error("Fetch question error", error)
    } finally {
      setIsLoadingQuestion(false);
    }
  };

  useEffect(() => {
    if (currentSession) {
      fetchNextQuestion();
    }
  }, [currentSession]); // Fetch first question when session is ready

  const handleAnswerSubmit = async (selectedOptionId) => {
    if (!question || !currentSession) return;
    setAnswered(true);

    const result = await submitUserAnswer(question.id, selectedOptionId);

    if (result) {
      setFeedback({
        message: result.result ? 'Correct!' : 'Incorrect!',
        isCorrect: result.result,
        correctOptionId: result.correct_option_id,
        explanation: result.explanation || question.options.find(o => o.id === result.correct_option_id)?.explanation
      });

      if (!result.session_active) {
        // Session ended (e.g., 3 mistakes in survival, or 50q in 5-topic)
        Alert.alert(
          'Session Ended',
          result.result ? 'Correct! Session Over.' : 'Incorrect! Session Over.',
          [{ text: 'View Summary', onPress: () => navigation.replace('Summary', { sessionId: currentSession.id }) }]
        );
      }
    } else {
      Alert.alert('Error', sessionError || 'Could not submit answer.');
      setAnswered(false); // Allow retry if submission failed for some reason
    }
  };

  const handleNextQuestion = () => {
    if (currentSession && currentSession.is_active) {
      fetchNextQuestion();
    } else {
      // Should have been navigated away by session end logic, but as a fallback:
      navigation.replace('Summary', { sessionId: currentSession?.id });
    }
  };
  
  const handleFinishInfiniteMode = async () => {
      if (currentSession && currentSession.mode === 'infinite') {
          endCurrentSession(); // This will also call summary on backend implicitly to end it
          navigation.replace('Summary', { sessionId: currentSession.id });
      }
  };

  if (!currentSession) {
    return <View style={styles.container}><Text>Loading session...</Text></View>;
  }

  if (isLoadingQuestion || isLoadingSession) {
    return <View style={styles.container}><ActivityIndicator size="large" color="#007bff" /></View>;
  }

  return (
    <ScrollView contentContainerStyle={styles.container}>
      <Text style={styles.modeText}>Mode: {gameMode}</Text>
      {currentSession && (
          <Text style={styles.scoreText}>
              Score: {currentSession.score} | Mistakes: {currentSession.mistakes_made}
              {gameMode === 'survival' && ` / ${currentSession.max_mistakes}`}
              {gameMode === 'five_topic' && ` | Question: ${currentSession.current_question_index} / ${currentSession.total_questions}`}
          </Text>
      )}

      <QuestionDisplay question={question} />

      {question && question.options.map((opt) => (
        <OptionButton
          key={opt.id}
          option={opt}
          onPress={handleAnswerSubmit}
          disabled={answered || isLoadingSession}
          style={[
            answered && opt.id === feedback.correctOptionId && styles.correctOption,
            answered && !feedback.isCorrect && opt.id === question.options.find(o => o.is_correct)?.id && styles.correctOption, // Highlight correct if user was wrong
            answered && !feedback.isCorrect && opt.id !== feedback.correctOptionId && opt.id === question.options.find(o => o.id === feedback.selectedOptionId) && styles.incorrectOption, // Highlight user's wrong selection
          ]}
        />
      ))}

      {answered && (
        <View style={styles.feedbackContainer}>
          <Text style={feedback.isCorrect ? styles.feedbackCorrect : styles.feedbackIncorrect}>
            {feedback.message}
          </Text>
          {feedback.explanation && <Text style={styles.explanationText}>{feedback.explanation}</Text>}
          {currentSession && currentSession.is_active && (
            <Button title="Next Question" onPress={handleNextQuestion} />
          )}
        </View>
      )}

      {currentSession && currentSession.mode === 'infinite' && currentSession.is_active && !answered && (
          <Button title="Finish Quiz" onPress={handleFinishInfiniteMode} color="#c00" />
      )}
      {sessionError && <Text style={styles.errorText}>{sessionError}</Text>}
    </ScrollView>
  );
};

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    padding: 20,
    backgroundColor: '#fff',
  },
  modeText: {
    fontSize: 18,
    fontWeight: 'bold',
    textAlign: 'center',
    marginBottom: 5,
  },
  scoreText: {
    fontSize: 16,
    textAlign: 'center',
    marginBottom: 20,
    color: '#555',
  },
  feedbackContainer: {
    marginTop: 20,
    padding: 10,
    borderRadius: 5,
    alignItems: 'center',
  },
  feedbackCorrect: {
    fontSize: 18,
    color: 'green',
    fontWeight: 'bold',
    marginBottom: 5,
  },
  feedbackIncorrect: {
    fontSize: 18,
    color: 'red',
    fontWeight: 'bold',
    marginBottom: 5,
  },
  explanationText: {
    fontSize: 15,
    color: '#333',
    textAlign: 'center',
    marginBottom: 15,
  },
  correctOption: {
    backgroundColor: '#c8e6c9', // Light green
    borderColor: 'green',
    borderWidth: 1,
  },
  incorrectOption: {
    backgroundColor: '#ffcdd2', // Light red
    borderColor: 'red',
    borderWidth: 1,
  },
  errorText: {
    color: 'red',
    textAlign: 'center',
    marginTop: 10,
  }
});

export default QuizScreen; 