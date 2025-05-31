import React, { useState, useEffect } from 'react';
import './App.css';
import DifficultyScreen from './components/DifficultyScreen';
import TopicScreen from './components/TopicScreen';
import QuestionScreen from './components/QuestionScreen';
import ResultScreen from './components/ResultScreen';
import LoginScreen from './components/LoginScreen';
import RegisterScreen from './components/RegisterScreen';
import UserScoresScreen from './components/UserScoresScreen';

type Screen = 'login' | 'register' | 'difficulty' | 'topic' | 'question' | 'result' | 'scores';

function App() {
  const [currentScreen, setCurrentScreen] = useState<Screen>('login');
  const [difficulty, setDifficulty] = useState<string>('');
  const [topic, setTopic] = useState<string>('');
  const [score, setScore] = useState(0);
  const [totalQuestions, setTotalQuestions] = useState(0);
  const [token, setToken] = useState<string | null>(localStorage.getItem('echotech_token'));
  const [userId, setUserId] = useState<string | null>(localStorage.getItem('echotech_userId'));
  const [userEmail, setUserEmail] = useState<string | null>(localStorage.getItem('echotech_userEmail'));

  useEffect(() => {
    if (token) {
      setCurrentScreen('difficulty'); // or 'scores' or a dashboard
    } else {
      setCurrentScreen('login');
    }
  }, [token]);

  const handleLoginSuccess = (newToken: string, newUserId: string, newUserEmail: string) => {
    localStorage.setItem('echotech_token', newToken);
    localStorage.setItem('echotech_userId', newUserId);
    localStorage.setItem('echotech_userEmail', newUserEmail);
    setToken(newToken);
    setUserId(newUserId);
    setUserEmail(newUserEmail);
    setCurrentScreen('difficulty');
  };

  const handleLogout = () => {
    localStorage.removeItem('echotech_token');
    localStorage.removeItem('echotech_userId');
    localStorage.removeItem('echotech_userEmail');
    setToken(null);
    setUserId(null);
    setUserEmail(null);
    setCurrentScreen('login');
    // Reset quiz state too
    setDifficulty('');
    setTopic('');
    setScore(0);
    setTotalQuestions(0);
  };

  const handleSelectDifficulty = (selectedDifficulty: string) => {
    setDifficulty(selectedDifficulty);
    setCurrentScreen('topic');
  };

  const handleSelectTopic = (selectedTopic: string) => {
    setTopic(selectedTopic);
    setCurrentScreen('question');
  };

  const handleQuizComplete = (finalScore: number, numQuestions: number) => {
    setScore(finalScore);
    setTotalQuestions(numQuestions);
    // Attempt to save score to backend if token exists
    if (token && difficulty && topic) {
      fetch('http://localhost:8080/api/v1/scores', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          topic: topic,
          difficulty: difficulty,
          points: finalScore,
          total: numQuestions,
        }),
      })
      .then(response => response.json())
      .then(data => {
        if (!data.error) {
          console.log('Score saved successfully:', data);
        } else {
          console.error('Failed to save score:', data.error);
        }
      })
      .catch(error => console.error('Error saving score:', error));
    }
    setCurrentScreen('result');
  };

  const handleRestartQuiz = () => {
    setDifficulty('');
    setTopic('');
    setScore(0);
    setTotalQuestions(0);
    setCurrentScreen('difficulty');
  };
  
  const handleNavigateToHome = () => {
    setCurrentScreen('difficulty');
  }

  const renderScreen = () => {
    switch (currentScreen) {
      case 'login':
        return <LoginScreen onLoginSuccess={handleLoginSuccess} onNavigateToRegister={() => setCurrentScreen('register')} />;
      case 'register':
        return <RegisterScreen onRegisterSuccess={() => setCurrentScreen('login')} onNavigateToLogin={() => setCurrentScreen('login')} />;
      case 'difficulty':
        return <DifficultyScreen onSelectDifficulty={handleSelectDifficulty} />;
      case 'topic':
        return <TopicScreen difficulty={difficulty} onSelectTopic={handleSelectTopic} onBack={() => setCurrentScreen('difficulty')} />;
      case 'question':
        return <QuestionScreen difficulty={difficulty} topic={topic} onQuizComplete={handleQuizComplete} token={token} />;
      case 'result':
        return <ResultScreen score={score} totalQuestions={totalQuestions} onRestart={handleRestartQuiz} />;
      case 'scores':
        return <UserScoresScreen token={token} onNavigateToQuiz={handleNavigateToHome}/>;
      default:
        return <LoginScreen onLoginSuccess={handleLoginSuccess} onNavigateToRegister={() => setCurrentScreen('register')} />;
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <h1>EchoTech Quiz</h1>
        {userEmail && <p className="user-greeting">Welcome, {userEmail}!</p>}
        {token && (
          <nav className="app-nav">
            <button onClick={handleNavigateToHome}>Quiz Home</button>
            <button onClick={() => setCurrentScreen('scores')}>My Scores</button>
            <button onClick={handleLogout}>Logout</button>
          </nav>
        )}
      </header>
      <main>
        {renderScreen()}
      </main>
      <footer>
        <p>&copy; {new Date().getFullYear()} EchoTech Learning. All rights reserved.</p>
      </footer>
    </div>
  );
}

export default App;
