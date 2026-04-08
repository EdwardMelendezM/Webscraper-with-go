#!/bin/bash

# Colores para la terminal
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo -e "${GREEN}Iniciando el despliegue de base de datos...${NC}"

# Levantar los contenedores en segundo plano (detached mode)
# --remove-orphans limpia contenedores antiguos que ya no están en el config
docker compose up -d --remove-orphans

echo -e "${GREEN}¡Todo listo! Los contenedores están corriendo.${NC}"
docker ps

# OJO: execute this: chmod +x execute.sh