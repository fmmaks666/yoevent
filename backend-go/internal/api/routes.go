package api

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/fmmaks666/yoevent-backend/internal/models"
	"github.com/fmmaks666/yoevent-backend/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Handler struct {
	db   *gorm.DB
	salt string
}

func (h *Handler) sendError(ctx *gin.Context, code int, err string) {
	ctx.JSON(code, gin.H{"error": err})
}

// /event?eventId
func (h *Handler) getEvent(ctx *gin.Context) {
	eventId := ctx.Query("event")
	if eventId == "" {
		h.sendError(ctx, 404, "Invalid EventID")
		return
	}
	event, err := gorm.G[models.Event](h.db).Where("public_id = ?", eventId).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return

	}
	dto := event.ToDTO()
	//buf, err := json.Marshal(dto)
	//if err != nil {
	//	h.sendError(ctx, 400, "Failed to encode the event")
	//	return
	//}
	ctx.JSON(200, dto)
}

func (h *Handler) getEventsRaw(ctx *gin.Context, hidePrivate bool) {
	var err error
	/* pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	} */
	visitor := ctx.Query("visitor")

	//offset := (page - 1) * models.PerPage
	// WHO THE FUCK WROTE THIS? OH IT WAS ME
	var events []models.EventQueryResult
	if visitor == "" {
		// I'll give you a chuu, GORM
		// MY LOVELY HACKS
		base := gorm.G[models.Event](h.db).Select("events.*, 0 as is_visited").Where("1 = 1") /* Offset(offset).Limit(models.PerPage).*/
		if hidePrivate {
			base = base.Where("is_private = ?", false)
		}
		err = base.Order("date ASC").Scan(ctx, &events)
	} else {
		base := gorm.G[models.Event](h.db).Select(
			"events.*, "+
				"(SELECT visits.visit_date FROM visits "+
				"LEFT JOIN visitors ON visits.visitor_id = visitors.id "+
				"WHERE visits.event_id = events.id AND visitors.hash = ? "+
				"ORDER BY visit_date DESC LIMIT 1) as last_visit ", visitor).Where("1 = 1") // HACKS, take my CHUU!
		if hidePrivate {
			base = base.Where("is_private = ?", false)
		}

		err = base.Scan(ctx, &events)
	}

	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}
	var dtos []models.EventDTO
	var dto models.EventDTO
	for _, e := range events {
		dto = e.ToDTO()
		dto.LastVisit = e.LastVisit
		dtos = append(dtos, dto)
	}

	ctx.JSON(200, dtos)
}

func (h *Handler) getEvents(ctx *gin.Context) {
	h.getEventsRaw(ctx, true)
}

func (h *Handler) getEventsAdmin(ctx *gin.Context) {
	h.getEventsRaw(ctx, false)
}

// /visit?eventId,visitor={}
// TODO: Return safe data
func (h *Handler) createVisit(ctx *gin.Context) {
	var req models.CreateVisitRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	// TODO: Check whether the event exists???
	eventId := req.EventID
	event, err := gorm.G[models.Event](h.db).Where("public_id = ?", eventId).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	// TODO: Unhardcode?

	// One day timeout
	// What the fuck

	loc := time.UTC

	timeout := time.Date(1970, time.January, 2, 0, 0, 0, 0, loc).Sub(time.Date(1970, time.January, 1, 0, 0, 0, 0, loc))
	if *event.IsOnetime && event.Date.Compare(time.Now()) == 1 || event.IsCancelled {
		h.sendError(ctx, 400, "The event didn't yet happen or is cancelled")
		return
	} /* else if *event.IsOnetime && time.Now().After(event.Date.Add(timeout)) {
		h.sendError(ctx, 400, "It's too late to add a visit")
		return
	} */

	// TODO: Holy fuck's sake optimize this
	var date time.Time
	if *event.IsOnetime {
		date = *event.Date
	} else {
		now := time.Now().UTC()
		date = utils.ClosestDate(now.Weekday(), now, *event.Weekday, *event.Time)
	}
	var v models.Visitor
	v, err = gorm.G[models.Visitor](h.db).Where("hash = ?", req.Hash).First(ctx)
	//v, err = h.FindOrCreateVisitor(ctx, req.VisitorEssentialDTO)
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}
	_, err = gorm.G[models.Visit](h.db).Where("event_id = ? AND visitor_id = ? AND visit_date = ?", event.ID, v.ID, date).First(ctx) // HUH, Okay gotta use this context

	if err == nil {
		h.sendError(ctx, 403, "The visit already exists")
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		h.sendError(ctx, 404, err.Error())
		return
	}

	err = gorm.G[models.Visit](h.db).Create(ctx, &models.Visit{
		EventID:   event.ID,
		VisitorID: v.ID,
		VisitDate: date,
	})
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}
	var vi models.Visit
	vi, err = gorm.G[models.Visit](h.db).Where("visitor_id = ? AND event_id = ?", v.ID, event.ID).Preload("Event", nil).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	var visit models.CreateVisitResponse
	visit.VisitID = vi.ID // TODO:: FOR THE SAKE OF GOD DON'T EXPOSE THE FING NUMBER
	visit.EventID = vi.Event.PublicID
	visit.Visitor = v.Hash
	visit.VisitDate = date
	visit.Event = vi.Event.ToDTO().EventEssentialDTO
	ctx.JSON(200, visit)
}

func (h *Handler) createVisitor(ctx *gin.Context) {
	var req models.CreateVisitorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	dto := req.VisitorEssentialDTO
	hash := utils.HashVisitor(&dto, h.salt)

	err := gorm.G[models.Visitor](h.db).Create(ctx, &models.Visitor{
		Hash:            hash,
		FirstName:       dto.FirstName,
		LastName:        dto.LastName,
		Patronymic:      dto.Patronymic,
		Birthdate:       dto.Birthdate,
		Sex:             dto.Sex,
		PhoneNumber:     dto.PhoneNumber,
		IsLocal:         *dto.IsLocal,
		Residence:       dto.Residence,
		IsDisabled:      *dto.IsDisabled,
		AgreedToPrivacy: *dto.AgreedToPrivacy,
	})

	if err != nil {
		var sqlErr sqlite3.Error
		// Risky af
		if !errors.As(err, &sqlErr) || sqlErr.Code != sqlite3.ErrNo(sqlite3.ErrConstraint) {

			h.sendError(ctx, 400, err.Error())
			return
		}
	}
	ctx.JSON(200, models.CreateVisitorResponse{Hash: hash, VisitorEssentialDTO: dto})
}

func (h *Handler) updateVisitor(ctx *gin.Context) {
	var req models.UpdateVisitorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	var err error

	newHash := utils.HashVisitor(&req.VisitorDTO, h.salt)
	_, err = gorm.G[models.Visitor](h.db).Where("hash = ?", req.Hash).Updates(ctx, models.Visitor{
		Hash:            newHash,
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Patronymic:      req.Patronymic,
		Birthdate:       req.Birthdate,
		Sex:             req.Sex,
		PhoneNumber:     req.PhoneNumber,
		IsLocal:         *req.IsLocal,
		Residence:       req.Residence,
		IsDisabled:      *req.IsDisabled,
		AgreedToPrivacy: *req.AgreedToPrivacy,
	})

	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}
	ctx.JSON(200, models.CreateVisitorResponse{
		Hash:                newHash,
		VisitorEssentialDTO: req.VisitorEssentialDTO,
	})
}

func (h *Handler) getVisitorVisits(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	hash := ctx.Query("visitor")
	if hash == "" {
		h.sendError(ctx, 400, "Invalid visitor")
		return

	}

	offset := (page - 1) * models.PerPage
	var visits []models.Visit
	// Order by EVENT.Date or by CreatedAt?
	//visits, err = gorm.G[models.Visit](h.db).Order("created_at asc").Joins(clause.LeftJoin.Association("visitor_id"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
	//		db.Where("hash = ?", hash)
	//		return nil
	//}).Offset(offset).Limit(models.PerPage).Find(ctx)
	// TODO: THIS SHIT'S UNREADABLE
	// TODO: Handle month/all
	visits, err = gorm.G[models.Visit](h.db).Joins(clause.InnerJoin.Association("Visitor"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
		db.Where("hash = ?", hash)
		return nil
	}).Joins(clause.InnerJoin.Association("Event"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
		return nil
	}).Preload("Event", nil).Offset(offset).Limit(models.PerPage).Find(ctx)

	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}
	var res []models.GetVisitorVisitResponse
	for _, v := range visits {
		res = append(res, models.GetVisitorVisitResponse{
			VisitDTO: models.VisitDTO{
				VisitID:   v.ID,
				VisitDate: v.VisitDate,
			},
			Visitor: hash,
			Event:   v.Event.ToDTO(),
		})
	}

	ctx.JSON(200, res)
}

func (h *Handler) checkAuth(ctx *gin.Context) {
	ctx.JSON(204, gin.H{"success": true})
}

func (h *Handler) getVisits(ctx *gin.Context) {
	var req models.GetVisitsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	var visits []models.VisitWithAge
	var err error
	// Order by EVENT.Date or by CreatedAt?
	//visits, err = gorm.G[models.Visit](h.db).Order("created_at asc").Joins(clause.LeftJoin.Association("visitor_id"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
	//		db.Where("hash = ?", hash)
	//		return nil
	//}).Offset(offset).Limit(models.PerPage).Find(ctx)

	base := gorm.G[models.Visit](h.db).Table("visits_with_age").Joins(clause.InnerJoin.Association("Event"), func(db gorm.JoinBuilder, joinTable, curTable clause.Table) error {
		db.Where("public_id = ?", req.EventID)
		return nil
	})

	if !*req.All {
		base = base.Where("strftime('%m', visit_date) = ? AND strftime('%Y', visit_date) = ?", fmt.Sprintf("%02d", req.Month), strconv.Itoa(req.Year))
	}

	err = base.Preload("Visitor", nil).Scan(ctx, &visits)

	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}
	buff := new(bytes.Buffer)

	writer := csv.NewWriter(buff)
	// NOTE: Remember to edit this, whilst I have this terrible architecture, icha icha shite

	writer.Write([]string{
		"Прізвище", "Ім'я", "По батькові",
		"Номер телефону",
		"Вік", "Стать", "Наявність інвалідності",
		"Соц. статус", "Місто проживання",
		"Згода на обр. даних", "Дата відвідування",
	})
	for _, vi := range visits {
		v := vi.Visitor
		sex := "Жінка"
		if v.Sex == utils.SexMale {
			sex = "Чоловік"
		}
		disability := "Ні"
		if v.IsDisabled {
			disability = "Так"
		}
		status := "Місцевий мешканець"
		if !v.IsLocal {
			status = "ВПО"
		}
		residence := "Не вказано"
		if v.Residence != nil && strings.Trim(*v.Residence, " ") != "" {
			residence = *v.Residence
		}
		privacy := "Ні"
		if v.AgreedToPrivacy {
			privacy = "Так"
		}
		// TODO: If an event is !onetime, check for visits for each possible date
		writer.Write([]string{v.LastName, v.FirstName, v.Patronymic, v.PhoneNumber, strconv.Itoa(vi.Age), sex, disability, status, residence, privacy, vi.VisitDate.Format(time.RFC822)})
	}

	writer.Flush()
	ctx.JSON(200, gin.H{
		"csv": buff.String(),
	})
}

func (h *Handler) getStats(ctx *gin.Context) {
	var req models.GetStatsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	event, err := gorm.G[models.Event](h.db).Where("public_id = ?", req.EventID).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 400, err.Error())
	}

	id := event.ID
	stats := models.GetEventStatsResponse{}

	base := gorm.G[models.Visit](h.db).Table("visits_with_age").Distinct("visitor_id").Where("event_id = ?", id).Joins(clause.LeftJoin.Association("Visitor"), nil)
	if !*req.All {
		base = base.Where("strftime('%m', visit_date) = ? AND strftime('%Y', visit_date) = ?", fmt.Sprintf("%02d", req.Month), strconv.Itoa(req.Year))
	}
	male := base.Where("sex = ?", utils.SexMale)
	female := base.Where("sex = ?", utils.SexFemale)

	// SAVE ME GOD FROM THESE ERRORS
	// THIS IS THE FUCKING BIGESST NIGHTMARE I WROTE IN MY DICK SHORT LIFE
	// Calculate age:
	// (VisitDate.YEAR - BirthDate.YEAR) - (VisitDate.Month.Day - BirthDay.Month.Day)
	stats.Men, err = male.Count(ctx, "visitor_id")
	stats.Women, err = female.Count(ctx, "visitor_id")
	stats.MenBelowThirty, err = male.Where("age < ?", 30).Count(ctx, "visitor_id")
	stats.WomenBelowThirty, err = female.Where("age < ?", 30).Count(ctx, "visitor_id")
	stats.DisabledMen, err = male.Where("is_disabled = ?", true).Count(ctx, "visitor_id")
	stats.DisabledWomen, err = female.Where("is_disabled = ?", true).Count(ctx, "visitor_id")
	stats.NonLocalMen, err = male.Where("is_local = ?", false).Count(ctx, "visitor_id")
	stats.NonLocalWomen, err = female.Where("is_local = ?", false).Count(ctx, "visitor_id")
	stats.AgeLessThanFifteen, err = base.Where("age < ?", 15).Count(ctx, "visitor_id")
	stats.AgeLessThanTwentyNine, err = base.Where("age >= ? AND age < ?", 15, 30).Count(ctx, "visitor_id")
	stats.AgeMoreThanThirty, err = base.Where("age >= ? AND age < ?", 30, 35).Count(ctx, "visitor_id")
	stats.AgeMoreThanThirtyFive, err = base.Where("age >= ?", 35).Count(ctx, "visitor_id")
	stats.Members, err = base.Count(ctx, "visitor_id")
	stats.Occurences, err = base.Distinct("visit_date").Count(ctx, "visit_date")
	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}

	ctx.JSON(200, stats)
}

func (h *Handler) getVisitors(ctx *gin.Context) {
	var err error
	/* pageStr := ctx.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	} */

	//offset := (page - 1) * models.PerPage
	var visitors []models.Visitor

	base := gorm.G[models.Visitor](h.db)
	err = base.Scan(ctx, &visitors)

	if err != nil {
		h.sendError(ctx, 400, err.Error())
		return
	}
	var dtos []models.VisitorDTO
	var dto models.VisitorDTO
	for _, e := range visitors {
		dto = e.ToDTO()
		dtos = append(dtos, dto)
	}

	ctx.JSON(200, dtos)
}

func (h *Handler) createEvent(ctx *gin.Context) {
	var req models.CreateEventRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	// TODO: Event more fking validation?
	if *req.IsOnetime && req.Date == nil {
		h.sendError(ctx, 404, "Event is onetime but not date specified")
		return
	} else if !*req.IsOnetime && req.Weekday == nil && req.Time == nil {
		h.sendError(ctx, 404, "Event is not one time but no weekday and/or time specified")
		return
	}

	// TODO: Use a Transaction
	dateStr := "regular"
	if *req.IsOnetime && req.Date != nil {
		dateStr = req.Date.Format(time.ANSIC)
	}
	id := utils.Slugify(req.Title + "-" + dateStr)
	date := req.Date
	if !*req.IsOnetime {
		date = nil
	}
	time := req.Time
	if *req.IsOnetime {
		time = nil
	}
	weekday := req.Weekday
	if *req.IsOnetime {
		weekday = nil
	}
	err := gorm.G[models.Event](h.db).Create(ctx, &models.Event{
		PublicID:          id,
		Title:             req.Title,
		Description:       req.Description,
		Date:              date,
		IsOnetime:         req.IsOnetime,
		Weekday:           weekday,
		Time:              time,
		IsPrivate:         req.IsPrivate,
		NeedsRegistration: *req.NeedsRegistration,
		IsCancelled:       *req.IsCancelled,
	}) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	event, err := gorm.G[models.Event](h.db).Where("public_id = ?", id).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	ctx.JSON(200, event.ToDTO())
}

func (h *Handler) updateEvent(ctx *gin.Context) {
	var req models.UpdateEventRequest

	var err error
	if err = ctx.ShouldBindJSON(&req); err != nil {
		h.sendError(ctx, 400, "Malformed request: "+err.Error())
		return
	}

	// TODO: Event more fking validation?
	if *req.IsOnetime && req.Date == nil {
		h.sendError(ctx, 404, "Event is onetime but not date specified")
		return
	} else if !*req.IsOnetime && req.Weekday == nil && req.Time == nil {
		h.sendError(ctx, 404, "Event is not one time but no weekday and/or time specified")
		return
	}

	// TODO: Use a Transaction
	date := req.Date
	if !*req.IsOnetime {
		date = nil
	}
	time := req.Time
	if *req.IsOnetime {
		time = nil
	}
	weekday := req.Weekday
	if *req.IsOnetime {
		weekday = nil
	}

	_, err = gorm.G[models.Event](h.db).Where("public_id = ?", req.PublicID).Updates(ctx, models.Event{
		PublicID:          req.PublicID,
		Title:             req.Title,
		Description:       req.Description,
		Date:              date,
		IsOnetime:         req.IsOnetime,
		Weekday:           weekday,
		Time:              time,
		IsPrivate:         req.IsPrivate,
		NeedsRegistration: *req.NeedsRegistration,
		IsCancelled:       *req.IsCancelled,
	}) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	event, err := gorm.G[models.Event](h.db).Where("public_id = ?", req.PublicID).First(ctx) // HUH, Okay gotta use this context
	if err != nil {
		h.sendError(ctx, 404, err.Error())
		return
	}

	ctx.JSON(200, event.ToDTO())
}

func Setup(db *gorm.DB, adminPass, frontendUrl, salt string) *gin.Engine {
	handler := Handler{db, salt}

	router := gin.Default()
	conf := cors.Config{
		AllowOrigins:     []string{frontendUrl}, // TODO: I Don't need to say this but the thing into an env file
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept", "Authorization"},
		AllowCredentials: true,
	}
	router.Use(cors.New(conf)) // TODO: Get this to be secure

	router.GET("/event", handler.getEvent)
	router.GET("/events", handler.getEvents)
	router.GET("/visits", handler.getVisitorVisits)
	router.POST("/visitor", handler.createVisitor)
	router.PUT("/visitor", handler.updateVisitor)
	router.POST("/visit", handler.createVisit)
	admin := router.Group("/admin")
	admin.Use(gin.BasicAuth(gin.Accounts{
		"admin": adminPass,
	}))
	admin.GET("/check", handler.checkAuth)
	admin.GET("/events", handler.getEventsAdmin)
	admin.GET("/visits", handler.getVisits)
	admin.GET("/stats", handler.getStats)
	admin.GET("/visitors", handler.getVisitors)
	admin.POST("/event", handler.createEvent)
	admin.PUT("/event", handler.updateEvent)

	return router
}
