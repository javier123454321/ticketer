# Invoicer Demo Application
## By Javier Gonzalez
This application is written in Go. I like go, so that's why. I have been recently building with it and felt it was either this or using JavaScript (or TS) but I kind of have a philosophical preference against using Js in the backend. This was the quickest to get something ready, and it has everything in the standard library. (I also considered PHP and Laravel but it felt like overkill)
## The Application
It is a very simple Web App that allows you to (read only) an invoice. I planned on creating a way to create through the UI, but I already sank too much time into it and am at the point of cutting my losses. There are two routes, home, which shows you the outstanding invoices, and the invoice page. Now every invoice DOES HAVE a user tied to it but didn't implement the UI showing you that info for the same reason I didn't do a lot of things, time. In total I spent around 4 - 6 hours but I didn't count them properly. I would just rather write a better README.
## The Architecture
It is built in a rough approximation of an MVC framework. The controllers and the routing is all in the main file, but the models and views are better organized. I also added a migration and a seeder which took time away from writing the actual application logic. 
It is a server rendered monolith in the sense that there is very little happening client side. The postgres database has 3 tables: Users, Invoices, and Line Items. a User has many Invoices and an Invoice has many Line Items.
The client side logic is quite narrow, and the index page is frankly quite ugly. However, I did sprinkle in some tailwind css to make it a bit more bearable to look at. I used alpine.js although even that seemed a bit of overkill. However, it is a nice companion to the types of server rendered apps like this is shaping out to be. I did not feel like creating an SPA for reasons that should be apparent in the scope of the challenge.
## What is missing
It is not yet a CRUD app, it is just a CR app, with the create being only handled with the seeder. Obvious next steps are adding create functionality for line items, for invoices, and for users. Oh, also I should probably show the user that is the payer of the invoice, would do that next time. The application also lacks testing, styling, and some more refactoring (I meant to make a config object for the server info, separate the router, and create a controller directory). 
## How to Run it
You need Go installed and a postgres database called 'invoicer' (no user or password). Then install our 2 dependencies with `go get .`. Then you can migrate the database and seed it by running these commands:
```
go run models/migrations/migrations.go
go run models/seeder/seeder.go
```
You can then run 
```
go run main.go
``` 
on the root directory and it should spin up a local server which you can connect in port :8080. It should be seeded and ready!