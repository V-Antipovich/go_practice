package main

import (
	"flag"
	"fmt"
	pb "go_prac/grpc_app/accounts"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Command struct {
	Cmd     string
	Name    string
	Amount  int
	NewName string
}

func main() {
	cmdVal := flag.String("cmd", "", "command")
	nameVal := flag.String("name", "", "name of the account")
	ammountVal := flag.Int("amount", 0, "ammount of the account")
	newNameVal := flag.String("newname", "", "new name of the account")
	flag.Parse()
	cmd := Command{
		Cmd:     *cmdVal,
		Name:    *nameVal,
		Amount:  *ammountVal,
		NewName: *newNameVal,
	}
	conn, err := grpc.NewClient("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Could not connect to the server: %v", err)
	}
	defer conn.Close()
	c := pb.NewBankClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	switch cmd.Cmd {
	case "create":
		r, err := c.CreateAccount(ctx, &pb.Account{
			Name:   cmd.Name,
			Amount: int64(cmd.Amount)})
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("Created account", r.GetName())
		}
	case "get":
		r, err := c.GetAccount(ctx, &pb.Name{
			Name: cmd.Name})
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("Account:", r.Name, r.Amount)
		}
	case "change":
		r, err := c.UpdateAccount(ctx, &pb.ChangeAccount{
			Name: cmd.Name, Newname: cmd.NewName})
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("New name set for the account:", r.Name)
		}
	case "patch":
		r, err := c.PatchAccount(ctx, &pb.Account{
			Name: cmd.Name, Amount: int64(cmd.Amount)})
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("Changed amount of account", r.Name)
		}
	case "delete":
		r, err := c.DeleteAccount(ctx, &pb.Name{Name: cmd.Name})
		if err != nil {
			fmt.Printf("Errors: %v\n", err)
		} else {
			fmt.Println("Deleted account", r.Name)
		}
	default:
		fmt.Printf("unknown command: %s", cmd.Cmd)
	}
}
