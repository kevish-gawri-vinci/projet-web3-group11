import "./Footer.css";

const Footer = ({ urlLogo }) => {
  return (
    <footer className="footer">
      <img src={urlLogo} alt="logo" />
    </footer>
  );
};

export default Footer;