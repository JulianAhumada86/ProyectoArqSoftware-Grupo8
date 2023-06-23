import axios from 'axios';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go

//Register

export const postUser = async (name,LastName,DNI,Password,Email,Admin) => {
  try {
    const response = await axios.post(`${API_URL}/addUsuario/${name}/${LastName}/${DNI}/${Password}/${Email}/0`);  
    return response;
    
  } catch (error) {
    if (error.response.status=400) {

      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario
      console.error(errorMessage)
      return error.response
    } else {
      // Error de red o solicitud cancelada
      console.error('Error en la solicitud:', error.message)
    }

  }
    
  throw new Error('Error al agregar reserva');
}

//LogIn

export const loginUser = async (email, password) => {
    const data = {
      email: email,
      password:password
    }

  try{
    const response = await axios.post(`${API_URL}/login`,data)
    return response;
  }catch(error){
    if (error.response.status=400) {
      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario
      console.error(errorMessage)
      return error.response

    } else {
      // Error de red o solicitud cancelada
      console.error('Error en la solicitud:', error.message)
    }

  }
};

//Reservation

export const agregarReservation = async (idHotel, inicio, final, idUser, habitacion) => {
  try {
    const response = await axios.post(`${API_URL}/agregarReservation/${idHotel}/${inicio}/${final}/${idUser}/${habitacion}`);
    return response;
  } catch (error) {
    return error.response
    
  }
};

export const getUsers = async () => {
  try {
    const response = await axios.get(`${API_URL}/users`);
    return response


  } catch (error) {
    console.error('Error al obtener los usuarios:', error);
  }
};

export const getReservations = async () => {
  try {
    const response = await axios.get(`${API_URL}/reservas`);
    return response


  } catch (error) {
    console.error('Error al obtener los usuarios:', error);
  }
};