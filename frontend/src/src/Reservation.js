import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { getHotels } from './api';
import { agregarReservation } from './api';
import Cookies from 'js-cookie';
import Swal from 'sweetalert2';

function Reservation() {
  var user = JSON;
  const [hoteles, setHoteles] = useState([]);
  const [formData, setFormData] = useState({
    option1: '',
    startDate: '',
    endDate: '',
    option2: '',
  });
  const [errorMessage, setErrorMessage] = useState('');
  const [showError, setShowError] = useState(false);

  const navigate = useNavigate();

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    // ... código del formulario

    try {
      const response = await agregarReservation(
        formData.option1,
        formData.startDate,
        formData.endDate,
        user.id,
        formData.option2,
        user.token
      );

      if (response.status === 200 || response.status === 201) {
        navigate('/hotel');
      } else {
        setErrorMessage('Error en los datos');
        setShowError(true);
      }
    } catch (error) {
      console.error('Error al realizar la reserva:', error);
      setErrorMessage('Algo salió mal');
      setShowError(true);
    }
  };

  const getHoteles = async () => {
    try {
      const response = await getHotels();
      const reservasData = response.data.hotels;
      setHoteles(reservasData);
    } catch (error) {
      console.error('Error al obtener hoteles:', error);
    }
  };

  useEffect(() => {
    getHoteles();
  }, []);

  return (
    <div className="container mt-5">
      <h1>Reserva</h1>
      <form onSubmit={handleSubmit}>
        <div className="row">
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="option1">Locación</label>
              <select
                onChange={handleChange}
                value={formData.option1}
                className="form-control"
                id="option1"
                name="option1"
              >
                <option value="0">Seleccionar el lugar de su estadía</option>
                {hoteles.map((hotel) => (
                  <option key={hotel.id} value={hotel.id}>
                    {hotel.name}
                  </option>
                ))}
              </select>
            </div>
          </div>
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="option2">Habitación</label>
              <select
                onChange={handleChange}
                value={formData.option2}
                className="form-control"
                id="option2"
                name="option2"
              >
                <option value="0">Seleccionar el tipo de habitación</option>
                {hoteles
                  .find((hotel) => hotel.id === parseInt(formData.option1, 10))
                  ?.habitaciones.map((habitacion) => (
                    <option key={habitacion.id} value={habitacion.id}>
                      {habitacion.tipo}
                    </option>
                  ))}
              </select>
            </div>
          </div>
        </div>

        <div className="row">
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="startDate">Fecha de inicio</label>
              <input
                onChange={handleChange}
                type="date"
                className="form-control"
                id="startDate"
                name="startDate"
                value={formData.startDate}
              />
            </div>
          </div>
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="endDate">Fecha de fin</label>
              <input
                onChange={handleChange}
                type="date"
                className="form-control"
                id="endDate"
                name="endDate"
                value={formData.endDate}
              />
            </div>
          </div>
        </div>

        {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
        <button type="submit" className="btn btn-primary">
          Reservar
        </button>
      </form>
    </div>
  );
}

export default Reservation;

/*

//esta función no se si tiene que ir aca, me parece que no
async function fetchReservaById() { //cambie el nombre para que no sea igual a la importada
  var id = document.getElementById("Rid").value;
  try {
    const response = await getReservaById(id);
    console.log(response.data);
  } catch (error) {
    console.log(error.message);
  }
}


function Reservation() { //esto es para un POST reserva
  const handleReservarClick = async () => { //PARA ESTO, CREO QUE HAY QUE AGREGAR UN ONCLICK EN EL BOTON RESERVAR
    const idHotel = document.getElementById('option1').value;
    const habitacion = document.getElementById('option2').value;
    const inicio = document.getElementById('startDate').value;
    const final = document.getElementById('endDate').value;
    const idUser = ''; // Obtén el ID de usuario según tus necesidades

    try {
      const response = await agregarReservation(idHotel, inicio, final, idUser, habitacion);
      // Manejar la respuesta exitosa, por ejemplo, mostrar un mensaje de éxito
      console.log('Reserva agregada exitosamente');
    } catch (error) {
      // Manejar el error, por ejemplo, mostrar un mensaje de error
      console.error('Error al agregar reserva:', error.message);
    }
  };
}
*/

