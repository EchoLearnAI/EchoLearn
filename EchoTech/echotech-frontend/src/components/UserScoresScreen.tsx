import React, { useState, useEffect } from 'react';

interface Score {
  id: number;
  userId: number;
  topic: string;
  difficulty: string;
  points: number;
  total: number;
  attemptedAt: string; // ISO date string
}

interface UserScoresScreenProps {
  token: string | null;
  onNavigateToQuiz: () => void;
}

const UserScoresScreen: React.FC<UserScoresScreenProps> = ({ token, onNavigateToQuiz }) => {
  const [scores, setScores] = useState<Score[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);
  const [filterTopic, setFilterTopic] = useState<string>("");

  // This list should ideally come from a shared source or API if dynamic
  const topics = [
    "Network", "Cybersecurity", "VMs", "Container", "Architecture", 
    "Databases", "Git", "Cache and CDN", "Monitoring", "Admin and Ops", 
    "Kubernetes", "Linux", "Pipelines and CI/CD", "APIs", "Terraform", 
    "Ansible", "Azure", "AWS", "GCP", "All"
  ];

  useEffect(() => {
    const fetchScores = async () => {
      if (!token) {
        setError("Not authenticated. Please login.");
        setLoading(false);
        return;
      }
      setLoading(true);
      setError(null);
      try {
        const response = await fetch('http://localhost:8080/api/v1/users/me/scores', {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });
        const data = await response.json();
        if (!response.ok) {
          throw new Error(data.error || 'Failed to fetch scores');
        }
        setScores(data || []);
      } catch (err: any) {
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };

    fetchScores();
  }, [token]);

  const filteredScores = filterTopic && filterTopic !== "All"
    ? scores.filter(score => score.topic === filterTopic)
    : scores;

  if (loading) return <div className="screen-container"><p>Loading scores...</p></div>;
  if (error) return <div className="screen-container"><p className="error-message">{error}</p></div>;

  return (
    <div className="screen-container score-history-screen">
      <h2>My Scores</h2>
      <div>
        <label htmlFor="topicFilter">Filter by Topic: </label>
        <select id="topicFilter" value={filterTopic} onChange={(e) => setFilterTopic(e.target.value)}>
            <option value="All">All Topics</option>
            {topics.filter(t => t !== "All").map(topic => (
                <option key={topic} value={topic}>{topic}</option>
            ))}
        </select>
      </div>
      {filteredScores.length === 0 ? (
        <p>No scores found for this topic yet. <button onClick={onNavigateToQuiz} className="link-button">Take a Quiz!</button></p>
      ) : (
        <table>
          <thead>
            <tr>
              <th>Date</th>
              <th>Topic</th>
              <th>Difficulty</th>
              <th>Score</th>
              <th>Percentage</th>
            </tr>
          </thead>
          <tbody>
            {filteredScores.map((score) => (
              <tr key={score.id}>
                <td>{new Date(score.attemptedAt).toLocaleDateString()}</td>
                <td>{score.topic}</td>
                <td>{score.difficulty}</td>
                <td>{score.points}/{score.total}</td>
                <td>{((score.points / score.total) * 100).toFixed(2)}%</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
       <div className="navigation-buttons">
        <button onClick={onNavigateToQuiz}>Back to Quiz Home</button>
      </div>
    </div>
  );
};

export default UserScoresScreen; 