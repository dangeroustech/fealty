#!/bin/bash

echo "###---MongoDB Connectivity Check---###"
mongosh mongodb://fealty:$MONGODB_FEALTY_PASS@localhost:27017/fealty?authSource=fealty --eval "db.getCollectionNames()"
