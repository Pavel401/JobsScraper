package models

import "gorm.io/gorm"

type GormDB struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Location  string `json:"location"`
	CreatedAt int64  `json:"createdAt"` // Use int64 for timestamp
	Company   string `json:"company"`
	ApplyURL  string `json:"applyUrl"`
	ImageUrl  string `json:"imageUrl"`
}

func ConvertToGormDB(job Job) GormDB {
	return GormDB{
		Title:     job.Title,
		Location:  job.Location,
		CreatedAt: job.CreatedAt,
		Company:   job.Company,
		ApplyURL:  job.ApplyURL,
		ImageUrl:  job.ImageUrl,
	}
}
func (GormDB) TableName() string {
	return "gorm_dbs"
}
