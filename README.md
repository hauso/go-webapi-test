## Basic install
1. Install go (docker and docker compose)
2. create local files from `.dist` files (.env, .gitignore, docker-compose.yaml)
3. go to api/ folder
4. create softlink to `.env` in `api/` folder (linux: ln -s ../.env .env)
5. run `go build webapi` (if some requirement occurred, install it)
6. go to main folder 
7. run `docker-compose build -d`

## Temporary requirements (personal infos)
- require login to container and run api server
 `docker exec -it api_go bash`
  when logged to docker in console run: 
  `go run webapi`
- maybe there could be a problem with db dir
  check if `docker/db/data` folder exists

## Check it run
api server: http://localhost:8080

db admin (pma): http://localhost:10081 
for credentials see `.env` file 

for basic operation @see 
`Go-WebAPI.postman_collection.json`
