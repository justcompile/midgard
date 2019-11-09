package dal

import (
	"context"
	"time"
)

type Indexed interface {
	Indexes() []string
}

type BuildStatus string

var (
	BuildStatusQueued  BuildStatus = "QUEUED"
	BuildStatusStarted BuildStatus = "STARTED"
)

// Project defines the database model for a `project` which is the root level model for Midgard
type Project struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created" sql:"default:now()"`
	UpdatedAt time.Time `json:"updated"`
}

// BeforeInsert is a Model hook which is invoked before a Project is saved
func (p *Project) BeforeInsert(ctx context.Context) (context.Context, error) {
	p.UpdatedAt = time.Now()
	return ctx, nil
}

// BeforeUpdate is a Model hook which is invoked before a Project is saved
func (p *Project) BeforeUpdate(ctx context.Context) (context.Context, error) {
	p.UpdatedAt = time.Now()
	return ctx, nil
}

func (p *Project) Indexes() []string {
	return []string{
		`CREATE UNIQUE INDEX IF NOT EXISTS name_unique_idx on ?TableName (LOWER(name))`,
	}
}

// Build ...
type Build struct {
	Id       int64       `json:"id"`
	BuildId  int64       `json:"build_id"`
	Project  *Project    `json:"project"`
	Created  time.Time   `json:"created"`
	Started  time.Time   `json:"started"`
	Finished *time.Time  `json:"finished"`
	Status   BuildStatus `json:"status"`
}

// BeforeInsert is a Model hook which is invoked before a Build is saved
func (b *Build) BeforeInsert(ctx context.Context) (context.Context, error) {
	b.Created = time.Now()
	return ctx, nil
}
