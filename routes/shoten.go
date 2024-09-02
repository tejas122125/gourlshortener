package routes

import (
	"time"
	"urlshortener/helpers"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string `json : "url"`
	CustomShort string `json : "short"`
	Expiry      time.Duration `json : "expiry"`
}

type response struct {
	URL string `json : url`
	CustomShort string `json : "customshort"`
	Expiry time.Duration `json : "expiry"`
	XRateRemaining int `jsom : "xrateremaining"`
	XRateLimitReset time.Duration `json:"xratelimitreset"`
}

func ShortenUrl( c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(&body); err!=nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "cannot parse the request body"})
	}
	// implementing rate  limimting

	// checking url is real or not

	if !govalidator.IsUrl(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "Invalid URL"})
	}
// cannot give same domain name as url shortener
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error" : "cannot process the same domain give another domain"})
	}
	//  enforce https and SSL

	body.URL =  helpers.EnforceHTTP(body.URL)
}