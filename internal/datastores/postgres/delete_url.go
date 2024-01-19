package postgres

import (
	"context"
	"errors"
	"fmt"
)

func (ds *DataStore) DeleteURL(shortCode string) error {
	sqlQuery := `DELETE FROM links WHERE short_code = $1;`

	result, err := ds.conn.Exec(context.Background(), sqlQuery, shortCode)
	if err != nil {
		return fmt.Errorf("error deleting URL: %w", err)
	}
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("URL not found")
	}

	return nil
}
