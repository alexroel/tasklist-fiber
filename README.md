# tasklist-fiber

## Herramientas necesarias

- **Tener instalado Go**

  - Configuración adicional en Linux:

    ```bash
    # Variables de entorno para Go
    export GOPATH=/home/alexroel/go
    export GOBIN=$GOPATH/bin
    export GOROOT=/usr/local/go

    # Configurar el PATH
    export PATH=$PATH:$GOBIN:$GOROOT/bin
    ```

- **Tener instalado VSCode**

  - Extensiones recomendadas para Go:
    - Go (golang.go)
    - Go Template Support

- **Instalar Postman** (para probar las rutas de la API)

- **Instalar PostgreSQL** (para la base de datos)

---

## Instalaciones de Fiber y GORM

1. Instalar Fiber:

   ```bash
   go get github.com/gofiber/fiber/v2
   ```

2. Instalar GORM y el driver de PostgreSQL:

   ```bash
   go get -u gorm.io/gorm
   go get -u gorm.io/driver/postgres
   ```

3. Instalar `air` para recargar automáticamente la aplicación durante el desarrollo:

   ```bash
   go install github.com/air-verse/air@latest

   air init

   air
   ```

4. Instalar soporte para variables de entorno:
   ```bash
   go get github.com/joho/godotenv
   ```

---

## Configuración de variables de entorno

Crea un archivo `.env` en el directorio raíz del proyecto con el siguiente contenido:

```
DB_DSN = host=localhost user=alexroel password=123456 dbname=tasks_db port=5432 sslmode=disable TimeZone=UTC
```

---

## Crear usuario y base de datos en PostgreSQL

Ejecuta los siguientes comandos en PostgreSQL para configurar la base de datos:

```sql
-- Crear un usuario
create user alexroel with password '123456';

-- Otorgar permisos de superusuario
alter user alexroel with superuser;

-- Crear la base de datos
create database tasks_db owner alexroel;

-- Verificar las bases de datos
\l

-- Conectarse a la base de datos
\c tasks_db;

-- Verificar las tablas
\dt
```

---

## Migrar los modelos de la aplicación

Después de configurar los modelos en tu aplicación, ejecuta las migraciones para crear las tablas en la base de datos. Luego, verifica que las tablas se hayan creado correctamente:

```sql
\dt
        Listado de relaciones
 Esquema | Nombre | Tipo  |  Dueño
---------+--------+-------+----------
 public  | tasks  | tabla | alexroel
 public  | users  | tabla | alexroel
(2 filas)
```

---

## Encriptar la contraseña

Para encriptar las contraseñas de los usuarios, instala la librería `bcrypt`:

```bash
go get golang.org/x/crypto/bcrypt
```

En el código, asegúrate de encriptar las contraseñas antes de guardarlas en la base de datos. Por ejemplo:

```go
import "golang.org/x/crypto/bcrypt"

hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
if err != nil {
    // Manejar el error
}
user.Password = string(hashedPassword)
```

---

## Ejecutar la aplicación

1. Inicia el servidor con `air`:

   ```bash
   air
   ```

2. Prueba las rutas de la API utilizando Postman o cURL.

---

## Notas adicionales

- Asegúrate de que las dependencias estén correctamente instaladas ejecutando:

  ```bash
  go mod tidy
  ```

- Si encuentras problemas con la conexión a la base de datos, verifica que las credenciales en el archivo `.env` sean correctas y que el servidor de PostgreSQL esté en ejecución.
