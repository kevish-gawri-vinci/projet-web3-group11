import { useParams } from "react-router-dom";
import { useState, useEffect } from "react";
import "./ArticleWithId.css";
import { UserContext } from "../Context/UserContext";
import React, { useContext } from "react";


const ArticleWithId = () => {
  const { articleId } = useParams();
  const [article, setArticle] = useState(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);
  const [message, setMessage] = useState("");
  const [quantity, setQuantity] = useState(1); // État pour la quantité
  const { auth, handleLogout } = useContext(UserContext);


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
    // document.getElementById("infoMessage").classList.remove("errorMessage")
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
        const err = await response.json()
        console.log(err.error)
        throw new Error(err.error)
      } 

      setMessage("Article ajouté au panier avec succès !");
      const classes = document.getElementById("infoMessage").classList
      classes.contains("errorMessage") ? classes.remove("errorMessage") : undefined
      classes.add("goodMessage")
    } catch (err) {
      console.log(err)
      setMessage(`${err}`);
      const classes = document.getElementById("infoMessage").classList
      classes.contains("goodMessage") ? classes.remove("goodMessage") : undefined
      classes.add("errorMessage")
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
          <h6>{article.price.toFixed(2)} €</h6>
        </div>
      </div>
      {(auth.role === "user" || auth.role === "admin") && (
        <><div id="quantity-div">
          {/* Champ pour personnaliser la quantité */}
          <label htmlFor="quantity">Quantity : </label>
          <input
            id="quantity"
            type="number"
            min="1"
            value={quantity}
            onChange={(e) => setQuantity(Number(e.target.value))} // Mise à jour de la quantité
          />
        </div><div>
            <button onClick={addToBasket}>Ajouter au Panier</button>
          </div>
        {message && <p id="infoMessage" className="">{message}</p>}</>
      )}
    </div>
  );
};

export default ArticleWithId;
