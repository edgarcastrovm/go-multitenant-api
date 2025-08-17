package middleware

import (
	"context"
	. "my-app-tx/utils/constants"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

// CustomFormatter personaliza el formato de los logs
type CustomFormatter struct{}

// Intercepta las peticiones http
func RequestIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener TenantID para setear el dueño de la petición
		tenantId := r.Header.Get(HEADER_TENANT_ID)
		if tenantId == "" {
			tenantId = "undefined"
		}

		// Obtener RequestID del header o generar uno nuevo
		requestId := r.Header.Get(HEADER_REQUEST_ID)
		// Generar UUID si no hay RequestID
		if requestId == "" {
			requestId = uuid.New().String()
		}

		// Agregar RequestID al contexto
		ctx := r.Context()
		ctx = context.WithValue(ctx, HEADER_TENANT_ID, tenantId)
		ctx = context.WithValue(ctx, HEADER_REQUEST_ID, requestId)
		r = r.WithContext(ctx)

		// Llamar al siguiente manejador
		next.ServeHTTP(w, r)
	})
}

func RequestIDMiddlewareGin(c *gin.Context) {

	tenantId := c.GetHeader(HEADER_TENANT_ID)
	if tenantId == "" {
		tenantId = "unknown_tenant"
	}

	requestId := c.GetHeader(HEADER_REQUEST_ID)
	if requestId == "" {
		requestId = uuid.New().String()
	}

	// Insertar en el contexto base
	ctx := c.Request.Context()
	ctx = context.WithValue(ctx, HEADER_TENANT_ID, tenantId)
	ctx = context.WithValue(ctx, HEADER_REQUEST_ID, requestId)

	// Guardar en gin.Context
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
