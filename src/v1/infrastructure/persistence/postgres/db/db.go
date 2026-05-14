package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/EdwardMelendezM/api-info-shared/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Client and ClientV2 are kept for compatibility with existing call sites.
var (
	Client   *sql.DB
	ClientV2 *sql.DB
)

type Options struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	SSLMode         string
}

func InitClients(cfg config.Configuration) error {
	return InitClientsWithOptions(cfg, Options{})
}

func InitClientsWithOptions(cfg config.Configuration, opts Options) error {
	dsn := buildDSN(cfg.DB, opts.SSLMode)

	client, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	applyPoolOptions(client, opts)

	if err := client.Ping(); err != nil {
		_ = client.Close()
		return err
	}

	Client = client
	ClientV2 = client
	return nil
}

func buildDSN(dbCfg config.DB, sslMode string) string {
	mode := sslMode
	if mode == "" {
		mode = "disable"
	}
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbCfg.DbHost,
		dbCfg.DbPort,
		dbCfg.DbUsername,
		dbCfg.DbPassword,
		dbCfg.DbDatabase,
		mode,
	)
}

func applyPoolOptions(client *sql.DB, opts Options) {
	if client == nil {
		return
	}
	if opts.MaxOpenConns > 0 {
		client.SetMaxOpenConns(opts.MaxOpenConns)
	}
	if opts.MaxIdleConns > 0 {
		client.SetMaxIdleConns(opts.MaxIdleConns)
	}
	if opts.ConnMaxLifetime > 0 {
		client.SetConnMaxLifetime(opts.ConnMaxLifetime)
	}
	if opts.ConnMaxIdleTime > 0 {
		client.SetConnMaxIdleTime(opts.ConnMaxIdleTime)
	}
}
