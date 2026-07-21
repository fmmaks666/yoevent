package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"strconv"
	"strings"
	"time"

	"github.com/gosimple/slug"
)

type Sex uint8

const (
	SexOther Sex = iota
	SexMale
	SexFemale
)

// This fucking Go wants me dead
type HashableVisitor interface {
	GetFirstName() string
	GetLastName() string
	GetPatronymic() string
	GetBirthdate() time.Time
	GetSex() Sex
	GetPhoneNumber() string
	GetIsLocal() bool
	GetIsDisabled() bool
	GetAgreedToPrivacy() bool
	// firstTimer bool
}

// These two functions will be in my worst nightmares
// TODO: Make an interface to FUCK these duplicates
func HashVisitor(a HashableVisitor, salt string) string {
	b := strings.Builder{}
	b.WriteString(a.GetFirstName())
	b.WriteString(a.GetLastName())
	b.WriteString(a.GetPatronymic())
	b.WriteString(a.GetBirthdate().String())
	b.WriteString(string(a.GetPhoneNumber()))
	b.WriteString(string(a.GetSex()))
	b.WriteString(strconv.FormatBool(a.GetIsLocal()))
	b.WriteString(strconv.FormatBool(a.GetIsDisabled()))
	b.WriteString(strconv.FormatBool(a.GetAgreedToPrivacy()))
	b.WriteString(salt)

	hash := sha1.Sum([]byte(b.String()))
	return hex.EncodeToString(hash[:])
}

func TimeToSeconds(t time.Time) int {
	var h, m int
	h, m, _ = t.Clock()

	return (h * 60 * 60) + (m * 60)
}

func ClosestDate(w time.Weekday, t time.Time, targetW time.Weekday, targetT time.Time) time.Time {
	now := time.Now()
	days := time.Duration((now.Weekday()-targetW+6)%7) + 1
	timeNow := TimeToSeconds(now)
	targetTime := TimeToSeconds(targetT)
	if now.Weekday() == targetW && timeNow >= targetTime {
		days = 0
	}
	date := now.Add(-days * 24 * time.Hour)
	prevDate := time.Date(
		date.Year(), date.Month(), date.Day(),
		targetT.Hour(), targetT.Minute(), targetT.Second(), targetT.Nanosecond(),
		targetT.Location(),
	)
	return prevDate
}

func Slugify(s string) string {
	return slug.Make(s)
}
