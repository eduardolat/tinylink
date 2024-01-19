package postgres

import (
	"context"
	"fmt"
)

func (ds *DataStore) IsURLAlreadyStored(originalURL string) (bool, error) {
	sqlQuery := `SELECT EXISTS(SELECT 1 FROM links WHERE original_url = $1);`
	var exists bool
	err := ds.conn.QueryRow(context.Background(), sqlQuery, originalURL).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking URL existence: %w", err)
	}
	return exists, nil
}
