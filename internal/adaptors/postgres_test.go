package adaptors

import (
	"testing"
)

var psql DB

func TestConnection(t *testing.T) {
	var err error
	psql, err = NewPostgres()
	if err != nil {
		t.Errorf("Error when connecting to DB: %v", err)
		return
	}
	t.Logf("DB connected")
}
