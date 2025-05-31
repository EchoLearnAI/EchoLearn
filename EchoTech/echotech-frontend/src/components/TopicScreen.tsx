import React, { useState, useEffect } from 'react';

interface TopicScreenProps {
  selectedDifficulty: string;
  onSelectTopic: (topic: string) => void;
}

const TopicScreen: React.FC<TopicScreenProps> = ({ selectedDifficulty, onSelectTopic }) => {
  const [topics, setTopics] = useState<string[]>([]);

  useEffect(() => {
    // In a real app, fetch topics from the backend: GET /api/v1/topics
    // For now, using mock data similar to backend
    const fetchTopics = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/v1/topics');
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            const data = await response.json();
            setTopics(data.topics || []); 
        } catch (error) {
            console.error("Failed to fetch topics:", error);
            // Fallback to mock data if API fails
            setTopics([
                "Network", "Cybersecurity", "VMs", "Container", "Architecture", 
                "Databases", "Git", "Cache and CDN", "Monitoring", "Admin and Ops", 
                "Kubernetes", "Linux", "Pipelines and CI/CD", "APIs", "Terraform", 
                "Ansible", "Azure", "AWS", "GCP",
              ]);
        }
    };

    fetchTopics();
  }, []);

  return (
    <div className="screen-container">
      <h1>Choose Topic for {selectedDifficulty.charAt(0).toUpperCase() + selectedDifficulty.slice(1)}</h1>
      {topics.length === 0 ? (
        <p>Loading topics...</p>
      ) : (
        <div className="topic-list">
          {topics.map(topic => (
            <button key={topic} onClick={() => onSelectTopic(topic)}>
              {topic}
            </button>
          ))}
        </div>
      )}
    </div>
  );
};

export default TopicScreen; 