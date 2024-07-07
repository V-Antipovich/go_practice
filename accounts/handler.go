package accounts

import (
	// "fmt"
	"go_prac/accounts/dto"
	"go_prac/accounts/models"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

func New() *Handler {
	return &Handler{
		accounts: make(map[string]*models.Account),
		guard:    &sync.RWMutex{},
	}
}

// Создать аккаунт
func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	// fmt.Println(c.Query("name"))
	req := dto.ChangeAccountRequest{}
	if err := c.BodyParser(&req); err != nil {
		// return err
		c.Context().Logger().Printf("error: %s\n", err)
	}

	h.guard.Lock()
	// Если имя пустое или уже имеется
	if len(req.Name) == 0 {
		return c.SendString("Can't have empty name")
	}
	if _, ok := h.accounts[req.Name]; ok {
		h.guard.Unlock()
		return c.SendString("already have")
	}
	h.accounts[req.Name] = &models.Account{
		Name:   req.Name,
		Amount: req.Amount,
	}
	h.guard.Unlock()
	return c.SendStatus(http.StatusCreated)
}

// Получить аккаунт
func (h *Handler) GetAccount(c *fiber.Ctx) error {
	name := c.Query("name")
	h.guard.RLock()
	acc, ok := h.accounts[name]
	h.guard.RUnlock()
	if !ok {
		return c.Status(http.StatusNotFound).SendString("account not found")
	}
	resp := dto.GetAccountResponse{
		Name:   acc.Name,
		Amount: acc.Amount,
	}
	return c.JSON(resp)
}

// Изменить имя аккаунта
func (h *Handler) ChangeAccount(c *fiber.Ctx) error {
	panic("implement me")
}

// Изменить баланс
func (h *Handler) PathAccount(c *fiber.Ctx) error {
	panic("implement me")
}

func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	panic("implement me")
}
