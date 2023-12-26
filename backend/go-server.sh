#!/bin/bash
mkdir -p ./storage
chmod 777 -R ./storage


# Use the PORT variable, or default to 8080 if not set
gin --appPort 8080 --immediate
