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
    habitaciones: [
      { tipo: '', cantidad: 0 },
      { tipo: '', cantidad: 0 },
      { tipo: '', cantidad: 0 },
      { tipo: '', cantidad: 0 },
    ],
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

  const handleHabitacionesChange = (index, field, value) => {
    setFormData((prevFormData) => {
      const updatedHabitaciones = [...prevFormData.habitaciones];
      updatedHabitaciones[index][field] = value;
      return { ...prevFormData, habitaciones: updatedHabitaciones };
    });
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
                style={{ height: '160px' }}
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
          <Form.Group style={{marginTop:'10px'}}>
            <Form.Label>Tipo y Cantidad de Habitaciones</Form.Label>
            {[0, 1, 2, 3].map((index) => (
              <Row key={index}>
                <Col md={4}>
                  <Form.Control
                    as="select"
                    value={formData.habitaciones[index].tipo}
                    onChange={(e) => handleHabitacionesChange(index, 'tipo', e.target.value)}
                  >
                    {/* Opciones de tipo */}
                    <option value="">Seleccionar tipo</option>
                    {/* Agrega más opciones según tus necesidades */}
                  </Form.Control>
                </Col>
                <Col md={4}>
                  <Form.Control
                    type="number"
                    min={0}
                    max={800}
                    value={formData.habitaciones[index].cantidad}
                    onChange={(e) => handleHabitacionesChange(index, 'cantidad', e.target.value)}
                  />
                </Col>
              </Row>
            ))}
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
