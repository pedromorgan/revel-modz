#!/bin/bash

# install grunt module dependencies
npm init
npm install grunt
npm install grunt-contrib-jshint grunt-contrib-uglify grunt-contrib-sass

grunt

# the init.sh.template file move has been excluded here as it has been theoretically done already
