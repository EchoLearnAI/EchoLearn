import React, { useState, useEffect } from 'react';
import { Category, Topic } from '../types'; // Import Category and Topic
import '../App.css';

interface TopicScreenProps {
  category: Category; // Changed from difficulty: string
  onSelectTopic: (topic: Topic) => void; // Changed from (topic: string)
  onBack: () => void;
}

const TopicScreen: React.FC<TopicScreenProps> = ({ category, onSelectTopic, onBack }) => {
  const [topics, setTopics] = useState<Topic[]>([]); // State now holds Topic objects
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!category || !category.slug) return;

    const fetchTopics = async () => {
      try {
        setLoading(true);
        setError(null);
        // Fetch topics for the selected category
        const response = await fetch(`/api/v1/topics?category_slug=${category.slug}`);
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setTopics(data.topics || []); // Assuming backend returns { topics: Topic[] }
      } catch (e) {
        if (e instanceof Error) {
          setError(e.message);
        } else {
          setError('An unexpected error occurred.');
        }
        console.error(`Failed to fetch topics for category ${category.name}:`, e);
        setTopics([]); // Clear topics on error or provide fallback
      }
      setLoading(false);
    };

    fetchTopics();
  }, [category]); // Re-fetch if category changes

  if (loading) return <div className="loading-message">Loading topics for {category.name}...</div>;
  if (error) return <div className="error-message">Error fetching topics: {error}</div>;
  
  return (
    <div className="screen topic-screen">
      <button onClick={onBack} className="back-button">&larr; Back to Categories</button>
      <h2>Select a Topic in {category.name}</h2>
      {topics.length === 0 ? (
        <div className="info-message">No topics found for {category.name}.</div>
      ) : (
        <div className="item-grid">
          {topics.map(topic => (
            <button 
              key={topic.slug} 
              onClick={() => onSelectTopic(topic)} 
              className="item-button topic-button"
            >
              {topic.name}
            </button>
          ))}
        </div>
      )}
    </div>
  );
};

export default TopicScreen; 