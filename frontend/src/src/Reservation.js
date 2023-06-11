import React from 'react';

function Reservation() {
  return (
    <div className="container mt-5">
      <h1>Reserva</h1>

      <div className="row">
        <div className="col-md-6">
          <div className="form-group">
            <label htmlFor="option1">Locación</label>
            <select className="form-control" id="option1">
              <option value="">Seleccionar el lugar de su estadía</option>
              <option value="option1-1">Maldron Dublín</option>
              <option value="option1-2">Maldron Buenos Aires</option>
            </select>
          </div>
        </div>
        <div className="col-md-6">
          <div className="form-group">
            <label htmlFor="option2">Habitación</label>
            <select className="form-control" id="option2">
              <option value="">Seleccionar el tipo de habitación</option>
              <option value="option2-1">1 Cama Matrimonial</option>
              <option value="option2-1">2 Camas Matrimoniales</option>
              <option value="option2-2">1 Cucheta</option>
              <option value="option2-3">1 Cama Matrimonial y 1 Cucheta</option>
            </select>
          </div>
        </div>
      </div>

      <div className="row">
        <div className="col-md-6">
          <div className="form-group">
            <label htmlFor="startDate">Fecha de inicio</label>
            <input type="date" className="form-control" id="startDate" />
          </div>
        </div>
        <div className="col-md-6">
          <div className="form-group">
            <label htmlFor="endDate">Fecha de fin</label>
            <input type="date" className="form-control" id="endDate" />
          </div>
        </div>
      </div>
      
      <button className="btn btn-primary">Reservar</button>
    </div>
  );
}

export default Reservation;