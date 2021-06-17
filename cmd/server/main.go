package main

//go:generate protoc -I ../../api ../../api/karlchen.proto --go_out=plugins=grpc:../../api

import (
	crand "crypto/rand"
	"encoding/binary"
	"github.com/spf13/cobra"
	"github.com/supermihi/karlchencloud/server"
	"github.com/supermihi/karlchencloud/server/tables"
	u "github.com/supermihi/karlchencloud/server/users"
	"log"
	"math/rand"
	"net"
	"strconv"
)

var (
	port  int
	noWeb bool
	seed  int64
)
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "karlchencloud server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var v int64
		err := binary.Read(crand.Reader, binary.BigEndian, &v)
		if err != nil {
			log.Fatal(err)
		}
		rand.Seed(v)
		users, err := u.NewSqlUserDatabase("users.sqlite")
		if err != nil {
			log.Fatalf("error creating users: %v", err)
		}
		config := server.ServerConfig{Tables: server.TablesConfig{InputSeed: seed}}
		srv := server.CreateServer(users, tables.NewTables(), config)

		if noWeb {
			log.Printf("starting bare grpc server")
			lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			if err := srv.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		} else {
			httpServer := server.WrapServer(srv)
			lisHttp, err := net.Listen("tcp", ":"+strconv.Itoa(port))
			if err != nil {
				log.Fatalf("failed to listen HTTP: %v", err)
			}
			log.Printf("starting HTTP proxy server")
			if err := httpServer.Serve(lisHttp); err != nil {
				log.Fatalf("failed to serve HTTP: %v", err)
			}
		}
	}}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolVar(&noWeb, "no-web", false, "don't enable grpc-web wrapper")
	rootCmd.Flags().IntVarP(&port, "port", "p", 50501, "gRPC server port")
	rootCmd.Flags().Int64Var(&seed, "seed", 0, "random seed")
}
