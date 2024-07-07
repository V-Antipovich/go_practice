package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go_prac/accounts/dto"
	"io"
	"net/http"
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
	default:
		return fmt.Errorf("unknown command: %s", c.Cmd)
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
	if err != nil {
		return fmt.Errorf("http post failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	// check errors, status codes and io read body
	if resp.StatusCode == http.StatusAccepted {
		return nil
	}

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body failed: %w", err)
	}
	return nil
	// return fmt.Sprintf("response error %s", string(body))
	// return fmt.Errorf("response error %s", string(body))
}

func (cmd *Command) getQ() error {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/account?name=%s", cmd.Host, cmd.Port, cmd.Name))
	if err != nil {
		return fmt.Errorf("get request failed: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
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
