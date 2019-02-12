package sheets

import (
	"github.com/mesg-foundation/core/client/service"
)

const readOnlyAPI = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Sheets exposes functionalities from Google Sheets API.
type Sheets struct {
	service *service.Service
}

// New creates a new Sheets.
func New() (*Sheets, error) {
	s := &Sheets{}
	var err error
	s.service, err = service.New()
	return s, err
}

// Listen listens for tasks.
func (s *Sheets) Listen() error {
	return s.service.Listen(
		service.Task("get", s.get),
	)
}

// Close gracefully closes the service.
func (s *Sheets) Close() error {
	return s.service.Close()
}
