package internal

import (
	"database/sql/driver"
	"time"
)

// NullByteArray struct
type NullByteArray struct {
	ByteArray []byte
	Valid     bool
}

// Scan implements the Scanner interface.
func (nb *NullByteArray) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		nb.ByteArray = append(v[:0:0], v...)
		nb.Valid = true
	default:
		nb.Valid = false
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nb NullByteArray) Value() (driver.Value, error) {
	if !nb.Valid {
		return nil, nil
	}
	return nb.ByteArray, nil
}

// NullTime struct
type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// NullInt32 struct
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt32) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		n.Int32, n.Valid = int32(v), true
		return nil
	case int32:
		n.Int32, n.Valid = v, true
		return nil
	case uint8:
		n.Int32, n.Valid = int32(v), true
		return nil
	}

	n.Valid = false

	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int32, nil
}

// NullInt16 struct
type NullInt16 struct {
	Int16 int16
	Valid bool // Valid is true if Int64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt16) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		n.Int16, n.Valid = int16(v), true
		return nil
	case int16:
		n.Int16, n.Valid = v, true
		return nil
	case uint8:
		n.Int16, n.Valid = int16(v), true
		return nil
	}

	n.Valid = false

	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int16, nil
}

// NullFloat32 struct
type NullFloat32 struct {
	Float32 float32
	Valid   bool // Valid is true if Int64 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullFloat32) Scan(value interface{}) error {
	switch v := value.(type) {
	case float64:
		n.Float32, n.Valid = float32(v), true
		return nil
	case float32:
		n.Float32, n.Valid = v, true
		return nil
	}

	n.Valid = false

	return nil
}

// Value implements the driver Valuer interface.
func (n NullFloat32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float32, nil
}
