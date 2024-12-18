import './NavBar.css'
import { Link } from 'react-router-dom'
import { UserContext } from "../Context/UserContext";
import React, { useContext } from "react";

const NavBar = () => {
  const { auth, handleLogout } = useContext(UserContext);
  console.log("role role ", auth.role)
  return (
    <>
      {auth.role === "guest" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">Home</button>
            </Link>
          </div>
          <div className="right-section">
            <Link to="/login">
              <button className="navbar-buttons">Login</button>
            </Link>
            <Link to="/signup">
              <button className="navbar-buttons">Signup</button>
            </Link>
          </div>
        </div>
      )}
      {auth.role === "user" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">Home</button>
            </Link>
          </div>
          <div className="center-section">
              <p>Hello {auth.username}</p>
          </div>
          <div className="right-section">
            <Link to="/panier">
              <button className="navbar-buttons">Panier</button>
            </Link>
            <Link to="/commande">
              <button className="navbar-buttons">Commande</button>
            </Link>
            <button className="navbar-buttons" onClick={handleLogout}>
              Logout
            </button>
          </div>
        </div>
      )}
      {auth.role === "admin" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">Home</button>
            </Link>
            <Link to="/addArticle">
              <button className="navbar-buttons">Add Articles</button>
            </Link>
          </div>
          <div className="right-section">
          <Link to="/panier">
              <button className="navbar-buttons">Panier</button>
            </Link>
            <Link to="/commande">
              <button className="navbar-buttons">Commande</button>
            </Link>
            <button className="navbar-buttons" onClick={handleLogout}>
              Logout
            </button>
          </div>
        </div>
      )}
    </>
  );
};


export default NavBar;