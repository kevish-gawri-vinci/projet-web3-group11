import './NavBar.css'
import { Link } from 'react-router-dom'
import { UserContext } from "../Context/UserContext";
import React, { useContext } from "react";

const NavBar = () => {

    const { user, handleLogout } = useContext(UserContext);

    return (
      <>
      {user.role === "guest" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">home</button>
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
      {user.role === "user" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">home</button>
            </Link>
          </div>
          <div className="right-section">
            <Link to="/logout">
              <button className="navbar-buttons">Logout</button>
            </Link>
          </div>
        </div>
      )}
      {user.role === "admin" && (
        <div id="navbar-wrapper">
          <div className="left-section">
            <Link to="/">
              <button className="navbar-buttons">home</button>
            </Link>
            <Link to="/addArticle">
              <button className="navbar-buttons">Add Articles</button>
            </Link>
          </div>
          <div className="right-section">
            <Link to="/logout">
              <button className="navbar-buttons" onClick={handleLogout}>Logout</button>
            </Link>
          </div>
        </div>
      )}
      </>

    );

  };

export default NavBar;