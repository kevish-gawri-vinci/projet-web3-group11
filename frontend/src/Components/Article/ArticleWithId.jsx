import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import "./ArticleWithId.css";

const ArticleWithId = () => {
  const { articleId } = useParams();
  const [article, setArticle] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [message, setMessage] = useState("");
  const [quantity, setQuantity] = useState(1); // État pour la quantité

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
          Authorization: localStorage.getItem("token"),
        },
        body: JSON.stringify({
          articleid: article.id, // Article à ajouter
          quantity: quantity, // Quantité choisie
        }),
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
            src={article.imgurl} // Affiche l'image
            alt={article.name} // Remplacer `article.name` par `article.title` si nécessaire
            className="card-img-top"
          />
        </div>
        <h5 className="card-title">{article.name}</h5>
        <p className="card-text">{article.description}</p>
        <div>
          <h6>{article.price} €</h6>
        </div>
      </div>
      <div id="quantity-div">
        {/* Champ pour personnaliser la quantité */}
        <label htmlFor="quantity">Quantity : </label>
        <input
          id="quantity"
          type="number"
          min="1"
          value={quantity}
          onChange={(e) => setQuantity(Number(e.target.value))} // Mise à jour de la quantité
        />
      </div>
      <div>
        <button onClick={addToBasket}>Ajouter au Panier</button>
      </div>
      {message && <p>{message}</p>}
    </div>
  );
};

export default ArticleWithId;
