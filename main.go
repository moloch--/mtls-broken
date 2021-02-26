/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ./helloworld --go_out=plugins=grpc:./helloworld ./helloworld/helloworld.proto

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	pb "github.com/moloch--/mtls-broken/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {

	tlsConfig := getTLSConfig()
	creds := credentials.NewTLS(tlsConfig)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterGreeterServer(s, &server{})

	fmt.Printf("Starting helloworld server ...\n")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getTLSConfig() *tls.Config {

	caCertPEM, _, _ := getCertificateAuthorityPEM()

	certBlock, _ := pem.Decode(caCertPEM)
	caCertPtr, _ := x509.ParseCertificate(certBlock.Bytes)

	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(caCertPtr)

	certPEM, keyPEM, _ := getLeafCertificatePEM()
	leafCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		fmt.Printf("Error loading server certificate: %v", err)
	}

	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
		Certificates: []tls.Certificate{leafCert},
		// PreferServerCipherSuites: true,
		// MinVersion:               tls.VersionTLS13,
	}
	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}

func getCertificateAuthorityPEM() ([]byte, []byte, error) {

	caCertPath := "ca-cert.pem"
	caKeyPath := "ca-key.pem"

	certPEM, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	keyPEM, err := ioutil.ReadFile(caKeyPath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return certPEM, keyPEM, nil
}

func getLeafCertificatePEM() ([]byte, []byte, error) {

	serverCertPath := "server-cert.pem"
	serverKeyPath := "server-key.pem"

	certPEM, err := ioutil.ReadFile(serverCertPath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	keyPEM, err := ioutil.ReadFile(serverKeyPath)
	if err != nil {
		log.Fatalf("%s", err)
	}

	return certPEM, keyPEM, nil
}
