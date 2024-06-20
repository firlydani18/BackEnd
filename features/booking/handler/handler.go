package handler

import (
	"KosKita/features/booking"
	"KosKita/utils/middlewares"
	"KosKita/utils/responses"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// _midtransController "BE-REPO-20/features/midtrans/controller"

type BookingHandler struct {
	bookingService booking.BookingServiceInterface
}

func NewBooking(service booking.BookingServiceInterface) *BookingHandler {
	return &BookingHandler{
		bookingService: service,
	}
}

func (handler *BookingHandler) CreateBooking(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	newBooking := BookingRequest{}
	errBind := c.Bind(&newBooking)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data booking not valid", nil))
	}

	bookingCore := RequestToCoreBooking(newBooking)
	payment, errInsert := handler.bookingService.PostBooking(uint(idJWT), bookingCore)
	if errInsert != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error insert booking "+errInsert.Error(), nil))
	}

	results := CoreToResponse(*payment)

	return c.JSON(http.StatusOK, responses.WebResponse("Success get booking.", results))
}

func (handler *BookingHandler) GetBookings(c echo.Context) error {
	idJWT := middlewares.ExtractTokenUserId(c)
	if idJWT == 0 {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("unauthorized or jwt expired", nil))
	}

	results, err := handler.bookingService.GetBookings(uint(idJWT))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("Error booking. "+err.Error(), nil))
	}

	var bookingResults []BookingHistoryResponse
	for _, result := range results {
		bookingResults = append(bookingResults, CoreToResponseBookingHistory(result))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("Success get booking.", bookingResults))
}

func (handler *BookingHandler) CancelBookingById(c echo.Context) error {
	userIdLogin := middlewares.ExtractTokenUserId(c)
	if userIdLogin == 0 {
		return c.JSON(http.StatusUnauthorized, responses.WebResponse("Unauthorized user", nil))
	}

	bookingId := c.Param("booking_id")

	updateBookingStatus := CancelBookingRequest{}
	errBind := c.Bind(&updateBookingStatus)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}
	fmt.Println("Booking id::", bookingId)
	bookingCore := CancelRequestToCoreBooking(updateBookingStatus)
	errCancel := handler.bookingService.CancelBooking(userIdLogin, bookingId, bookingCore)
	if errCancel != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error cancel booking "+errCancel.Error(), nil))
	}

	return c.JSON(http.StatusOK, responses.WebResponse("success cancel booking", nil))
}

func (handler *BookingHandler) WebhoocksNotification(c echo.Context) error {

	var webhoocksReq = WebhoocksRequest{}
	errBind := c.Bind(&webhoocksReq)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, responses.WebResponse("error bind data. data not valid", nil))
	}

	bookingCore := WebhoocksRequestToCore(webhoocksReq)
	err := handler.bookingService.WebhoocksService(bookingCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.WebResponse("error Notif "+err.Error(), nil))
	}

	log.Println("transaction success")
	return c.JSON(http.StatusOK, responses.WebResponse("transaction success", nil))
}
