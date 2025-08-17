package middleware

import (
	"context"
	"fmt"
	. "my-app-tx/utils/constants"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type logEntryKey string

const logKey logEntryKey = API_LOGKEY

// CustomFormatter personaliza el formato de los logs
type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	pkg, _ := entry.Data["pkg"].(string)
	file, _ := entry.Data["file"].(string)
	if pkg == "" || file == "" {
		pkg = "unknown"
		file = "unknown"
	}

	requestID, _ := entry.Data["request_id"].(string)
	tenantID, _ := entry.Data["tenant_id"].(string)
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := entry.Level.String()
	message := entry.Message

	formatted := fmt.Sprintf("%s-[%s]-[%s]-[%s]- %s/%s %s\n",
		timestamp, tenantID, requestID, level, pkg, file, message)

	return []byte(formatted), nil
}

// GetLogger obtiene el logger del contexto
func GetLogger(r *http.Request) *logrus.Entry {
	if logger, ok := r.Context().Value(logKey).(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithField("request_id", "unknown")
}

// Intercepta las peticiones http
func RequestIDMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener TenantID para setear el dueño de la petición
		tenant_id := r.Header.Get("x-tenant-id")
		if tenant_id == "" {
			tenant_id = "undefined"
		}

		// Obtener RequestID del header o generar uno nuevo
		requestId := r.Header.Get("x-request-id")
		// Generar UUID si no hay RequestID
		if requestId == "" {
			requestId = uuid.New().String()
		}
		// Configurar logger con RequestID
		logger := logrus.New()
		logger.SetFormatter(&CustomFormatter{})
		logger.SetLevel(logrus.InfoLevel)
		logEntry := logger.WithFields(logrus.Fields{"request_id": requestId, "tenant_id": tenant_id})
		// Agregar RequestID al contexto
		ctx := context.WithValue(r.Context(), logKey, logEntry)
		r = r.WithContext(ctx)

		// Llamar al siguiente manejador
		next.ServeHTTP(w, r)
	})
}

func GetLoggerGin(c *gin.Context) *logrus.Entry {
	if logger, ok := c.Get(API_LOGKEY); ok {
		if entry, ok := logger.(*logrus.Entry); ok {
			return entry
		}
	}
	return logrus.WithField("request_id", "unknown")
}
func RequestIDMiddlewareGin(c *gin.Context) {
	// Establecer el ID en el contexto para que esté disponible en los manejadores

	// Obtener TenantID para setear el dueño de la petición
	tenant_id := c.GetHeader("x-tenant-id")
	if tenant_id == "" {
		tenant_id = "undefined"
	}
	// Agregar el ID al encabezado de la respuesta
	requestId := c.GetHeader("x-request-id")
	// Generar UUID si no hay RequestID
	if requestId == "" {
		requestId = uuid.New().String()
	}
	// Configurar logger con RequestID
	logger := logrus.New()
	logger.SetFormatter(&CustomFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	logEntry := logger.WithFields(logrus.Fields{"request_id": requestId, "tenant_id": tenant_id})

	c.Set(API_LOGKEY, logEntry)
	// Continuar con la cadena de middleware
	c.Next()
}

func GetLog(pkg, file string, log *logrus.Entry) *logrus.Entry {
	return log.WithFields(logrus.Fields{"pkg": pkg, "file": file})
}
