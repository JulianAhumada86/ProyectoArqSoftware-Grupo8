import React, {useEffect, useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { loginUser } from './api';
import Cookies from 'js-cookie';


const API_URL = 'http://localhost:8000';





const Admin = () => {
    const [reservas, setReservas] = useState([]);
    const [users, setUsers] = useState([]);
    const navigate = useNavigate();
    const [errorMessage, setErrorMessage] = useState('');
    const [showError, setShowError] = useState(false);
    const [showLogin,setShowLogin] = useState(true);
    const [showOptions,setShowOption] = useState(false);
    const [showUsers,setShowUsers] = useState(false);
    const [showReservations,setShowReservations] = useState(false);

    const [formData, setFormData] = useState({
      email: '',
      password: '',
     });

     useEffect(() => {

     
      try{
        const userData = Cookies.get('userData');
        const user = JSON.parse(userData);
        if (user.admin===1) {
          setShowLogin(false);
          setShowOption(true);
          setShowError(false)
          
        }else{
          alert("No sos administrador")
          navigate("/")
        }

      }catch{}
    }, []);


    const handleChange = (event) => {//Cuando se te cante el culo 
      const { name, value } = event.target;
      setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));  
    };

    const handleSubmit = async (event) => { //Cuando clickeas el boton
      event.preventDefault();
      
      
      try {
        const response = await loginUser(formData.email, formData.password);
         if (response.status === 200) {
          console.log(response)
          
          const user = {
            email: response.data.Email,
            name: response.data.Name,
            lastName: response.data.LastName,
            dni: response.data.DNI,
            id: response.data.Id,
            admin: response.data.Admin
          } ;
          Cookies.set('userData', JSON.stringify(user));
          if(user.admin==1){
            setShowLogin(false);
            setShowOption(true);
            setShowError(false)
          }else{
            alert("No tenes permiso de administracion");
            navigate("/");
          }
       
        } else if (response.status === 400) {
          setErrorMessage('El usuario no existe o la contraseña es incorrecta');
          setShowError(true);
        } else{
          setErrorMessage('Error al iniciar sesión');
          setShowError(true);
        }  
          
        } catch (error) {
          setErrorMessage('Error al iniciar sesión');
          setShowError(true);
          console.error(error);
      }
    };


    return (
      <div>
        {showOptions && 
        <div id="Usuarios">
        <h1>Usuarios</h1>
        <ul>
        <button className="btn btn-primary">
          Reservaciones
        </button>
        <button className="btn btn-primary">
          Usuarios
        </button>
        {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
          
        </ul>
        </div>}
        
        {showLogin &&
        <div className="container">
        <h1 id="h1">Log In</h1>
        <form>
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
            name ="password"
            value={formData.password}
            onChange={handleChange}
          />
        </div>
        {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
        <button type="button" className="btn btn-primary" onClick={handleSubmit}>
          Iniciar sesión
        </button>
        </form>
    
        </div>
            
          }

          {showReservations && <div>




          </div>}



          {showUsers && <div>



          </div>}
    </div>  
    
    
    );
  };

  export default Admin;


  
  