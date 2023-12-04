package middleware

import (
	"mostafaba29/models"

	"gorm.io/gorm"
)

type SQLiteStore struct {
	db *gorm.DB
}

// Save saves the session data to the database, associating it with the user
func (s *SQLiteStore) Save(sid, data string, userID uint) error {
	session := &models.Session{
		ID:     sid,
		UserID: userID,
		Data:   data,
	}

	// Create or update the session record
	return s.db.Create(session).Error
}

// Load loads the session data from the database
func (s *SQLiteStore) Load(sid string) (string, uint, error) {
	var session models.Session
	err := s.db.First(&session, "id = ?", sid).Error
	if err != nil {
		return "", 0, err
	}
	return session.Data, session.UserID, nil
}

// Delete deletes the session data from the database
func (s *SQLiteStore) Delete(sid string) error {
	return s.db.Delete(&models.Session{}, "id = ?", sid).Error
}

// Close closes the GORM database connection
func (s *SQLiteStore) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
