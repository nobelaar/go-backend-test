# API de autenticacion (Go + SQLite)

Servicio HTTP minimalista para registrar y autenticar usuarios con SQLite y contrasenas protegidas con bcrypt. Expone endpoints JSON simples para `/register` y `/login`.

## Requisitos
- Go 1.24 o superior
- Nada mas: el driver SQLite se instala con `go mod download`

## Puesta en marcha
1) Instala dependencias (solo la primera vez):  
   `go mod download`
2) Ejecuta el servidor en puerto 8080:  
   `go run cmd/api/main.go`

La base `app.db` se crea en la raiz si no existe.

## Endpoints
- `POST /register`  
  Body: `{"username": "alice", "password": "secret"}`  
  Respuestas: `201` exito, `409` usuario ya existe, `400` JSON invalido.

- `POST /login`  
  Body: igual a register  
  Respuestas: `200` exito, `401` credenciales invalidas, `400` JSON invalido.

Ejemplos rapidos:
```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"secret"}'

curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"alice","password":"secret"}'
```

## Estructura del proyecto
- `cmd/api/main.go`: arranque del servidor HTTP y wiring de dependencias.
- `internal/delivery/http`: handlers y router de los endpoints.
- `internal/usecase`: logica de negocio para registrar y autenticar.
- `internal/infra/sqlite`: inicializacion de la DB y repositorio de usuarios.
- `internal/infra/crypto`: hasher bcrypt.
- `internal/domain`: entidad `User` y errores de dominio.
- `internal/repository`: interfaces de repositorio.

## Notas
- Contrasenas se almacenan con bcrypt por defecto.
- Para reiniciar datos, elimina `app.db` antes de volver a ejecutar.
