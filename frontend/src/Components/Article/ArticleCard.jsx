import { useState } from 'react'
import './ArticleCard.css'

const ArticleCard = ({ article }) => {
    return (
        <div className="card">
            <div className="card-body">
                <div>
                    <img
                        src={article.imgurl} // Affiche la première (et unique) image
                        alt={article.title} // Remplacer `article.name` par `article.title` si nécessaire
                        className="card-img-top"
                    />
                </div>
                <h5 className="card-title">{article.name}</h5>
                <p className="card-text">{article.description}</p>
                <div>
                    <h6>{article.price} $</h6>
                </div>
            </div>
        </div>
    );
};

export default ArticleCard;