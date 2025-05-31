import React, { useState, useEffect } from 'react';

// Mirror backend structures for clarity
interface Option {
  id: string;
  text: string;
  isCorrect?: boolean; // isCorrect might not be sent to client initially
  explanation?: string; // Explanation sent after an answer
}

interface Question {
  id: string;
  topic: string;
  difficulty: string;
  text: string;
  options: Option[];
}

interface QuestionScreenProps {
  difficulty: string;
  topic: string;
  onQuizComplete: (score: number, totalQuestions: number) => void;
  token: string | null;
}

// Define the expected structure of the API response from /submit
interface SubmitResponse {
    questionId: string;
    selectedOptionId: string;
    isCorrect: boolean;
    options: Option[]; // Array of all options with their details
}

const QuestionScreen: React.FC<QuestionScreenProps> = ({ difficulty, topic, onQuizComplete, token }) => {
  const [questions, setQuestions] = useState<Question[]>([]);
  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0);
  const [selectedAnswerId, setSelectedAnswerId] = useState<string | null>(null);
  const [isAnswered, setIsAnswered] = useState(false);
  const [score, setScore] = useState(0);
  // Store the full options with feedback details from the API
  const [answeredOptions, setAnsweredOptions] = useState<Option[]>([]); 

  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        const response = await fetch(`http://localhost:8080/api/v1/questions/${difficulty}/${topic}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        // Initialize questions without feedback details initially
        const initialQuestions = data.questions?.map((q: Question) => ({
            ...q,
            options: q.options.map(opt => ({ ...opt, isCorrect: undefined, explanation: undefined }))
        })) || [];
        setQuestions(initialQuestions);
        setAnsweredOptions([]); // Clear previous answer details

      } catch (error) {
        console.error("Failed to fetch questions:", error);
        setQuestions([]); 
        setAnsweredOptions([]);
      }
    };
    fetchQuestions();
  }, [difficulty, topic]);

  const handleAnswerSubmit = async () => {
    if (!selectedAnswerId || !questions[currentQuestionIndex]) return;

    const questionId = questions[currentQuestionIndex].id;

    try {
        const response = await fetch('http://localhost:8080/api/v1/submit', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ questionId: questionId, answerId: selectedAnswerId }),
        });
        if (!response.ok) {
            throw new Error('Failed to submit answer');
        }
        const result: SubmitResponse = await response.json();

        setIsAnswered(true);
        setAnsweredOptions(result.options); // Store all options with their feedback

        if (result.isCorrect) {
            setScore(prevScore => prevScore + 1);
        }

    } catch (error) {
        console.error("Error submitting answer:", error);
    }
  };

  const handleNextQuestion = () => {
    setSelectedAnswerId(null);
    setIsAnswered(false);
    setAnsweredOptions([]); // Clear feedback for the next question
    if (currentQuestionIndex < questions.length - 1) {
      setCurrentQuestionIndex(prevIndex => prevIndex + 1);
    } else {
      onQuizComplete(score, questions.length);
    }
  };

  if (questions.length === 0 || !questions[currentQuestionIndex]) {
    return <div className="screen-container"><p>Loading questions or no questions available for this topic/difficulty...</p></div>;
  }

  const currentQuestion = questions[currentQuestionIndex];
  // Use answeredOptions if available (after submission), otherwise use currentQuestion.options
  const displayOptions = isAnswered ? answeredOptions : currentQuestion.options;

  return (
    <div className="screen-container">
      <h2>{currentQuestion.topic} - Question {currentQuestionIndex + 1} of {questions.length}</h2>
      <div className="question-container">
        <p>{currentQuestion.text}</p>
        <ul className="options-list">
          {displayOptions.map(option => {
            let buttonClass = '';
            if (isAnswered) {
                if (option.isCorrect) {
                    buttonClass = 'correct';
                } else if (option.id === selectedAnswerId) {
                    buttonClass = 'incorrect';
                }
            } else if (option.id === selectedAnswerId) {
                buttonClass = 'selected';
            }

            return (
                <li key={option.id}>
                <button
                    onClick={() => !isAnswered && setSelectedAnswerId(option.id)}
                    className={buttonClass}
                    disabled={isAnswered}
                >
                    {option.text}
                </button>
                {isAnswered && option.explanation && (
                    <div className={`explanation ${option.isCorrect ? 'correct' : (option.id === selectedAnswerId ? 'incorrect' : '')}`.trim()}>
                        <p><strong>Explanation:</strong> {option.explanation}</p>
                    </div>
                )}
                </li>
            );
          })}
        </ul>
      </div>
      <div className="navigation-buttons">
        {!isAnswered ? (
          <button onClick={handleAnswerSubmit} disabled={!selectedAnswerId || questions.length === 0}>
            Submit Answer
          </button>
        ) : (
          <button onClick={handleNextQuestion}>
            {currentQuestionIndex < questions.length - 1 ? 'Next Question' : 'Show Results'}
          </button>
        )}
      </div>
      <p>Score: {score}</p>
    </div>
  );
};

export default QuestionScreen; 