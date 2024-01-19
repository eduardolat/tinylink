package postgres

import (
	"context"
	"errors"
	"fmt"
)

func (ds *DataStore) IncrementClicks(shortCode string) error {
	sqlQuery := `UPDATE links SET clicks = clicks + 1, updated_at = NOW() WHERE short_code = $1;`

	result, err := ds.conn.Exec(context.Background(), sqlQuery, shortCode)
	if err != nil {
		return fmt.Errorf("error incrementing clicks: %w", err)
	}
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return errors.New("Link not found")
	}

	return nil
}
