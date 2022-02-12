package main

import (
	"net/http"
	"os"

	"github.com/adipatiarya/apis/api"
	"github.com/gin-gonic/gin"
)

type RequestOtpParam struct {
	Msisdn string `json:"msisdn" binding:"required"`
}
type LoginOtpParam struct {
	Msisdn  string `json:"msisdn" binding:"required"`
	OtpCode string `json:"otp_code" binding:"required"`
}
type PackageParam struct {
	SeviceId string `json:"pkgid" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

func main() {

	port := os.Getenv("PORT")
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/otp", func(c *gin.Context) {
		var json RequestOtpParam
		if err := c.ShouldBindJSON(&json); err == nil {
			otp, err := api.RequestOtp(json.Msisdn)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
			c.JSON(http.StatusOK, gin.H{"code": otp.StatusCode})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	router.POST("/loginotp", func(c *gin.Context) {

		var json LoginOtpParam
		if err := c.ShouldBindJSON(&json); err == nil {
			login, err := api.LoginOtp(json.Msisdn, json.OtpCode)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			c.JSON(http.StatusOK, gin.H{
				"status":        login.Status,
				"status_code":   login.StatusCode,
				"error_message": login.ErrorMessage,
				"token":         login.Data,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.POST("/package/buy", func(c *gin.Context) {
		var json PackageParam

		if err := c.ShouldBindJSON(&json); err == nil {
			resp, err := api.BuyPackage(json.Token, json.SeviceId)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
	})
	router.POST("/package/claim", func(c *gin.Context) {
		var json PackageParam

		if err := c.ShouldBindJSON(&json); err == nil {
			resp, err := api.ClaimPackage(json.Token, "CLAIM", json.SeviceId)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
	})
	router.POST("/package/card", func(c *gin.Context) {
		var json PackageParam

		if err := c.ShouldBindJSON(&json); err == nil {
			resp, err := api.ClaimPackage(json.Token, "CARD", json.SeviceId)
			if err != nil {
				return
			}
			c.JSON(http.StatusOK, resp)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}
	})

	router.Run(":" + port)
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
