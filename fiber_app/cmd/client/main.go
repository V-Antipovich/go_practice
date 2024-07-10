package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go_prac/fiber_app/accounts/dto"

	"github.com/gofiber/fiber/v2"
)

type Command struct {
	Port    int
	Host    string
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	case "get":
		return c.getQ()
	case "delete":
		return c.del()
	case "patch":
		return c.patch()
	case "change":
		return c.change()
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
	}
}

func main() {
	portVal := flag.Int("port", 3000, "server port")
	hostVal := flag.String("host", "0.0.0.0", "server host")
	cmdVal := flag.String("cmd", "", "command")
	nameVal := flag.String("name", "", "name of the account")
	ammountVal := flag.Int("amount", 0, "ammount of the account")
	newNameVal := flag.String("newname", "", "new name of the account")
	flag.Parse()
	cmd := Command{
		Port:    *portVal,
		Host:    *hostVal,
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *ammountVal,
		NewName: *newNameVal,
	}
	if err := cmd.Do(); err != nil {
		panic(err)
	}
}

func (cmd *Command) create() error {
	reqBody := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	// json marshall
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	a := fiber.AcquireAgent()
	a.Add("Content-Type", "application/json")
	req := a.Request()
	req.Header.SetMethod(fiber.MethodPost)
	a.Body(data)
	req.SetRequestURI(fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port))
	if err := a.Parse(); err != nil {
		return fmt.Errorf("req issues: %w", err)
	}
	code, respBody, errs := a.Bytes()

	if len(errs) > 0 {
		return fmt.Errorf("errors: %s", errs)
	}
	if code == fiber.StatusCreated {
		return nil
	}
	defer func() {
		_ = req.CloseBodyStream()
	}()
	return fmt.Errorf("response error %s", respBody)
}

func (cmd *Command) getQ() error {
	agent := fiber.Get(fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name))
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return fmt.Errorf("request failed: %s", errs)
	}
	if statusCode != fiber.StatusOK {
		return fmt.Errorf("response error: %s", string(body))
	}
	defer func() {
		_ = agent.Request().CloseBodyStream()
	}()
	var response dto.GetAccountResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}
	fmt.Printf("Response account name: %s, amount: %d\n", response.Name, response.Amount)
	return nil
}

func (cmd *Command) patch() error {
	reqBody := dto.PatchAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	agent := fiber.AcquireAgent()
	agent.Add("Content-Type", "application/json")
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPatch)
	agent.Body(data)
	req.SetRequestURI(fmt.Sprintf("http://%s:%d/account/patch", cmd.Host, cmd.Port))
	if err := agent.Parse(); err != nil {
		return fmt.Errorf("req issues: %w", err)
	}
	code, respBody, errs := agent.Bytes()
	if len(errs) > 0 {
		return fmt.Errorf("errors: %s", errs)
	}
	if code == fiber.StatusOK {
		return nil
	}
	defer func() {
		_ = agent.Request().CloseBodyStream()
	}()
	return fmt.Errorf("response error %s", respBody)
}

func (cmd *Command) change() error {
	reqBody := dto.ChangeAccountRequest{
		Name:    cmd.Name,
		NewName: cmd.NewName,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	agent := fiber.AcquireAgent()
	agent.Add("Content-Type", "application/json")
	req := agent.Request()
	req.Header.SetMethod(fiber.MethodPost)
	agent.Body(data)
	req.SetRequestURI(fmt.Sprintf("http://%s:%d/account/change", cmd.Host, cmd.Port))
	if err := agent.Parse(); err != nil {
		return fmt.Errorf("req issues: %w", err)
	}
	code, respBody, errs := agent.Bytes()
	if len(errs) > 0 {
		return fmt.Errorf("errors: %s", errs)
	}
	if code == fiber.StatusOK {
		return nil
	}
	defer func() {
		_ = agent.Request().CloseBodyStream()
	}()
	return fmt.Errorf("response error %s", respBody)
}

func (cmd *Command) del() error {
	reqBody := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}
	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	a := fiber.AcquireAgent()
	a.Add("Content-Type", "application/json")
	req := a.Request()
	req.Header.SetMethod(fiber.MethodDelete)
	a.Body(data)
	req.SetRequestURI(fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port))
	if err := a.Parse(); err != nil {
		return fmt.Errorf("req issues: %w", err)
	}
	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return fmt.Errorf("errors: %s", errs)
	}
	if statusCode == fiber.StatusNoContent {
		return nil
	}
	defer func() {
		_ = req.CloseBodyStream()
	}()
	return fmt.Errorf("response error: %s", body)
}
