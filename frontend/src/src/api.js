import axios from 'axios';
import Cookies from 'js-cookie';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go

//Register

export const postUser = async (name,LastName,DNI,Password,Email,Admin) => {
  try {
    const response = await axios.post(`${API_URL}/addUsuario/${name}/${LastName}/${DNI}/${Password}/${Email}`);  
    return response;
    
  } catch (error) {
    if (error.response.status===400) {

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
    if (error.response.status===400) {
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

export const agregarReservation = async (idHotel, inicio, final, idUser, habitacion,token) => {
  try { 
    axios.defaults.headers.common['Authorization'] = token
    const response = await axios.post(`${API_URL}/usuario/agregarReservation/${idHotel}/${inicio}/${final}/${idUser}/${habitacion}`);
    return response;
  } catch (error) {
    return error.response
    
  }
};

export const getUsers = async () => {
  try {
    const userData = Cookies.get('userData');
    const user = JSON.parse(userData);
    
    axios.defaults.headers.common['Authorization'] = user.token
    const response = await axios.get(`${API_URL}/admin/users`);
    return response


  } catch (error) {
    console.error('Error al obtener los usuarios:', error);
  }
};

export const getReservations = async () => {
  try {
    const userData = Cookies.get('userData');
    const user = JSON.parse(userData);
    
    axios.defaults.headers.common['Authorization'] = user.token

    const response = await axios.get(`${API_URL}/admin/reservas`);
    return response


  } catch (error) {
    console.error('Error al obtener los usuarios:', error);
  }
};


export const getReservationsByUser = async () => {
  try {
    const userData = Cookies.get('userData');
    const user = JSON.parse(userData);
    
    axios.defaults.headers.common['Authorization'] = user.token
    const id = user.id
    
    const response = await axios.get(`${API_URL}/usuario/reservaByUserId/${id}`);
    return response


  } catch (error) {
    console.error('Error al obtener los usuarios:', error);
  }
};