package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:5432/deliverydb?sslmode=disable"
	}

	var err error
	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("Успешное подключение к базе данных!")

	DB.MustExec(`
		-- Создание таблицы курьеров
		CREATE TABLE couriers (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			surname VARCHAR(100) NOT NULL,
			transport VARCHAR(20) NOT NULL CHECK (transport IN ('on_foot', 'scooter', 'car')),
			email VARCHAR(255) UNIQUE NOT NULL,
			phone VARCHAR(20) UNIQUE NOT NULL,
			rating FLOAT DEFAULT 0.0 CHECK (rating >= 0 AND rating <= 5),
			current_order_id INTEGER,
			is_available BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			password TEXT NOT NULL,
			
			CONSTRAINT fk_current_order 
				FOREIGN KEY (current_order_id) 
				REFERENCES orders(id) 
				ON DELETE SET NULL
		);
		
		-- Создание таблицы заказов
		CREATE TABLE orders (
			id SERIAL PRIMARY KEY,
			courier_id INTEGER,
			status VARCHAR(20) NOT NULL 
				DEFAULT 'created' 
				CHECK (status IN ('created', 'in_transit', 'delivered', 'canceled')),
			receiver_name VARCHAR(100) NOT NULL,
			receiver_phone VARCHAR(20) NOT NULL,
			address_from TEXT NOT NULL,
			address_to TEXT NOT NULL,
			cost INTEGER NOT NULL CHECK (cost > 0),
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			assigned_at TIMESTAMP WITH TIME ZONE,
			finished_at TIMESTAMP WITH TIME ZONE,
			
			CONSTRAINT fk_courier 
				FOREIGN KEY (courier_id) 
				REFERENCES couriers(id) 
				ON DELETE SET NULL
		);
		
		-- Создание индексов для ускорения запросов
		CREATE INDEX idx_couriers_availability ON couriers(is_available);
		CREATE INDEX idx_orders_status ON orders(status);
		CREATE INDEX idx_orders_courier ON orders(courier_id);
		CREATE INDEX idx_orders_created ON orders(created_at);`)

	return nil
}
