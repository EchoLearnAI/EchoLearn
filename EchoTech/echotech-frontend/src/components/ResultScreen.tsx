import React from 'react';

interface ResultScreenProps {
  score: number;
  totalQuestions: number;
  onRestart: () => void;
}

const ResultScreen: React.FC<ResultScreenProps> = ({ score, totalQuestions, onRestart }) => {
  const percentage = totalQuestions > 0 ? (score / totalQuestions) * 100 : 0;

  return (
    <div className="screen-container score-screen">
      <h1>Quiz Results</h1>
      <p>You scored {score} out of {totalQuestions}!</p>
      <p>Percentage: {percentage.toFixed(2)}%</p>
      <div className="navigation-buttons">
        <button onClick={onRestart}>Play Again</button>
      </div>
    </div>
  );
};

export default ResultScreen; 