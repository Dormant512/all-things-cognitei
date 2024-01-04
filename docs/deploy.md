# Deployment
## What to do
1. `ssh` to API host
2. `docker --version`, make sure it is installed
3. `git clone https://github.com/Dormant512/all-things-cognitei.git`
4. `cd all-things-cognitei`
5. `vim build/.env`, enter environment variables
6. `make run`

## .env file contents
```dotenv
DB_USER=placeholder_user
DB_PASSWORD=placeholder_pass
DB_NAME=""
DB_HOST=db-moc-things
DB_PORT=27017
DB_TYPE=mongodb
MG_COLLECTION=placeholder_collection
MG_SAVEFILE=placeholder_backup_filename
APP_PORT=3001
```

## Useful Makefile commands
- `make mongosh` enters db container with mongo shell
- `make logs-app` prints logs for app container
- `make delete-data` deletes volumes for mongo container, please back up first and run this only when no containers are running
- `make save-json` saves the data from mongo container to json backup file
- `make from-json` imports the backup json into mongo container

## Maintenance
No jobs were configured to regularly back up the database, so running `make save-json` and `scp` for retrieving the file from host to local is a necessity.
