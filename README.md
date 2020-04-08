# Healthmeter


# run it all together! with Docker 
> docker-compose build && docker-compose up -d

# Opening shell of mongo db running within our docker container
> docker exec -it mongo mongo -u " --- (MONGO_ADMIN) --- " -p " --- (MONGO_ADMIN_PSWD) --- " --authenticationDatabase admin
# Changing to (creating) a needed db
> use --- (MONGO_DB_NAME) ---
# Creating a super user
> db.createUser({user: ' --- (MONGO_DB_USER) --- ', pwd: '---- (MONGO_DB_PSWD) ----', roles:[{role:'dbOwner', db:'--- (MONGO_DB_NAME) ---'}]})
# Inserting test data to test collection
> db.test_collection.insert({ test: "test" })
# Displaying all collections of our previously created db in order to make sure that our test collection was successfully created
> show collections
# Saying goodbye to mongo shell
> exit

# stop the corona Api container 
docker stop <api_contrainer_id> 

# rerun the corona Api container 
docker-compose build && docker-compose up -d