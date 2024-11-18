import { useState } from 'react'
import './ArticleCard.css'

const ArticleCard = ({ article }) => {
    const [currentImage, setCurrentImage] = useState(0);

    const nextImage = () => {
        setCurrentImage((prevIndex) =>
            (prevIndex + 1) % article.imageUrls.length // Utilisez `currentImage`
        );
    };

    const prevImage = () => {
        setCurrentImage((prevIndex) =>
            (prevIndex - 1 + article.imageUrls.length) % article.imageUrls.length // Utilisez `currentImage`
        );
    };

    return (
        <div className="card">
            <div className="card-body">
                <div>
                    {article.imageUrls && article.imageUrls.length > 0 && (
                        <img
                            src={article.imageUrls[currentImage]}
                            alt={article.title}
                            className="card-img-top"
                        />
                    )}
                    <div>
                        <button onClick={prevImage}></button>
                        <button onClick={nextImage}></button>
                    </div>
                </div>                   
                <h5 className="card-title">{article.title}</h5>
                <p className="card-text">{article.description}</p>
                <div>
                    <h6>{article.score}</h6>
                    <h6>{article.prix}</h6>
                </div>
            </div>
        </div>
    );
    };

export default ArticleCard;