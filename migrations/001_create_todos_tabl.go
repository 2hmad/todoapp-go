package migrations

import (
	"context"
	"github.com/todoapp-pulse/config"
)

func M001CreateTodosTable(ctx context.Context) {
	db, err := config.Connect()

	// Drop the table if it already exists
	_, err = db.ExecContext(ctx, "DROP TABLE IF EXISTS todos")
	if err != nil {
		panic(err)
	}

	// Run the migration
	_, err = db.ExecContext(ctx, `
        CREATE TABLE todos (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title TEXT NOT NULL,
            done BOOLEAN DEFAULT false,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        );
    `)
	if err != nil {
		panic(err)
	}

}
