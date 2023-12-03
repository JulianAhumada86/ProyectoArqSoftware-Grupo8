import React, { useState, useEffect } from 'react';
import './App.css';
import { BrowserRouter as Router, Routes, Route, Link, Navigate } from 'react-router-dom';
import Navbar from 'react-bootstrap/Navbar';
import Nav from 'react-bootstrap/Nav';
import Carousel from 'react-bootstrap/Carousel';
import 'bootstrap/dist/css/bootstrap.min.css';
import Register from './Register';
import Reservation from './Reservation';
import MiCuenta from './MiCuenta';
import LogIn from './LogIn';
import Hotel from './Hotel';
import Confirmation from './Confirmation';
import Cookies from 'js-cookie';
//import jsonServerProvider from 'ra-data-json-server';
import Admin from './Admin';
import CreateHotel from './CreateHotel';
import AddImages from './AddImages';

function App() {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [accountName, setAccountName] = useState('');
  const [userData, setUserData] = useState(null);

  const userDatax = Cookies.get('userData');
  const user = userDatax ? JSON.parse(userDatax) : null;

  useEffect(() => {
    const userDataCookie = Cookies.get('userData');
    if (userDataCookie) {
      const user = JSON.parse(userDataCookie);
      setIsLoggedIn(true);
      setAccountName(`${user.name} ${user.lastName}`);
      setUserData(user);
    }
  }, []);

  const handleLogin = (name, data) => {
    setIsLoggedIn(true);
    setAccountName(name);
    setUserData(data);
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    setAccountName('');
    setUserData(null);
    Cookies.remove('userData');
  };
  
  const MaldronLogo = 'https://r7c7u2r3.rocketcdn.me/wp-content/uploads/2017/12/Maldron-Hotels-and-Partners.png';

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
          <Navbar.Brand style={{marginLeft:'20px'}}><img src={MaldronLogo}/></Navbar.Brand>
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
            <Nav className="ml-auto" style={{marginRight:'20px'}}>
              {isLoggedIn ? (
                <>
                  <Nav.Link as={Link} to="/micuenta">
                    {user.name} {user.lastName}
                  </Nav.Link>
                  <Nav.Link onClick={handleLogout}>Cerrar sesión</Nav.Link>
                </>
              ) : (
                <Nav.Link as={Link} to="/login">
                  Iniciar sesión
                </Nav.Link>
              )}
            </Nav>
        </Navbar>
        <div className="container mt-5">
          <Routes>
            <Route
              path="/"
              element={
                <Carousel> 
                  <Carousel.Item>
                    <img
                      className="d-block w-100"
                      src="https://mcaleer-rushe.co.uk/site/wp-content/uploads/2019/08/Maldron-Glasgow-I.jpg"
                      alt="Imagen 1"
                    />
                    <Link to="/hotel">
                      <Carousel.Caption>
                        <h3>Maldron Hotel Glasgow</h3>
                        <p>Se encuentra en el centro de la capital escocesa</p>
                      </Carousel.Caption>
                    </Link>  
                  </Carousel.Item>
                  <Carousel.Item>
                    <img
                      className="d-block w-100"
                      src="https://mcaleer-rushe.co.uk/site/wp-content/uploads/2019/05/Maldron-Hotel-Belfast-IntAirport-I.jpg"
                      alt="Imagen 2"
                    />
                    <Link to="/hotel">
                      <Carousel.Caption>
                        <h3>Maldron Hotel Belfast</h3>
                        <p>El primer hotel de nuestra cadena fuera del Reino Unido</p>
                      </Carousel.Caption>
                    </Link>  
                  </Carousel.Item>
                  <Carousel.Item>
                    <img 
                      className="d-block w-100"
                      src="https://apithegreeks.gumlet.io/hotels/maldron-hotel-manchester-city-centre/cover.jpg?w=1900&h=&v=1687503002&site=thegreeks"
                      alt="Imagen 3"
                    />
                    <Link to="/hotel">
                      <Carousel.Caption>
                        <h3>Maldron Hotel Dublín</h3>
                        <p>Ubicado en el aeropuerto más importante de toda Irlanda</p>
                      </Carousel.Caption>
                    </Link>
                  </Carousel.Item>
                </Carousel>
              }
            />
            <Route
              path="/micuenta"
              element={<MiCuenta usuario={userData} />}
            />
            <Route
              path="/registro"
              element={<Register onLogin={handleLogin} />}
            />
            <Route path="/login" element={<LogIn onLogin={handleLogin} />} />
            <Route path="/reserva" element={<Reservation />} />
            <Route path="/admin" element={<Admin />} />
            <Route path="/admin/crearHotel" element={<CreateHotel />} />
            <Route path="/admin/crearHotel/imagenes" element={<AddImages/>}/>
            <Route path="/hotel" element={<Hotel/>}/>
            <Route path="/confirmacion" element={<Confirmation/>}/>
          </Routes>
        </div>
        <Footer />
      </div>
    </Router>
  
  );
}

export default App;