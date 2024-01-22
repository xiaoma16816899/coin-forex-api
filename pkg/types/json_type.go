package types

import (
	"database/sql/driver"

	"errors"
	"fmt"

	"github.com/goccy/go-json"
)

type JSON json.RawMessage

// Convert Json string to bytes,
func NewJSON(value any) (JSON, error) {
	bytes, err := json.Marshal(value)
	return JSON(bytes), err
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

// Convert Struct to Bytes
func (j *JSON) Set(value interface{}) error {
	bytes, err := json.Marshal(value)
	*j = JSON(bytes)
	return err
}

// Convert Bytes
func (j JSON) Out(out interface{}) error {
	return json.Unmarshal(j, &out)
}

func (j JSON) String() string {
	return string(j)
}

func (j JSON) ArrayInt() []int {
	res := make([]int, 0)
	json.Unmarshal(j, &res)
	return res
}

func (j JSON) Object() map[string]any {
	res := make(map[string]any, 0)
	json.Unmarshal(j, &res)
	return res
}

func (j *JSON) UnmarshalJSON(b []byte) error {
	result := json.RawMessage{}
	err := result.UnmarshalJSON(b)
	*j = JSON(result)
	return err
}

func (j JSON) MarshalJSON() ([]byte, error) {
	return json.RawMessage(j).MarshalJSON()
}

func (j JSON) IsValid() bool {
	return json.Valid(j)
}
