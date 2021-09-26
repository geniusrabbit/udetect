package protocol

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// UUIDFromString converts string representation of UUID to the binary
func UUIDFromString(suuid string) (*UUID, error) {
	_uuid, err := uuid.Parse(suuid)
	if err != nil {
		return nil, err
	}
	return UUIDFrom(_uuid)
}

// UUIDFrom from type UUID
func UUIDFrom(uuid uuid.UUID) (*UUID, error) {
	data, err := uuid.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return &UUID{Value: data}, nil
}

// UUIDObject returns object of UUID
func (ud *UUID) UUIDObject() (uuid.UUID, error) {
	return uuid.FromBytes(ud.Value)
}

// MarshalJSON implements json.Marshaler interface
func (ud *UUID) MarshalJSON() ([]byte, error) {
	_uuid, err := ud.UUIDObject()
	if err != nil {
		return nil, err
	}
	return []byte(`"` + _uuid.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler interface
func (ud *UUID) UnmarshalJSON(data []byte) error {
	ud.Value = ud.Value[:]
	if len(data) < 32 {
		return nil
	}
	if bytes.HasPrefix(data, []byte(`"`)) {
		data = data[1 : len(data)-1]
	}
	_uuid, err := uuid.Parse(string(data))
	if err != nil {
		return err
	}
	data, err = _uuid.MarshalBinary()
	if err != nil {
		return err
	}
	if ud.Value == nil {
		ud.Value = make([]byte, 16)
	}
	copy(ud.Value[:], data)
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (ud *UUID) MarshalText() ([]byte, error) {
	_uuid, err := ud.UUIDObject()
	if err != nil {
		return nil, err
	}
	return _uuid.MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (ud *UUID) UnmarshalText(data []byte) error {
	_uuid, err := uuid.ParseBytes(data)
	if err != nil {
		return err
	}
	data, err = _uuid.MarshalBinary()
	if err != nil {
		return err
	}
	if ud.Value == nil {
		ud.Value = make([]byte, 16)
	}
	copy(ud.Value[:], data)
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (ud *UUID) MarshalBinary() ([]byte, error) {
	return ud.Value, nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (ud *UUID) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	if ud.Value == nil {
		ud.Value = make([]byte, 16)
	}
	copy(ud.Value[:], data)
	return nil
}

var (
	_ json.Marshaler   = (*UUID)(nil)
	_ json.Unmarshaler = (*UUID)(nil)
)
