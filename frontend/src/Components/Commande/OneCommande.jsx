import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import "./OneCommande.css";

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
      setData(data.response);
      setIsLoading(false); // Indique que le chargement est terminé
    } catch (err) {
      console.log("e ", err)
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
  console.log(data)
  return (
    <div>
      <h1>Commande #{data.orderid}</h1>
      <div className="articles-container">
        {data.articles.map((article) => (
          <div className="orderid-card" key={article.articledetail.id}>
            <img
              src={article.articledetail.imgurl}
              alt={article.articledetail.name}
              className="orderid-image"
            />
            <div className="article-info">
              <h2>{article.articledetail.name}</h2>
              <p>{article.articledetail.description}</p>
            </div>  
              <p className="orderid-quantity">Quantité : {article.articleline.quantity}</p>
              <p className="orderid-priceOne">Prix unitaire : {article.articledetail.price.toFixed(2)}</p>
              <p className="orderid-totalPrice">Prix total : {article.articleline.price.toFixed(2)} $</p>
            
          </div>
        ))}
      </div>
      <div id="separator"></div>
      <div id="ordertotal-wrapper">
        <p id="ordertotal-title">Total de la commande</p>
        <p id="ordertotal-quantity">{data.totalprice.toFixed(2)} $</p>
      </div>
    </div>
  );
};
export default Commande;
