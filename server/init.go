package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arya-bhanu/intensive-go/service/product"
	"github.com/arya-bhanu/intensive-go/service/user"
	"github.com/arya-bhanu/intensive-go/utils"
)

// gRPC service registration helpers
func RegisterUserServiceServer(s *grpc.Server, srv *user.UserServiceServer) {
	// In a real implementation, this would be generated code
	// For this challenge, we'll manually handle the registration
}

func RegisterProductServiceServer(s *grpc.Server, srv *product.ProductServiceServer) {
	// In a real implementation, this would be generated code
	// For this challenge, we'll manually handle the registration
}

func StartUserService(port string) (*grpc.Server, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(utils.LoggingInterceptor))
	userServer := user.NewUserServiceServer()

	// Register HTTP handlers for gRPC methods
	mux := http.NewServeMux()
	mux.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("id")
		userID, _ := strconv.ParseInt(userIDStr, 10, 64)

		user, err := userServer.GetUser(r.Context(), userID)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	mux.HandleFunc("/user/validate", func(w http.ResponseWriter, r *http.Request) {
		userIDStr := r.URL.Query().Get("id")
		userID, _ := strconv.ParseInt(userIDStr, 10, 64)

		valid, err := userServer.ValidateUser(r.Context(), userID)
		if err != nil {
			if status.Code(err) == codes.NotFound {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]bool{"valid": valid})
	})

	go func() {
		log.Printf("User service HTTP server listening on %s", port)
		if err := http.Serve(lis, mux); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	return s, nil
}

// StartProductService starts the product service on the given port
func StartProductService(port string) (*grpc.Server, error) {
	// TODO: Implement this function
	// Hint: create listener, gRPC server with interceptor, register service, serve
	lis, err := net.Listen("tcp", port)
	var opts []grpc.ServerOption
	if err != nil {
		return nil, err
	}

	opts = append(opts, grpc.UnaryInterceptor(utils.LoggingInterceptor))

	grpcServer := grpc.NewServer(opts...)
	srv := product.NewProductServiceServer()

	RegisterProductServiceServer(grpcServer, srv)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Printf("grpc server error: %v", err)
		}
	}()

	return grpcServer, nil
}
