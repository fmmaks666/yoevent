package models

import "time"

type CreateVisitRequest struct {
	EventID string `json:"event_id" binding:"required" form:"event_id"`
	Hash    string `json:"visitor" binding:"required" form:"visitor"`
}

type CreateEventRequest struct {
	EventEssentialDTO
}

type CreateVisitorRequest struct {
	VisitorEssentialDTO
}

type CreateRegistrationRequest struct {
	EventID          string     `json:"event_id" binding:"required" form:"event_id"`
	Title            string     `json:"title" binding:"required" form:"title"`
	Description      string     `json:"description" binding:"required" form:"description"`
	MaxRegistrations uint       `json:"max_registrations" binding:"required" form:"max_registrations"`
	Until            *time.Time `json:"until" binding:"required" form:"until"`
	IsPrivate        *bool      `json:"is_private" binding:"required" form:"is_private"`
}

type UpdateRegistrationRequest = CreateRegistrationRequest

type GetVisitorVisitsRequest struct {
	Hash string `json:"hash"`
	VisitorEssentialDTO
}

/*
type UpdateVisitorRequest struct {
	Hash            string    `json:"hash"`
	FirstName       string    `json:"first_name" form:"first_name"`
	LastName        string    `json:"last_name" form:"last_name"`
	Patronymic      string    `json:"patronymic" form:"patronymic"`
	Age             int       `json:"age" form:"age"`
	Sex             utils.Sex `json:"sex" form:"sex"`
	PhoneNumber     string    `json:"phone_number" form:"phone_number"`
	IsLocal         *bool     `json:"is_local" form:"is_local"`
	IsDisabled      *bool     `json:"is_disabled" form:"is_disabled"`
	AgreedToPrivacy *bool     `json:"agreed_to_privacy" form:"agreed_to_privacy"`
}
*/

type UpdateVisitorRequest struct {
	VisitorDTO
}

type UpdateEventRequest struct {
	EventDTO
}

type GetStatsRequest struct {
	EventID string `json:"event_id" binding:"required" form:"event_id"`
	All     *bool  `json:"all" binding:"required" form:"all"`
	DateRange
}

type GetVisitsRequest struct {
	EventID string `json:"event_id" binding:"required" form:"event_id"`
	All     *bool  `json:"all" binding:"required" form:"all"`
	DateRange
}

type DeleteEvent struct {
	EventID string `json:"event_id" binding:"required" form:"event_id"`
}
