### For idea intellij

go modules -> enable go module integration 


### Check the generate_proto.sh


### Start mongoo
docker run -d  --name my-mongo -p 27017:27017   -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=admin mongo  

mydb
user: mydbadmin
pass: admin

mongosh -u mydbadmin -p admin --authenticationDatabase mydb

