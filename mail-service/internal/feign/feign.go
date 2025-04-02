package feign

import (
	"net/http"
	"time"

	"github.com/hailsayan/sophocles/pkg/utils/jwtutils"
)

var (
	feignClient = &http.Client{
		Timeout: 10 * time.Second,
	}
	jwtUtil = jwtutils.NewJwtUtil()
)

const (
	APPLICATION_JSON = "application/json"
)
