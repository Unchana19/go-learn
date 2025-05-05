package main

import (
	"log"
	"time"

	"github.com/Unchana19/go-learn/internal/auth"
	"github.com/Unchana19/go-learn/internal/db"
	"github.com/Unchana19/go-learn/internal/env"
	"github.com/Unchana19/go-learn/internal/mailer"
	"github.com/Unchana19/go-learn/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const version = "1.0.0"

//	@title			Social Media API
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := config{
		addr:        env.GetString("ADDR", ":9000"),
		apiURL:      env.GetString("EXTERNAL_URL", "localhost:9000"),
		frontendURL: env.GetString("FRONTEND_URL", "http://localhost:3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://postgres:password@localhost:5432/golearn?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 10),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 5),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "10s"),
		},
		env: env.GetString("ENV", "development"),
		mail: mailConfig{
			exp:       time.Hour * 24 * 3,
			fromEmail: env.GetString("SENDGRID_FROM_EMAIL", ""),
			sendGrid: sendGridConfig{
				apiKey: env.GetString("SENDGRID_API_KEY", ""),
			},
			mailTrap: mailTrapConfig{
				apiKey: env.GetString("MAILTRAP_API_KEY", ""),
			},
		},
		auth: authConfig{
			basic: basicAuthConfig{
				user: env.GetString("BASIC_AUTH_USER", "admin"),
				pass: env.GetString("BASIC_AUTH_PASS", "admin"),
			},
			token: tokenConfig{
				secret:   env.GetString("JWT_SECRET", "secret"),
				exp:      time.Hour * 24 * 3,
				issuer:   env.GetString("JWT_ISSUER", "http://localhost:9000"),
				audience: env.GetString("JWT_AUDIENCE", "http://localhost:9000"),
			},
		},
	}

	// Logger
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	// Database
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Panic(err)
	}

	defer db.Close()
	logger.Info("Connected to database: %s\n", cfg.db.addr)

	// Store
	store := store.NewStorage(db)

	mailer := mailer.NewSendGridMailer(cfg.mail.fromEmail, cfg.mail.sendGrid.apiKey)

	jwtAuthenticator := auth.NewJWTAuthenticator(cfg.auth.token.secret, cfg.auth.token.audience, cfg.auth.token.issuer)

	app := &application{
		config:        cfg,
		store:         store,
		logger:        logger,
		mailer:        mailer,
		authenticator: jwtAuthenticator,
	}

	mux := app.mount()

	logger.Fatal(app.run(mux))
}
