package views

import (
	"lenslocked.com/models"
	"log"
)

// Data is the top level structure that views expect data
// to come from
type Data struct {
	Alert *Alert
	User  *models.User
	Owner bool
	Yield interface{}
}

// Alert is used to render Bootstrap alert messages in templates
type Alert struct {
	Level   string
	Message string
}

const (
	AlertLvlError   = "danger"
	AlertLvlWarning = "warning"
	AlertLvlInfo    = "info"
	AlertLvlSuccess = "success"

	// AlertMsgGeneric is displayed when any random error
	// is encountered by our backend.
	AlertMsgGeneric = "Something went wrong. Please try " +
		"again, and contact us if the problem persists."
)

func (d *Data) SetAlert(err error) {
	var msg string

	if publicError, ok := err.(PublicError); ok {
		msg = publicError.Public()
	} else {
		log.Println(err)
		msg = AlertMsgGeneric
	}

	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLvlError,
		Message: msg,
	}
}

type PublicError interface {
	error
	Public() string
}
