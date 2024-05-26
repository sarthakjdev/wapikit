package main

import (
	"database/sql"
	"log/slog"
	"os"

	koanf "github.com/knadh/koanf/v2"
	"github.com/sarthakjdev/wapikit/database"
)

// because this will be a single binary, we will be providing the flags here
// 1. --install to install the setup the app, but it will be idempotent
// 2. --migrate to apply the migration to the database
// 3. --config to setup the config files
// 4. --version to check the version of the application running
// 5. --help to check the

var (
	// Global variables
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	koa    = koanf.New(".")
)

type App struct {
	db     *sql.DB
	logger slog.Logger
	koa    *koanf.Koanf
}

func init() {
	initFlags()
	loadConfigFiles(koa.Strings("config"), koa)

	if koa.Bool("version") {

	} else if koa.Bool("install") {

		// ! should be idempotent

		// setup the database
		// 1. run database migration
		// 2. setup the config files
		// 3. install the app
		// 4, setup the filesystem required

	} else if koa.Bool("upgrade") {

		// ! should not upgrade without asking for thr permission, because database migration can be destructive
		// upgrade handler
	} else {
		// do nothing
		// ** NOTE: if no flag is provided, then let the app move to the main function and start the server
	}

}

func main() {

	logger.Info("Starting the application")

	app := &App{
		logger: *logger,
		db:     database.GetDbInstance(),
		koa:    koa,
	}

	initHTTPServer(app)
}
