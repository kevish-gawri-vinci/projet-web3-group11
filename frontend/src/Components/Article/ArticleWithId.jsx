import { useParams } from 'react-router-dom';
import { useState, useEffect } from "react";

const ArticleWtihId = ({ Article }) => {
    const { articleId } = useParams();
    const [article, setArticle] = useState(null);
    const [currentImageIndex, setCurrentImageIndex] = useState(0);

    useEffect(() => {
        const selectedArticle = articles.find((a) => a.id === parseInt(articleId));
        setArticle(selectedArticle);
    }, [articleId, articles]);

    const nextImage = () => {
        setCurrentImageIndex((prevIndex) =>
            (prevIndex + 1) % article.imageUrls.length
        );
    };

    const prevImage = () => {
        setCurrentImageIndex((prevIndex) =>
            (prevIndex - 1 + article.imageUrls.length) % article.imageUrls.length
        );
    };

    if (!article) {
        return <p>Chargement...</p>;
    }

    return (
        <div className="card">
            <div className='card-body'>
                <div>
                    {article.imageUrls && article.imageUrls.length > 0 && (
                        <img
                            src={article.imageUrls[currentImageIndex]}
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
}
export default ArticleWtihId;