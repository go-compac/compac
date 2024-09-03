package example

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-compac/compac"
	"github.com/stretchr/testify/assert"
)

// db mock of DB query interface
type dbConn struct{}

func (db *dbConn) Query(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	return nil, nil
}

// PersonInfo data struct which describes a person.
// We must know an ID of person, but more info may be provided by nullable fields
type personInfo struct {
	ID        int
	FirstName compac.Nl[string]
	LastName  compac.Nl[string]
	Age       compac.Nl[int]
}

// TestQueryPersonInfo - build sql query filter by nullable fields
func TestQueryPersonInfo(t *testing.T) {
	age := 26
	agePtr := &age

	person := personInfo{
		ID:        0,
		FirstName: compac.NlFromValue("Ivan"),
		LastName: compac.Nl[string]{
			Data:  "Revushkin",
			Valid: true,
		},
		Age: compac.NlFromPtr(agePtr),
	}

	query := `SELECT * FROM persons`

	var filter []string
	if person.FirstName.Valid {
		filter = append(filter, fmt.Sprintf("first_name = '%v'", person.FirstName.Data))
	}

	if person.LastName.Valid {
		filter = append(filter, fmt.Sprintf("last_name = '%v'", person.LastName.Data))
	}

	if person.Age.Valid {
		filter = append(filter, fmt.Sprintf("age = %v", person.Age.Data))
	}

	if len(filter) > 0 {
		query += " WHERE " + strings.Join(filter, " AND ")
	}

	wantQuery := "SELECT * FROM persons WHERE first_name = 'Ivan' AND last_name = 'Revushkin' AND age = 26"
	assert.Equal(t, wantQuery, query)
}
