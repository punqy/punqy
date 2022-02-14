package storage

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Entity struct {
	ID        uuid.UUID `json:"id" db:"id"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

func (u *User) SetUpdated() {
	u.UpdatedAt = time.Now().UTC()
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

func (e *Entity) Init() error {
	e.CreatedAt = time.Now()
	e.UpdatedAt = e.CreatedAt
	id, err := uuid.NewRandom()
	e.ID = id
	return err
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
