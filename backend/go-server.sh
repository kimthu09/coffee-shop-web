#!/bin/bash
mkdir -p ./storage
chmod 777 -R ./storage

# Load PORT variable from .env file
if [ -f .env ]; then
    export $(cat .env | xargs)
fi

# Use the PORT variable, or default to 8080 if not set
gin --appPort ${PORT:-8080} --immediate
