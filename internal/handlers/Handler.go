package handlers

import "github.com/jackc/pgx/v5/pgxpool"

type Handler struct {
	DB *pgxpool.Pool
}
