// src/UserContext.js
import React, { createContext, useState, useEffect } from "react";
import axios from "axios";


export const UserContext = createContext();

export const AuthProvider = ({ children }) => {
  const [auth, setAuth] = useState({
    isAuthenticated: false,
    role: "guest", // "guest", "user", "admin"
    loading: true,
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
        console.log("dzfjdef " + response.data.is_admin)
        setAuth({
          isAuthenticated: true,
          role: response.data.is_admin ? "admin" : "user",
        });
      })
      .catch(() => {
        setAuth({ isAuthenticated: false, role: "guest" });
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

