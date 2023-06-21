import React from 'react';
import { Create, SimpleForm, TextInput, EmailInput } from 'react-admin';

const UserCreate = (props) => (
  <Create {...props}>
    <SimpleForm>
      <TextInput source="name" />
      <TextInput source="email" />
    </SimpleForm>
  </Create>
);

export default UserCreate;