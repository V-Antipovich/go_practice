package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	pb "go_prac/grpc_app/accounts"

	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Он будет имплементировать интерфейс сервера из pb
type server struct {
	pb.UnimplementedBankServer
}

var DB *sql.DB

const (
	connectString = "host=0.0.0.0 port=5432 dbname=godb user=vitalii password=1234 sslmode=disable"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	DB, err = sql.Open("postgres", connectString)
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
	defer DB.Close()
	if err := DB.Ping(); err != nil {
		log.Fatalf("Could not access db: %v", err)
	}
	s := server{}
	grpcServer := grpc.NewServer()
	pb.RegisterBankServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (s *server) CreateAccount(ctx context.Context, ac *pb.Account) (*pb.Name, error) {
	if len(ac.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	_, err := DB.ExecContext(ctx, "INSERT INTO accounts(name, amount) VALUES($1, $2);", ac.Name, ac.Amount)
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Problems while creating a new user: %v", err))
	}
	return &pb.Name{Name: ac.Name}, nil
}

func (s *server) GetAccount(ctx context.Context, name *pb.Name) (*pb.Account, error) {
	if len(name.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	accName := ""
	accAmount := 0
	res := DB.QueryRowContext(ctx, "SELECT name, amount FROM accounts WHERE name = $1;", name.Name)
	err1 := res.Scan(&accName, &accAmount)
	if err1 == sql.ErrNoRows {
		return nil, status.Error(codes.NotFound, "No such entry")
	} else if err1 != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Error: %v", err1))
	}
	return &pb.Account{Name: accName, Amount: int64(accAmount)}, nil
}

func (s *server) UpdateAccount(ctx context.Context, ac *pb.ChangeAccount) (*pb.Account, error) {
	if len(ac.Name) == 0 || len(ac.Newname) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Neither old nor new name can be empty")
	}
	_, err := DB.ExecContext(ctx, "UPDATE accounts SET name=$1 WHERE name=$2", ac.Newname, ac.Name)
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not update: %v", err))
	}
	return &pb.Account{Name: ac.Newname}, nil
}

func (s *server) PatchAccount(ctx context.Context, ac *pb.Account) (*pb.Name, error) {
	if len(ac.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	_, err := DB.ExecContext(ctx, "UPDATE accounts SET amount=$1 WHERE name=$2", ac.Amount, ac.Name)
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not patch account: %v", err))
	}
	return &pb.Name{Name: ac.Name}, nil
}

func (s *server) DeleteAccount(ctx context.Context, name *pb.Name) (*pb.Name, error) {
	if len(name.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	rb, err := DB.Query("SELECT COUNT(*) FROM accounts")
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not delete account: %v", err))
	}
	defer rb.Close()
	before := 0
	for rb.Next() {
		if err := rb.Scan(&before); err != nil {
			return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not delete account: %v", err))
		}
	}
	_, errD := DB.ExecContext(ctx, "DELETE FROM accounts WHERE name=$1", name.Name)

	ra, err := DB.Query("SELECT COUNT(*) FROM accounts")
	if err != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not delete account: %v", err))
	}
	defer ra.Close()
	after := 0
	for ra.Next() {
		if err := ra.Scan(&after); err != nil {
			return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not delete account: %v", err))
		}
	}
	log.Println(before, after)
	if after == before {
		return nil, status.Error(codes.NotFound, "No such entry")
	} else if errD != nil {
		return nil, status.Error(codes.Canceled, fmt.Sprintf("Could not delete account: %v", err))
	}
	return &pb.Name{Name: name.Name}, nil
}
