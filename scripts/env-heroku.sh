#!/bin/sh

heroku config:set $(cat ./../deployment/docker/.env | sed '/^$/d; /#[[:print:]]*$/d') --app goway