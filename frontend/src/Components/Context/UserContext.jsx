// src/UserContext.js
import React, { createContext, useState, useEffect } from "react";
import axios from "axios";


export const UserContext = createContext();

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState({
    isAuthenticated: false,
    role: "guest", // "guest", "user", "admin"
    loading: true,
    username: ""
  });

  useEffect(() => {
    // Appel à l'API pour récupérer le rôle
    axios
      .get("http://localhost:8080/auth/user-role", {
        headers: {
          Authorization: `${localStorage.getItem("token")}`,
        },
      })
      .then((response) => {
        setAuth({
          isAuthenticated: true,
          role: response.data.response.isadmin ? "admin" : "user",
          username: response.data.response.username?? undefined
        });
      })
      .catch(() => {
        setAuth({ isAuthenticated: false, role: "guest" });
        localStorage.removeItem("token")
      });
  }, []);

  const handleLogout = () => {
    localStorage.removeItem("token"); // Supprime le token
    setAuth({ isAuthenticated: false, role: "guest" }); // Réinitialise l'état
  };

  return (
    <UserContext.Provider value={{ auth, setAuth, handleLogout }}>
      {children}
    </UserContext.Provider>
  );
};

