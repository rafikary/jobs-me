package handlers

import (
	"jobsme/config"
	"jobsme/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetDashboard renders the dashboard page
func GetDashboard(c *fiber.Ctx) error {
	var jobs []models.JobApplication
	var stats struct {
		Total         int64
		BelumApply    int64
		Applied       int64
		InterviewHRD  int64
		InterviewUser int64
		Ditolak       int64
		Diterima      int64
	}

	// Get all jobs
	config.DB.Order("created_at DESC").Find(&jobs)

	// Get statistics
	config.DB.Model(&models.JobApplication{}).Count(&stats.Total)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusBelumApply).Count(&stats.BelumApply)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusApplied).Count(&stats.Applied)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusInterviewHRD).Count(&stats.InterviewHRD)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusInterviewUser).Count(&stats.InterviewUser)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusDitolak).Count(&stats.Ditolak)
	config.DB.Model(&models.JobApplication{}).Where("status = ?", models.StatusDiterima).Count(&stats.Diterima)

	return c.Render("dashboard", fiber.Map{
		"Title": "Job Application Tracker",
		"Jobs":  jobs,
		"Stats": stats,
	})
}

// GetAllJobs returns all job applications as JSON
func GetAllJobs(c *fiber.Ctx) error {
	var jobs []models.JobApplication

	status := c.Query("status")
	platform := c.Query("platform")
	search := c.Query("search")

	query := config.DB.Model(&models.JobApplication{})

	if status != "" {
		query = query.Where("status = ?", status)
	}
	if platform != "" {
		query = query.Where("platform = ?", platform)
	}
	if search != "" {
		query = query.Where("company ILIKE ? OR position ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Order("created_at DESC").Find(&jobs)

	return c.JSON(jobs)
}

// GetJobByID returns a single job application
func GetJobByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var job models.JobApplication

	if err := config.DB.First(&job, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Job application not found",
		})
	}

	return c.JSON(job)
}

// CreateJob creates a new job application
func CreateJob(c *fiber.Ctx) error {
	job := new(models.JobApplication)

	if err := c.BodyParser(job); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Set applied date if status is not "Belum Apply"
	if job.Status != models.StatusBelumApply {
		now := time.Now()
		job.AppliedDate = &now
	}

	if err := config.DB.Create(&job).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create job application",
		})
	}

	return c.Status(201).JSON(job)
}

// UpdateJob updates an existing job application
func UpdateJob(c *fiber.Ctx) error {
	id := c.Params("id")
	var job models.JobApplication

	if err := config.DB.First(&job, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Job application not found",
		})
	}

	oldStatus := job.Status

	if err := c.BodyParser(&job); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Update applied date if status changed from "Belum Apply"
	if oldStatus == models.StatusBelumApply && job.Status != models.StatusBelumApply {
		now := time.Now()
		job.AppliedDate = &now
	}

	if err := config.DB.Save(&job).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update job application",
		})
	}

	return c.JSON(job)
}

// DeleteJob deletes a job application
func DeleteJob(c *fiber.Ctx) error {
	id := c.Params("id")
	var job models.JobApplication

	if err := config.DB.First(&job, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Job application not found",
		})
	}

	if err := config.DB.Delete(&job).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete job application",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Job application deleted successfully",
	})
}
