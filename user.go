package udetect

import "github.com/google/uuid"

// User object info
type User struct {
	UUID          uuid.UUID          `json:"uuid"`
	SessionID     string             `json:"sessid"`
	FingerPrintID string             `json:"fpid,omitempty"`
	ETag          string             `json:"etag,omitempty"`
	AgeStart      int                `json:"age_start,omitempty"` // Year of birth from
	AgeEnd        int                `json:"age_end,omitempty"`   // Year of birth to
	Keywords      string             `json:"keywords,omitempty"`  // Comma separated list of keywords, interests, or intent
	Interests     map[string]float64 `json:"interests,omitempty"` //
	Sex           map[int]float64    `json:"sex,omitempty"`       // 0 – undefined, 1 – man, 2 – woman
}

// MostPossibleSex detection
func (u *User) MostPossibleSex() int {
	if u == nil || u.Sex == nil {
		return -1
	}
	var (
		sex int
		val float64
	)
	for i, v := range u.Sex {
		if v > val {
			sex, val = i, v
		}
	} // end for
	if val >= .3 {
		return sex
	}
	return -1
}

// Age middle of user
func (u *User) Age() int {
	if u == nil || (u.AgeStart < 1 && u.AgeEnd < 1) {
		return -1
	}
	if u.AgeEnd > u.AgeStart {
		return (u.AgeEnd + u.AgeStart) / 2
	}
	return u.AgeStart
}
