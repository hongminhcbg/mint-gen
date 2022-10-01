package erp

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrGeneric error = status.Error(codes.Code(500000), "Internal server error")
