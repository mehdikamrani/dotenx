package postgresql

import (
	"database/sql"
	"log"
	"strings"

	"github.com/utopiops/automated-ops/ao-api/models"
)

var seeds = []string{
	insertTaskStatusValues,
}

func Seed(db *sql.DB) error {
	for _, seed := range seeds {
		log.Println(seed)
		if _, err := db.Exec(seed); err != nil {
			return err
		}
	}
	return nil
}

func format(values []string) string {
	var str strings.Builder
	for i, v := range values {
		str.WriteString(`('` + v + `')`)
		if i < len(values)-1 {
			str.WriteString(",")
		}
	}
	return str.String()
}

var insertTaskStatusValues = `
INSERT INTO task_status (name) VALUES ` + format(models.TaskStatusValues()) + ` 
ON CONFLICT (name) 
DO NOTHING;
`
