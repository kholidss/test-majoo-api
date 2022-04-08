#!/bin/bash

echo "Downloading packages needed ..."
go get firebase.google.com/go
go get github.com/go-ozzo/ozzo-validation/v4
go get github.com/jinzhu/gorm
go get github.com/joho/godotenv
go get github.com/labstack/echo/v4

echo "Setup initial git hooks ..."
cp pre-commit .git/hooks
chmod +x .git/hooks/pre-commit
