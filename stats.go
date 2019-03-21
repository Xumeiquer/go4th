package go4th

import (
	"fmt"
)

// Stats represents a stats based on a query
type Stats struct {
	Query interface{} `json:"query,omitempty"`
	Stats []*Stat     `json:"stats,omitempty"`
}

// Stat defines a stat object
type Stat struct {
	Agg    string              `json:"_agg,omitempty"`
	Field  string              `json:"_field,omitempty"`
	Select []map[string]string `json:"_select,omitempty"`
	Order  []string            `json:"_order,omitempty"`
	Size   int                 `json:"_size,omitempty"`
}

// ObservableStats represents stats for observables. [Missing some fields]
type ObservableStats struct {
	IP struct {
		Count int `json:"count,omitempty"`
	} `json:"IP,omitempty"`
	URL struct {
		Count int `json:"count,omitempty"`
	} `json:"URL,omitempty"`
	Regex struct {
		Count int `json:"count,omitempty"`
	} `json:"regex,omitempty"`
	Count int `json:"count,omitempty"`
}

// NewStats returns a new Stats object
func NewStats(q interface{}, stats ...*Stat) *Stats {
	st := &Stats{}
	st.Query = q
	for _, stat := range stats {
		st.Stats = append(st.Stats, stat)
	}
	return st
}

// NewStat returns a new Stat object
func NewStat() *Stat {
	return &Stat{}
}

// SetAgg sets Agg
func (s *Stat) SetAgg(agg string) error {
	if agg == "" {
		return fmt.Errorf("agg could not be empty")
	}
	s.Agg = agg
	return nil
}

// SetField sets Field
func (s *Stat) SetField(field string) error {
	if field == "" {
		return fmt.Errorf("field could not be empty")
	}
	s.Field = field
	return nil
}

// AddSelect sets Select
func (s *Stat) AddSelect(key, value string) error {
	if key == "" || value == "" {
		return fmt.Errorf("key or value could not be empty")
	}
	gotKey := false
	for _, sel := range s.Select {
		if _, ok := sel[key]; ok {
			gotKey = true
		}
	}
	if gotKey {
		return fmt.Errorf("key already in select")
	}
	var entry map[string]string
	entry = make(map[string]string)
	entry[key] = value
	s.Select = append(s.Select, entry)
	return nil
}

// AddOrder sets Order
func (s *Stat) AddOrder(order string) error {
	if order == "" {
		return fmt.Errorf("order could not be empty")
	}
	for _, o := range s.Order {
		if o == order {
			return fmt.Errorf("order already in order list")
		}
	}
	s.Order = append(s.Order, order)
	return nil
}

// SetSize sets Size
func (s *Stat) SetSize(n int) error {
	s.Size = n
	return nil
}
