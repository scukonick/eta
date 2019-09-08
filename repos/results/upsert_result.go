package results

import (
	"context"

	"github.com/pkg/errors"

	"github.com/scukonick/eta/repos/structs"
)

func (r *Repo) UpsertResult(ctx context.Context, result structs.Result) error {
	q := `INSERT INTO results (id, status, eta)
	VALUEs (:id, :status, :eta)
	ON CONFLICT (id)
	DO UPDATE SET status = :status, eta = :eta`

	_, err := r.db.NamedExecContext(ctx, q, &result)
	if err != nil {
		return errors.Wrap(err, "failed to exec upsert")
	}

	return nil
}
