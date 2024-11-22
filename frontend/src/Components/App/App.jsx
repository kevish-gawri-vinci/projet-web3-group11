import Footer from '../Footer/Footer';
import Header from '../Header/Header';
import NavBar from '../NavBar/NavBar';
import { Outlet } from 'react-router-dom';
import './App.css';

const App = () => {
  return (
    <div className="page">
      <Header>
        <NavBar />
      </Header>
      <main>
        <Outlet />
      </main>
      <Footer>
        <p>Footer</p>
      </Footer>
    </div>
  );
};

export default App;
