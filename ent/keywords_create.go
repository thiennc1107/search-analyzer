// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/keywords"
)

// KeywordsCreate is the builder for creating a Keywords entity.
type KeywordsCreate struct {
	config
	mutation *KeywordsMutation
	hooks    []Hook
}

// SetKeyword sets the "keyword" field.
func (kc *KeywordsCreate) SetKeyword(s string) *KeywordsCreate {
	kc.mutation.SetKeyword(s)
	return kc
}

// SetStatus sets the "status" field.
func (kc *KeywordsCreate) SetStatus(k keywords.Status) *KeywordsCreate {
	kc.mutation.SetStatus(k)
	return kc
}

// SetAdsAmount sets the "ads_amount" field.
func (kc *KeywordsCreate) SetAdsAmount(i int) *KeywordsCreate {
	kc.mutation.SetAdsAmount(i)
	return kc
}

// SetLinksAmount sets the "links_amount" field.
func (kc *KeywordsCreate) SetLinksAmount(i int) *KeywordsCreate {
	kc.mutation.SetLinksAmount(i)
	return kc
}

// SetSearchResultAmount sets the "search_result_amount" field.
func (kc *KeywordsCreate) SetSearchResultAmount(i int) *KeywordsCreate {
	kc.mutation.SetSearchResultAmount(i)
	return kc
}

// SetHTMLCode sets the "html_code" field.
func (kc *KeywordsCreate) SetHTMLCode(s string) *KeywordsCreate {
	kc.mutation.SetHTMLCode(s)
	return kc
}

// Mutation returns the KeywordsMutation object of the builder.
func (kc *KeywordsCreate) Mutation() *KeywordsMutation {
	return kc.mutation
}

// Save creates the Keywords in the database.
func (kc *KeywordsCreate) Save(ctx context.Context) (*Keywords, error) {
	return withHooks(ctx, kc.sqlSave, kc.mutation, kc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kc *KeywordsCreate) SaveX(ctx context.Context) *Keywords {
	v, err := kc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kc *KeywordsCreate) Exec(ctx context.Context) error {
	_, err := kc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kc *KeywordsCreate) ExecX(ctx context.Context) {
	if err := kc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kc *KeywordsCreate) check() error {
	if _, ok := kc.mutation.Keyword(); !ok {
		return &ValidationError{Name: "keyword", err: errors.New(`ent: missing required field "Keywords.keyword"`)}
	}
	if _, ok := kc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Keywords.status"`)}
	}
	if v, ok := kc.mutation.Status(); ok {
		if err := keywords.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Keywords.status": %w`, err)}
		}
	}
	if _, ok := kc.mutation.AdsAmount(); !ok {
		return &ValidationError{Name: "ads_amount", err: errors.New(`ent: missing required field "Keywords.ads_amount"`)}
	}
	if _, ok := kc.mutation.LinksAmount(); !ok {
		return &ValidationError{Name: "links_amount", err: errors.New(`ent: missing required field "Keywords.links_amount"`)}
	}
	if _, ok := kc.mutation.SearchResultAmount(); !ok {
		return &ValidationError{Name: "search_result_amount", err: errors.New(`ent: missing required field "Keywords.search_result_amount"`)}
	}
	if _, ok := kc.mutation.HTMLCode(); !ok {
		return &ValidationError{Name: "html_code", err: errors.New(`ent: missing required field "Keywords.html_code"`)}
	}
	return nil
}

func (kc *KeywordsCreate) sqlSave(ctx context.Context) (*Keywords, error) {
	if err := kc.check(); err != nil {
		return nil, err
	}
	_node, _spec := kc.createSpec()
	if err := sqlgraph.CreateNode(ctx, kc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	kc.mutation.id = &_node.ID
	kc.mutation.done = true
	return _node, nil
}

func (kc *KeywordsCreate) createSpec() (*Keywords, *sqlgraph.CreateSpec) {
	var (
		_node = &Keywords{config: kc.config}
		_spec = sqlgraph.NewCreateSpec(keywords.Table, sqlgraph.NewFieldSpec(keywords.FieldID, field.TypeInt))
	)
	if value, ok := kc.mutation.Keyword(); ok {
		_spec.SetField(keywords.FieldKeyword, field.TypeString, value)
		_node.Keyword = value
	}
	if value, ok := kc.mutation.Status(); ok {
		_spec.SetField(keywords.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := kc.mutation.AdsAmount(); ok {
		_spec.SetField(keywords.FieldAdsAmount, field.TypeInt, value)
		_node.AdsAmount = value
	}
	if value, ok := kc.mutation.LinksAmount(); ok {
		_spec.SetField(keywords.FieldLinksAmount, field.TypeInt, value)
		_node.LinksAmount = value
	}
	if value, ok := kc.mutation.SearchResultAmount(); ok {
		_spec.SetField(keywords.FieldSearchResultAmount, field.TypeInt, value)
		_node.SearchResultAmount = value
	}
	if value, ok := kc.mutation.HTMLCode(); ok {
		_spec.SetField(keywords.FieldHTMLCode, field.TypeString, value)
		_node.HTMLCode = value
	}
	return _node, _spec
}

// KeywordsCreateBulk is the builder for creating many Keywords entities in bulk.
type KeywordsCreateBulk struct {
	config
	err      error
	builders []*KeywordsCreate
}

// Save creates the Keywords entities in the database.
func (kcb *KeywordsCreateBulk) Save(ctx context.Context) ([]*Keywords, error) {
	if kcb.err != nil {
		return nil, kcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kcb.builders))
	nodes := make([]*Keywords, len(kcb.builders))
	mutators := make([]Mutator, len(kcb.builders))
	for i := range kcb.builders {
		func(i int, root context.Context) {
			builder := kcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KeywordsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kcb *KeywordsCreateBulk) SaveX(ctx context.Context) []*Keywords {
	v, err := kcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kcb *KeywordsCreateBulk) Exec(ctx context.Context) error {
	_, err := kcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kcb *KeywordsCreateBulk) ExecX(ctx context.Context) {
	if err := kcb.Exec(ctx); err != nil {
		panic(err)
	}
}
