import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { getUsers } from './api';
import { postUser } from './api';
import { Axios } from 'axios';

const Register = ({ onLogin }) => {
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    firstName: '',
    lastName: '',
    dni: '',
    email: '',
    password: '',
    confirmPassword: '',
  });

  const [errorMessage, setErrorMessage] = useState('');
  const [showError, setShowError] = useState(false);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      // Las contraseñas no coinciden, muestra un mensaje de error
      setErrorMessage('Las contraseñas no coinciden');
      setShowError(true);
      return;
    }

    try {
      const response = await postUser(
        formData.firstName,
        formData.lastName,
        formData.dni,
        formData.password,
        formData.email,
        0
      );
      onLogin(formData.firstName, formData); // Llama a la función onLogin pasando el nombre del usuario registrado y los datos del formulario
      navigate('/'); // Redirige al usuario a la página principal después de registrar exitosamente
    } catch (error) {
      // Manejo de error
    }

    // Restablece los valores y oculta el mensaje de error
    setFormData({
      firstName: '',
      lastName: '',
      dni: '',
      email: '',
      password: '',
      confirmPassword: '',
    });
    setShowError(false);
  };

  return (
    <div className="container">
      <h1 id="h1">Sign In</h1>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="firstName">Nombre</label>
          <input
            type="text"
            className="form-control"
            id="firstName"
            name="firstName"
            value={formData.firstName}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="lastName">Apellido</label>
          <input
            type="text"
            className="form-control"
            id="lastName"
            name="lastName"
            value={formData.lastName}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="dni">DNI</label>
          <input
            type="text"
            className="form-control"
            id="dni"
            name="dni"
            value={formData.dni}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            className="form-control"
            id="email"
            name="email"
            value={formData.email}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="password">Contraseña</label>
          <input
            type="password"
            className="form-control"
            id="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
          />
        </div>
        <div className="form-group">
          <label htmlFor="confirmPassword">Confirmar Contraseña</label>
          <input
            type="password"
            className="form-control"
            id="confirmPassword"
            name="confirmPassword"
            value={formData.confirmPassword}
            onChange={handleChange}
          />
        </div>
        {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
        <button type="submit" className="btn btn-primary">
          Registrarse
        </button>
      </form>
    </div>
  );
};

export default Register;


async function postuser(){
  var dni = document.getElementById("dni").value
  var lastName = document.getElementById("lastName").value
  var firstName = document.getElementById("firstName").value
  var email = document.getElementById("email").value
  var password = document.getElementById("password").value
  var cpassword =document.getElementById("confirmPassword").value

  try{
    const response = await postUser(firstName,lastName,dni,password,email,0)
    
  }catch(error){
    
  }
}

//ACA HAY QUE ARERGLAR LO DEL NOMBRE QUE EN REALIDAD ES USER BY ID
/*Primera funcion para obtener datos
async function getuser() { //Aca deberia llamarse get user by id
  var id = document.getElementById("dnix").value
  var last = document.getElementById("h1")
  try {
    const response = await getUserbyId(id); //response
    console.log(response.data.LastName)
    last.innerText = response.data.LastName;
  }
  catch(error){
    console.log(error.message)
  }

}
*/