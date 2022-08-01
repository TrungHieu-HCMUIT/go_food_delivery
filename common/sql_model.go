package common

import "time"

type SQLModel struct {
	Id        int       `json:"-" gorm:"column:id;"`
	FakeId    *UID      `json:"id" gorm:"-"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
	Status    int       `json:"status" gorm:"status;default:1;"`
}

func (m *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(m.Id), dbType, 1)
	m.FakeId = &uid
}
