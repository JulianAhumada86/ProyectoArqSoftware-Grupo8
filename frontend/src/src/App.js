import React, { useState } from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link, Redirect } from 'react-router-dom';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Carousel from 'react-bootstrap/Carousel';
import 'bootstrap/dist/css/bootstrap.min.css';
import AdminPanel from './AdminPanel';
import Register from './Register';
import Reservation from './Reservation';
import MiCuenta from './MiCuenta';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [accountName, setAccountName] = useState('');
  const [userData, setUserData] = useState(null);

  const handleLogin = (name, data) => {
    setIsLoggedIn(true);
    setAccountName(name);
    setUserData(data);
  };

  const Footer = () => {
    return (
      <footer className="footer mt-5">
        <div className="container text-center">
          <span>© 2023 Maldron Web Page. All Rights Reserved.</span>
        </div>
      </footer>
    );
  };

  return (
    <Router>
      <div>
        <Navbar bg="light" expand="lg">
          <Navbar.Brand>Maldron</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="mr-auto">
              <Nav.Link as={Link} to="/">
                Inicio
              </Nav.Link>
              <Nav.Link as={Link} to="/reserva">
                Reserva
              </Nav.Link>
            </Nav>
          </Navbar.Collapse>
          <Nav className="ml-auto">
            {isLoggedIn ? (
              <Nav.Link as={Link} to="/micuenta">
                Mi Cuenta ({accountName})
              </Nav.Link>
            ) : (
              <Nav.Link as={Link} to="/registro">
                Sign In
              </Nav.Link>
            )}
          </Nav>
        </Navbar>
        <div className="container mt-5">
          <Routes>
            <Route path="/" element={
              <Carousel>
                <Carousel.Item>
                  <img
                    className="d-block w-100"
                    src="https://mcaleer-rushe.co.uk/site/wp-content/uploads/2019/05/Maldron-Hotel-Belfast-IntAirport-I.jpg"
                    alt="Imagen 1"
                  />
                  <Carousel.Caption>
                    <h3>Maldron Hotel</h3>
                    <p>Ubicado en Dublín</p>
                  </Carousel.Caption>
                </Carousel.Item>
                <Carousel.Item>
                  <img
                    className="d-block w-100"
                    src="https://www.maldronhotelnewcastle.com/wp-content/uploads/sites/25/2017/10/Room-Double-Single-Maldron-Newcastle-1-1680x860.jpg"
                    alt="Imagen 2"
                  />
                  <Carousel.Caption>
                    <h3>Mejores Habitaciones</h3>
                    <p>Las más deluxe de toda Irlanda!</p>
                  </Carousel.Caption>
                </Carousel.Item>
                <Carousel.Item>
                  <img
                    className="d-block w-100"
                    src="https://www.mac-group.com/wp-content/uploads/2018/03/800x400-2.jpg"
                    alt="Imagen 3"
                  />
                  <Carousel.Caption>
                    <h3>Salón de Eventos</h3>
                    <p>Con acomodaciones para celebrar las mejores fiestas</p>
                  </Carousel.Caption>
                </Carousel.Item>
              </Carousel>
            } />
            <Route
              path="/micuenta"
              element={<MiCuenta usuario={userData} />}
            />
            <Route
              path="/registro"
              element={<Register onLogin={handleLogin} />}
            />
            <Route path="/reserva" element={<Reservation />} />
            <Route path="/admin" element={<AdminPanel />} />
            <Route path="/*" element={<Redirect to="/" />} />
          </Routes>
        </div>
        <Footer />
      </div>
    </Router>
  );
}

export default App;