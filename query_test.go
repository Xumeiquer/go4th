package go4th

import (
	"encoding/json"
	"fmt"
	"testing"
)

func getJSON(e interface{}) (string, error) {
	d, err := json.Marshal(&e)
	return string(d), err
}

func TestEq(t *testing.T) {
	q := NewQuery()

	field := "filed"
	value := "value"

	eq, err := q.Eq(field, value)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if eq.Field != field {
		t.Errorf("expected field to be %s, but found %s", field, eq.Field)
	}
	if eq.Value != value {
		t.Errorf("expected value to be %s, but found %s", value, eq.Value)
	}

	_, err = q.Eq("", value)
	if err == nil {
		t.Errorf("expected error to be field could not be empty, but none found")
	}

	jq, err := getJSON(eq)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_field\":\"%s\",\"_value\":\"%s\"}", field, value)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestGt(t *testing.T) {
	q := NewQuery()

	field := "filed"
	value := "value"

	gt, err := q.Gt(field, value)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if _, ok := gt.Gt[field]; !ok {
		t.Errorf("expected field to be %s, but none found", field)
	}
	if gt.Gt[field] != value {
		t.Errorf("expected value to be %s, but found %v", value, gt.Gt[field])
	}

	_, err = q.Gt("", value)
	if err == nil {
		t.Errorf("expected error to be field could not be empty, but none found")
	}

	jq, err := getJSON(gt)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_gt\":{\"%s\":\"%s\"}}", field, value)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestGte(t *testing.T) {
	q := NewQuery()

	field := "filed"
	value := "value"

	gte, err := q.Gte(field, value)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if _, ok := gte.Gte[field]; !ok {
		t.Errorf("expected field to be %s, but none found", field)
	}
	if gte.Gte[field] != value {
		t.Errorf("expected value to be %s, but found %v", value, gte.Gte[field])
	}

	_, err = q.Gte("", value)
	if err == nil {
		t.Errorf("expected error to be field could not be empty, but none found")
	}

	jq, err := getJSON(gte)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_gte\":{\"%s\":\"%s\"}}", field, value)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestLt(t *testing.T) {
	q := NewQuery()

	field := "filed"
	value := "value"

	lt, err := q.Lt(field, value)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if _, ok := lt.Lt[field]; !ok {
		t.Errorf("expected field to be %s, but none found", field)
	}
	if lt.Lt[field] != value {
		t.Errorf("expected value to be %s, but found %v", value, lt.Lt[field])
	}

	_, err = q.Lt("", value)
	if err == nil {
		t.Errorf("expected error to be field could not be empty, but none found")
	}

	jq, err := getJSON(lt)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_lt\":{\"%s\":\"%s\"}}", field, value)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestLte(t *testing.T) {
	q := NewQuery()

	field := "filed"
	value := "value"

	lte, err := q.Lte(field, value)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if _, ok := lte.Lte[field]; !ok {
		t.Errorf("expected field to be %s, but none found", field)
	}
	if lte.Lte[field] != value {
		t.Errorf("expected value to be %s, but found %v", value, lte.Lte[field])
	}

	_, err = q.Lte("", value)
	if err == nil {
		t.Errorf("expected error to be field could not be empty, but none found")
	}

	jq, err := getJSON(lte)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_lte\":{\"%s\":\"%s\"}}", field, value)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestAnd(t *testing.T) {
	q := NewQuery()
	elems := 2
	fieldA := "filed"
	fieldB := "filed"
	valueA := "value"
	valueB := "value"

	eq, _ := q.Eq(fieldA, valueA)
	lt, _ := q.Lt(fieldB, valueB)

	and, err := q.And(eq, lt)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if len(and.And) != elems {
		t.Errorf("expected to have %d elems, but found %d", elems, len(and.And))
	}

	if and.And[0].(Eq).Field != fieldA {
		t.Errorf("expected first element to have filed %s, but found %s", fieldA, and.And[0].(Eq).Field)
	}
	if and.And[0].(Eq).Value != valueA {
		t.Errorf("expected first element to have filed %s, but found %s", valueA, and.And[0].(Eq).Value)
	}
	if _, ok := and.And[1].(Lt).Lt[fieldB]; !ok {
		t.Errorf("expected second element to be Lt with field %s, but found %v", fieldB, and.And[1].(Lt).Lt[fieldB])
	}
	if and.And[1].(Lt).Lt[fieldB] != valueB {
		t.Errorf("expected second element to have filed %s, but found %s", valueB, and.And[1].(Eq).Value)
	}

	jq, err := getJSON(and)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_and\":[{\"_field\":\"%s\",\"_value\":\"%s\"},{\"_lt\":{\"%s\":\"%s\"}}]}", fieldA, valueA, fieldB, valueB)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestOr(t *testing.T) {
	q := NewQuery()
	elems := 2
	fieldA := "filed"
	fieldB := "filed"
	valueA := "value"
	valueB := "value"

	eq, _ := q.Eq(fieldA, valueA)
	lt, _ := q.Lt(fieldB, valueB)

	or, err := q.Or(eq, lt)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if len(or.Or) != elems {
		t.Errorf("expected to have %d elems, but found %d", elems, len(or.Or))
	}

	if or.Or[0].(Eq).Field != fieldA {
		t.Errorf("expected first element to have filed %s, but found %s", fieldA, or.Or[0].(Eq).Field)
	}
	if or.Or[0].(Eq).Value != valueA {
		t.Errorf("expected first element to have filed %s, but found %s", valueA, or.Or[0].(Eq).Value)
	}
	if _, ok := or.Or[1].(Lt).Lt[fieldB]; !ok {
		t.Errorf("expected second element to be Lt with field %s, but found %v", fieldB, or.Or[1].(Lt).Lt[fieldB])
	}
	if or.Or[1].(Lt).Lt[fieldB] != valueB {
		t.Errorf("expected second element to have filed %s, but found %s", valueB, or.Or[1].(Eq).Value)
	}

	jq, err := getJSON(or)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_or\":[{\"_field\":\"%s\",\"_value\":\"%s\"},{\"_lt\":{\"%s\":\"%s\"}}]}", fieldA, valueA, fieldB, valueB)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestNot(t *testing.T) {
	q := NewQuery()

	valueA := "value"

	not, err := q.Not(valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	jq, err := getJSON(not)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_not\":\"%s\"}", valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestIn(t *testing.T) {
	q := NewQuery()

	fieldA := "field"
	valueA := "value"

	in, err := q.In(fieldA, valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if in.In.Field != fieldA {
		t.Errorf("expected to be %s, but found %s", fieldA, in.In.Field)
	}
	if in.In.Value != valueA {
		t.Errorf("expected to be %s, but found %s", valueA, in.In.Value)
	}

	jq, err := getJSON(in)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_in\":{\"_field\":\"%s\",\"_value\":\"%s\"}}", fieldA, valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestContains(t *testing.T) {
	q := NewQuery()

	valueA := "value"

	contains, err := q.Contains(valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if contains.Contains != valueA {
		t.Errorf("expected contains to be %s, but found %s", valueA, contains.Contains)
	}

	jq, err := getJSON(contains)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}

	jqStr := fmt.Sprintf("{\"_contains\":\"%s\"}", valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestID(t *testing.T) {
	q := NewQuery()

	valueA := "value"

	id, err := q.ID(valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if id.ID != valueA {
		t.Errorf("expected id to be %s, but found %s", valueA, id.ID)
	}

	jq, err := getJSON(id)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}

	jqStr := fmt.Sprintf("{\"_id\":\"%s\"}", valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestBetween(t *testing.T) {
	q := NewQuery()

	fieldA := "field"
	fromA := "from"
	toA := "to"

	between, err := q.Between(fieldA, fromA, toA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if between.Between.Field != fieldA {
		t.Errorf("expected to be %s, but found %s", fieldA, between.Between.Field)
	}
	if between.Between.From != fromA {
		t.Errorf("expected to be %s, but found %s", fromA, between.Between.From)
	}
	if between.Between.To != toA {
		t.Errorf("expected to be %s, but found %s", toA, between.Between.To)
	}

	jq, err := getJSON(between)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_between\":{\"_field\":\"%s\",\"_from\":\"%s\",\"_to\":\"%s\"}}", fieldA, fromA, toA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestParentID(t *testing.T) {
	q := NewQuery()

	typeA := "type"
	idA := "id"

	parentID, err := q.ParentID(typeA, idA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if parentID.ParentID.Type != typeA {
		t.Errorf("expected to be %s, but found %s", typeA, parentID.ParentID.Type)
	}
	if parentID.ParentID.ID != idA {
		t.Errorf("expected to be %s, but found %s", idA, parentID.ParentID.ID)
	}

	jq, err := getJSON(parentID)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_parent\":{\"_type\":\"%s\",\"_id\":\"%s\"}}", typeA, idA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestParent(t *testing.T) {
	q := NewQuery()

	typeA := "type"
	queryA := "id"

	parent, err := q.Parent(typeA, queryA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if parent.Parent.Type != typeA {
		t.Errorf("expected to be %s, but found %s", typeA, parent.Parent.Type)
	}
	if parent.Parent.Query != queryA {
		t.Errorf("expected to be %s, but found %s", queryA, parent.Parent.Query)
	}

	jq, err := getJSON(parent)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_parent\":{\"_type\":\"%s\",\"_query\":\"%s\"}}", typeA, queryA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestChild(t *testing.T) {
	q := NewQuery()

	typeA := "type"
	queryA := "id"

	child, err := q.Child(typeA, queryA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}

	if child.Child.Type != typeA {
		t.Errorf("expected to be %s, but found %s", typeA, child.Child.Type)
	}
	if child.Child.Query != queryA {
		t.Errorf("expected to be %s, but found %s", queryA, child.Child.Query)
	}

	jq, err := getJSON(child)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}
	jqStr := fmt.Sprintf("{\"_child\":{\"_type\":\"%s\",\"_query\":\"%s\"}}", typeA, queryA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestType(t *testing.T) {
	q := NewQuery()

	valueA := "value"

	typ, err := q.Type(valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if typ.Type != valueA {
		t.Errorf("expected type to be %s, but found %s", valueA, typ.Type)
	}

	jq, err := getJSON(typ)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}

	jqStr := fmt.Sprintf("{\"_type\":\"%s\"}", valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestString(t *testing.T) {
	q := NewQuery()

	valueA := "value"

	str, err := q.String(valueA)
	if err != nil {
		t.Errorf("expected error to be nil, but found %s", err.Error())
	}
	if str.String != valueA {
		t.Errorf("expected string to be %s, but found %s", valueA, str.String)
	}

	jq, err := getJSON(str)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}

	jqStr := fmt.Sprintf("{\"_string\":\"%s\"}", valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}

func TestBuildQuery(t *testing.T) {
	q := NewQuery()

	fieldA := "field"
	valueA := "value"

	eq, _ := q.Eq(fieldA, valueA)

	bq := BuildQuery(eq)

	jq, err := getJSON(bq)
	if err != nil {
		t.Errorf("unexpected error, %s", err.Error())
	}

	jqStr := fmt.Sprintf("{\"query\":{\"_field\":\"%s\",\"_value\":\"%s\"}}", fieldA, valueA)
	if jq != jqStr {
		t.Errorf("expecting string to be %s, but found %s", jqStr, jq)
	}
}
