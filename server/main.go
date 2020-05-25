package main

import (
	"context"
	"fmt"
	"helloWorld/proto"
	"net"

	//will be used later
	"database/sql"
	_ "database/sql"

	//will be used later
	_ "fmt"

	// import needed for postgrtes
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Printf("Failed to listen to port 4040")
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterCRUDSreviceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		fmt.Printf("Failed to serve on listner")
		panic(e)
	}
}

//create emp
func (s *server) CreateEmp(ctx context.Context, create *proto.Create) (*proto.BoolResult, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("in error")
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	test, err := db.Prepare("Insert into COMPANY Values ($1,$2,$3,$4,$5);")
	if err != nil {
		fmt.Printf("/n Error fetching row /n")
		panic(err)
	}
	_, e := test.Exec(create.GetId(), create.GetName(), create.GetAge(), create.GetAddress(), create.GetAge())
	if e != nil {
		return &proto.BoolResult{Done: false}, e
	}
	return &proto.BoolResult{Done: true}, nil
}

// Implement after confirming client side for create
func (s *server) RetrieveEmp(ctx context.Context, retrieve *proto.Retrieve) (*proto.Create, error) {
	// implementation pending
	return &proto.Create{Id: 1, Name: "TEST", Address: "t", Age: 12, Salary: 12}, nil
}

func (s *server) UpdateEmp(ctx context.Context, update *proto.Create) (*proto.BoolResult, error) {
	// implementation pending
	return &proto.BoolResult{Done: false}, nil
}

func (s *server) DeleteEmp(ctx context.Context, delete *proto.Retrieve) (*proto.BoolResult, error) {
	// implementation pending
	return &proto.BoolResult{Done: false}, nil
}

const (
	host     = "localhost"
	port     = 5432
	user     = "dbuser"
	password = "password"
	dbname   = "test"
)

var (
	id      int
	name    string
	age     int
	address string
	salary  int
)
