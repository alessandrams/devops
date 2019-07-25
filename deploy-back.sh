sudo su
# Entra no contexto do usuario do back-end
su pyback

# Mata o processo anterior
pkill python3.6

# Entra na pasta com o repositorio, atualiza e instala as dependencias
cd /home/pyback/back-end
git pull
sudo pip3 install -e .

# Inicializa a aplicacao com o nohup
nohup python3.6 run.py &

# Verifica a saida para conferir se nao ouve erros 
tail -f nohup.out

