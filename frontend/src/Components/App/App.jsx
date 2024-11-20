import Footer from '../Footer/Footer';
import Header from '../Header/Header';
import NavBar from '../NavBar/NavBar';
import { Outlet, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';


const App = () => {
  const [articles, setArticles] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Simulation de chargement des articles par défaut
    const simulateFetch = async () => {
      try {

        // Ajout d'un délai simulé pour le chargement
        // await new Promise(resolve => setTimeout(resolve, 500)); 
        let response = await axios.get("http://localhost:8080/article/getall");
        setArticles(response.data.response);
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
