# API go-multitenant-api 

API en **Go** para la gestión de transacciones financieras multi-empresa (**multitenant**), implementando el **patrón Strategy** en los servicios.

La aplicación permite definir estrategias de transacciones según el banco (**Banco A**, **Banco B**, etc.), soportando múltiples controladores HTTP (Gorilla Mux y Gin).

> ⚠️ **Nota:** Este proyecto tiene un propósito **ilustrativo y educativo**. Cualquier sugerencia o mejora es bienvenida.

---

## 🚀 Características

* Arquitectura **multitenant**: soporta múltiples empresas.
* Implementa el **patrón Strategy** en los servicios (`service/`) para desacoplar la lógica de negocio por banco.
* Soporta **Gorilla Mux** y **Gin** como frameworks HTTP.
    * `go get -u github.com/gorilla/mux`
    * `go get -u github.com/gin-gonic/gin`
* Conexión a **PostgreSQL** para persistencia de datos.
    * `go get github.com/lib/pq`
* Manejo de logs con **Zap**.
    * `go get -u go.uber.org/zap`
* Variables de entorno con **godotenv**.
    * `go get github.com/joho/godotenv`
* Validación de modelos con **ozzo-validation**.
    * `go get github.com/go-ozzo/ozzo-validation`
    * `go get github.com/go-ozzo/ozzo-validation/is`
* Generación de identificadores únicos con **UUID**.
    * `go get github.com/google/uuid`

---

## 📦 Instalación

### 1. Clonar el repositorio

```sh
git clone git@github.com:edgarcastrovm/go-multitenant-api.git
cd go-multitenant-api
```

### 2. Instalar dependencias

```sh
go mod tidy
```

---

## 🗄️ Base de datos (PostgreSQL)

Ejecutar el siguiente script SQL para crear la tabla principal de transacciones:

```sql
CREATE TABLE transacciones (
    id BIGSERIAL PRIMARY KEY,
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

Configurar las credenciales en el archivo `my-app-tx/.env`:

```properties
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=finanzas
DB_SSLMODE=disable
```

### 🐘 Uso postgres con Docker Compose

Ejemplo de levantar PostgreSQL y PgAdmin:

```yml
services:
  db-postgres17:
    image: postgres:17
    container_name: postgres_17
    restart: always
    environment:
      POSTGRES_PASSWORD: clavedb
    ports:
      - "5434:5432"
    volumes:
      - ./data17:/var/lib/postgresql/data
    networks:
      postgres:
        ipv4_address: 172.19.0.4

  pgadmin4:
    image: dpage/pgadmin4:9.6.0
    container_name: pgadmin4
    environment:
      PGADMIN_DEFAULT_EMAIL: "email@dominio.com"
      PGADMIN_DEFAULT_PASSWORD: "clavepgadmin"
    ports:
      - "81:80"
    depends_on:
      - db-postgres17
    volumes:
      - ./backups_pgadmin:/var/lib/pgadmin/backups
    networks:
      postgres:
        ipv4_address: 172.19.0.10

networks:
  postgres:
    ipam:
      driver: default
      config:
        - subnet: 172.19.0.0/16
```

---

## ▶️ Ejecución

Puedes iniciar la app con **Gin** o con **Gorilla/Mux**:

Inicia el servidor con `gin`:

```sh
go run main.go --router=gin
```

Inicia el servidor con `mux`:

```sh
go run main.go --router=mux
```

---

## 📂 Estructura del proyecto

```
.
├── controller
│   ├── gin                  # Controladores con Gin
│   └── mux                  # Controladores con Gorilla Mux
├── data                     # Persistencia y almacenamiento
├── env                      # Configuración de entorno
├── service                  # Lógica de negocio (patrón Strategy)
│   ├── banco_a_trx_service.go
│   ├── banco_b_trx_service.go
│   └── itrx_service.go
├── utils
│   ├── constants            # Constantes generales
│   ├── http                 # Respuestas HTTP estandarizadas
│   ├── logger               # Configuración de logs
│   ├── middleware           # Middlewares globales
│   └── models               # Modelos de datos
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## 🏗️ Patrón Strategy en la API

El **patrón Strategy** se implementa en `service/`:

* `itrx_service.go`: define la **interfaz** de transacciones.
* `banco_a_trx_service.go`: implementación específica para **Banco A**.
* `banco_b_trx_service.go`: implementación específica para **Banco B**.

En `itrx_service.go` se integra un **almacén de estrategias** donde se registran las implementaciones.  
Desde el controlador, se toma el header `x-tenant-id` para seleccionar el servicio correspondiente.

---

## 📖 Ejemplos de uso

Obtener todas las transacciones de un tenant:

```bash
curl --request GET   --url http://localhost:8080/transacciones   --header 'x-tenant-id: banco_b'
```

Crear una transacción:

```bash
curl --request POST   --url http://localhost:8080/transacciones   --header 'Content-Type: application/json'   --header 'x-tenant-id: banco_a'   --data '{
    "cuenta": "12345",
    "cuenta_destino": "67890",
    "monto": 250.75,
    "tipo": "INGRESO",
    "descripcion": "Pago mensual"
  }'
```

---

## 🔧 Próximas mejoras

* Pruebas unitarias con `testing`.
* Documentación OpenAPI (Swagger).
* Implementar un ORM (GORM o SQLC).
* Integración con contenedores Docker.
* Agregar soporte a otras bases de datos (ej. Oracle).
* CI/CD para despliegue automático.

---
