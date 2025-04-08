package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func NewDatabaseConnection() (*sql.DB, error) {
    config := DatabaseConfig{
        Host:     getEnvWithDefault("DB_HOST", "localhost"),
        Port:     getEnvWithDefault("DB_PORT", "5432"),
        User:     getEnvWithDefault("DB_USER", "mvp_user"),  // Changed default to mvp_user
        Password: getEnvWithDefault("DB_PASSWORD", "352535"),
        Name:     getEnvWithDefault("DB_NAME", "mvp_db"),
    }

	if config.Password == "" {
		return nil, fmt.Errorf("database password is required")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}

func main() {
	fmt.Println("Starting MultyLink API server...")

	// Initialize database connection
	db, err := NewDatabaseConnection()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	defer db.Close()

	// Test connection with retries
	if err := retryDBConnection(db); err != nil {
		log.Fatalf("Failed to connect to database after retries: %v", err)
	}
	log.Println("Successfully connected to database")

	// TODO: Initialize router
	// TODO: Initialize services

	// Simple health check endpoint for now
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Add root route handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("MultyLink API is running"))
	})

	// Add register route handler
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	    if r.Method != "POST" {
	        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	        return
	    }
	
	    // TODO: Add actual registration logic
	    w.Write([]byte(`{"status":"registration endpoint works"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutting down...")

	// TODO: Close database connection and other resources

	log.Println("Server exited properly")
}

func retryDBConnection(db *sql.DB) error {
	var dbErr error
	for i := 0; i < 5; i++ {
		dbErr = db.Ping()
		if dbErr == nil {
			return nil
		}
		log.Printf("Attempt %d: Database ping failed: %v", i+1, dbErr)
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("database connection failed after retries: %w", dbErr)
}

func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
