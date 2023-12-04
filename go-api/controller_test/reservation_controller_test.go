package controller_test

import (
	rdto "go-api/dto/reservations_dto"
	"go-api/errors"
	"go-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

type TestReservations struct {}

func (t *TestReservations) InsertBooking(resDto rdto.ReservationDto)(rdto.ReservationDto, errors.ErrorApi){
	if resDto.UserId == 0 {
		return rdto.ReservationDto{}, errors.NewErrorApi("Error al insertar la reserva", "reservation_insert_error", http.StatusInternalServerError, nil)
	}

	return rdto.ReservationDto{}, nil
}

func (t *TestReservations) GetReservationByUserId(id int) (rdto.ReservationDto, errors.ErrorApi) {
	if id == 1 {
		return rdto.ReservationDto {
			Id: 1,
			InitialDate: "20230418",
			FinalDate: "20220126",
			UserId: 1,
			HotelId: 2,
		}, nil
	}
	return rdto.ReservationDto{}, errors.NewNotFoundErrorApi("Reserva no encontrada")
}

func (t *TestReservations) GetReservationById(id int) (rdto.ReservationDto, errors.ErrorApi){
	if id == 1 {
		return rdto.ReservationDto{
			Id: 1,
			InitialDate: "20230418",
			FinalDate: "20220126",
			UserId: 1,
			HotelId: 2,
		}, nil
	}
	return rdto.ReservationDto{}, errors.NewErrorApi("Reserva no encontrada", "reservation_not_found", http.StatusNotFound, nil)
}

func (t *TestReservations) GetReservations() (rdto.ReservationsDto, errors.ErrorApi){
	return rdto.ReservationsDto{}, nil
}

func (t *TestReservations) GetReservationsByUserId(id int)(rdto.ReservationsDto, errors.ErrorApi){
	return rdto.ReservationsDto{}, nil 
}

func (t *TestReservations) Dispoibilidad_de_reserva(reserva rdto.ReservationCreateDto) error {
	parseInitial, err := time.Parse("20060102", reserva.InitialDate)
	if err != nil {
		return errors.NewBadRequestErrorApi("Error al analizar la fecha de inicio")
	}

	parseFinal, err := time.Parse("20060102", reserva.FinalDate)
	if err != nil {
		return errors.NewBadRequestErrorApi("Error al analizar la fecha de fin")
	}

	// Si la fecha de inicio es después de la fecha de fin, devuelve un error
	if parseFinal.Before(parseInitial) {
		return errors.NewBadRequestErrorApi("Fecha de inicio después de la fecha de fin")
	}
	return nil
}


func TestInsertReservation(t *testing.T) {
	// Configuramos el servicio con el mock
	services.ReservationService = &TestReservations{}
	
	// Creamos un router de prueba
	router := gin.Default()
	
	// Definimos la ruta de inserción de reserva
	router.POST("/reservations", func(c *gin.Context) {
		var request rdto.ReservationCreateDto
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Verificamos la disponibilidad de reserva
		if err := services.ReservationService.Dispoibilidad_de_reserva(request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Insertamos la reserva (esto es lo que estás probando)
		reservation, err := services.ReservationService.NewReserva(request)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		
		// Respondemos con la reserva creada
		c.JSON(http.StatusOK, reservation)
	})
	
	// Creamos un JSON de prueba para la solicitud
	jsonData := `{"userId": 1, "hotelId": 2, "initialDate": "20230418", "finalDate": "20220126", "habitacion": "Habitacion 1"}`
	
	// Creamos una solicitud HTTP de tipo POST con el JSON de prueba
	req, err := http.NewRequest("POST", "/reservations", strings.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	// Creamos un ResponseRecorder para grabar la respuesta de la solicitud
	rr := httptest.NewRecorder()
	
	// Ejecutamos la solicitud HTTP
	router.ServeHTTP(rr, req)
	
	// Verificamos el código de estado y la respuesta esperada
	if rr.Code != http.StatusOK {
		t.Errorf("El código de estado esperado %v, pero obtuvo %v", http.StatusOK, rr.Code)
	}
	// Puedes agregar más verificaciones según tus necesidades
}

