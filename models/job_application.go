package models

import (
	"time"

	"gorm.io/gorm"
)

type ApplicationStatus string

const (
	StatusBelumApply    ApplicationStatus = "Belum Apply"
	StatusApplied       ApplicationStatus = "Applied"
	StatusInterviewHRD  ApplicationStatus = "Interview HRD"
	StatusInterviewUser ApplicationStatus = "Interview User"
	StatusDitolak       ApplicationStatus = "Ditolak"
	StatusDiterima      ApplicationStatus = "Diterima"
)

type ApplicationPlatform string

const (
	PlatformLinkedIn    ApplicationPlatform = "LinkedIn"
	PlatformGlints      ApplicationPlatform = "Glints"
	PlatformJobstreet   ApplicationPlatform = "Jobstreet"
	PlatformCompanySite ApplicationPlatform = "Company Site"
	PlatformLainnya     ApplicationPlatform = "Lainnya"
)

type JobApplication struct {
	ID          uint                `gorm:"primaryKey" json:"id"`
	Company     string              `gorm:"not null" json:"company"`
	Position    string              `gorm:"not null" json:"position"`
	Platform    ApplicationPlatform `gorm:"not null" json:"platform"`
	Status      ApplicationStatus   `gorm:"not null;default:'Belum Apply'" json:"status"`
	JobLink     string              `json:"job_link"`
	Notes       string              `gorm:"type:text" json:"notes"`
	AppliedDate *time.Time          `json:"applied_date"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt      `gorm:"index" json:"-"`
}

// TableName overrides the table name
func (JobApplication) TableName() string {
	return "job_applications"
}
