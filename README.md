# go-ecommerce
An e-commerce website built with Go

## To run the application:
- make run

### Comando para rodar o container do MySQL:
- docker run --name mysqldatabase -p 3307:3306 -e MYSQL_ROOT_PASSWORD=mypassword123! -d mysql

### Comando para executar comando no container do banco de dados:
- docker exec -it mysqldatabase bash
- mysql -uroot -p -A
- Now, just run some command like: CREATE DATABASE `ecom`