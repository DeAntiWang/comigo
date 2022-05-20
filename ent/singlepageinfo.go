// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/yumenaka/comi/ent/singlepageinfo"
)

// SinglePageInfo is the model entity for the SinglePageInfo schema.
type SinglePageInfo struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SinglePageInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case singlepageinfo.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SinglePageInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SinglePageInfo fields.
func (spi *SinglePageInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case singlepageinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			spi.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this SinglePageInfo.
// Note that you need to call SinglePageInfo.Unwrap() before calling this method if this SinglePageInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (spi *SinglePageInfo) Update() *SinglePageInfoUpdateOne {
	return (&SinglePageInfoClient{config: spi.config}).UpdateOne(spi)
}

// Unwrap unwraps the SinglePageInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (spi *SinglePageInfo) Unwrap() *SinglePageInfo {
	tx, ok := spi.config.driver.(*txDriver)
	if !ok {
		panic("ent: SinglePageInfo is not a transactional entity")
	}
	spi.config.driver = tx.drv
	return spi
}

// String implements the fmt.Stringer.
func (spi *SinglePageInfo) String() string {
	var builder strings.Builder
	builder.WriteString("SinglePageInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", spi.ID))
	builder.WriteByte(')')
	return builder.String()
}

// SinglePageInfos is a parsable slice of SinglePageInfo.
type SinglePageInfos []*SinglePageInfo

func (spi SinglePageInfos) config(cfg config) {
	for _i := range spi {
		spi[_i].config = cfg
	}
}
