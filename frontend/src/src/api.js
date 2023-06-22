import axios from 'axios';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go


export const getUserbyId = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/userId/${id}`);
    return response;
    
  } catch (error) {
    throw new Error('Error al obtener usuario');
  }
}

export const getUsers = async () => {
  try {
    const response = await axios.get(`${API_URL}/users`);
    return response.data;
  } catch (error) {
    throw new Error('Error al obtener usuarios');
  }
};

export const postUser = async (name,LastName,DNI,Password,Email,Admin) => {
  try {
    const response = await axios.post(`${API_URL}/addUsuario/${name}/${LastName}/${DNI}/${Password}/${Email}/${Admin}`);  
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
      console.log("Entre al else, me rompi")
    }

  }
    
  throw new Error('Error al agregar reserva');
}

export const updateUser = async (userId, userData) => {
  try {
    const response = await axios.put(`${API_URL}/users/${userId}`, userData);
    return response.data;
  } catch (error) {
    throw new Error('Error al actualizar usuario');
  }
};

export const deleteUser = async (userId) => {
  try {
    const response = await axios.delete(`${API_URL}/users/${userId}`);
    return response.data;
  } catch (error) {
    throw new Error('Error al eliminar usuario');
  }
};

export const loginUser = async (email, password) => {
    const data = {
      email: email,
      password:password
    }

  try{
    const response = await axios.post(`${API_URL}/login`,data)
    console.log(response.status)
    return response;
  }catch(error){
    throw new Error('Error al logear usuario');
  }
};


export const getHotelById = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/hotelId/${id}`);
    return response.data;
  } catch (error) {
    throw new Error('Error al obtener hotel por ID');
  }
};

export const getReservaById = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/reserva/${id}`);
    return response.data;
  } catch (error) {
    throw new Error('Error al obtener reserva por ID');
  }
};

export const insertHotel = async (name, Nroom, descr) => {
  try {
    const response = await axios.post(`${API_URL}/insertHotel/${name}/${Nroom}/${descr}`);
    return response.data;
  } catch (error) {
    if (error.response) {
      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario

      //console.error(errorMessage);
      return error
    } else {
      // Error de red 
      console.error('Error en la solicitud:', error.message);
    }
    throw new Error('Error al insertar hotel');
  }
};

export const agregarReservation = async (idHotel, inicio, final, idUser, habitacion) => {
  try {
    const response = await axios.post(`${API_URL}/agregarReservation/${idHotel}/${inicio}/${final}/${idUser}/${habitacion}`);
    return response.data;
  } catch (error) {
    if (error.response) {
      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario
      console.error(errorMessage);
    } else {
      // Error de red o solicitud cancelada
      console.error('Error en la solicitud:', error.message);
    }
    throw new Error('Error al agregar reserva');
  }
};

export const insertAmenitie = async (name) => {
  try {
    const response = await axios.post(`${API_URL}/insertAmenitie/${name}`);
    return response.data;
  } catch (error) {
    if (error.response) {
      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario
      console.error(errorMessage);
    } else {
      // Error de red o solicitud cancelada
      console.error('Error en la solicitud:', error.message);
    }
    throw new Error('Error al insertar amenidad');
  }
}
