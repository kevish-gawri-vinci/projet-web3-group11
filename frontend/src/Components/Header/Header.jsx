import "./Header.css";

const Header = ({ children }) => {
  return (
    <header className="header">
      <div>{children}</div>
    </header>
  );
};
export default Header;