// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/ent/record"
)

// Record is the model entity for the Record schema.
type Record struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastUpdatedAt holds the value of the "last_updated_at" field.
	LastUpdatedAt time.Time `json:"last_updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			values[i] = new(sql.NullInt64)
		case record.FieldCreatedAt, record.FieldLastUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Record", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case record.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case record.FieldLastUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_updated_at", values[i])
			} else if value.Valid {
				r.LastUpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return (&RecordClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", last_updated_at=")
	builder.WriteString(r.LastUpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Records is a parsable slice of Record.
type Records []*Record

func (r Records) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
