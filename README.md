# Prueba técnica t-evolvers

Crud para realizar el seguimiento de órdenes de servicio

## Develop

There is no a docker-compose for develop. The api was developed running with the
vsCode debugger.

You can provide postgres and redis from docker exposing ports to attach locally,
and provide env variables to the app using a `.vscode/launch.json`

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

Setting env variable `env` as **dev** you can see logs for debug

```

## Deploy

Execute using docker-compose

Required environment variables to run from docker-compose

- DB_USER
- DB_PASSWORD
- DB_NAME

you can create an `.env` file to provide them

```
DB_USER=edwin
DB_PASSWORD=helloworld
DB_NAME=enerbit
```
