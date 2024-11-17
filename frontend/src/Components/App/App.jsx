import sac1 from '../../assets/sac1.jpg';
import sac2 from '../../assets/sac2.jpg';
import sac3 from '../../assets/sac3.jpg';
import Footer from '../Footer/Footer';
import Header from '../Header/Header';
import NavBar from '../NavBar/NavBar';
import { Outlet, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

const defaultArticles = [
  {
      id: 1,
      title: 'Article 1',
      description: 'Description 1',
      score: 1,
      prix: 1,
      imageUrls: [sac1, sac2, sac3]
  },
  {
      id: 2,
      title: 'Article 2',
      description: 'Description 2',
      score: 2,
      prix: 2,
      imageUrls: [sac1, sac2, sac3]
  },
  {
      id: 3,
      title: 'Article 3',
      description: 'Description 3',
      score: 3,
      prix: 3,
      imageUrls: [sac1, sac2, sac3]
  }
];

const App = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Simulation de chargement des articles par défaut
    const simulateFetch = async () => {
      try {
        // Ajout d'un délai simulé pour le chargement
        await new Promise(resolve => setTimeout(resolve, 500)); 
        setArticles(defaultArticles);
      } catch (err) {
        console.error('Erreur simulée :', err);
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    simulateFetch();
  }, []);

  if (loading) {
    return <p>Chargement des articles...</p>;
  }

  if (error){
   return <p>Erreur : {error.message}</p>;
  }  


  const articleContext = {
      articles,
  };

  return (
    <div className="page">
      <Header>
        <NavBar />
      </Header>
      <main>
        <Outlet context={articleContext} />       
      </main>
      
      <Footer>
        <p>Footer</p>
      </Footer>
    </div>
  );
};

export default App;
