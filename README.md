# registration-web-server
___
This is a registration form implemented in Go

## Start
___

1.Install and run [Docker](https://www.docker.com/get-started)  
2.Open terminal, in project folder   
3.Use this command to start database:

````
make build
````
4.Use this command to run migrations:

````
make migrate-up
````

5.Open a new console window and run this command:

````
make start
````