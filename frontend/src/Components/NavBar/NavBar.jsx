import './NavBar.css'
import { Link } from 'react-router-dom'

const NavBar = () => {
    return (
        <header id="navbar-wrapper">
            <Link to="/">Home</Link>
            <Link to="/signup"><button className="navbar-buttons">Signup</button></Link>
        </header>
    )
}

export default NavBar