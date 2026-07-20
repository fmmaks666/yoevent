package models

import "time"

// TODO: This is a HUH
type CreateVisitResponse struct {
	VisitID   uint      `json:"visit_id"`
	EventID   string    `json:"event_id"`
	Visitor   string    `json:"visitor"`
	VisitDate time.Time `json:"visit_date"`

	Event EventEssentialDTO `json:"event"`
}

type GetVisitorVisitResponse struct {
	Visitor string `json:"visitor"`
	VisitDTO

	Event EventDTO `json:"event"`
}

type CreateVisitorResponse struct {
	Hash string `json:"visitor"`
	VisitorEssentialDTO
}

type GetEventStatsResponse struct {
	Organizers            int64 `json:"organizers"`
	Members               int64 `json:"members"`
	Men                   int64 `json:"men"`
	Women                 int64 `json:"women"`
	WomenBelowThirty      int64 `json:"women_below_30"`
	MenBelowThirty        int64 `json:"men_below_30"`
	DisabledWomen         int64 `json:"disabled_women"`
	DisabledMen           int64 `json:"disabled_men"`
	NonLocalWomen         int64 `json:"nonlocal_women"`
	NonLocalMen           int64 `json:"nonlocal_men"`
	AgeLessThanFifteen    int64 `json:"age_less_than_15"`
	AgeLessThanTwentyNine int64 `json:"age_less_than_29"`
	AgeMoreThanThirty     int64 `json:"age_more_than_30"`
	AgeMoreThanThirtyFive int64 `json:"age_more_than_35"`
	Occurences            int64 `json:"occurences"`
}
