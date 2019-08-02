#!/bin/bash

docker run --rm --name mysql -v mysql_data:/var/lib/mysql -e MYSQL_USER=root -e MYSQL_PASSWORD=newhacker -e MYSQL_DATABASE=grab -p 3306:3306 -d mysql
cat User_Table.sql | docker exec -i mysql mysql -uroot -pnewhacker grab
