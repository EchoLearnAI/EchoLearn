export interface Category {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  slug: string;
}

export interface Topic {
  id: number;
  createdAt: string;
  updatedAt: string;
  name: string;
  slug: string;
  categoryId: number;
  category?: Category; // Included if preloaded by backend
}

export interface Option {
  id: string;       // This is the original option ID like "opt1"
  text: string;
  isCorrect?: boolean; // Backend might only send this on submit, or always for review
  explanation?: string;
}

export interface Question {
  id: string;         // This is the OriginalID from backend, e.g., "net-e-q1"
  topic: string;      // Topic name or slug (ensure consistency with how backend sends it)
  difficulty: string;
  text: string;
  options: Option[];
  // DB specific fields like numeric id, TopicID, CreatedAt, UpdatedAt are not always needed by frontend for display
}

// For score submission and display, if needed more formally on frontend
export interface Score {
  id?: number;
  userId?: number;
  topicId: number; // Or topicSlug string, depending on API for submission
  topicName?: string; // For display
  difficulty: string;
  points: number;
  total: number;
  attemptedAt?: string;
} 