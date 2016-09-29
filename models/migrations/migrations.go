package migrations

import (
    "fmt"

    "github.com/jinzhu/gorm"

    "github.com/zeuxisoo/go-goonui/modules/log"
)

// Migration version table
type MigrationVersion struct {
    ID      int64 `gorm:"primary_key;AUTO_INCREMENT"`
    Version int64
}

// Migration object
type Migration interface {
    Description()           string
    Migrate(db *gorm.DB)    error
}

type migration struct {
    description     string
    migrate         func(db *gorm.DB) error
}

func NewMigration(description string, migrate func(db *gorm.DB) error) Migration {
    return &migration{
        description: description,
        migrate    : migrate,
    }
}

func (m *migration) Description() string {
    return m.description
}

func (m *migration) Migrate(db *gorm.DB) error {
    return m.migrate(db)
}

// Migration list
var migrations = []Migration{
    NewMigration("Testing A", testingA),
    NewMigration("Testing B", testingB),
    NewMigration("Testing C", testingC),
}

func testingA(db *gorm.DB) error {

    return nil
}

func testingB(db *gorm.DB) error {

    return nil
}

func testingC(db *gorm.DB) error {

    return nil
}

// Do all migrations
func Migrate(db *gorm.DB) error {
    // Create migration version table
    if err := db.AutoMigrate(&MigrationVersion{}).Error; err != nil {
        return fmt.Errorf("Faield to migration version table: %v", err)
    }

    // Add current version to migration version table when first record not exists
    var migrationVersion MigrationVersion

    if db.First(&migrationVersion, 1).Error == gorm.ErrRecordNotFound {
        migrationVersion.Version = int64(len(migrations))

        if err := db.Create(&migrationVersion).Error; err != nil {
            fmt.Errorf("Failed to add current migration version: %v", err)
        }
    }

    // Execute new migrations
    currentVersion := migrationVersion.Version

    for i, migration := range migrations[currentVersion:] {
        log.Infof("Migration: %s\n", migration.Description())

        if err := migration.Migrate(db); err != nil {
            return fmt.Errorf("Faield to execute migration: %v", err)
        }

        // Update current version to migration version table
        migrationVersion.Version = currentVersion + int64(i) + 1

        if err := db.Save(&migrationVersion).Error; err != nil {
            return err
        }
    }

    return nil
}
