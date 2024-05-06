package eid

import (
	"database/sql/driver"
	"fmt"
)

// Value implements the driver.Valuer interface
func (e *EID) Value() (driver.Value, error) {
	err := e.Valid()
	if err != nil {
		return nil, err
	}
	return e.String(), nil
}

// Scan implements the sql.Scanner interface
func (e *EID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	eidString, ok := src.(string)
	if !ok {
		return fmt.Errorf("expected string when scanning EID field")
	}
	out, err := Parse(eidString)
	if err != nil {
		return err
	}
	e.ID = out.ID
	e.Type = out.Type
	return nil
}
