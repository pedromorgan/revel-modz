#!/bin/bash

# npminit.sh

npm init
npm install grunt
npm install grunt-contrib-jshint grunt-contrib-concat grunt-contrib-uglify grunt-contrib-coffee grunt-markdown grunt-contrib-sass grunt-contrib-less

grunt
# the init.sh.template file move has been excluded here as it has been theoretically done already