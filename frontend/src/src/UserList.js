import React from 'react';
import { List, Datagrid, TextField, EmailField, EditButton } from 'react-admin';

const UserList = (props) => (
  <List {...props}>
    <Datagrid>
      <TextField source="id" />
      <TextField source="name" />
      <EmailField source="email" />
      <EditButton />
    </Datagrid>
  </List>
);

export default UserList;