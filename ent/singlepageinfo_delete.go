// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yumenaka/comi/ent/predicate"
	"github.com/yumenaka/comi/ent/singlepageinfo"
)

// SinglePageInfoDelete is the builder for deleting a SinglePageInfo entity.
type SinglePageInfoDelete struct {
	config
	hooks    []Hook
	mutation *SinglePageInfoMutation
}

// Where appends a list predicates to the SinglePageInfoDelete builder.
func (spid *SinglePageInfoDelete) Where(ps ...predicate.SinglePageInfo) *SinglePageInfoDelete {
	spid.mutation.Where(ps...)
	return spid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (spid *SinglePageInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(spid.hooks) == 0 {
		affected, err = spid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SinglePageInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			spid.mutation = mutation
			affected, err = spid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(spid.hooks) - 1; i >= 0; i-- {
			if spid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, spid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (spid *SinglePageInfoDelete) ExecX(ctx context.Context) int {
	n, err := spid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (spid *SinglePageInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: singlepageinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: singlepageinfo.FieldID,
			},
		},
	}
	if ps := spid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, spid.driver, _spec)
}

// SinglePageInfoDeleteOne is the builder for deleting a single SinglePageInfo entity.
type SinglePageInfoDeleteOne struct {
	spid *SinglePageInfoDelete
}

// Exec executes the deletion query.
func (spido *SinglePageInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := spido.spid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{singlepageinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (spido *SinglePageInfoDeleteOne) ExecX(ctx context.Context) {
	spido.spid.ExecX(ctx)
}
