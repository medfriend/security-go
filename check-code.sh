#!/bin/bash

##echo "🔍 Ejecutando análisis con golangci-lint..."
##golangci-lint run ./...
##echo

echo "🔐 Ejecutando análisis de seguridad con gosec..."
gosec ./...
echo

##echo "🧠 Ejecutando análisis estático con staticcheck..."
##staticcheck ./...
##echo

##echo "✅ Todos los análisis han sido ejecutados."
