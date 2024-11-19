import React, { useState, useEffect } from "react";
import "./Panier.css";

const Panier = () => {
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const updateQuantity = async (articleId, newQuantity, url) => {
    try {
      const response = await fetch(url, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          Authorization: `${localStorage.getItem("token")}`,
        },
        body: JSON.stringify({
          articleId: articleId,
          quantity: newQuantity,
        }),
      });

      if (!response.ok) {
        throw new Error("Erreur lors de la mise à jour de la quantité.");
      }

      // Recharger les données du panier après la mise à jour
      fetchBasket();
    } catch (err) {
      console.error(err.message);
    }
  };

  const increaseQuantity = (articleId, currentQuantity) => {
    const newQuantity = currentQuantity + 1;
    console.log(newQuantity)
    updateQuantity(articleId, 1, "http://localhost:8080/basket/increase-quantity");
  };

  const decreaseQuantity = (articleId, currentQuantity) => {
    const newQuantity = currentQuantity - 1;
      updateQuantity(articleId, 1, "http://localhost:8080/basket/decrease-quantity");

  };

  const fetchBasket = async () => {
    try {
      const response = await fetch("http://localhost:8080/basket/get", {
        headers: {
          Authorization: `${localStorage.getItem("token")}`,
        },
      });

      if (!response.ok) {
        throw new Error(`Erreur serveur : ${response.statusText}`);
      }

      const data = await response.json();
      setData(data.response.articles); // Met à jour le panier
      setIsLoading(false); // Indique que le chargement est terminé
    } catch (err) {
      setError(err);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchBasket();
  }, []);

  if (isLoading) {
    return <p>Chargement...</p>;
  }

  if (error) {
    return <p>Erreur : {error.message}</p>;
  }

  if (data == null || data.length === 0) {
    return <p>Aucun article dans le panier.</p>;
  }

  return (
    <div>
      <h1>Panier</h1>
      <div>
        {data.map((article) => (
          <div className="basket-line-div" key={article.article.id}>
            <h2>{article.article.name}</h2>
            <img className="basket-line-img" src={article.article.imgurl} alt={article.article.name} />
            <div className="basket-line-price-section">
                <p className="basket-line-lineprice">Prix total : {article.lineprice} €</p>
                <p>({article.quantity} x {article.article.price} €)</p>
                <div className="basket-line-buttons">
                {/* Bouton pour diminuer la quantité */}
                <button
                    onClick={() => decreaseQuantity(article.article.id, article.quantity)}
                    className="decrease-btn"
                >
                    -
                </button>
                {/* Bouton pour augmenter la quantité */}
                <button
                    onClick={() => increaseQuantity(article.article.id, article.quantity)}
                    className="increase-btn"
                >
                    +
                </button>
                </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Panier;
