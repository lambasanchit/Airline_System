package main

import (
	"airline-system/handlers"
	"airline-system/middleware"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all routes with corresponding handlers and applies middleware
func RegisterRoutes(
	r *mux.Router,
	authHandler *handlers.AuthHandler,
	flightHandler *handlers.FlightHandler,
	bookingHandler *handlers.BookingHandler,
	paymentHandler *handlers.PaymentHandler,
) {
	// Public routes
	r.HandleFunc("/register", authHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", authHandler.LoginUser).Methods("POST")

	// Flights routes (protected)
	flightRouter := r.PathPrefix("/flights").Subrouter()
	flightRouter.Use(middleware.JWTAuthMiddleware)
	flightRouter.HandleFunc("", flightHandler.AddFlight).Methods("POST")
	flightRouter.HandleFunc("/search", flightHandler.SearchFlights).Methods("GET")

	// Booking routes (protected)
	bookingRouter := r.PathPrefix("/bookings").Subrouter()
	bookingRouter.Use(middleware.JWTAuthMiddleware)
	bookingRouter.HandleFunc("", bookingHandler.CreateBooking).Methods("POST")
	bookingRouter.HandleFunc("/{id}/cancel", bookingHandler.CancelBooking).Methods("POST")

	// Payment routes (protected)
	paymentRouter := r.PathPrefix("/payment").Subrouter()
	paymentRouter.Use(middleware.JWTAuthMiddleware)
	paymentRouter.HandleFunc("/pay", paymentHandler.MakePayment).Methods("POST")
}
