import "./Header.css";
import NavBar from "../NavBar/NavBar";

const Header = ({ children }) => {
  return (
    <header className="header">
      <div className="logo">
        <img src="/path-to-logo.png" alt="Logo" className="logo" />
      </div>
      <div className="nav">
        <NavBar />
      </div>
    </header>

  );
};
export default Header;