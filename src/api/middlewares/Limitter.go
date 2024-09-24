package middlewares

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"github.com/miladshalikar/golang-clean-web-api/src/api/helper"
	"golang.org/x/time/rate"
	"time"
)

//func LimitByRequest() fiber.Handler {
//	lmt := tollbooth.NewLimiter(1, nil)
//	return func(c fiber.Ctx) error {
//
//		req := &http.Request{
//			Method: c.Method(),
//			URL:    &url.URL{Path: c.Request().URI().String()},
//			Header: make(http.Header),
//		}
//
//		c.Request().Header.VisitAll(func(key, value []byte) {
//			req.Header.Add(string(key), string(value))
//		})
//
//		rw := httptest.NewRecorder()
//
//		err := tollbooth.LimitByRequest(lmt, rw, req)
//
//		if err != nil {
//			return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
//				"error": err.Error(),
//			})
//		} else {
//			return c.Next()
//		}
//	}
//}

func RateLimiter() fiber.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Second), 1)
	return func(c fiber.Ctx) error {
		if !limiter.Allow() {
			return c.Status(fiber.StatusTooManyRequests).JSON(
				helper.GenerateBaseResponseWithError(nil, false, -100, errors.New("Too Many Requests")))
		}
		return c.Next()
	}
}
