import axios from 'axios';

const API_URL = 'http://localhost:8000'; // Reemplaza con la URL de tu API de Go

export const getUsers = async () => {
  try {
    const response = await axios.get(`${API_URL}/users`);
    return response.data;
  } catch (error) {
    throw new Error('Error al obtener usuarios');
  }
};