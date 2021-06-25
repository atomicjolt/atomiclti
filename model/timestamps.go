package model

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
)

// Timestamps contains shared timestamps, and implements a hook to keep
// UpdatedAt accurate
type Timestamps struct {
	CreatedAt time.Time `pg:",notnull,default:now()"`
	UpdatedAt time.Time `pg:",notnull,default:now()"`
}

// This ensures that BeforeUpdate is compiled
var _ pg.BeforeUpdateHook = (*Timestamps)(nil)

// BeforeUpdate is run before an updated record is saved to the database
func (m *Timestamps) BeforeUpdate(ctx context.Context) (context.Context, error) {
	m.UpdatedAt = time.Now()

	return ctx, nil
}
