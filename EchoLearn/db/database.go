package db

import (
	"log"
	"os"
	"sync"

	"github.com/yourusername/echolearn/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Connect initializes the database connection.
// It ensures that the connection is established only once.
func Connect() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open("echolearn.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // Log SQL queries
		})
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
			return
		}

		log.Println("Database connection established.")

		// Auto-migrate schemas
		err = db.AutoMigrate(
			&models.User{},
			&models.GrammarRule{},
			&models.Option{},
			&models.Question{},
			&models.GameSession{},
		)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
		log.Println("Database migrated.")
	})
	return db, err
}

// GetDB returns the existing database instance.
// It panics if Connect() has not been called first.
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not connected. Call db.Connect() first.")
	}
	return db
}

// SeedDatabase checks if the database needs seeding and seeds it if necessary.
func SeedDatabase(seedFunc func(*gorm.DB)) {
	dbInstance := GetDB()

	// Simple check: if there are no questions, assume it needs seeding.
	// A more robust check might involve a separate 'versions' table or similar.
	var questionCount int64
	dbInstance.Model(&models.Question{}).Count(&questionCount)

	if questionCount == 0 {
		log.Println("No questions found, seeding database...")
		seedFunc(dbInstance)
		log.Println("Database seeded.")
	} else {
		log.Println("Database already contains data, skipping seed.")
	}
}

// ResetDatabase drops all tables and re-migrates. Useful for testing.
func ResetDatabase() error {
	dbInstance := GetDB()

	tables := []string{"game_sessions", "questions", "options", "grammar_rules", "users"}
	for _, table := range tables {
		if err := dbInstance.Migrator().DropTable(table); err != nil {
			log.Printf("Warning: could not drop table %s: %v", table, err)
		}
	}

	// Re-migrate
	err := dbInstance.AutoMigrate(
		&models.User{},
		&models.GrammarRule{},
		&models.Option{},
		&models.Question{},
		&models.GameSession{},
	)
	if err != nil {
		return err
	}
	log.Println("Database reset and migrated.")
	return nil
}

// IsErrNotFound checks if GORM error is record not found
func IsErrNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

// InitTestDB initializes an in-memory SQLite database for testing.
func InitTestDB() (*gorm.DB, error) {
	testDb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Keep tests quiet
	})
	if err != nil {
		return nil, err
	}

	err = testDb.AutoMigrate(
		&models.User{},
		&models.GrammarRule{},
		&models.Option{},
		&models.Question{},
		&models.GameSession{},
	)
	if err != nil {
		return nil, err
	}
	return testDb, nil
}

// CloseTestDB closes the test database connection.
// For in-memory SQLite, this might not be strictly necessary but good practice.
func CloseTestDB(testDb *gorm.DB) {
	sqlDB, err := testDb.DB()
	if err != nil {
		log.Printf("Error getting SQL DB from GORM for closing: %v", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		log.Printf("Error closing test database: %v", err)
	}
}

// Check if the database file exists
func DatabaseExists(filePath string) bool {
    _, err := os.Stat(filePath)
    return !os.IsNotExist(err)
} 