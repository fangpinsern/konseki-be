# Konseki App Backend

Written in Go.

## Development Principle
```
Make it work. Refactor later.
```

Though this is not good principles, it works for a single member team.
However, once it works, it should be refactored to make it more readable as
well as follows good software engineering principles - especially DRY

## Current structure of code

We use the controller-services model.

Controller - interacts with service to obtain information and manipulate it to 
fit feature needs.

Services - interacts with database to obtain information. Primary purpose is to cast
to relevant structs for controllers use. Handles low level get and set functions. This 
provides and abstracting layer to help with querying the database.

## Backend Stack

### Language

We are using `Go` as our programming language of choice. Go supports
concurrency in a simple way and is a strictly typed language. Most importantly, it is a language
I am less familiar with hence a good learning opportunity.

### Database

For database, we are using `Firebase` to store our information. It uses a 
collection-document model to store information. It is `NoSQL` and most importantly, `Free`.
Main reason for choosing this over mongo or cassandra is due to the authentication support
firebase provides. It is known that it is best to not build your own authentication system.
I do know how to do it, but I think it would be the best if my time is spent on building the
features.

### Logging

Currently we use logrus (a go package for logging to the log file). It does support remote 
logging which we can tap into in the future.