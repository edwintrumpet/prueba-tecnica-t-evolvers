# Prueba técnica t-evolvers

Crud para realizar el seguimiento de órdenes de servicio

Disponible en http://35.175.235.82
Documentación http://35.175.235.82/swagger/index.html

## Desarrollo

Instalar **swag** para construir la documentación

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Construir la documentación de swagger

```bash
swag init
```

Proveer una base de datos postgres desde docker pasándole las variables de entorno

```bash
docker run --name db -e POSTGRES_PASSWORD=helloworld -e POSTGRES_USER=edwin -e POTGRES_DB=enerbit -p 5432:5432 postgres
```

Proveer un redis desde docker

```bash
docker run --name rdb -p 6379:6379
```

Crear un archivo `.vscode/launch.json` para ejecutar la aplicación usando el
debugger de Visual Studio Code y que contenga las variables de entorno
necesarias para el funcionamiento y conexión a la base de datos.

```json
{
    // Use IntelliSense para saber los atributos posibles.
    // Mantenga el puntero para ver las descripciones de los existentes atributos.
    // Para más información, visite: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "enerbit",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "DB_USER": "edwin",
                "DB_PASSWORD": "helloworld",
                "DB_NAME": "enerbit",
                "ENV": "dev"
            }
        }
    ]
}
```

Poniendo la variable de entorno `env` como **dev** puedes ver los logs para debug

## Despliegue

Ejecute usando docker-compose
```bash
docker-compose up -d
```

Variables requeridas para ejecutar desde docker-compose

- DB_USER
- DB_PASSWORD
- DB_NAME

Puedes crear un archivo `.env` para proveerlas

```
DB_USER=edwin
DB_PASSWORD=helloworld
DB_NAME=enerbit
```
