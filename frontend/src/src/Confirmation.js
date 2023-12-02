import React from 'react';
import { Container, Card, Carousel, Button, Row, Col } from 'react-bootstrap';
import { FaWifi, FaSwimmingPool, FaParking, FaDumbbell, FaSpa } from 'react-icons/fa';

function Confirmation() {
  const hotel = {
    nombre: 'Hotel Ejemplo',
    descripcion: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation.',
    fotos: [
      'https://www.maldronhotelnewcastle.com/wp-content/uploads/sites/25/2017/10/Room-Double-Single-Maldron-Newcastle-1-1680x860.jpg',
      'https://mcaleer-rushe.co.uk/site/wp-content/uploads/2019/05/Maldron-Hotel-Belfast-IntAirport-I.jpg',
      'https://www.mac-group.com/wp-content/uploads/2018/03/800x400-2.jpg',
    ],
    amenities: [
      { nombre: 'Wifi Gratis', icono: <FaWifi /> },
      { nombre: 'Piscina', icono: <FaSwimmingPool /> },
      { nombre: 'Estacionamiento', icono: <FaParking /> },
      { nombre: 'Gimnasio', icono: <FaDumbbell /> },
      { nombre: 'Spa', icono: <FaSpa /> },
    ],
  };

  return (
    <Container className="mt-5">
      <Row>
        <Col md={6}>
              <h1>{hotel.nombre}</h1>
              <p>{hotel.descripcion}</p>
            
              <div className="mt-3">
                <h5>Amenidades</h5>
                <div className="d-flex flex-wrap">
                  {hotel.amenities.map((amenity, index) => (
                    <Card key={index} className="m-2" style={{ width: '4rem', height: '4rem' }}>
                      <Card.Body style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center', height: '100%' }}>
                        <div style={{ fontSize: '1.75rem' }}>{amenity.icono}</div>
                      </Card.Body>
                    </Card>
                  ))}
                </div>
              </div>
        </Col>
        <Col md={6}>
          <Carousel style={{ maxWidth: '100%', height: '100%' }}>
            {hotel.fotos.map((foto, index) => (
              <Carousel.Item key={index}>
                <img src={foto} alt={`Imagen ${index}`} className="d-block w-100" style={{ objectFit: 'cover', height: '100%' }} />
              </Carousel.Item>
            ))}
          </Carousel>
        </Col>
      </Row>
      
    </Container>
  );
  
}

export default Confirmation;