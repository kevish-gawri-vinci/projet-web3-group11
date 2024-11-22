import React, { useState, useEffect } from 'react';
import ArticleList from '../Article/ArticleList';
import axios from "axios"

const HomePage = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Chargement des articles depuis l'API
    const fetchArticles = async () => {
      try {
        const response = await axios.get("http://localhost:8080/article/getall");
        setArticles(response.data.response);
      } catch (err) {
        console.error('Erreur lors du chargement des articles :', err);
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchArticles();
  }, []);

  if (loading) {
    return <p>Chargement des articles...</p>;
  }

  if (error) {
    return <p>Erreur : {error.message}</p>;
  }

  return (
    <div>
      <h1>Home Page</h1>
      <div>
        <ArticleList articles={articles} />
      </div>
    </div>
  );
};

export default HomePage;
