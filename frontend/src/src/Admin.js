import React, { useEffect } from 'react';
import { Admin, Resource, ListGuesser, EditGuesser } from 'react-admin';
import { getUsers } from './api'; // Importa las funciones de tu archivo api.js

const AdminPanel = () => {
  useEffect(() => {
    const userId = 123; // Define el ID del usuario que deseas obtener
    getUsers(userId)
      .then(response => {
        console.log('Usuarios:', response.data);
      })
      .catch(error => {
        console.error('Error al obtener usuarios:', error);
      });
  }, []);

  return (
    <Admin>
      <Resource name="users" list={ListGuesser} edit={EditGuesser} />
    </Admin>
  );
};

export default AdminPanel;