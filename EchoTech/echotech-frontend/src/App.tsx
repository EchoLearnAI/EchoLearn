import React, { useState } from 'react';
import './App.css';
import DifficultyScreen from './components/DifficultyScreen';
import TopicScreen from './components/TopicScreen';
import QuestionScreen from './components/QuestionScreen';
import ResultScreen from './components/ResultScreen';

type Screen = 'difficulty' | 'topic' | 'question' | 'result';

function App() {
  const [currentScreen, setCurrentScreen] = useState<Screen>('difficulty');
  const [selectedDifficulty, setSelectedDifficulty] = useState<string>('');
  const [selectedTopic, setSelectedTopic] = useState<string>('');
  const [finalScore, setFinalScore] = useState<number>(0);
  const [totalQuestions, setTotalQuestions] = useState<number>(0);

  const handleSelectDifficulty = (difficulty: string) => {
    setSelectedDifficulty(difficulty);
    setCurrentScreen('topic');
  };

  const handleSelectTopic = (topic: string) => {
    setSelectedTopic(topic);
    setCurrentScreen('question');
  };

  const handleQuizComplete = (score: number, total: number) => {
    setFinalScore(score);
    setTotalQuestions(total);
    setCurrentScreen('result');
  };

  const handleRestart = () => {
    setCurrentScreen('difficulty');
    setSelectedDifficulty('');
    setSelectedTopic('');
    setFinalScore(0);
    setTotalQuestions(0);
  };

  const renderScreen = () => {
    switch (currentScreen) {
      case 'difficulty':
        return <DifficultyScreen onSelectDifficulty={handleSelectDifficulty} />;
      case 'topic':
        return <TopicScreen selectedDifficulty={selectedDifficulty} onSelectTopic={handleSelectTopic} />;
      case 'question':
        return <QuestionScreen selectedDifficulty={selectedDifficulty} selectedTopic={selectedTopic} onQuizComplete={handleQuizComplete} />;
      case 'result':
        return <ResultScreen score={finalScore} totalQuestions={totalQuestions} onRestart={handleRestart} />;
      default:
        return <DifficultyScreen onSelectDifficulty={handleSelectDifficulty} />;
    }
  };

  return (
    <div className="App">
      <header>
        <h1>EcoLearners Tech Quiz</h1>
      </header>
      <main>
        {renderScreen()}
      </main>
    </div>
  );
}

export default App;
