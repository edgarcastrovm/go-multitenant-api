## Inicializa el modulo principal

```sh
go mod init my-app-tx
```

##  Installa en adminitrador de rutas

```sh
# mux http
go get -u github.com/gorilla/mux
# Gin framework
go get -u github.com/gin-gonic/gin
```

##  Ejecuta el programa

```
go run main.go 
```

## Integramos con postgres

##  Depencdencias db para postgres

```
go get github.com/lib/pq
```

```sql
CREATE TABLE transacciones (
    id INTEGER PRIMARY KEY,
    cuenta VARCHAR(20) NOT NULL,
    cuenta_destino VARCHAR(20) NOT NULL,
    monto DECIMAL(15,2) NOT NULL CHECK (monto >= 0),
    tipo VARCHAR(10) NOT NULL CHECK (tipo IN ('INGRESO', 'EGRESO')),
    fecha TIMESTAMP WITH TIME ZONE NOT NULL,
    descripcion VARCHAR(100) NOT NULL,
    estado INTEGER NOT NULL CHECK (estado IN (0, 1, 2)),
    empresa VARCHAR(25) NOT NULL
);

```

```
go get github.com/joho/godotenv
```

```properties
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=finanzas
DB_SSLMODE=disable
```

## Generador codigos únicos

```
go get github.com/google/uuid
```

## Log

```
go get github.com/sirupsen/logrus
```

## Validador

```
go get github.com/go-ozzo/ozzo-validation
go get github.com/go-ozzo/ozzo-validation/is
```
