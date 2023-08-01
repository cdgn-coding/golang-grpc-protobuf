package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
)

func NewRecoveryInterceptor() grpc.UnaryServerInterceptor {
	return recovery.UnaryServerInterceptor()
}
