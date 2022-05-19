package runner

import (
	"cinema-search/internal/app"
	"cinema-search/internal/app/query"
	"cinema-search/internal/config"
	"cinema-search/internal/dao"
	"cinema-search/internal/port/http"
	"cinema-search/internal/server"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func Start(configDir string) {
	cfg := newConfig(configDir)
	db := initDB(cfg)
	application := newApplication(db)
	startServer(cfg, application)
}

func newConfig(configDir string) *config.Config {
	cfg, err := config.New(configDir)
	if err != nil {
		log.Panicln(err)
	}

	return cfg
}

func initDB(cfg *config.Config) *sql.DB {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Name)
	db, err := sql.Open("postgres", dbInfo)
	//conn, er := pgx.Connect(context.Background(), dbInfo)
	//defer conn.Close(context.Background())
	//if conn != nil && er == nil {
	//	fmt.Println("wew")
	//}

	if err != nil {
		fmt.Println("DB err")
	}
	//defer db.Close()
	return db
}

func newApplication(db *sql.DB) app.Application {
	movieRepository := dao.NewMovieRepository(db)

	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetMovie: query.NewGetMovieHandler(movieRepository),
		},
	}
}

func startServer(cfg *config.Config, application app.Application) {
	logrus.Info(fmt.Sprintf("Starting HTTP server on address: %s", cfg.HTTP.Port))

	httpServer := server.New(cfg, http.NewHandler(application))

	err := httpServer.Run()

	logrus.WithError(err).Fatal("HTTP server stopped")
}
