import React, { useState }  from 'react';
import { useNavigate } from 'react-router-dom';
import { getReservaById } from './api';
import { agregarReservation } from './api';
import Cookies from 'js-cookie';
import Swal from 'sweetalert2';

function Reservation() {

  var user=JSON;
  const [formData, setFormData] = useState({
    option1: '',
    startDate: '',
    endDate: '',
    option2: '',
  });

  const navigate = useNavigate();
  const [errorMessage, setErrorMessage] = useState('');
  const [showError, setShowError] = useState(false);

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
    console.log(formData.option1)
  };

  const handleSubmit = async (event) => {
    event.preventDefault(); 

    try{
      const userData = Cookies.get('userData');
      user = JSON.parse(userData);

    }catch(error){
      console.log(error)
      alert("No podes reservar sin estar registrado")
      navigate("/login")
    } 

    if(formData.option1===0||formData.option2===0){
      setErrorMessage('Debe completar todas las opciones de selccion multiple');
      setShowError(true);

    }else if(formData.startDate ==="" || formData.endDate ===""){
      Swal.fire({
        title:"Hotel Completo",
        text:"Nuestras habitaciones están ocupadas, prueba otra fecha",
        icon: "info",
        showCancelButton: false,
        confirmButtonText: "OK",
      }).then((result) =>{
            if (result.isConfirmed){
              console.log('OK');  
              setFormData ({
                startDate: '',
                endDate: '',
              });
            }
      });
      //setErrorMessage('Debe ingresar fechas validas');
      //setShowError(true);
    } else{
    
      try {
      
      const response = await agregarReservation(
        
        formData.option1,
        formData.startDate,
        formData.endDate,
        user.id,
        formData.option2,
        user.token
      );

      if (response.status===200 ||response.status===201 ){
        navigate("/hotel")
        //setShowError(false);
        //alert("Reserva registrada el dia " + formData.startDate)
        //navigate("/")

      }else if (response.status===400){
        //
        setErrorMessage('Algo no está funcionando');
        setShowError(true)

      }else{

        setErrorMessage('Error en los datos');
        setShowError(true)
      }

    }catch(error){

    }
  }
  };
  
  return (
    <div className="container mt-5">
      
      <h1>Reserva</h1>
      <form onSubmit={handleSubmit}>
        <div className="row">
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="option1">Locación</label>
              <select onChange={handleChange} className="form-control" id="option1" name="option1">
                <option value="0">Seleccionar el lugar de su estadía</option>
                <option value="1">Maldron Dublín</option>
                <option value="2">Maldron Buenos Aires</option>
              </select>
            </div>
          </div>
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="option2">Habitación</label>
              <select onChange={handleChange} className="form-control" id="option2" name="option2">
                <option  value="0">Seleccionar el tipo de habitación</option>
                <option value="1 Cama Matrimonial">1 Cama Matrimonial</option>
                <option value="2 Camas Matrimoniales">2 Camas Matrimoniales</option>
                <option value="1 Cucheta">1 Cucheta</option>
                <option value="1 Cama Matrimonial y 1 Cucheta">1 Cama Matrimonial y 1 Cucheta</option>
              </select>
            </div>
          </div>
        </div>

        <div className="row">
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="startDate">Fecha de inicio</label>
              <input onChange={handleChange} type="date" className="form-control" id="startDate" name="startDate" />
            </div>
          </div>
          <div className="col-md-6">
            <div className="form-group">
              <label htmlFor="endDate">Fecha de fin</label>
              <input onChange={handleChange} type="date" className="form-control" id="endDate" name="endDate"/>
            </div>
          </div>
        </div>
        {showError && <p style={{ color: 'red' }}>{errorMessage}</p>}
        <button type="submit" className="btn btn-primary" >Reservar</button>
        </form>
      </div>
    
  );
}
export default Reservation;


/*

//esta función no se si tiene que ir aca, me parece que no
async function fetchReservaById() { //cambie el nombre para que no sea igual a la importada
  var id = document.getElementById("Rid").value;
  try {
    const response = await getReservaById(id);
    console.log(response.data);
  } catch (error) {
    console.log(error.message);
  }
}


function Reservation() { //esto es para un POST reserva
  const handleReservarClick = async () => { //PARA ESTO, CREO QUE HAY QUE AGREGAR UN ONCLICK EN EL BOTON RESERVAR
    const idHotel = document.getElementById('option1').value;
    const habitacion = document.getElementById('option2').value;
    const inicio = document.getElementById('startDate').value;
    const final = document.getElementById('endDate').value;
    const idUser = ''; // Obtén el ID de usuario según tus necesidades

    try {
      const response = await agregarReservation(idHotel, inicio, final, idUser, habitacion);
      // Manejar la respuesta exitosa, por ejemplo, mostrar un mensaje de éxito
      console.log('Reserva agregada exitosamente');
    } catch (error) {
      // Manejar el error, por ejemplo, mostrar un mensaje de error
      console.error('Error al agregar reserva:', error.message);
    }
  };
}
*/

