package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"jobsme/config"
	"jobsme/models"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup test database
func setupTestDB() {
	var err error
	config.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	config.DB.AutoMigrate(&models.JobApplication{})
}

// Test GetAllJobs handler
func TestGetAllJobs(t *testing.T) {
	setupTestDB()

	// Create test data
	testJob := models.JobApplication{
		Company:  "Test Company",
		Position: "Test Position",
		Platform: models.PlatformLinkedIn,
		Status:   models.StatusApplied,
	}
	config.DB.Create(&testJob)

	// Create Fiber app
	app := fiber.New()
	app.Get("/api/jobs", GetAllJobs)

	// Create request
	req := httptest.NewRequest("GET", "/api/jobs", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error testing GetAllJobs: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Read response body
	body, _ := io.ReadAll(resp.Body)
	var jobs []models.JobApplication
	json.Unmarshal(body, &jobs)

	if len(jobs) != 1 {
		t.Errorf("Expected 1 job, got %d", len(jobs))
	}

	if jobs[0].Company != "Test Company" {
		t.Errorf("Expected company 'Test Company', got '%s'", jobs[0].Company)
	}
}

// Test CreateJob handler
func TestCreateJob(t *testing.T) {
	setupTestDB()

	app := fiber.New()
	app.Post("/api/jobs", CreateJob)

	newJob := models.JobApplication{
		Company:  "New Company",
		Position: "Developer",
		Platform: models.PlatformGlints,
		Status:   models.StatusBelumApply,
		JobLink:  "https://example.com",
		Notes:    "Test notes",
	}

	jsonData, _ := json.Marshal(newJob)
	req := httptest.NewRequest("POST", "/api/jobs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error testing CreateJob: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Errorf("Expected status 201, got %d", resp.StatusCode)
	}

	// Verify job was created
	var job models.JobApplication
	config.DB.First(&job)

	if job.Company != "New Company" {
		t.Errorf("Expected company 'New Company', got '%s'", job.Company)
	}
}

// Test UpdateJob handler
func TestUpdateJob(t *testing.T) {
	setupTestDB()

	// Create initial job
	testJob := models.JobApplication{
		Company:  "Original Company",
		Position: "Original Position",
		Platform: models.PlatformLinkedIn,
		Status:   models.StatusBelumApply,
	}
	config.DB.Create(&testJob)

	app := fiber.New()
	app.Put("/api/jobs/:id", UpdateJob)

	// Update data
	updateData := models.JobApplication{
		Company:  "Updated Company",
		Position: "Updated Position",
		Platform: models.PlatformJobstreet,
		Status:   models.StatusApplied,
	}

	jsonData, _ := json.Marshal(updateData)
	req := httptest.NewRequest("PUT", "/api/jobs/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error testing UpdateJob: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Verify job was updated
	var job models.JobApplication
	config.DB.First(&job, 1)

	if job.Company != "Updated Company" {
		t.Errorf("Expected company 'Updated Company', got '%s'", job.Company)
	}

	// Check AppliedDate was set
	if job.AppliedDate == nil {
		t.Error("AppliedDate should be set when status changed to Applied")
	}
}

// Test DeleteJob handler
func TestDeleteJob(t *testing.T) {
	setupTestDB()

	// Create test job
	testJob := models.JobApplication{
		Company:  "Delete Me",
		Position: "Test",
		Platform: models.PlatformLinkedIn,
		Status:   models.StatusApplied,
	}
	config.DB.Create(&testJob)

	app := fiber.New()
	app.Delete("/api/jobs/:id", DeleteJob)

	req := httptest.NewRequest("DELETE", "/api/jobs/1", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Error testing DeleteJob: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Verify job was deleted (soft delete)
	var job models.JobApplication
	result := config.DB.Unscoped().First(&job, 1)

	if result.Error != nil {
		t.Error("Job should still exist (soft deleted)")
	}

	if job.DeletedAt.Time.IsZero() {
		t.Error("DeletedAt should be set")
	}
}

// Test status auto-date setting
func TestAppliedDateAutoSet(t *testing.T) {
	setupTestDB()

	app := fiber.New()
	app.Post("/api/jobs", CreateJob)

	// Create job with status "Applied"
	newJob := models.JobApplication{
		Company:  "Test Company",
		Position: "Developer",
		Platform: models.PlatformLinkedIn,
		Status:   models.StatusApplied,
	}

	jsonData, _ := json.Marshal(newJob)
	req := httptest.NewRequest("POST", "/api/jobs", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	app.Test(req)

	// Verify AppliedDate was set
	var job models.JobApplication
	config.DB.First(&job)

	if job.AppliedDate == nil {
		t.Error("AppliedDate should be automatically set when status is not 'Belum Apply'")
	}

	// Check it's today
	if !job.AppliedDate.Equal(time.Now().Truncate(24 * time.Hour)) {
		// Allow some time difference
		diff := time.Since(*job.AppliedDate)
		if diff > time.Hour {
			t.Error("AppliedDate should be set to current time")
		}
	}
}
