package piloto

import (
	"time"

	"github.com/andromeda-fie/gatekeeper/src/handler/astrolog"
	. "github.com/zoedsoupe/exo/changeset"
)

type Credito int64

type Localizacaot struct {
	Planeta string    `json:"planeta"`
	Quando  time.Time `json:"quando"`
}

type Piloto struct {
	Id            string         `json:"id"`
	Nome          string         `json:"nome"`
	Idade         int32          `json:"idade"`
	Certificacao  string         `json:"certificacao"`
	Creditos      Credito        `json:"credito"`
	Aposentadoria bool           `json:"aposentadoria"`
	Localizacao   []Localizacaot `json:"localizacao"`
}

// Neste serviço, apenas mapeamos um mapa JSON para o struct
// Validações devem ser feitas no serviço Astrolog.
func changeset(piloto *Piloto, attrs map[string]interface{}) error {
	required := []string{"Id", "Nome", "Idade", "Certificacao", "Creditos"}
	c := Cast[Piloto](attrs).ValidateRequired(required)
	return Apply[Piloto](piloto, c)
}

func Find(id string) (Piloto, error) {
	var piloto Piloto

	attrs, err := astrolog.FindPiloto(id)
	if err != nil {
		return piloto, err
	}

	err = changeset(&piloto, attrs)
	if err != nil {
		return piloto, err
	}

	return piloto, nil
}
