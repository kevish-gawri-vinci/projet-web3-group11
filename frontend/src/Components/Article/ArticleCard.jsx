import { useState } from "react";
import { Link } from "react-router-dom";
import "./ArticleCard.css";

const ArticleCard = ({ article }) => {
  return (
    <div className="card">
      <Link to={`/${article.id}`}>
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
            <h6>{article.price} $</h6>
          </div>
        </div>
      </Link>6
    </div>
  );
};

export default ArticleCard;
