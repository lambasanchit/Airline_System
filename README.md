# Airline Booking System

A simple REST API built in Go for managing users, flights, and bookings with JWT authentication.  
This system allows users to register, login, add flights, search flights, and book seats on flights securely.

---

## Features

- User registration and login with JWT authentication  
- Add new flights with details like source, destination, departure time, and total seats  
- Search for flights by source and destination  
- Book seats on available flights  
- All protected endpoints require a valid JWT token  

---

## Installation

1. Clone the repository:
   git clone <your-repository-url>
   cd airline-system
Run the Go server:

go run main.go routes.go
The server will start and listen on http://localhost:8080.

API Usage
1. Register a New User
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123"
  }'
Response example:

json
{
  "id": "user-uuid",
  "name": "John Doe",
  "email": "john@example.com"
}
2. Login and Get JWT Token
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
Response example:

json
{
  "token": "your.jwt.token.here"
}
Save the token for authenticated requests below.

3. Add a New Flight (Requires Authentication)
curl -X POST http://localhost:8080/flights \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{
    "source": "DEL",
    "destination": "BOM",
    "departure_time": "2025-07-06T10:00:00Z",
    "total_seats": 5
  }'
Response example:

json
{
  "id": "flight-uuid",
  "source": "del",
  "destination": "bom",
  "departure_time": "2025-07-06T10:00:00Z",
  "total_seats": 5,
  "available_seats": 5
}
4. Search Flights (Requires Authentication)
curl -X GET "http://localhost:8080/flights/search?source=DEL&destination=BOM" \
  -H "Authorization: Bearer <JWT_TOKEN>"
Response example:

json
[
  {
    "id": "flight-uuid",
    "source": "del",
    "destination": "bom",
    "departure_time": "2025-07-06T10:00:00Z",
    "total_seats": 5,
    "available_seats": 5
  }
]
5. Book Seats on a Flight (Requires Authentication)
curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -d '{
    "flight_id": "<FLIGHT_ID>",
    "seats": 2
  }'
Response example:

json
{
  "id": "booking-uuid",
  "user_id": "user-uuid",
  "flight_id": "<FLIGHT_ID>",
  "seat_no": 1,
  "status": "Pending",
  "created_at": "timestamp"
}
Testing
Use the above curl commands step-by-step to test your API.

Always use the JWT token received from login for all authenticated endpoints.

Make sure your server is running on localhost:8080.

Code Structure :
main.go - Entry point, initializes services and routes

routes.go - Defines HTTP routes and handlers

handlers/ - Contains HTTP handlers for users, flights, bookings

services/ - Business logic for users, flights, bookings

JWT middleware for protecting routes

.

Just replace placeholders like `<your-repository-url>`, `<JWT_TOKEN>`, and `<FLIGHT_ID>` with your actual values.  

