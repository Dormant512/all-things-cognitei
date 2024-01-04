ifneq (,$(wildcard ./build/.env))
    include build/.env
    export
endif

.PHONY: run
run:
	docker compose -f ./build/docker-compose.yml rm && \
	docker compose -f ./build/docker-compose.yml build --no-cache && \
	docker compose -f ./build/docker-compose.yml up -d


.PHONY: rerun
rerun:
	docker compose -f ./build/docker-compose.yml up -d


.PHONY: destroy
destroy:
	docker compose -f ./build/docker-compose.yml down


.PHONY: check-config
check-config:
	docker compose -f ./build/docker-compose.yml config


.PHONY: mongosh
mongosh:
	docker exec -it db-moc-things mongosh


.PHONY: databash
databash:
	docker exec -it db-moc-things bash


.PHONY: logs-app
logs-app:
	docker logs app-moc-things


.PHONY: save-json
save-json:
	docker exec -i db-moc-things mongoexport --port $(DB_PORT) --db admin \
	-u $(DB_USER) -p $(DB_PASSWORD) --collection $(MG_COLLECTION) --pretty \
	--jsonArray | tee backup/$(MG_SAVEFILE).json backup/$(MG_SAVEFILE)_$$(date +%Y-%m-%d_%H-%M-%S).json > /dev/null


.PHONY: from-json
from-json:
	docker cp backup/$(MG_SAVEFILE).json db-moc-things:retrieve_data.json && \
	docker exec -i db-moc-things mongoimport --port $(DB_PORT) --db admin \
    -u $(DB_USER) -p $(DB_PASSWORD) --collection $(MG_COLLECTION) --jsonArray retrieve_data.json


.PHONY: delete-data
delete-data:
	docker volume rm build_mongo-data


#.PHONY: save-potions
#save-potions:
#	docker exec -i db-moc-things mongoexport --port $(DB_PORT) --db admin \
#	-u $(DB_USER) -p $(DB_PASSWORD) --collection $(MG_COLLECTION) --pretty \
#	--jsonArray > backup/$(MG_SAVEFILE)_$$(date +%Y-%m-%d_%H-%M-%S).json

