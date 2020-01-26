package validateInput

import (
	"github.com/nattigy/parentschoolcommunicationsystem/models"
	"regexp"
)

var PhoneRX = regexp.MustCompile("(^\\+[0-9]{2}|^\\+[0-9]{2}\\(0\\)|^\\(\\+[0-9]{2}\\)\\(0\\)|^00[0-9]{2}|^0)([0-9]{9}$|[0-9\\-\\s]{10}$)")

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var StringRX = regexp.MustCompile("^[a-zA-Z0-9 _]*[A-Za-z0-9][A-Za-z0-9 _]*$")

type ValidationErrors map[string][]string

type Input struct {
	VErrors ValidationErrors
}

func (inVal Input) MatchesPattern(field string, pattern *regexp.Regexp) {
	if field == "" {
		return
	}
	if !pattern.MatchString(field) {
		inVal.VErrors.Add(field, "The value entered is invalid")
	}
}

func (ve ValidationErrors) Add(field, message string) {
	ve[field] = append(ve[field], message)
}

func (inVal Input) CSRFCheck(field string, sess models.Session) {
	if field == "" {
		return
	}
	if field != sess.Uuid {
		inVal.VErrors.Add(field, "Invalid CSRF token")
	}
}
