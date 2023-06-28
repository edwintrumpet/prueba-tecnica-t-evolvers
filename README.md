# Prueba técnica t-evolvers

Crud para realizar el seguimiento de órdenes de servicio

## Desarrollo

No hay docker-compose para develop. La api fue desarrollada ejecutando con el
debugger de vsCode.

Puedes proveer postgres y redis desde docker exponiendo los puertos en local y
ejecutar la api con el debugger de vsCode usando un `.vscode/launch.json` para
proveer las variables de entorno

```json
{
    // Use IntelliSense para saber los atributos posibles.
    // Mantenga el puntero para ver las descripciones de los existentes atributos.
    // Para más información, visite: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "DB_USER": "edwin",
                "DB_PASSWORD": "helloworld",
                "DB_NAME": "enerbit",
                "DB_PORT": "5432",
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
