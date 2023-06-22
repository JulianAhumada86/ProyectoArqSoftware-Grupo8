import React from 'react';
import './MiCuenta.css';
import Cookies from 'js-cookie';

const MiCuenta = () => {
  const userData = Cookies.get('userData');
  
  if (!userData) {
    return (
      <div className="container">
        <h1>Mi Cuenta</h1>
        <p>No se encontraron datos de usuario.</p>
      </div>
    );
  }

  const user = JSON.parse(userData);
  

  return (
    <div className="container">
      <h1>Mi Cuenta</h1>
      <div className="user-details">
        <div className="user-image">
          <img
            src="https://definicion.de/wp-content/uploads/2019/07/perfil-de-usuario.png"
            alt="Foto de perfil"
            width="250"
            height="250"
          />
        </div>
        <div className="user-info">
          <p className="user-info-line">
            <span className="label">Nombre:</span> <span className="value">{user.name}</span>
          </p>
          <p className="user-info-line">
            <span className="label">Apellido:</span> <span className="value">{user.lastName}</span>
          </p>
          <p className="user-info-line">
            <span className="label">Correo electrónico:</span> <span className="value">{user.email}</span>
          </p>
          <p className="user-info-line">
            <span className="label">DNI:</span> <span className="value">{user.dni}</span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default MiCuenta;