### install goose:
```
curl -fsSL https://raw.githubusercontent.com/pressly/goose/master/install.sh | GOOSE_INSTALL=$HOME/.goose sh

alias goose="$HOME/.goose/bin/goose"
```

### create migration:
`goose --dir migrations/ create MIGRATION-NAME sql`

### Up migrations
`goose --dir migrations/ postgres "postgresql://${postgres_user}:${postgres_password}@${postgres_host}:${postgres_port}/${postgres_db_name}?sslmode=${postgres_sslmode}" up`

### Down migrations
`goose --dir migrations/ postgres "postgresql://${postgres_user}:${postgres_password}@${postgres_host}:${postgres_port}/${postgres_db_name}?sslmode=${postgres_sslmode}` down

### Reset migrations
`goose --dir migrations/ postgres "postgresql://${postgres_user}:${postgres_password}@${postgres_host}:${postgres_port}/${postgres_db_name}?sslmode=${postgres_sslmode}` reset

### Status migrations
`goose --dir migrations/ postgres "postgresql://${postgres_user}:${postgres_password}@${postgres_host}:${postgres_port}/${postgres_db_name}?sslmode=${postgres_sslmode}` status