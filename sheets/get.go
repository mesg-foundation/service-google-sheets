package sheets

import (
	"context"

	"github.com/mesg-foundation/core/client/service"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type getInputs struct {
	ServiceAccountCredentials string `json:"serviceAccountCredentials"`
	SpreadsheetID             string `json:"spreadsheetID"`
	Range                     string `json:"range"`
}

type getRowsOutput struct {
	Rows []RowContainer `json:"rows"`
}

type RowContainer struct {
	Row []interface{} `json:"row"`
}

func (s *Sheets) get(execution *service.Execution) (string, interface{}) {
	var inputs getInputs
	if err := execution.Data(&inputs); err != nil {
		return errorOutputFrom(err)
	}

	config, err := google.JWTConfigFromJSON([]byte(inputs.ServiceAccountCredentials), readOnlyAPI)
	if err != nil {
		return errorOutputFrom(err)
	}

	client := config.Client(context.Background())
	srv, err := sheets.New(client)
	if err != nil {
		return errorOutputFrom(err)
	}

	resp, err := srv.Spreadsheets.Values.Get(inputs.SpreadsheetID, inputs.Range).Do()
	if err != nil {
		return errorOutputFrom(err)
	}

	var rows []RowContainer
	for _, row := range resp.Values {
		rows = append(rows, RowContainer{Row: row})
	}

	return "success", getRowsOutput{rows}
}
