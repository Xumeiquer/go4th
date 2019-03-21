package go4th

import (
	"fmt"
)

// Query is a global object just to give a method scope
type Query struct {
	Query interface{} `json:"query,omitempty"`
}

// Eq will construct a equal query
type Eq struct {
	Field string      `json:"_field,omitempty"`
	Value interface{} `json:"_value,omitempty"`
}

// Gt will construct a greater than query
type Gt struct {
	Gt map[string]interface{} `json:"_gt,omitempty"`
}

// Gte will construct a greater than equal query
type Gte struct {
	Gte map[string]interface{} `json:"_gte,omitempty"`
}

// Lt will construct a less than query
type Lt struct {
	Lt map[string]interface{} `json:"_lt,omitempty"`
}

// Lte will construct a less than equal query
type Lte struct {
	Lte map[string]interface{} `json:"_lte,omitempty"`
}

// And will construct a logical and query
type And struct {
	And []interface{} `json:"_and,omitempty"`
}

// Or will construct a logical or query
type Or struct {
	Or []interface{} `json:"_or,omitempty"`
}

// Not will construct a logical not query
type Not struct {
	Not interface{} `json:"_not,omitempty"`
}

// In will construct a in query
type In struct {
	In struct {
		Field string      `json:"_field,omitempty"`
		Value interface{} `json:"_value,omitempty"`
	} `json:"_in,omitempty"`
}

// Contains will construct a contains query
type Contains struct {
	Contains interface{} `json:"_contains,omitempty"`
}

// ID will construct a query to filter by ID
type ID struct {
	ID interface{} `json:"_id,omitempty"`
}

// IBetweenD will construct a query to filter a field
type Between struct {
	Between struct {
		Field string      `json:"_field,omitempty"`
		From  interface{} `json:"_from,omitempty"`
		To    interface{} `json:"_to,omitempty"`
	} `json:"_between,omitempty"`
}

// ParentID will construct a query to filter by parent ID
type ParentID struct {
	ParentID struct {
		Type interface{} `json:"_type,omitempty"`
		ID   interface{} `json:"_id,omitempty"`
	} `json:"_parent,omitempty"`
}

// Parent will construct a query to filter by parent
type Parent struct {
	Parent struct {
		Type  interface{} `json:"_type,omitempty"`
		Query interface{} `json:"_query,omitempty"`
	} `json:"_parent,omitempty"`
}

// Child will construct a query to filter by child
type Child struct {
	Child struct {
		Type  interface{} `json:"_type,omitempty"`
		Query interface{} `json:"_query,omitempty"`
	} `json:"_child,omitempty"`
}

// Type will construct a query to filter by type
type Type struct {
	Type interface{} `json:"_type,omitempty"`
}

// Status will construct a query to filter by status
type Status struct {
	Status interface{} `json:"status,omitempty"`
}

// String will construct a query to filter by string
type String struct {
	String interface{} `json:"_string,omitempty"`
}

// NewQuery returns a new pointer to Query whichs is used to build up a query
func NewQuery() *Query {
	return &Query{}
}

// BuildQuery returns a new query ready to be used
func BuildQuery(q interface{}) *Query {
	return &Query{Query: q}
}

// Eq returns a Eq object
func (q *Query) Eq(field string, value interface{}) (Eq, error) {
	if field == "" {
		return Eq{}, fmt.Errorf("field could not be empty")
	}
	return Eq{field, value}, nil
}

// Gt returns a Gt object
func (q *Query) Gt(field string, value interface{}) (Gt, error) {
	if field == "" {
		return Gt{}, fmt.Errorf("field could not be empty")
	}
	gt := Gt{}
	gt.Gt = make(map[string]interface{})
	gt.Gt[field] = value
	return gt, nil
}

// Gte returns a Gte object
func (q *Query) Gte(field string, value interface{}) (Gte, error) {
	if field == "" {
		return Gte{}, fmt.Errorf("field could not be empty")
	}
	gte := Gte{}
	gte.Gte = make(map[string]interface{})
	gte.Gte[field] = value
	return gte, nil
}

// Lt returns a Lt object
func (q *Query) Lt(field string, value interface{}) (Lt, error) {
	if field == "" {
		return Lt{}, fmt.Errorf("field could not be empty")
	}
	lt := Lt{}
	lt.Lt = make(map[string]interface{})
	lt.Lt[field] = value
	return lt, nil
}

// Lte returns a Lte object
func (q *Query) Lte(field string, value interface{}) (Lte, error) {
	if field == "" {
		return Lte{}, fmt.Errorf("field could not be empty")
	}
	lte := Lte{}
	lte.Lte = make(map[string]interface{})
	lte.Lte[field] = value
	return lte, nil
}

// And returns a And object
func (q *Query) And(query ...interface{}) (And, error) {
	and := And{}
	for _, q := range query {
		and.And = append(and.And, q)
	}
	return and, nil
}

// Or returns a Or object
func (q *Query) Or(query ...interface{}) (Or, error) {
	or := Or{}
	for _, q := range query {
		or.Or = append(or.Or, q)
	}
	return or, nil
}

// Not returns a Not object
func (q *Query) Not(not interface{}) (Not, error) {
	return Not{not}, nil
}

// In returns a In object
func (q *Query) In(field string, value interface{}) (In, error) {
	if field == "" {
		return In{}, fmt.Errorf("field could not be empty")
	}
	var subIn struct {
		Field string      `json:"_field,omitempty"`
		Value interface{} `json:"_value,omitempty"`
	}
	subIn.Field = field
	subIn.Value = value
	return In{subIn}, nil
}

// Contains returns a Contains object
func (q *Query) Contains(contains interface{}) (Contains, error) {
	return Contains{contains}, nil
}

// ID returns a ID object
func (q *Query) ID(id interface{}) (ID, error) {
	return ID{id}, nil
}

// Between returns a Between object
func (q *Query) Between(field string, from, to interface{}) (Between, error) {
	if field == "" {
		return Between{}, fmt.Errorf("field could not be empty")
	}
	var subBetween struct {
		Field string      `json:"_field,omitempty"`
		From  interface{} `json:"_from,omitempty"`
		To    interface{} `json:"_to,omitempty"`
	}
	subBetween.Field = field
	subBetween.From = from
	subBetween.To = to
	return Between{subBetween}, nil
}

// ParentID returns a ParentID object
func (q *Query) ParentID(typ, id interface{}) (ParentID, error) {
	if typ == nil {
		return ParentID{}, fmt.Errorf("type could not be nil")
	}
	var subParentID struct {
		Type interface{} `json:"_type,omitempty"`
		ID   interface{} `json:"_id,omitempty"`
	}
	subParentID.Type = typ
	subParentID.ID = id
	return ParentID{subParentID}, nil
}

// Parent returns a Parent object
func (q *Query) Parent(typ, query interface{}) (Parent, error) {
	if typ == nil {
		return Parent{}, fmt.Errorf("type could not be nil")
	}
	var subParent struct {
		Type  interface{} `json:"_type,omitempty"`
		Query interface{} `json:"_query,omitempty"`
	}
	subParent.Type = typ
	subParent.Query = query
	return Parent{subParent}, nil
}

// Child returns a Child object
func (q *Query) Child(typ, query interface{}) (Child, error) {
	if typ == nil {
		return Child{}, fmt.Errorf("type could not be nil")
	}
	var subChild struct {
		Type  interface{} `json:"_type,omitempty"`
		Query interface{} `json:"_query,omitempty"`
	}
	subChild.Type = typ
	subChild.Query = query
	return Child{subChild}, nil
}

// Type returns a Type object
func (q *Query) Type(typ interface{}) (Type, error) {
	if typ == nil {
		return Type{}, fmt.Errorf("type could not be nil")
	}
	return Type{typ}, nil
}

// Status returns a Status object
func (q *Query) Status(str interface{}) (Status, error) {
	return Status{str}, nil
}

// String returns a String object
func (q *Query) String(str interface{}) (String, error) {
	return String{str}, nil
}
