package storage

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Entity struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type ColVal struct {
	Column string
	Value  interface{}
}

func NewEntity() (Entity, error) {
	e := Entity{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return e, err
	}
	e.ID = id
	return e, nil
}

func (e *Entity) NewId() error {
	uid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	e.ID = uid
	return nil
}

func (e Entity) GetColVal() ([]ColVal, error) {
	var rows []ColVal
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		col := field.Tag.Get("db")
		if col == "" {
			col = field.Name
		}
		val, err := driver.DefaultParameterConverter.ConvertValue(v.Field(i).Interface())
		if err != nil {
			return rows, err
		}
		rows = append(rows, ColVal{Column: col, Value: val})
	}
	return rows, nil
}

type StringList []string

func (r StringList) Value() (driver.Value, error) {
	bytes, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}
	return driver.Value(bytes), nil
}

func (r *StringList) Scan(v interface{}) error {
	if v == nil {
		return nil
	}
	recs := &StringList{}
	switch v.(type) {
	case []byte:
		if fmt.Sprintf("%s", v) != "null" {
			if err := json.Unmarshal(v.([]byte), &recs); err != nil {
				return err
			}
		}
	}
	*r = *recs
	return nil
}
