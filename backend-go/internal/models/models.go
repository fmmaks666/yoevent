package models

import (
	"time"

	"github.com/fmmaks666/yoevent-backend/internal/utils"
	"gorm.io/gorm"
)

const PerPage = 10

type Event struct {
	gorm.Model
	PublicID          string        `gorm:"unique;size:30;not null"`
	Title             string        `gorm:"not null"`
	Description       string        `gorm:"not null"`
	Date              *time.Time    `gorm:"nullable"`
	IsOnetime         *bool         `gorm:"not null"`
	Weekday           *time.Weekday `gorm:"nullable"`
	Time              *time.Time    `gorm:"nullable"`
	IsPrivate         *bool         `gorm:"not null"`
	NeedsRegistration bool          `gorm:"not null"`
	IsCancelled       bool          `gorm:"not null"`
	Visits            []Visit       `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE"` //`gorm:"foreignKey:EventID;references:ID"`
}

// For the sake of god, use JSON
func (e *Event) ToDTO() EventDTO {
	return EventDTO{
		PublicID: e.PublicID,
		EventEssentialDTO: EventEssentialDTO{
			Title:             e.Title,
			Description:       e.Description,
			Date:              e.Date,
			IsOnetime:         e.IsOnetime,
			Weekday:           e.Weekday,
			Time:              e.Time,
			IsPrivate:         e.IsPrivate,
			NeedsRegistration: &e.NeedsRegistration,
			IsCancelled:       &e.IsCancelled,
		},
	}
}

type EventQueryResult struct {
	Event
	LastVisit *time.Time
}

type Visitor struct {
	gorm.Model
	Hash            string    `gorm:"unique;size:256;not null;secondaryKey"`
	FirstName       string    `gorm:"not null"`
	LastName        string    `gorm:"not null"`
	Patronymic      string    `gorm:"not null"`
	Birthdate       time.Time `gorm:"not null"`
	Sex             utils.Sex `gorm:"type:tinyint;default:0;not null"`
	PhoneNumber     string    `gorm:"not null"`
	IsLocal         bool      `gorm:"not null"`
	IsDisabled      bool      `gorm:"not null"`
	AgreedToPrivacy bool      `gorm:"not null;default:false"`
	Visits          []Visit   `gorm:"foreignKey:VisitorID;constraint:OnDelete:CASCADE,"` //`gorm:"foreignkey:VisitorID;references:ID"`
	// firstTimer bool
}

func (v *Visitor) ToDTO() VisitorDTO {
	return VisitorDTO{
		Hash: v.Hash,
		VisitorEssentialDTO: VisitorEssentialDTO{
			FirstName:       v.FirstName,
			LastName:        v.LastName,
			Patronymic:      v.Patronymic,
			Birthdate:       v.Birthdate,
			Sex:             v.Sex,
			PhoneNumber:     v.PhoneNumber,
			IsLocal:         &v.IsLocal,
			IsDisabled:      &v.IsDisabled,
			AgreedToPrivacy: &v.AgreedToPrivacy,
		},
	}
}

func (v *Visitor) GetFirstName() string    { return v.FirstName }
func (v *Visitor) GetLastName() string     { return v.LastName }
func (v *Visitor) GetPatronymic() string   { return v.Patronymic }
func (v *Visitor) GetBirthdate() time.Time { return v.Birthdate }
func (v *Visitor) GetSex() utils.Sex       { return v.Sex }
func (v *Visitor) GetPhoneNumber() string  { return v.PhoneNumber }
func (v *Visitor) GetIsLocal() bool        { return v.IsLocal }
func (v *Visitor) GetIsDisabled() bool     { return v.IsDisabled }

type Visit struct {
	gorm.Model
	EventID   uint      `gorm:"not null"`
	VisitorID uint      `gorm:"not null"`
	VisitDate time.Time `gorm:"not null"`

	Event   Event   `gorm:"foreignKey:EventID;"`
	Visitor Visitor `gorm:"foreignKey:VisitorID;"`
}

func (vi *Visit) ToDTO() VisitDTO {
	return VisitDTO{
		VisitID:   vi.Model.ID, // UMMMM
		EventID:   vi.EventID,
		VisitorID: vi.VisitorID,
		VisitDate: vi.VisitDate,
	}
}

type VisitWithAge struct {
	Visit
	Age int
}

func createVisitsView(db *gorm.DB) {
	// LOVE Hardcoding table names lol
	db.Exec(`CREATE VIEW IF NOT EXISTS visits_with_age AS
		SELECT vi.*, 
		(
			(strftime('%Y', vi.visit_date) - strftime('%Y', v.birthdate)) -
			(strftime('%m-%d', vi.visit_date) < strftime('%m-%d', v.birthdate))
		) as age
		FROM visits vi
		LEFT JOIN visitors v ON v.id = vi.visitor_id`)
}

func Setup(db *gorm.DB) {
	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Visitor{})
	db.AutoMigrate(&Visit{})
	createVisitsView(db)
}
