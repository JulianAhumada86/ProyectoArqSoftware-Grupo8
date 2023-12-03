import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { Container, Row, Col, Form, Button } from 'react-bootstrap';
import Select from 'react-select';
import { postHotel } from './api';

function CreateHotel() {
  const [formData, setFormData] = useState({
    nombre: '',
    descripcion: '',
    amenities: [],
    numHabitaciones: 100,
  });

  const navigate = useNavigate();
  const [errorMessage, setErrorMessage] = useState('');
  const [showError, setShowError] = useState(false);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleAmenitiesChange = (selectedOptions) => {
    // Actualizar el estado de amenities
    setFormData((prevFormData) => ({
      ...prevFormData,
      amenities: selectedOptions.map((option) => option.value),
    }));
  };

  const handleIncrementarHabitaciones = () => {
    // Incrementar el número de habitaciones por 10
    setFormData((prevFormData) => ({
      ...prevFormData,
      numHabitaciones: Math.max(110, prevFormData.numHabitaciones + 10),
    }));
  };

  const handleDecrementarHabitaciones = () => {
    // Decrementar el número de habitaciones por 10, asegurándote de que no sea menos de 100
    setFormData((prevFormData) => ({
      ...prevFormData,
      numHabitaciones: Math.max(100, prevFormData.numHabitaciones - 10),
    }));
  };

  const handleCrearHotel = async (event) => {
      event.preventDefault(); 

      if(formData.nombre ==="" || formData.descripcion ===""){
        setErrorMessage('Debe ingresar datos');
        setShowError(true);
      } else{
      
      try {
        
        const response = await postHotel(
          
          formData.nombre,
          formData.numHabitaciones,
          formData.descripcion
        );

        if (response.status===200 ||response.status===201 ){
          setShowError(false);
          navigate("/admin/crearHotel/imagenes")

        }else if (response.status===400){
          setErrorMessage('Algo no está funcionando');
          setShowError(true)
        }else{
          setErrorMessage('Error en los datos');
          setShowError(true)
        }
      }catch(error){
        setErrorMessage('caca');
        setShowError(true)
      }
    }
  };

  return (
    <Container className="mt-5">
      <h2>Crear Hotel</h2>
      <p>Información básica</p>
      <Row>
        <Col md={6}>
          <Form>
            <Form.Group>
              <Form.Label>Nombre</Form.Label>
              <Form.Control
                type="text"
                name="nombre"
                value={formData.nombre}
                onChange={handleChange}
              />
            </Form.Group>
            <Form.Group style={{ marginTop: '10px' }}>
              <Form.Label>Descripción</Form.Label>
              <Form.Control
                as="textarea"
                rows={4}
                name="descripcion"
                value={formData.descripcion}
                onChange={handleChange}
                style={{ height: '80px' }}
              />
            </Form.Group>
          </Form>
        </Col>
        <Col md={6}>
          <Form.Group>
            <Form.Label>Amenities</Form.Label>
            <Select
              isMulti
              options={[
                { value: 'Piscina', label: 'Piscina' },
                { value: 'Wifi gratuito', label: 'Wifi gratuito' },
                { value: 'Estacionamiento', label: 'Estacionamiento' },
                // Agrega más opciones según tus necesidades
              ]}
              onChange={handleAmenitiesChange}
            />
          </Form.Group>
          <Form.Group style={{ marginTop: '10px' }}>
            <Form.Label>Número de Habitaciones</Form.Label>
            <div className="d-flex">
              <Button
                variant="outline-primary"
                onClick={handleDecrementarHabitaciones}
              >
                -
              </Button>
              <Form.Control
                type="text"
                name="numHabitaciones"
                value={formData.numHabitaciones}
                readOnly
                style={{ margin: '0 5px', textAlign: 'center' }}
              />
              <Button
                variant="outline-primary"
                onClick={handleIncrementarHabitaciones}
              >
                +
              </Button>
            </div>
          </Form.Group>
        </Col>
      </Row>
      <Row className="mt-3">
        <Col md={6}>
          {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
        </Col>
        <Col md={6} className="d-flex justify-content-end">
          <Button variant="primary" onClick={handleCrearHotel}>
            Siguiente
          </Button>
        </Col>
      </Row>
    </Container>
  );
}

export default CreateHotel;
