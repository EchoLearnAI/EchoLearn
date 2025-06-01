import React, { useState, useEffect } from 'react';
import './App.css';
import { Category, Topic } from './types';
import CategoryScreen from './components/CategoryScreen';
import DifficultyScreen from './components/DifficultyScreen';
import TopicScreen from './components/TopicScreen';
import QuestionScreen from './components/QuestionScreen';
import ResultScreen from './components/ResultScreen';
import LoginScreen from './components/LoginScreen';
import RegisterScreen from './components/RegisterScreen';
import UserScoresScreen from './components/UserScoresScreen';

type Screen = 'login' | 'register' | 'category' | 'topic' | 'difficulty' | 'question' | 'result' | 'scores';

function App() {
  const [currentScreen, setCurrentScreen] = useState<Screen>('login');
  const [selectedCategory, setSelectedCategory] = useState<Category | null>(null);
  const [selectedTopic, setSelectedTopic] = useState<Topic | null>(null);
  const [selectedDifficulty, setSelectedDifficulty] = useState<string>('');
  const [score, setScore] = useState(0);
  const [totalQuestions, setTotalQuestions] = useState(0);
  const [token, setToken] = useState<string | null>(localStorage.getItem('echotech_token'));
  const [userId, setUserId] = useState<string | null>(localStorage.getItem('echotech_userId'));
  const [userEmail, setUserEmail] = useState<string | null>(localStorage.getItem('echotech_userEmail'));

  useEffect(() => {
    if (token) {
      setCurrentScreen('category');
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
    setCurrentScreen('category');
  };

  const handleLogout = () => {
    localStorage.removeItem('echotech_token');
    localStorage.removeItem('echotech_userId');
    localStorage.removeItem('echotech_userEmail');
    setToken(null);
    setUserId(null);
    setUserEmail(null);
    setCurrentScreen('login');
    setSelectedCategory(null);
    setSelectedTopic(null);
    setSelectedDifficulty('');
    setScore(0);
    setTotalQuestions(0);
  };

  const handleCategorySelect = (category: Category) => {
    setSelectedCategory(category);
    setCurrentScreen('topic');
  };

  const handleTopicSelect = (topic: Topic) => {
    setSelectedTopic(topic);
    setCurrentScreen('difficulty');
  };
  
  const handleDifficultySelect = (difficulty: string) => {
    setSelectedDifficulty(difficulty);
    setCurrentScreen('question');
  };

  const handleQuizComplete = (finalScore: number, numQuestions: number) => {
    setScore(finalScore);
    setTotalQuestions(numQuestions);
    if (token && selectedTopic && selectedDifficulty) {
      fetch('/api/v1/scores', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          topic: selectedTopic.slug,
          difficulty: selectedDifficulty,
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
    setSelectedCategory(null);
    setSelectedTopic(null);
    setSelectedDifficulty('');
    setScore(0);
    setTotalQuestions(0);
    setCurrentScreen('category');
  };
  
  const handleNavigateToHome = () => {
    setCurrentScreen('category');
  }

  const renderScreen = () => {
    switch (currentScreen) {
      case 'login':
        return <LoginScreen onLoginSuccess={handleLoginSuccess} onNavigateToRegister={() => setCurrentScreen('register')} />;
      case 'register':
        return <RegisterScreen onRegisterSuccess={() => setCurrentScreen('login')} onNavigateToLogin={() => setCurrentScreen('login')} />;
      case 'category':
        return <CategoryScreen onCategorySelect={handleCategorySelect} currentDifficulty={null} />;
      case 'topic':
        if (!selectedCategory) return <CategoryScreen onCategorySelect={handleCategorySelect} currentDifficulty={null} />;
        return <TopicScreen category={selectedCategory} onSelectTopic={handleTopicSelect} onBack={() => setCurrentScreen('category')} />;
      case 'difficulty':
        if (!selectedCategory || !selectedTopic) return <CategoryScreen onCategorySelect={handleCategorySelect} currentDifficulty={null} />;
        return <DifficultyScreen onSelectDifficulty={handleDifficultySelect} onBack={() => setCurrentScreen('topic')} />;
      case 'question':
        if (!selectedTopic || !selectedDifficulty) return <CategoryScreen onCategorySelect={handleCategorySelect} currentDifficulty={null} />;
        return <QuestionScreen difficulty={selectedDifficulty} topic={selectedTopic.slug} onQuizComplete={handleQuizComplete} token={token} />;
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
