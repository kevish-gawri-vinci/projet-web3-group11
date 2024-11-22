import React, { useState } from "react";
import "./Speedtest.css";
import { Link } from 'react-router-dom'


const Speedtest = () => {
    const [responseTime, setResponseTime] = useState(null); // État pour stocker le temps de réponse
    const [error, setError] = useState(null); // État pour gérer les erreurs

    const testApiSpeed = async () => {
        setError(null); // Réinitialise les erreurs
        setResponseTime(null); // Réinitialise le temps précédent

        const startTime = performance.now(); // Capturer le temps de début
        try {
            const response = await fetch("http://localhost:8080/ping"); // Requête vers l'API
            if (!response.ok) {
                throw new Error(`Erreur: ${response.statusText}`);
            }
            const endTime = performance.now(); // Capturer le temps de fin
            setResponseTime(endTime - startTime); // Calculer la différence
        } catch (err) {
            setError(err.message); // Capturer l'erreur si la requête échoue
        }
    };

    return (
        <div className="speedtest-container">
            <button onClick={testApiSpeed}>Tester si notre API tourne</button>
            {responseTime !== null && (
                <><p>L'API tourne bien ! Temps de réponse : {responseTime.toFixed(2)} ms</p><div>
                    <Link to="/">
                        <button>Page d'accueil</button>
                    </Link>
                </div></>
            )}
            {error && <p className="error">Erreur : {error}</p>}
        </div>
    );
};

export default Speedtest;
