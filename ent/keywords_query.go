// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/keywords"
	"github.com/mikestefanello/pagoda/ent/predicate"
)

// KeywordsQuery is the builder for querying Keywords entities.
type KeywordsQuery struct {
	config
	ctx        *QueryContext
	order      []keywords.OrderOption
	inters     []Interceptor
	predicates []predicate.Keywords
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the KeywordsQuery builder.
func (kq *KeywordsQuery) Where(ps ...predicate.Keywords) *KeywordsQuery {
	kq.predicates = append(kq.predicates, ps...)
	return kq
}

// Limit the number of records to be returned by this query.
func (kq *KeywordsQuery) Limit(limit int) *KeywordsQuery {
	kq.ctx.Limit = &limit
	return kq
}

// Offset to start from.
func (kq *KeywordsQuery) Offset(offset int) *KeywordsQuery {
	kq.ctx.Offset = &offset
	return kq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (kq *KeywordsQuery) Unique(unique bool) *KeywordsQuery {
	kq.ctx.Unique = &unique
	return kq
}

// Order specifies how the records should be ordered.
func (kq *KeywordsQuery) Order(o ...keywords.OrderOption) *KeywordsQuery {
	kq.order = append(kq.order, o...)
	return kq
}

// First returns the first Keywords entity from the query.
// Returns a *NotFoundError when no Keywords was found.
func (kq *KeywordsQuery) First(ctx context.Context) (*Keywords, error) {
	nodes, err := kq.Limit(1).All(setContextOp(ctx, kq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{keywords.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (kq *KeywordsQuery) FirstX(ctx context.Context) *Keywords {
	node, err := kq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Keywords ID from the query.
// Returns a *NotFoundError when no Keywords ID was found.
func (kq *KeywordsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kq.Limit(1).IDs(setContextOp(ctx, kq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{keywords.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (kq *KeywordsQuery) FirstIDX(ctx context.Context) int {
	id, err := kq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Keywords entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Keywords entity is found.
// Returns a *NotFoundError when no Keywords entities are found.
func (kq *KeywordsQuery) Only(ctx context.Context) (*Keywords, error) {
	nodes, err := kq.Limit(2).All(setContextOp(ctx, kq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{keywords.Label}
	default:
		return nil, &NotSingularError{keywords.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (kq *KeywordsQuery) OnlyX(ctx context.Context) *Keywords {
	node, err := kq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Keywords ID in the query.
// Returns a *NotSingularError when more than one Keywords ID is found.
// Returns a *NotFoundError when no entities are found.
func (kq *KeywordsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = kq.Limit(2).IDs(setContextOp(ctx, kq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{keywords.Label}
	default:
		err = &NotSingularError{keywords.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (kq *KeywordsQuery) OnlyIDX(ctx context.Context) int {
	id, err := kq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of KeywordsSlice.
func (kq *KeywordsQuery) All(ctx context.Context) ([]*Keywords, error) {
	ctx = setContextOp(ctx, kq.ctx, "All")
	if err := kq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Keywords, *KeywordsQuery]()
	return withInterceptors[[]*Keywords](ctx, kq, qr, kq.inters)
}

// AllX is like All, but panics if an error occurs.
func (kq *KeywordsQuery) AllX(ctx context.Context) []*Keywords {
	nodes, err := kq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Keywords IDs.
func (kq *KeywordsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if kq.ctx.Unique == nil && kq.path != nil {
		kq.Unique(true)
	}
	ctx = setContextOp(ctx, kq.ctx, "IDs")
	if err = kq.Select(keywords.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (kq *KeywordsQuery) IDsX(ctx context.Context) []int {
	ids, err := kq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (kq *KeywordsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, kq.ctx, "Count")
	if err := kq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, kq, querierCount[*KeywordsQuery](), kq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (kq *KeywordsQuery) CountX(ctx context.Context) int {
	count, err := kq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (kq *KeywordsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, kq.ctx, "Exist")
	switch _, err := kq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (kq *KeywordsQuery) ExistX(ctx context.Context) bool {
	exist, err := kq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the KeywordsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (kq *KeywordsQuery) Clone() *KeywordsQuery {
	if kq == nil {
		return nil
	}
	return &KeywordsQuery{
		config:     kq.config,
		ctx:        kq.ctx.Clone(),
		order:      append([]keywords.OrderOption{}, kq.order...),
		inters:     append([]Interceptor{}, kq.inters...),
		predicates: append([]predicate.Keywords{}, kq.predicates...),
		// clone intermediate query.
		sql:  kq.sql.Clone(),
		path: kq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status keywords.Status `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Keywords.Query().
//		GroupBy(keywords.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (kq *KeywordsQuery) GroupBy(field string, fields ...string) *KeywordsGroupBy {
	kq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &KeywordsGroupBy{build: kq}
	grbuild.flds = &kq.ctx.Fields
	grbuild.label = keywords.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status keywords.Status `json:"status,omitempty"`
//	}
//
//	client.Keywords.Query().
//		Select(keywords.FieldStatus).
//		Scan(ctx, &v)
func (kq *KeywordsQuery) Select(fields ...string) *KeywordsSelect {
	kq.ctx.Fields = append(kq.ctx.Fields, fields...)
	sbuild := &KeywordsSelect{KeywordsQuery: kq}
	sbuild.label = keywords.Label
	sbuild.flds, sbuild.scan = &kq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a KeywordsSelect configured with the given aggregations.
func (kq *KeywordsQuery) Aggregate(fns ...AggregateFunc) *KeywordsSelect {
	return kq.Select().Aggregate(fns...)
}

func (kq *KeywordsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range kq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, kq); err != nil {
				return err
			}
		}
	}
	for _, f := range kq.ctx.Fields {
		if !keywords.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if kq.path != nil {
		prev, err := kq.path(ctx)
		if err != nil {
			return err
		}
		kq.sql = prev
	}
	return nil
}

func (kq *KeywordsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Keywords, error) {
	var (
		nodes = []*Keywords{}
		_spec = kq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Keywords).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Keywords{config: kq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, kq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (kq *KeywordsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := kq.querySpec()
	_spec.Node.Columns = kq.ctx.Fields
	if len(kq.ctx.Fields) > 0 {
		_spec.Unique = kq.ctx.Unique != nil && *kq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, kq.driver, _spec)
}

func (kq *KeywordsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(keywords.Table, keywords.Columns, sqlgraph.NewFieldSpec(keywords.FieldID, field.TypeInt))
	_spec.From = kq.sql
	if unique := kq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if kq.path != nil {
		_spec.Unique = true
	}
	if fields := kq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, keywords.FieldID)
		for i := range fields {
			if fields[i] != keywords.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := kq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := kq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := kq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := kq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (kq *KeywordsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(kq.driver.Dialect())
	t1 := builder.Table(keywords.Table)
	columns := kq.ctx.Fields
	if len(columns) == 0 {
		columns = keywords.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if kq.sql != nil {
		selector = kq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if kq.ctx.Unique != nil && *kq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range kq.predicates {
		p(selector)
	}
	for _, p := range kq.order {
		p(selector)
	}
	if offset := kq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := kq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// KeywordsGroupBy is the group-by builder for Keywords entities.
type KeywordsGroupBy struct {
	selector
	build *KeywordsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (kgb *KeywordsGroupBy) Aggregate(fns ...AggregateFunc) *KeywordsGroupBy {
	kgb.fns = append(kgb.fns, fns...)
	return kgb
}

// Scan applies the selector query and scans the result into the given value.
func (kgb *KeywordsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, kgb.build.ctx, "GroupBy")
	if err := kgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeywordsQuery, *KeywordsGroupBy](ctx, kgb.build, kgb, kgb.build.inters, v)
}

func (kgb *KeywordsGroupBy) sqlScan(ctx context.Context, root *KeywordsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(kgb.fns))
	for _, fn := range kgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*kgb.flds)+len(kgb.fns))
		for _, f := range *kgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*kgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := kgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// KeywordsSelect is the builder for selecting fields of Keywords entities.
type KeywordsSelect struct {
	*KeywordsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ks *KeywordsSelect) Aggregate(fns ...AggregateFunc) *KeywordsSelect {
	ks.fns = append(ks.fns, fns...)
	return ks
}

// Scan applies the selector query and scans the result into the given value.
func (ks *KeywordsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ks.ctx, "Select")
	if err := ks.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*KeywordsQuery, *KeywordsSelect](ctx, ks.KeywordsQuery, ks, ks.inters, v)
}

func (ks *KeywordsSelect) sqlScan(ctx context.Context, root *KeywordsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ks.fns))
	for _, fn := range ks.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ks.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ks.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}