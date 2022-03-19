package controllers

import "konseki-be/services"

type CreateEventRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateEventResponse struct {
	Name string `json:"name"`
	Id    string `json:"id"`
}

type GetEventsResponse struct {
	Events []services.Event `json:"events"`
}

type ResponseMessage struct {
	Id string `json:"id"`
	Message string `json:"message"`
	ExposureDate int64 `json:"exposure_date"`
	MessageType string `json:"message_type"`
	CreatedDate int64 `json:"created_date"`
	IsImportant bool `json:"is_important"`
}

type GetMessagesResponse struct {
	Messages []ResponseMessage `json:"messages"`
}

type GetProfileInfoResponse struct {
	Email string `json:"email"`
	Name string `json:"name"`
	Id string `json:"id"`
	Bio string `json:"bio"`
}

type JoinEventRequest struct {
	Id string `json:"id" binding:"required"`
}

type JoinEventResponse struct {
	IsSuccess bool `json:"is_success"`
	EventName string `json:"event_name"`
	Id string `json:"id"`
}

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`
}

type RegisterResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Id    string `json:"id"`
}

type Profile struct {
	Id string `json:"id"`
	Name string `json:"name"`
	IsInfected bool `json:"is_infected"`
	Email string `json:"email"`
	Bio string `json:"bio"`
}

type UpdateStatusRequest struct {
	Date int64 `json:"date" binding:"required"`
}

type UpdateStatusResponse struct {
	IsSuccess bool `json:"is_success"`
	Id string `json:"id"`
	Date int64 `json:"date"`
}

type Infections struct {
	UserId string `json:"user_id"`
	Date int64 `json:"date"`
}
