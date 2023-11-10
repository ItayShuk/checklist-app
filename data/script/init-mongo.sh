#!/bin/bash
echo "@@@@@@@@@@@@@"

# Variables
db_name="DB"
collection_name="COLLECTION"
data_file="./data/db/COLLECTION.json" # Adjust the path as needed

# Wait for MongoDB to be ready

# Import data into MongoDB
mongoimport --db $db_name --collection $collection_name --file $data_file --jsonArray
#mongoimport --db $db_name --collection $collection_name --file $data_file --jsonArray
echo "Data import completed!"