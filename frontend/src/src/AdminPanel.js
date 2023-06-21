
import React, { useEffect } from 'react';
import { Admin, Resource, List, Edit, Datagrid, TextField, TextInput, EditButton, DeleteButton } from 'react-admin';
import { getUsers, updateUser, deleteUser } from './api';
import UserCreate from './UserCreate';
import UserEdit from './UserEdit';

const AdminPanel = () => {
  useEffect(() => {
    getUsers()
      .then(response => {
        console.log('Usuarios:', response);
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
        <EditButton basePath="/admin/users" />
        <DeleteButton basePath="/admin/users" />
      </Datagrid>
    </List>
  );

  return (
    <Admin>
      <Resource name="users" list={UserList} edit={UserEdit} create={UserCreate} />
    </Admin>
  );
};

export default AdminPanel;