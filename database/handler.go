package database

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type IDatabase interface {
	Insert(ctx context.Context, tableName string, data interface{}) error
	Update(ctx context.Context, tableName string, data interface{}, condition string, args ...interface{}) error
	GetOne(ctx context.Context, tableName string, result interface{}, condition string, args ...interface{}) error
	Delete(ctx context.Context, tableName string, condition string, args ...interface{}) error
}

func NewDatabase() (*Database, error) {
	db, err := gorm.Open(sqlite.Open("crud.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Database{
		DB: db,
	}, nil
}

func (db *Database) Insert(ctx context.Context, tableName string, data interface{}) error {
	if err := db.DB.Table(tableName).Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) Update(ctx context.Context, tableName string, data interface{}, condition string, args ...interface{}) error {
	result := db.DB.Table(tableName).Where(condition, args...).Updates(data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Database) GetOne(ctx context.Context, tableName string, result interface{}, condition string, args ...interface{}) error {
	if err := db.DB.Table(tableName).Where(condition, args...).First(result).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) Delete(ctx context.Context, tableName string, condition string, args ...interface{}) error {
	if err := db.DB.Table(tableName).Where(condition, args...).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
