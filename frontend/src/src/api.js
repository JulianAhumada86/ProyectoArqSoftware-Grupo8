import axios from 'axios';
import Cookies from 'js-cookie';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go

//Register

export const postUser = async (name,LastName,DNI,Password,Email) => {
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
    
  throw new Error('Error al agregar usuario');
}

//Hotel

export const postHotel = async (name,Nroom,descr) => {
  try {
    const response = await axios.post(`${API_URL}/insertHhotel/${name}/${Nroom}/${descr}`);  
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
    
  throw new Error('Error al agregar hotel');
}

//Imagen
export const postImage = async (image,idHotel) => {
  try {
    const response = await axios.post(`${API_URL}/${image}/${idHotel}`);  
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
    
  throw new Error('Error al agregar imagenes');
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

export const getHotels = async () => {
  try {
    const userData = Cookies.get('userData');
    const user = JSON.parse(userData);
    
    axios.defaults.headers.common['Authorization'] = user.token
    const response = await axios.get(`${API_URL}/admin/hotels`);
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


export const InsertHotel = async (data) => {
  try {
    const userData = Cookies.get('userData');
    const user = JSON.parse(userData);

    axios.defaults.headers.common['Authorization'] = user.token;
    
    const response = await axios.post(`${API_URL}admin/InsertHotel`, data, {
      headers: {
        'Content-Type': 'application/json',
        // Otros encabezados si es necesario
      },
    });

    return response.data;  // Retorna los datos de la respuesta

  } catch (error) {
    console.error('Error al insertar el hotel:', error);
    throw error;  // Lanza el error para que pueda ser manejado por el código que llama a InsertHotel
  }
};