import React, { useState, useEffect } from "react";
import { Link } from "react-router-dom";
import "./Commande.css";

const Commande = () => {
  const [data, setData] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchOrderAll = async () => {
    try {
      const response = await fetch("http://localhost:8080/order/getall", {
        headers: {
          Authorization: `${localStorage.getItem("token")}`,
        },
      });

      if (!response.ok) {
        throw new Error(`Erreur serveur : ${response.statusText}`);
      }

      const data = await response.json();
      setData(data.response); 
      setIsLoading(false); // Indique que le chargement est terminé
    } catch (err) {
      setError(err);
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchOrderAll();
  }, []);

  if (isLoading) {
    return <p>Chargement...</p>;
  }
  if (error) {
    return <p>Erreur : {error.message}</p>;
  }
  if (data == null || data.length === 0) {
    return <p>Aucun commande.</p>;
  }
  return (
    <div>
      <h1>Liste des Commandes</h1>
      <div className="orders-container">
        {data.map((order) => (
          <div className="order-card" key={order.orderid}>
            <h2>Commande #{order.orderid}</h2>
            <p>Total : {order.totalprice.toFixed(2)} $</p>
            <p>Quantité totale : {order.totalquantity}</p>
            <Link to={`/commande/${order.orderid}`} className="details-link">
              Voir les détails
            </Link>
          </div>
        ))}
      </div>
    </div>
  );
};
export default Commande;