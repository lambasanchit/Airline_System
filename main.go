package main

import (
	"log"
	"net/http"

	"airline-system/handlers"
	"airline-system/repository"
	"airline-system/services"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize repositories
	userRepo := repository.NewUserRepository()
	flightRepo := repository.NewFlightRepository()
	bookingRepo := repository.NewBookingRepository()
	paymentRepo := repository.NewPaymentRepository()

	// Initialize services
	authService := services.NewAuthService(userRepo)
	flightService := services.NewFlightService(flightRepo)
	bookingService := services.NewBookingService(bookingRepo, flightRepo)
	paymentService := services.NewPaymentService(bookingRepo, paymentRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	flightHandler := handlers.NewFlightHandler(flightService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	// Setup router
	r := mux.NewRouter()

	// Register routes
	RegisterRoutes(r, authHandler, flightHandler, bookingHandler, paymentHandler)

	// Start server
	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
