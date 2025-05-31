import React from 'react';

interface DifficultyScreenProps {
  onSelectDifficulty: (difficulty: string) => void;
}

const DifficultyScreen: React.FC<DifficultyScreenProps> = ({ onSelectDifficulty }) => {
  const difficulties = ['easy', 'medium', 'hard'];

  return (
    <div className="screen-container">
      <h1>Choose Difficulty</h1>
      <div className="button-container">
        {difficulties.map(diff => (
          <button key={diff} onClick={() => onSelectDifficulty(diff)}>
            {diff.charAt(0).toUpperCase() + diff.slice(1)}
          </button>
        ))}
      </div>
    </div>
  );
};

export default DifficultyScreen; 