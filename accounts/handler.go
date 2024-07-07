package accounts

import (
	"go_prac/accounts/models"
	"sync"
	// "github.com/gofiber/fiber/v2"
)

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}
