package repository

import (
	"errors"
	"log/slog"

	// Импортируем твой пакет с зашитыми файлами
	// Замени "твоя_папка_проекта" на имя из твоего go.mod
	"github.com/Alexeyts0Y/test_task_em/migrations"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func RunMigrations(dbURL string) error {
	d, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithSourceInstance("iofs", d, dbURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			slog.Info("No new migrations to apply")
			return nil
		}
		return err
	}

	slog.Info("Migrations applied successfully")
	return nil
}
