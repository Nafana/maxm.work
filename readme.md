# A Golang URL Shortner

`maxm.work/<slug>`

Very simple URL shortner to make links on my resume appear more unique / brandable. Slugs to links can only be added directly to the mongo database, I have no use for implementing an API for myself (that being said in the future maybe I'll add it in)

## Configuring

Simply configure the `.env` file after creating it

```
cp .env.sample .env
```

## Building

If you want to build the code locally simply run `make`

To build with Docker run

```
docker build -t max.work .
```

## Running

To run locally `./bin/max.work`

To run Docker image detached

```
docker run -d --env-file ./.env -p 3030:3030 max.work:latest
```

To run Docker image with STDIN attached

```
docker run -it --env-file ./.env -p 3030:3030 max.work:latest
```
