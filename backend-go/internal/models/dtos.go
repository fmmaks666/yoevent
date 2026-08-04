package models

import (
	//"encoding/json"
	"time"

	"github.com/fmmaks666/yoevent-backend/internal/utils"
)

type EventDTO struct {
	PublicID string `json:"event_id" binding:"required" form:"event_id"`
	EventEssentialDTO
}

type EventEssentialDTO struct {
	Title             string        `json:"title" binding:"required" form:"title"`
	Description       string        `json:"description" binding:"required" form:"description"`
	Date              *time.Time    `json:"date" form:"date"`
	IsOnetime         *bool         `json:"is_onetime" binding:"required" form:"is_onetime"`
	Weekday           *time.Weekday `json:"weekday" form:"weekday"`
	Time              *time.Time    `json:"time" form:"time"`
	IsPrivate         *bool         `json:"is_private" binding:"required" form:"is_private"`
	NeedsRegistration *bool         `json:"needs_registration" binding:"required" form:"needs_registration"`
	IsCancelled       *bool         `json:"is_cancelled" binding:"required" form:"is_cancelled"`
	LastVisit         *time.Time    `json:"last_visit"` // A meh thingy, just for one place
}

type VisitorDTO struct {
	Hash string `json:"hash" binding:"required" form:"hash"`
	// firstTimer bool `json:"first_timer"`
	VisitorEssentialDTO
}

type VisitorEssentialDTO struct {
	FirstName       string    `json:"first_name" binding:"required" form:"first_name"`
	LastName        string    `json:"last_name" binding:"required" form:"last_name"`
	Patronymic      string    `json:"patronymic" binding:"required" form:"patronymic"`
	Birthdate       time.Time `json:"birthdate" binding:"required" form:"birthdate"`
	Sex             utils.Sex `json:"sex" binding:"required" form:"sex"`
	PhoneNumber     string    `json:"phone_number" binding:"required" form:"phone_number"`
	IsLocal         *bool     `json:"is_local" binding:"required" form:"is_local"`
	Residence       *string   `json:"residence" form:"residence"`
	IsDisabled      *bool     `json:"is_disabled" binding:"required" form:"is_disabled"`
	AgreedToPrivacy *bool     `json:"agreed_to_privacy" binding:"required" form:"agreed_to_privacy"`
}

func (v *VisitorEssentialDTO) GetFirstName() string     { return v.FirstName }
func (v *VisitorEssentialDTO) GetLastName() string      { return v.LastName }
func (v *VisitorEssentialDTO) GetPatronymic() string    { return v.Patronymic }
func (v *VisitorEssentialDTO) GetBirthdate() time.Time  { return v.Birthdate }
func (v *VisitorEssentialDTO) GetSex() utils.Sex        { return v.Sex }
func (v *VisitorEssentialDTO) GetPhoneNumber() string   { return v.PhoneNumber }
func (v *VisitorEssentialDTO) GetIsLocal() bool         { return *v.IsLocal }
func (v *VisitorEssentialDTO) GetResidence() *string    { return v.Residence }
func (v *VisitorEssentialDTO) GetIsDisabled() bool      { return *v.IsDisabled }
func (v *VisitorEssentialDTO) GetAgreedToPrivacy() bool { return *v.AgreedToPrivacy }

type VisitDTO struct {
	VisitID   uint      `json:"visit_id"`
	EventID   uint      `json:"event_id"`
	VisitorID uint      `json:"visitor_id"`
	VisitDate time.Time `json:"visit_date"`

	Event EventDTO `json:"event"`
}

type DateRange struct {
	Month time.Month `json:"month" form:"month"`
	Year  int        `json:"year" form:"year"`
}
