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

//	func (cmd *Command) create(ctx *context.Context) error {
//		r, err :=
//		return nil
//	}
//
//	func (cmd *Command) getQ(ctx *context.Context) error {
//		return nil
//	}
//
//	func (cmd *Command) patch(ctx *context.Context) error {
//		return nil
//	}
//
//	func (cmd *Command) change(ctx *context.Context) error {
//		return nil
//	}
//
//	func (cmd *Command) del(ctx *context.Context) error {
//		return nil
//	}
func main() {
	cmdVal := flag.String("cmd", "", "command")
	nameVal := flag.String("name", "", "name of the account")
	ammountVal := flag.Int("amount", 0, "ammount of the account")
	newNameVal := flag.String("newname", "", "new name of the account")
	flag.Parse()
	cmd := Command{
		// Port:    *portVal,
		// Host:    *hostVal,
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
	// if err := cmd.Do(ctx); err != nil {
	// 	panic(err)
	// }
	defer cancel()
	switch cmd.Cmd {
	case "create":
		// return c.create(ctx *context.Context)
		r, err := c.CreateAccount(ctx, &pb.Account{
			Name:   cmd.Name,
			Amount: int64(cmd.Amount)})
		if err != nil {
			fmt.Printf("Errors: %v", err)
		} else {
			fmt.Println(r.GetName())
		}
	case "get":
		r, err := c.GetAccount(ctx, &pb.Name{
			Name: cmd.Name})
		if err != nil {
			fmt.Printf("Errors: %v", err)
		} else {
			fmt.Println("Account:", r.Name, r.Amount)
		}
		// return c.getQ(ctx *context.Context)
	// case "delete":
	// 	return c.del(ctx *context.Context)
	// case "patch":
	// 	return c.patch(ctx *context.Context)
	// case "change":
	// 	return c.change(ctx *context.Context)
	default:
		fmt.Printf("unknown command: %s", cmd.Cmd)
	}
}
