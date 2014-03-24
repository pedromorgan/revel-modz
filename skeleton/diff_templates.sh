#!/bin/bash

templates=(
	init.sh
	Gruntfile.js

	conf/app.conf
	app/init.go

	app/controllers/admin.go
	app/controllers/auth.go
	app/controllers/maillist.go
	app/controllers/signup.go
	app/controllers/user.go

	app/views/templates/header.html
)

for file in ${templates[@]}
do
	echo "${file}"
	echo "------------------------"
	colordiff ${file}.template ../sample/${file} 
	echo
	echo
	echo
done
