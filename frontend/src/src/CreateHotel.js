import React, { useState } from 'react';
import { Container, Row, Col, Form, Button, Carousel } from 'react-bootstrap';
import Select from 'react-select';

function CreateHotel() {
  const [nombre, setNombre] = useState('');
  const [descripcion, setDescripcion] = useState('');
  const [imagenes, setImagenes] = useState([]);
  const [amenities, setAmenities] = useState([]);
  const [previewImages, setPreviewImages] = useState([]);

  const handleImagenesChange = (e) => {
    const files = e.target.files;

    // Actualizar el estado de las imágenes y generar la previsualización
    setImagenes(files);
    const preview = [];
    for (let i = 0; i < files.length; i++) {
      const url = URL.createObjectURL(files[i]);
      preview.push(url);
    }
    setPreviewImages(preview);
  };

  const handleAmenitiesChange = (e) => {
    // Actualizar el estado de amenities
    setAmenities(Array.from(e.target.selectedOptions, (option) => option.value));
  };

  const handleCrearHotel = () => {
    // Lógica para enviar los datos del hotel al backend
    // Puedes usar estos datos: nombre, descripcion, imagenes, amenities
    console.log({
      nombre,
      descripcion,
      imagenes,
      amenities,
    });
  };

  return (
    <Container className="mt-5">
      <h1>Crear Hotel</h1>
      <Row>
        <Col md={6}>
          <Form>
            <Form.Group>
              <Form.Label>Nombre</Form.Label>
              <Form.Control
                type="text"
                value={nombre}
                onChange={(e) => setNombre(e.target.value)}
              />
            </Form.Group>
          </Form>
            <Carousel style={{ maxWidth: '100%', maxHeight: '270px', overflow: 'hidden', marginTop: '10px' }}>
                {previewImages.map((url, index) => (
                    <Carousel.Item key={index}>
                    <img
                        src={url}
                        alt={`Imagen ${index}`}
                        className="d-block w-100"
                        style={{ objectFit: 'cover', maxHeight: '270px', height: '270px' }}
                    />
                    </Carousel.Item>
                ))}
            </Carousel>
        </Col>
        <Col md={6}>
            <Form.Group>
                <Form.Label>Descripción</Form.Label>
                <Form.Control
                as="textarea"
                rows={4}
                value={descripcion}
                onChange={(e) => setDescripcion(e.target.value)}
                style={{height:'160px'}}
                />
            </Form.Group>
            <Form.Group style={{marginTop:'10px'}}>
                <Form.Label>Imágenes</Form.Label>
                <Form.Control type="file" multiple onChange={handleImagenesChange} />
            </Form.Group>
            <Form.Group style={{marginTop:'10px'}}>
                <Form.Label>Amenities</Form.Label>
                <Select
                    isMulti
                    options={[
                    { value: 'Piscina', label: 'Piscina' },
                    { value: 'Wifi gratuito', label: 'Wifi gratuito' },
                    { value: 'Estacionamiento', label: 'Estacionamiento' },
                    // Agrega más opciones según tus necesidades
                    ]}
                    onChange={(selectedOptions) => setAmenities(selectedOptions.map(option => option.value))}
                />
            </Form.Group>
        </Col>
      </Row>
      <Row className="mt-3">
        <Col md={6}>
          {/* Puedes agregar más campos según sea necesario */}
        </Col>
        <Col md={6} className="d-flex justify-content-end">
          <Button variant="primary" onClick={handleCrearHotel}>
            Crear
          </Button>
        </Col>
      </Row>
    </Container>
  );
}

export default CreateHotel;
