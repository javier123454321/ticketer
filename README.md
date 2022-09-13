# Ticketer Demo Application
## By Javier Gonzalez
This application is written in Go. I like go, so that's why. I have been recently building with it and felt it was either this or using JavaScript (or TS) but I kind of have a philosophical preference against using Js in the backend. This was the quickest to get something ready, and I used a boilerplate I had already made.
## The Application
It is a very simple Web App that allows you to read and create tickets.There are three routes, home, create, and show tickets.

## The Architecture
It is built in a rough approximation of an MVC framework. The routing is all in the main file, but the models controllers and views are in their respective packages. There is also a migration and seeding utility which will help you get set up. 
It is a server rendered monolith in the sense that there is very little happening client side. The postgres database has 2 tables: Users and Tickets. A user has many tickets and a ticket has one user.
The client side logic is quite narrow and the styling is quite secondary. However, I did sprinkle in some tailwind css to make it a bit more bearable to look at. I used alpine.js for client side functionality.It is a nice companion to the types of server rendered apps like this is shaping out to be. I did not feel like creating an SPA for reasons that should be apparent in the scope of the challenge. This is mostly relying on native web and browser technologies - a true server side application.
## What is missing
It is not yet a CRUD app, it is just a Create Read app. Obvious next steps are adding create functionality for users, and update for both users and tickets. The application also lacks testing, styling, and some more refactoring (I meant to make a config object for the server info, separate the router, for example). 
## How to Run it
You need Go installed and a postgres database called 'ticketer' (no user or password). Then install our 2 dependencies with `go get .`. Then you can migrate the database and seed it by running these commands:
```
go run models/migrations/migrations.go
go run models/seeder/seeder.go
```
You can then run 
```
go run main.go
``` 
on the root directory and it should spin up a local server which you can connect in port :8080. It should be seeded and ready!