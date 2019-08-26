#!/bin/sh

heroku config:set $(cat ./heroku.env | sed '/^$/d; /#[[:print:]]*$/d') --app goway