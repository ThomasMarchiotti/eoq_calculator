#!/bin/bash

# Esegui il backend
cd backend

# Crea la directory build e la sottocartella frontend nella root del progetto se non esistono
mkdir -p ../build/frontend

# Compila il file server.go e posiziona il file compilato nella directory build
go build -o ../build/server src/server.go

# Copia il file home.html nella directory build
cp frontend/home.html ../build/frontend

# Naviga nella directory build
cd ../build

# Esegui il file compilato
./server