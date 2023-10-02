package entities

import "time"

type BaseTime struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

func GetCreatedAtCurrentTime() BaseTime {
	return BaseTime{
		CreatedAt: time.Now(),
	}
}

func GetUpdatedCurrentTime() BaseTime {
	return BaseTime{
		UpdatedAt: time.Now(),
	}
}

func GetDeletedAtCurrentTime() BaseTime {
	return BaseTime{
		DeletedAt: time.Now(),
	}
}
