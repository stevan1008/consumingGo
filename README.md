# TESTING 1/ TWW.

#  Contrucción
* Golang
* Docker
* Postgresql (plus: PgAdmin4)

## Subconstrucción
Qué recursos alternos a la documentación oficial de Golang se usaron para la construcción del proyecto:

* Gorm: ORM de Golang y su driver especifico de postgres.
* Fiber: Framework de Golang basado en Express para el entorno de ejecución de NodeJs.
* JWT

#  Instalación y Ejecución
* Primero ejecutar los siguiente comandos en el ordén mostrado dentro de la raiz del proyecto:
    - `docker-compose build`
    - `docker-compose up -d`

Con lo anterior ya se puede proceder testear dentro del reverse Proxy.

#  Testing y Consumo

- Antes de comezar el consumo usar cualquier plataforma de Testing de API's (Postman, Insomnia, etc).
- El token se almacenará en cookies cuando se haga el logueo de sesión.
- Cuando se haga Logout el token se removerá de las cookies.

* Registro | POST Method

    URL `http://localhost/api/register`
      
    Siguiendo el siguiente ejemplo: 
    ```
    {
        "Name": "yourName",
        "Email": "your@email",
        "Password": "YourPassword"
    }
    ```

* Login | POST Method

    URL `http://localhost/api/login`
    ```
        {
            "Email": "your@email",
            "Password": "YourPassword"
        }
    ```

* Logout | POST Method

    URL `http://localhost/api/logout`
    
    No Body

* User | GET Method

    URL `http://localhost/api/user`

* Filtro | GET Method

    - El token deberá de existir en cookies para aplicar el filtro.
    - Query Params permitidos:
        - name (nombre de la cancion)
        - artist (artista musical)
        - album (album del artista)

    URL `http://localhost/api/filter?name=Ragga&artist=String`

