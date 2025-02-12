// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// IPAddressQuery is the builder for querying IPAddress entities.
type IPAddressQuery struct {
	config
	ctx         *QueryContext
	order       []ipaddress.OrderOption
	inters      []Interceptor
	predicates  []predicate.IPAddress
	withIPBlock *IPBlockQuery
	modifiers   []func(*sql.Selector)
	loadTotal   []func(context.Context, []*IPAddress) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IPAddressQuery builder.
func (iaq *IPAddressQuery) Where(ps ...predicate.IPAddress) *IPAddressQuery {
	iaq.predicates = append(iaq.predicates, ps...)
	return iaq
}

// Limit the number of records to be returned by this query.
func (iaq *IPAddressQuery) Limit(limit int) *IPAddressQuery {
	iaq.ctx.Limit = &limit
	return iaq
}

// Offset to start from.
func (iaq *IPAddressQuery) Offset(offset int) *IPAddressQuery {
	iaq.ctx.Offset = &offset
	return iaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iaq *IPAddressQuery) Unique(unique bool) *IPAddressQuery {
	iaq.ctx.Unique = &unique
	return iaq
}

// Order specifies how the records should be ordered.
func (iaq *IPAddressQuery) Order(o ...ipaddress.OrderOption) *IPAddressQuery {
	iaq.order = append(iaq.order, o...)
	return iaq
}

// QueryIPBlock chains the current query on the "ip_block" edge.
func (iaq *IPAddressQuery) QueryIPBlock() *IPBlockQuery {
	query := (&IPBlockClient{config: iaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ipaddress.Table, ipaddress.FieldID, selector),
			sqlgraph.To(ipblock.Table, ipblock.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ipaddress.IPBlockTable, ipaddress.IPBlockColumn),
		)
		fromU = sqlgraph.SetNeighbors(iaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first IPAddress entity from the query.
// Returns a *NotFoundError when no IPAddress was found.
func (iaq *IPAddressQuery) First(ctx context.Context) (*IPAddress, error) {
	nodes, err := iaq.Limit(1).All(setContextOp(ctx, iaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ipaddress.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iaq *IPAddressQuery) FirstX(ctx context.Context) *IPAddress {
	node, err := iaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IPAddress ID from the query.
// Returns a *NotFoundError when no IPAddress ID was found.
func (iaq *IPAddressQuery) FirstID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = iaq.Limit(1).IDs(setContextOp(ctx, iaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ipaddress.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iaq *IPAddressQuery) FirstIDX(ctx context.Context) gidx.PrefixedID {
	id, err := iaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IPAddress entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IPAddress entity is found.
// Returns a *NotFoundError when no IPAddress entities are found.
func (iaq *IPAddressQuery) Only(ctx context.Context) (*IPAddress, error) {
	nodes, err := iaq.Limit(2).All(setContextOp(ctx, iaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ipaddress.Label}
	default:
		return nil, &NotSingularError{ipaddress.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iaq *IPAddressQuery) OnlyX(ctx context.Context) *IPAddress {
	node, err := iaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IPAddress ID in the query.
// Returns a *NotSingularError when more than one IPAddress ID is found.
// Returns a *NotFoundError when no entities are found.
func (iaq *IPAddressQuery) OnlyID(ctx context.Context) (id gidx.PrefixedID, err error) {
	var ids []gidx.PrefixedID
	if ids, err = iaq.Limit(2).IDs(setContextOp(ctx, iaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ipaddress.Label}
	default:
		err = &NotSingularError{ipaddress.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iaq *IPAddressQuery) OnlyIDX(ctx context.Context) gidx.PrefixedID {
	id, err := iaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IPAddresses.
func (iaq *IPAddressQuery) All(ctx context.Context) ([]*IPAddress, error) {
	ctx = setContextOp(ctx, iaq.ctx, "All")
	if err := iaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*IPAddress, *IPAddressQuery]()
	return withInterceptors[[]*IPAddress](ctx, iaq, qr, iaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iaq *IPAddressQuery) AllX(ctx context.Context) []*IPAddress {
	nodes, err := iaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IPAddress IDs.
func (iaq *IPAddressQuery) IDs(ctx context.Context) (ids []gidx.PrefixedID, err error) {
	if iaq.ctx.Unique == nil && iaq.path != nil {
		iaq.Unique(true)
	}
	ctx = setContextOp(ctx, iaq.ctx, "IDs")
	if err = iaq.Select(ipaddress.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iaq *IPAddressQuery) IDsX(ctx context.Context) []gidx.PrefixedID {
	ids, err := iaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iaq *IPAddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iaq.ctx, "Count")
	if err := iaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iaq, querierCount[*IPAddressQuery](), iaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iaq *IPAddressQuery) CountX(ctx context.Context) int {
	count, err := iaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iaq *IPAddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iaq.ctx, "Exist")
	switch _, err := iaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iaq *IPAddressQuery) ExistX(ctx context.Context) bool {
	exist, err := iaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IPAddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iaq *IPAddressQuery) Clone() *IPAddressQuery {
	if iaq == nil {
		return nil
	}
	return &IPAddressQuery{
		config:      iaq.config,
		ctx:         iaq.ctx.Clone(),
		order:       append([]ipaddress.OrderOption{}, iaq.order...),
		inters:      append([]Interceptor{}, iaq.inters...),
		predicates:  append([]predicate.IPAddress{}, iaq.predicates...),
		withIPBlock: iaq.withIPBlock.Clone(),
		// clone intermediate query.
		sql:  iaq.sql.Clone(),
		path: iaq.path,
	}
}

// WithIPBlock tells the query-builder to eager-load the nodes that are connected to
// the "ip_block" edge. The optional arguments are used to configure the query builder of the edge.
func (iaq *IPAddressQuery) WithIPBlock(opts ...func(*IPBlockQuery)) *IPAddressQuery {
	query := (&IPBlockClient{config: iaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iaq.withIPBlock = query
	return iaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IPAddress.Query().
//		GroupBy(ipaddress.FieldCreatedAt).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (iaq *IPAddressQuery) GroupBy(field string, fields ...string) *IPAddressGroupBy {
	iaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &IPAddressGroupBy{build: iaq}
	grbuild.flds = &iaq.ctx.Fields
	grbuild.label = ipaddress.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.IPAddress.Query().
//		Select(ipaddress.FieldCreatedAt).
//		Scan(ctx, &v)
func (iaq *IPAddressQuery) Select(fields ...string) *IPAddressSelect {
	iaq.ctx.Fields = append(iaq.ctx.Fields, fields...)
	sbuild := &IPAddressSelect{IPAddressQuery: iaq}
	sbuild.label = ipaddress.Label
	sbuild.flds, sbuild.scan = &iaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a IPAddressSelect configured with the given aggregations.
func (iaq *IPAddressQuery) Aggregate(fns ...AggregateFunc) *IPAddressSelect {
	return iaq.Select().Aggregate(fns...)
}

func (iaq *IPAddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iaq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iaq); err != nil {
				return err
			}
		}
	}
	for _, f := range iaq.ctx.Fields {
		if !ipaddress.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if iaq.path != nil {
		prev, err := iaq.path(ctx)
		if err != nil {
			return err
		}
		iaq.sql = prev
	}
	return nil
}

func (iaq *IPAddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IPAddress, error) {
	var (
		nodes       = []*IPAddress{}
		_spec       = iaq.querySpec()
		loadedTypes = [1]bool{
			iaq.withIPBlock != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*IPAddress).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &IPAddress{config: iaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(iaq.modifiers) > 0 {
		_spec.Modifiers = iaq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iaq.withIPBlock; query != nil {
		if err := iaq.loadIPBlock(ctx, query, nodes, nil,
			func(n *IPAddress, e *IPBlock) { n.Edges.IPBlock = e }); err != nil {
			return nil, err
		}
	}
	for i := range iaq.loadTotal {
		if err := iaq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iaq *IPAddressQuery) loadIPBlock(ctx context.Context, query *IPBlockQuery, nodes []*IPAddress, init func(*IPAddress), assign func(*IPAddress, *IPBlock)) error {
	ids := make([]gidx.PrefixedID, 0, len(nodes))
	nodeids := make(map[gidx.PrefixedID][]*IPAddress)
	for i := range nodes {
		fk := nodes[i].BlockID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ipblock.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "block_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (iaq *IPAddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iaq.querySpec()
	if len(iaq.modifiers) > 0 {
		_spec.Modifiers = iaq.modifiers
	}
	_spec.Node.Columns = iaq.ctx.Fields
	if len(iaq.ctx.Fields) > 0 {
		_spec.Unique = iaq.ctx.Unique != nil && *iaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iaq.driver, _spec)
}

func (iaq *IPAddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(ipaddress.Table, ipaddress.Columns, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString))
	_spec.From = iaq.sql
	if unique := iaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iaq.path != nil {
		_spec.Unique = true
	}
	if fields := iaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddress.FieldID)
		for i := range fields {
			if fields[i] != ipaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if iaq.withIPBlock != nil {
			_spec.Node.AddColumnOnce(ipaddress.FieldBlockID)
		}
	}
	if ps := iaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iaq *IPAddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iaq.driver.Dialect())
	t1 := builder.Table(ipaddress.Table)
	columns := iaq.ctx.Fields
	if len(columns) == 0 {
		columns = ipaddress.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iaq.sql != nil {
		selector = iaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iaq.ctx.Unique != nil && *iaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iaq.predicates {
		p(selector)
	}
	for _, p := range iaq.order {
		p(selector)
	}
	if offset := iaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IPAddressGroupBy is the group-by builder for IPAddress entities.
type IPAddressGroupBy struct {
	selector
	build *IPAddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (iagb *IPAddressGroupBy) Aggregate(fns ...AggregateFunc) *IPAddressGroupBy {
	iagb.fns = append(iagb.fns, fns...)
	return iagb
}

// Scan applies the selector query and scans the result into the given value.
func (iagb *IPAddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, iagb.build.ctx, "GroupBy")
	if err := iagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPAddressQuery, *IPAddressGroupBy](ctx, iagb.build, iagb, iagb.build.inters, v)
}

func (iagb *IPAddressGroupBy) sqlScan(ctx context.Context, root *IPAddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(iagb.fns))
	for _, fn := range iagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*iagb.flds)+len(iagb.fns))
		for _, f := range *iagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*iagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := iagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// IPAddressSelect is the builder for selecting fields of IPAddress entities.
type IPAddressSelect struct {
	*IPAddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ias *IPAddressSelect) Aggregate(fns ...AggregateFunc) *IPAddressSelect {
	ias.fns = append(ias.fns, fns...)
	return ias
}

// Scan applies the selector query and scans the result into the given value.
func (ias *IPAddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ias.ctx, "Select")
	if err := ias.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*IPAddressQuery, *IPAddressSelect](ctx, ias.IPAddressQuery, ias, ias.inters, v)
}

func (ias *IPAddressSelect) sqlScan(ctx context.Context, root *IPAddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ias.fns))
	for _, fn := range ias.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ias.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ias.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
