import React from 'react';
import '../App.css'; // Import App.css for styling

interface DifficultyScreenProps {
  onSelectDifficulty: (difficulty: string) => void;
  onBack: () => void; // Added onBack prop
}

const DifficultyScreen: React.FC<DifficultyScreenProps> = ({ onSelectDifficulty, onBack }) => {
  const difficulties = ['easy', 'medium', 'hard'];

  // TODO: Optionally, receive selectedCategory and selectedTopic as props to customize the title
  // e.g., <h2>Select Difficulty for {topic.name} in {category.name}</h2>

  return (
    <div className="screen difficulty-screen">
      <button onClick={onBack} className="back-button">&larr; Back to Topics</button>
      <h2>Select Difficulty</h2>
      <div className="item-grid">
        {difficulties.map(diff => (
          <button 
            key={diff} 
            onClick={() => onSelectDifficulty(diff)} 
            className="item-button difficulty-button"
          >
            {diff.charAt(0).toUpperCase() + diff.slice(1)}
          </button>
        ))}
      </div>
    </div>
  );
};

export default DifficultyScreen; 