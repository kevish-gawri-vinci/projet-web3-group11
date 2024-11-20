import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import "./Commande.css";

const Commande = () => {
  const { orderId } = useParams();
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchOrder = async () => {
    try {
      const response = await fetch(
        `http://localhost:8080/order/get/${orderId}`,
        {
          headers: {
            Authorization: `${localStorage.getItem("token")}`,
          },
        }
      );

      if (!response.ok) {
        throw new Error(`Erreur serveur : ${response.statusText}`);
      }

      const data = await response.json();
      setData(data.response.articles);
      setIsLoading(false); // Indique que le chargement est terminé
    } catch (err) {
      setError(err);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchOrder();
  }, [orderId]);

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
      <h1>Commande</h1>
      <div className="articles-container">
        {data.map((article) => (
          <div className="article-card" key={article.articledetail.id}>
            <img
              src={article.articledetail.imgurl}
              alt={article.articledetail.name}
              className="article-image"
            />
            <div className="article-info">
              <h2>{article.articledetail.name}</h2>
              <p>{article.articledetail.description}</p>
              <p>Quantité : {article.articleline.quantity}</p>
              <p>Prix total : {article.articleline.price.toFixed(2)} $</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};
export default Commande;
