package astrolog

import (
	"fmt"

	"github.com/andromeda-fie/gatekeeper/src/handler"
)

var FindPilotoURL = "https://astrolog.net/pilotos/%s"

func FindPiloto(id string) (map[string]interface{}, error) {
	url := fmt.Sprintf(FindPilotoURL, id)
	return handler.MakeGet(url)
}
