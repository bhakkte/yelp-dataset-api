# Yelp Dataset Api

## Introduction

This repo contains two separate parts:

An __importer__ which takes Yelps academic dataset, parses it and stores it nice
and tidy in collections in MongoDB.

The __API__ part is actual endpoint that allow you to query this data via a simple
and intuitive Restful HTTP interface.

## Importer: Using it

Step one is dowloading Yelp's academic dataset on your computer and un-tar it, be sure you have the flat `.json` files, one per object type.

Next ensure you have mongodb installed and running locally (or remotly if you want to fill in a production db)

Now the fun begins, we will start by compiling the binary

```bash
go build -o yack-importer ./yack-importer
```

Now we'll run this binary, it expects a few params, to check on those you can always run

```bash
./yack-importer -h
```

But what we really want is point the importer to a json file, hint it with a type, and specify the target mongodb.

```bash
./yack-importer -mongo-url mongodb://localhost:27017/yack -type user -file <../dataset/users.json>
```

Little extra tip: it is fun to know how much time it took to import, for this you can simply prefix your invocation with the handy `time` command.

```bash
time ./yack-importer ...
```

## Importer: Performance

First implementation of the importer was single "coroutined", if we can say so, and well, had few problems.

After 5 minutes of running it only imported 18 000 lines, making the math importing the 1 600 000 reviews would'have taken ~7.5 hours.

So, now it's using 30 go coroutines with one coroutine reading the file for them and one coroutine handling error.

Whooooho, wellcome speed and performance, the importer now goes at a very rough ~1000 row / sec

| Type       | Line count | Time  |
|------------|-----------:|-------|
| Businesses | 61 000     | 35s   |
| Users      | 336 000    | ?     |
| Reviews    | 1 600 000  | 4m10s |

## API: Using it

```bash
go run service/* -port 8080 -mongo-url=localhost/yack
```