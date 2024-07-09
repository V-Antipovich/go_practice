package accounts

import (
	// "fmt"
	"fmt"
	"go_prac/accounts/dto"
	"go_prac/accounts/models"
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

// TODO: дублирующийся код

// Создать аккаунт
func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	req := dto.CreateAccountRequest{}
	if err := c.BodyParser(&req); err != nil {
		c.Context().Logger().Printf("error: %s\n", err)
	}
	h.guard.Lock()
	// Если имя пустое или уже имеется
	// fmt.Println(req, req.Name, "q name", c.Query("name"), "q amount", c.Query("amount"))
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
	return c.SendStatus(fiber.StatusCreated)
}

// Получить аккаунт
func (h *Handler) GetAccount(c *fiber.Ctx) error {
	name := c.Query("name")
	h.guard.RLock()
	acc, ok := h.accounts[name]
	h.guard.RUnlock()
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("account not found")
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
		c.Context().Logger().Printf("error: %s\n", err)
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("no such entry: %s", req.Name))
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
		c.Context().Logger().Printf("error: %s\n", err)
	}
	h.guard.Lock()
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("no such entry: %s", req.Name))
	}
	h.accounts[req.Name].Amount = req.Amount
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusOK)
	// panic("implement me")
}

func (h *Handler) DeleteAccount(c *fiber.Ctx) error {
	req := new(dto.DeleteAccountRequest)
	// name := c.Params("name")
	// log.Println(req)
	// fmt.Println(req)
	if err := c.BodyParser(&req); err != nil {
		c.Context().Logger().Printf("error: %s\n", err)
	}
	h.guard.Lock()
	// fmt.Println(name)
	if _, ok := h.accounts[req.Name]; !ok {
		h.guard.Unlock()
		return c.SendString(fmt.Sprintf("no such entry: %s", req.Name))
	}
	delete(h.accounts, req.Name)
	h.guard.Unlock()
	return c.SendStatus(fiber.StatusNoContent)
	// panic("implement me")
}
