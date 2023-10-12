package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"os"

	logrpc "github.com/izaakdale/service-logrpc/api/v1"
	"github.com/izaakdale/service-logrpc/internal/dao"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func NewServer(client dao.LogStore) *Server {
	return &Server{client, logrpc.UnimplementedLoggingServiceServer{}}
}

func Run() {
	crt, err := tls.LoadX509KeyPair(os.Getenv("SERVER_CRT"), os.Getenv("SERVER_KEY"))
	if err != nil {
		panic(err)
	}
	gsrv := grpc.NewServer(grpc.Creds(credentials.NewServerTLSFromCert(&crt)))
	reflection.Register(gsrv)

	connectionString := os.Getenv("MONGO_CONN")
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	srv := NewServer(dao.NewMongoConn(client))
	logrpc.RegisterLoggingServiceServer(gsrv, srv)
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", os.Getenv("GRPC_HOST"), os.Getenv("GRPC_PORT")))
	if err != nil {
		panic(err)
	}

	gsrv.Serve(l)
}
