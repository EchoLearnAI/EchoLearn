import React, { useState, useEffect } from 'react';
import { Category } from '../types';
import '../App.css'; // Assuming App.css is in src/

interface CategoryScreenProps {
  onCategorySelect: (category: Category) => void;
  currentDifficulty: string | null; // To pass along if needed, or integrate difficulty selection here
}

const CategoryScreen: React.FC<CategoryScreenProps> = ({ onCategorySelect, currentDifficulty }) => {
  const [categories, setCategories] = useState<Category[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        setLoading(true);
        const response = await fetch('/api/v1/categories'); // Adjust API endpoint if your proxy isn't set up
        if (!response.ok) {
          throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();
        setCategories(data.categories || []); // Assuming backend returns { categories: [...] }
      } catch (e) {
        if (e instanceof Error) {
          setError(e.message);
        } else {
          setError('An unexpected error occurred.');
        }
        console.error("Failed to fetch categories:", e);
      }
      setLoading(false);
    };

    fetchCategories();
  }, []);

  if (loading) return <div className="loading-message">Loading categories...</div>;
  if (error) return <div className="error-message">Error fetching categories: {error}</div>;
  if (categories.length === 0) return <div className="info-message">No categories available.</div>;

  return (
    <div className="screen category-screen">
      <h2>Select a Category</h2>
      <div className="item-grid">
        {categories.map((category) => (
          <button 
            key={category.slug} 
            onClick={() => onCategorySelect(category)} 
            className="item-button category-button"
          >
            {category.name}
          </button>
        ))}
      </div>
    </div>
  );
};

export default CategoryScreen; 