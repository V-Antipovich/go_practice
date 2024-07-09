package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go_prac/accounts/dto"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/client"
)

type Command struct {
	Port   int
	Host   string
	Cmd    string
	Name   string
	Amount int
}

func (c *Command) Do() error {
	switch c.Cmd {
	case "create":
		return c.create()
	case "get":
		return c.getQ()
	case "delete":
		return c.del()
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
	flag.Parse()
	cmd := Command{
		Port:   *portVal,
		Host:   *hostVal,
		Cmd:    *cmdVal,
		Name:   *nameVal,
		Amount: *ammountVal,
	}
	if err := cmd.Do(); err != nil {
		panic(err)
	}
}

func (cmd *Command) create() error {
	req := dto.CreateAccountRequest{
		Name:   cmd.Name,
		Amount: cmd.Amount,
	}
	// json marshall
	data, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	// make post request
	resp, err := http.Post(
		fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
		"application/json",
		bytes.NewReader(data),
	)
	// fmt.Println(fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
	// 	"application/json",
	// 	bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	// check errors, status codes and io read body
	if resp.StatusCode == fiber.StatusCreated {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	// return nil
	// return fmt.Sprintf("response error %s", string(body))
	return fmt.Errorf("response error %s", body)
}

func (cmd *Command) getQ() error {
	// TODO: use fiber methods
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name))
	if err != nil {
		return fmt.Errorf("get request failed: %w", err)
	}
	fmt.Println(resp.Header)
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != fiber.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("read body failed: %w", err)
		}
		return fmt.Errorf("response error: %s", string(body))
	}
	var response dto.GetAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("json decode failed: %w", err)
	}
	fmt.Printf("Response account name: %s, amount: %d\n", response.Name, response.Amount)
	return nil
}

func (cmd *Command) del() error {
	c := &http.Client{}
	reqBody := dto.DeleteAccountRequest{
		Name: cmd.Name,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("json marshall failed: %w", err)
	}
	req, err := http.NewRequest("DELETE",
		fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port),
		bytes.NewReader(data),
	)
	if err != nil {
		fmt.Println("if err != nil ONE")
		return fmt.Errorf("req issues: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("if err != nil TWO")
		return fmt.Errorf("resp issues: %w", err)
	}
	// c := fiber.Client{}

	// fmt.Println("got here")

	// check errors, status codes and io read body
	if resp.StatusCode == fiber.StatusNoContent {
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	// return nil
	// return fmt.Sprintf("response error %s", string(body))
	defer func() {
		_ = resp.Body.Close()
	}()
	return fmt.Errorf("response error %s", body)

	// resp, err := fiber
	// resp, err := http.
	// respBody := fmt.Sprintf("http://%s:%d/account/delete", cmd.Host, cmd.Port)
	// resp, err := http.Post(
	// 	fmt.Sprintf("http://%s:%d/account/create", cmd.Host, cmd.Port),
	// 	"application/json",
	// 	bytes.NewReader(data),
	// )
	// fmt.Println(respBody)

	// agent := fiber.Delete(fmt.Sprintf("http://%s:%d/account/delete/%s", cmd.Host, cmd.Port, cmd.Name))
	// statusCode, agentBody, errs := agent.Bytes()
	// if len(errs) > 0 {
	// 	return fmt.Errorf("errors: %s", errs)
	// }
	// if statusCode == fiber.StatusNoContent {
	// 	return nil
	// }
	// return fmt.Errorf("response error: %s", string(agentBody))
}
