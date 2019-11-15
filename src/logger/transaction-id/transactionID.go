package txID

import (
	"net/http"

	"github.com/google/uuid"
)

const Key = "txID"

func FromRequest(r *http.Request) string {
	txID := r.Header.Get(Key)

	if txID == "" {
		txID = uuid.New().String()
	}

	return txID
}
