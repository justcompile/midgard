package dal

import (
	"time"

	"github.com/go-pg/pg/v9"
)

// BuildQueue is a service which workers can use to pull builds to carry out
type BuildQueue struct {
	db *pg.DB
}

func (q *BuildQueue) All() []*Build {
	var builds []*Build

	q.db.Model(builds).
		Where("status = ?", BuildStatusQueued).
		Order("created ASC").
		Select()

	return builds
}

func (q *BuildQueue) Pop() *Build {
	var build *Build

	var id int64

	q.db.Model((*Build)(nil)).
		Column("id").
		Where("status = ?", BuildStatusQueued).
		Order("created ASC").
		Limit(1).
		Select(&id)

	if id == 0 {
		return nil
	}

	q.db.RunInTransaction(func(tx *pg.Tx) error {
		q.db.Model(&build).
			Where("id = ?", id).
			For("UPDATE").
			Select()

		if build == nil || build.Status != BuildStatusQueued {
			return nil
		}

		build.Status = BuildStatusStarted
		build.Started = time.Now()

		return q.db.Update(build)
	})

	return build
}
