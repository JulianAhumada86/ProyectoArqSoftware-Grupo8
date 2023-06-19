import axios from 'axios';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go


export const getUsers = async (id) => {
  try {
    const response = await axios.get(`${API_URL}/userId/${id}`);
    return response;
    
  } catch (error) {
    throw new Error('Error al obtener usuarios');
  }
}


export const postUser = async (name,LastName,DNI,Password,Email,Admin) => {
  try {
    const response = await axios.post(`${API_URL}/addUsuario/${name}/${LastName}/${DNI}/${Password}/${Email}/${Admin}`);
    return response;
    
  } catch(error ) {
    if (error.response) {
      // El servidor respondió con un código de estado de error
      const errorMessage = error.response.data;
      // Manejar el mensaje de error, por ejemplo, mostrarlo en la interfaz de usuario
      console.error(errorMessage);
    } else {
      // Error de red o solicitud cancelada
      console.error('Error en la solicitud:', error.message);
  }
  }
}