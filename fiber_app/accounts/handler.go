package accounts

import (
	"fmt"
	"go_prac/fiber_app/accounts/dto"
	"go_prac/fiber_app/accounts/models"
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
	req := dto.CreateAccountRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.SendString(fmt.Sprintf("errors while parsing the request: %s", err))
	}
	if len(req.Name) == 0 {
		return c.SendString("Can't have empty name")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; ok {
		h.guard.Unlock()
		return c.SendString("Account is already present")
	}
	h.accounts[req.Name] = &models.Account{
		Name:   req.Name,
		Amount: req.Amount,
	}
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusCreated)
}

// Получить аккаунт
func (h *Handler) GetAccount(c *fiber.Ctx) error {
	name := c.Query("name")
	if len(name) == 0 {
		return c.SendString("Can't have empty name")
	}
	h.guard.RLock()
	acc, ok := h.accounts[name]
	h.guard.RUnlock()
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("Account not found")
	}
	resp := dto.GetAccountResponse{
		Name:   acc.Name,
		Amount: acc.Amount,
	}
	return c.JSON(resp)
}

// Изменить имя аккаунта
func (h *Handler) ChangeAccount(c *fiber.Ctx) error {
	req := new(dto.ChangeAccountRequest)
	if err := c.BodyParser(&req); err != nil {
		return c.SendString(fmt.Sprintf("errors while parsing the request: %s", err))
	}
	if len(req.Name) == 0 || len(req.NewName) == 0 {
		return c.SendString("Can't have empty name")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("Account '%s' not found", req.Name))
	}
	amount := h.accounts[req.Name].Amount
	delete(h.accounts, req.Name)
	h.accounts[req.NewName] = &models.Account{
		Name:   req.NewName,
		Amount: amount,
	}
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusOK)
}

// Изменить баланс
func (h *Handler) PatchAccount(c *fiber.Ctx) error {
	req := new(dto.PatchAccountRequest)
	if err := c.BodyParser(&req); err != nil {
		return c.SendString(fmt.Sprintf("errors while parsing the request: %s", err))
	}
	if len(req.Name) == 0 {
		return c.SendString("Can't have empty name")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("Account '%s' not found", req.Name))
	}
	h.accounts[req.Name].Amount = req.Amount
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusOK)
}

// Удалить аккаунт
func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	req := new(dto.DeleteAccountRequest)
	if err := c.BodyParser(&req); err != nil {
		return c.SendString(fmt.Sprintf("errors while parsing the request: %s", err))
	}
	if len(req.Name) == 0 {
		return c.SendString("Can't have empty name")
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("Account '%s' not found", req.Name))
	}
	delete(h.accounts, req.Name)
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusNoContent)
}
