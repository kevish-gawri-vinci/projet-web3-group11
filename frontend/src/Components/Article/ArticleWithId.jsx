import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import "./ArticleWithId.css";

const ArticleWtihId = () => {
  const { articleId } = useParams();
  const [article, setArticle] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [message, setMessage] = useState("");

  useEffect(() => {
    // Appel au backend pour récupérer l'article
    const fetchArticle = async () => {
      try {
        const response = await fetch(
          `http://localhost:8080/article/get/${articleId}`
        );
        if (!response.ok) {
          throw new Error(`Erreur HTTP : ${response.status}`);
        }
        const data = await response.json();
        setArticle(data.response);
      } catch (err) {
        setError(err.message);
      } finally {
        setIsLoading(false);
      }
    };

    fetchArticle();
  }, [articleId]);

  const addToBasket = async () => {
    try {
      const response = await fetch("http://localhost:8080/basket/add", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(article), // Envoi de l'article complet
      });

      if (!response.ok) {
        throw new Error("Erreur lors de l'ajout au panier.");
      }

      setMessage("Article ajouté au panier avec succès !");
    } catch (err) {
      setMessage(`Erreur : ${err.message}`);
    }
  };

  if (isLoading) {
    return <p>Chargement...</p>;
  }

  if (error) {
    return <p>Erreur : {error}</p>;
  }

  if (!article) {
    return <p>Aucun article trouvé.</p>;
  }

  return (
    <div className="card">
      <div className="card-body">
        <div>
          <img
            src={article.imgurl} // Affiche la première (et unique) image
            alt={article.name} // Remplacer `article.name` par `article.title` si nécessaire
            className="card-img-top"
          />
        </div>
        <h5 className="card-title">{article.name}</h5>
        <p className="card-text">{article.description}</p>
        <div>
          <h6>{article.price}</h6>
        </div>
      </div>
      <div>
        <button onClick={addToBasket}>Add Panier</button>
      </div>
      {message && <p>{message}</p>}{" "}
    </div>
  );
};
export default ArticleWtihId;
