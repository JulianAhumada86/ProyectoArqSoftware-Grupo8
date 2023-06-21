import React, { useEffect } from 'react';
import { Admin, Resource, List, Edit, Datagrid, TextField, TextInput, EditButton, DeleteButton } from 'react-admin';
import { getUsers, updateUser, deleteUser } from './api';

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

  const UserList = (props) => (
    <List {...props}>
      <Datagrid>
        <TextField source="id" />
        <TextField source="name" label="Nombre" />
        <TextField source="email" label="Email" />
        <EditButton basePath="/users" />
        <DeleteButton basePath="/users" />
      </Datagrid>
    </List>
  );

  const UserEdit = (props) => (
    <Edit {...props}>
      <TextInput source="name" label="Nombre" />
      <TextInput source="email" label="Email" />
    </Edit>
  );

  return (
    <Admin>
      <Resource name="users" list={UserList} edit={UserEdit} />
      {/* Agrega más rutas personalizadas aquí si es necesario */}
    </Admin>
  );
};

export default AdminPanel;