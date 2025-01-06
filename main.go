package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var db *gorm.DB
var jwtKey = []byte("your_secret_key")

// Пользовательская структура
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

// Структура для тикетов
type Ticket struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

// JWT-структура для токена
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func initDatabase() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=helpdesk port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Миграция структуры
	err = db.AutoMigrate(&User{}, &Ticket{})
	if err != nil {
		panic("Failed to migrate database")
	}
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func authenticateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing token"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func main() {
	initDatabase()
	router := gin.Default()

	// Регистрация пользователя
	router.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	})

	// Логин пользователя
	router.POST("/login", func(c *gin.Context) {
		var user User
		var foundUser User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Where("username = ?", user.Username).First(&foundUser).Error; err != nil || foundUser.Password != user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
			return
		}
		token, err := generateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not generate token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Пример защищенного маршрута
	protected := router.Group("/tickets")
	protected.Use(authenticateMiddleware())
	{
		protected.GET("/", func(c *gin.Context) {
			var tickets []Ticket
			db.Find(&tickets)
			c.JSON(http.StatusOK, tickets)
		})
	}

	router.Run(":8080")
}
