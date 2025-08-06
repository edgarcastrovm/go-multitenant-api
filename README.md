-- inicializa el modulo principal
go mod init my-app-tx

-- Intalla en adminitrador de rutas
go get -u github.com/gorilla/mux

-- Ejecuta el programa
go run main.go 

##Integramos con postgres

-- Depencdencias db para postgres
go get github.com/lib/pq

go get github.com/joho/godotenv

```sql
CREATE TABLE transacciones (
    id SERIAL PRIMARY KEY,
    monto DOUBLE PRECISION NOT NULL,
    tipo VARCHAR(10) NOT NULL CHECK (tipo IN ('ingreso', 'gasto')),
    fecha TIMESTAMP NOT NULL,
    descripcion TEXT NOT NULL
);
```

