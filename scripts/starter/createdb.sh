#!/bin/sh
mysql -uroot -e 'CREATE DATABASE IF NOT EXISTS bookshelf;''
mysql -uroot -e 'CREATE USER bookshelf IDENTIFIED BY "hogehoge"; GRANT ALL ON bookshelf.* TO "bookshelf"@"localhost";'
