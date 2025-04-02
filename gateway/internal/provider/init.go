package provider

import (
	"github.com/hailsayan/sophocles/pkg/middleware"
	"github.com/hailsayan/sophocles/pkg/utils/jwtutils"
)

var (
	jwtUtil        jwtutils.JwtUtil
	authMiddleware *middleware.AuthMiddleware
)

func BootstrapGlobal() {
	jwtUtil = jwtutils.NewJwtUtil()
	authMiddleware = middleware.NewAuthMiddleware(jwtUtil)
}
