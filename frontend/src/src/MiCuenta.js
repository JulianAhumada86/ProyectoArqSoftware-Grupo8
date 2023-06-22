import React, { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import './MiCuenta.css';
import {User} from './Usuario.js';



export const MiCuenta = () => {
  const location = useLocation();
  
   // Inicializa con un objeto vacío
      /*
  useEffect(() => {
    const fetchUserData = async () => {

      try {
        const response = await getUserbyEmail(email);
        if (response.status === 200) {
          setUserData(response.data);
        } else {
          // Manejo de errores en caso de que no se pueda obtener los datos del usuario
          console.error('Error al obtener los datos del usuario');
        }
      } catch (error) {
        console.error('Error al obtener los datos del usuario:', error);
      }
    };

    fetchUserData();
  }, [email]);
*/
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
            <span className="label">Nombre:</span> <span className="value">{User.firstName}</span>
          </p>
          <p className="user-info-line">
            <span className="label">Apellido:</span> <span className="value">{User.lastName}</span>
          </p>
          <p className="user-info-line">
            <span className="label">Correo electrónico:</span> <span className="value">{User.email}</span>
          </p>
          <p className="user-info-line">
            <span className="label">DNI:</span> <span className="value">{User.dni}</span>
          </p>
        </div>
      </div>
    </div>
  );
};

export default MiCuenta;

