package results

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/repos/structs"
)

func (r *Repo) Get(ctx context.Context, id string) (*structs.Result, error) {
	result := &structs.Result{}

	q := `SELECT id, status, eta FROM results
	WHERE id = $1`

	err := r.db.GetContext(ctx, result, q, id)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(err, "db request failed")
	}

	return result, nil
}
