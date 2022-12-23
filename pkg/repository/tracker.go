package repository

import (
	"gorm.io/gorm"
	"time"
	"tweak_twitter/pkg/model"
)

type TrackerRepository struct {
	DB *gorm.DB
}

func (u *TrackerRepository) AddEntry(tracker *model.Tracker) (uint, error) {
	u.DB.Create(tracker)

	return tracker.ID, nil
}

func (u *TrackerRepository) AddBySourceAndValue(source string, value string) (uint, error) {

	entryId, err := u.AddEntry(&model.Tracker{
		Source:      source,
		SourceValue: value,
		CreatedAt:   time.Time{},
	})

	return entryId, err
}

func (u *TrackerRepository) GetBySourceAndValue(source string, value string) (model.Tracker, error) {
	result := &model.Tracker{
		Source:      source,
		SourceValue: value,
	}

	u.DB.Last(result)

	return *result, nil
}
