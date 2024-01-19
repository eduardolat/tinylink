package postgres

import (
	"context"
	"fmt"
)

func (ds *DataStore) IsShortCodeAvailable(shortCode string) (bool, error) {
	sqlQuery := `SELECT EXISTS(SELECT 1 FROM links WHERE short_code = $1);`
	var exists bool
	err := ds.conn.QueryRow(context.Background(), sqlQuery, shortCode).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking short code availability: %w", err)
	}
	return !exists, nil
}
