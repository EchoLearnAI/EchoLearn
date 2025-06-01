package main

import (
	"encoding/json"
	"log"

	"gorm.io/gorm"
)

// Define the category to topic mapping
var categoryTopicMap = map[string][]string{
	"Network":        {"Network", "Cache & CDN"},
	"Security":       {"Cybersecurity", "Monitoring"},
	"Virtualization": {"VMs", "Container"},
	"Cloud":          {"Azure", "AWS", "GCP"},
	"DevOps":         {"Kubernetes", "Git", "Pipelines & CI/CD", "Terraform", "Ansible"},
	"Data":           {"Databases", "APIs"},
	"Architecture":   {"Architecture"},
	"Operations":     {"Admin & Ops"},
	"Linux":          {"Linux"},
}

// TopicDetails holds the name and slug for a topic, used during seeding.
// We use this because the original `questions` array in `data.go` uses full names for topics.
var topicDetails = map[string]string{
	"Network":           "network",
	"Cache & CDN":       "cache-cdn",
	"Cybersecurity":     "cybersecurity",
	"Monitoring":        "monitoring",
	"VMs":               "vms",
	"Container":         "container",
	"Azure":             "azure",
	"AWS":               "aws",
	"GCP":               "gcp",
	"Kubernetes":        "kubernetes",
	"Git":               "git",
	"Pipelines & CI/CD": "pipelines-ci-cd",
	"Terraform":         "terraform",
	"Ansible":           "ansible",
	"Databases":         "databases",
	"APIs":              "apis",
	"Architecture":      "architecture",
	"Admin & Ops":       "admin-ops",
	"Linux":             "linux",
}

// CategoryDetails maps category names to their slugs.
var categoryDetails = map[string]string{
	"Network":        "network",
	"Security":       "security",
	"Virtualization": "virtualization",
	"Cloud":          "cloud",
	"DevOps":         "devops",
	"Data":           "data",
	"Architecture":   "architecture",
	"Operations":     "ops", // slug for "Admin & Ops" category is "ops"
	"Linux":          "linux",
}

func SeedDatabase(db *gorm.DB) {
	log.Println("Starting database seeding...")

	seededCategories := seedCategories(db)
	seedTopics(db, seededCategories)
	seedQuestions(db)

	log.Println("Database seeding completed.")
}

func seedCategories(db *gorm.DB) map[string]Category {
	log.Println("Seeding categories...")
	createdCategories := make(map[string]Category)

	for categoryName, categorySlug := range categoryDetails {
		var category Category
		result := db.FirstOrCreate(&category, Category{Name: categoryName, Slug: categorySlug})
		if result.Error != nil {
			log.Printf("Error seeding category %s: %v\n", categoryName, result.Error)
			continue
		}
		if result.RowsAffected > 0 {
			log.Printf("Seeded category: %s (Slug: %s)\n", category.Name, category.Slug)
		} else {
			// log.Printf("Category %s (Slug: %s) already exists.\n", category.Name, category.Slug)
		}
		createdCategories[categoryName] = category
	}
	return createdCategories
}

func seedTopics(db *gorm.DB, seededCategories map[string]Category) {
	log.Println("Seeding topics...")
	for categoryNameFromMap, topicNames := range categoryTopicMap {
		category, ok := seededCategories[categoryNameFromMap]
		if !ok {
			log.Printf("Category %s not found in seededCategories map. Skipping topics: %v\n", categoryNameFromMap, topicNames)
			continue
		}

		for _, topicName := range topicNames {
			topicSlug, slugExists := topicDetails[topicName]
			if !slugExists {
				log.Printf("Slug for topic '%s' not defined in topicDetails. Skipping.\n", topicName)
				continue
			}

			var topic Topic
			result := db.FirstOrCreate(&topic, Topic{Name: topicName, Slug: topicSlug, CategoryID: category.ID})
			if result.Error != nil {
				log.Printf("Error seeding topic %s: %v\n", topicName, result.Error)
				continue
			}
			if result.RowsAffected > 0 {
				log.Printf("Seeded topic: %s (Slug: %s) under Category: %s\n", topic.Name, topic.Slug, category.Name)
			} else {
				// log.Printf("Topic %s (Slug: %s) already exists.\n", topic.Name, topic.Slug)
			}
		}
	}
}

// seedQuestions populates the questions table from the mock data in data.go
func seedQuestions(db *gorm.DB) {
	log.Println("Seeding questions...")

	for _, qData := range questionsToSeed { // Iterate over questionsToSeed
		// 1. Find TopicID using the topic name from qData
		topicSlug, slugExists := topicDetails[qData.Topic] // qData.Topic is string from SeedQuestion
		if !slugExists {
			log.Printf("Slug for topic '%s' (from question ID %s) not defined in topicDetails. Skipping question.\n", qData.Topic, qData.ID)
			continue
		}

		var topic Topic // This is main.Topic (GORM model)
		if err := db.First(&topic, "slug = ?", topicSlug).Error; err != nil {
			log.Printf("Topic with slug '%s' (for question ID %s) not found in DB. Skipping question. Error: %v\n", topicSlug, qData.ID, err)
			continue
		}

		// 2. Marshal qData.Options (which is []SeedOption) to datatypes.JSON
		optionsJSON, err := json.Marshal(qData.Options)
		if err != nil {
			log.Printf("Error marshalling options for question ID %s: %v. Skipping question.\n", qData.ID, err)
			continue
		}

		// 3. Prepare the Question object for upsert
		dbQuestion := Question{ // This is main.Question (GORM model)
			TopicID:    topic.ID,
			Difficulty: qData.Difficulty,
			Text:       qData.Text,
			Options:    optionsJSON,
			OriginalID: qData.ID, // qData.ID is string from SeedQuestion
		}

		// Upsert logic: Assign attributes to update if found, otherwise create.
		// The Where clause specifies the unique key (OriginalID) for finding the record.
		// Assign specifies the fields to update if the record exists.
		// FirstOrCreate creates the record if it doesn't exist, or updates it (due to Assign) if it does.
		result := db.Where(Question{OriginalID: qData.ID}).Assign(dbQuestion).FirstOrCreate(&dbQuestion)

		if result.Error != nil {
			log.Printf("SEED_ERROR: Error processing question ID %s (Topic: %s, Difficulty: %s): %v\n", qData.ID, qData.Topic, qData.Difficulty, result.Error)
			continue
		}

		if result.RowsAffected > 0 {
			// Check if it was a create or an update. This is a bit tricky with FirstOrCreate alone.
			// A common way is to check CreatedAt vs UpdatedAt, but that requires the model to have been loaded.
			// For simplicity, we'll assume RowsAffected > 0 means it was either created or updated.
			log.Printf("SEED_SUCCESS: Processed (created/updated) question with OriginalID: %s (Topic: %s, Difficulty: %s, DB_ID: %d)\n", qData.ID, qData.Topic, qData.Difficulty, dbQuestion.ID)
		} else {
			log.Printf("SEED_INFO: Question with OriginalID %s (Topic: %s, Difficulty: %s) already exists and was not changed (DB_ID: %d).\n", qData.ID, qData.Topic, qData.Difficulty, dbQuestion.ID)
		}
	}
	log.Println("Finished seeding questions.")
}
