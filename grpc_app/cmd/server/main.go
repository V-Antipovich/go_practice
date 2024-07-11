package main

import (
	"log"
	"net"
	"sync"

	// "go_prac/fiber_app/accounts"
	pb "go_prac/grpc_app/accounts"
	"go_prac/grpc_app/accounts/models"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	// "google.golang.org/grpc/profiling/proto"
	"google.golang.org/grpc/status"
)

// Он будет имплементировать интерфейс сервера из pb
type server struct {
	pb.UnimplementedBankServer
}

// mustEmbedUnimplementedBankServer implements accounts.BankServer.
// func (s *server) mustEmbedUnimplementedBankServer() {
// 	panic("unimplemented")
// }

// mustEmbedUnimplementedBankServer implements accounts.BankServer.
// func (s *server) mustEmbedUnimplementedBankServer() {
// 	panic("unimplemented")
// }

type LocalStorage struct {
	accounts map[string]*models.Account
	guard    *sync.RWMutex
}

// func New() *LocalStorage {
// 	return &LocalStorage{
// 		accounts: make(map[string]*models.Account),
// 		guard:    &sync.RWMutex{},
// 	}
// }

var db = LocalStorage{
	accounts: make(map[string]*models.Account),
	guard:    &sync.RWMutex{},
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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
	db.guard.Lock()
	if _, ok := db.accounts[ac.Name]; ok {
		db.guard.Unlock()
		return nil, status.Error(codes.FailedPrecondition, "Account is already present")
	}
	db.accounts[ac.Name] = &models.Account{
		Name:   ac.Name,
		Amount: int(ac.Amount),
	}
	db.guard.Unlock()
	return &pb.Name{Name: ac.Name}, nil
}

func (s *server) GetAccount(ctx context.Context, name *pb.Name) (*pb.Account, error) {
	if len(name.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	db.guard.RLock()
	acc, ok := db.accounts[name.Name]
	db.guard.RUnlock()
	if !ok {
		return nil, status.Error(codes.NotFound, "No such entry")
	}
	return &pb.Account{Name: acc.Name, Amount: int64(acc.Amount)}, nil
}

// If newname already exists, throw error
func (s *server) UpdateAccount(ctx context.Context, ac *pb.ChangeAccount) (*pb.Account, error) {
	if len(ac.Name) == 0 || len(ac.Newname) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Neither old nor new name can be empty")
	}
	db.guard.Lock()
	if _, ok := db.accounts[ac.Name]; !ok {
		db.guard.Unlock()
		return nil, status.Error(codes.NotFound, "No such entry")
	}
	amount := db.accounts[ac.Name].Amount
	delete(db.accounts, ac.Name)
	db.accounts[ac.Newname] = &models.Account{
		Name: ac.Newname, Amount: amount,
	}
	db.guard.Unlock()
	return &pb.Account{Name: ac.Newname, Amount: int64(amount)}, nil
}

func (s *server) PatchAccount(ctx context.Context, ac *pb.Account) (*pb.Name, error) {
	if len(ac.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	db.guard.Lock()
	if _, ok := db.accounts[ac.Name]; !ok {
		db.guard.Unlock()
		return nil, status.Error(codes.NotFound, "No such entry")
	}
	db.accounts[ac.Name].Amount = int(ac.Amount)
	db.guard.Unlock()
	return &pb.Name{Name: ac.Name}, nil
}

func (s *server) DeleteAccount(ctx context.Context, name *pb.Name) (*pb.Name, error) {
	if len(name.Name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Can't have empty name")
	}
	db.guard.Lock()
	if _, ok := db.accounts[name.Name]; !ok {
		db.guard.Unlock()
		return nil, status.Error(codes.NotFound, "No such entry")
	}
	delete(db.accounts, name.Name)
	db.guard.Unlock()
	return &pb.Name{Name: name.Name}, nil
}
