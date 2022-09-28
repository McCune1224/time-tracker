package database

import (
	"backend/pkg/models"
)

// mock db for users
type dbUser map[string]models.User

// mock db for sessions to connect to string
type dbSession map[string]string
