import './NavBar.css'
import { Link } from 'react-router-dom'

const NavBar = () => {
    return (
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
          <Link to="/login">
            <button className="navbar-buttons">Login</button>
          </Link>
          <Link to="/signup">
            <button className="navbar-buttons">Signup</button>
          </Link>
        </div>
      </div>
    );
  };

export default NavBar