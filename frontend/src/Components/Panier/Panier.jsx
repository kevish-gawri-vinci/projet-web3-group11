const panier = () => {
    const [panier, setPanier] = useState([]);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    
    useEffect(() => {
        fetch("http://localhost:8080/panier")
        .then((response) => response.json())
        .then((data) => {
            setPanier(data);
            setIsLoading(false);
        })
        .catch((error) => {
            setError(error);
            setIsLoading(false);
        });
    }, []);
    
    if (isLoading) {
        return <p>Chargement...</p>;
    }
    
    if (error) {
        return <p>Erreur : {error}</p>;
    }
    
    if (panier.length === 0) {
        return <p>Aucun article dans le panier.</p>;
    }
    
    return (
        <div>
        <h1>Panier</h1>
        <ul>
            {panier.map((article) => (
            <li key={article.id}>
                <h2>{article.title}</h2>
                <p>{article.description}</p>
                <p>{article.prix} €</p>
            </li>
            ))}
        </ul>
        </div>
    );
    }