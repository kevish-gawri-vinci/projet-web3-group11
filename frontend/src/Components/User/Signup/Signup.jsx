import { useState } from "react";
import { useNavigate } from "react-router-dom"; // Pour naviguer entre les pages
import "./Signup.css";

const Signup = () => {
    const [formData, setFormData] = useState({ username: "", password: "" });
    const navigate = useNavigate(); // Hook pour rediriger
  const [error, setError] = useState("");


    // Gestion des changements dans les champs
    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData({ ...formData, [name]: value });
    };

    // Gestion de la soumission
    const handleSubmit = async (e) => {
        e.preventDefault(); // Empêcher la soumission native du formulaire

        try {
            const response = await fetch("http://localhost:8080/auth/signup", {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(formData),
            });

            if (response.ok) {
                // Si la requête est un succès, rediriger vers /login
                navigate("/login");
            } else {
                // Gérer les erreurs ici
                console.log("Erreur lors de la création de l'utilisateur ");
                const result = await response.json()
                console.log(result.error)
                throw Error(result.error)
            }
        } catch (err) {
            console.log("zbad ", err)
            setError(err ? `${err}` : "An error occurred.");
            console.log("Une erreur est survenue :", error);
        }
    };

    return (
        <div>
            {error && <p className="error-message">{error}</p>}
            <div id="signup-form-wrapper">
            
                <form onSubmit={handleSubmit}>
                    <label htmlFor="username">Name</label>
                    <input
                        type="text"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                        required
                    />

                    <label htmlFor="password">Password</label>
                    <input
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />

                    <input type="submit" value="Envoyer" id="signup-form-submit-btn" />
                </form>
            </div>
        </div>
    );
};

export default Signup;
