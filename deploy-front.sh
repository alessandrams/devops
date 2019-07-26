#!/bin/bash

# Entra no contexto do usuario do front-end
sudo -i -u pyback bash << EOF

# Mata o processo anterior
pkill node

# Entra na pasta com o repositorio, atualiza e instala as dependencias
cd /home/node/front-end
git pull
yarn
yarn build

# Inicializa a aplicacao com o serve
serve -s build -l 3000 &

EOF









