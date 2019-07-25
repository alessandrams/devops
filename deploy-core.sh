sudo su

# Entra no contexto do usuario do go
su gorunner

# Entra na pasta com o repositorio, atualiza e instala as dependencias
cd /home/gorunner/go/src/middleway
git checkout dev
git pull
go get
go build

# Remove o log antigo
rm -f /var/log/middleway.log
touch /var/log/middleway.log

# Inicia o servico de novo
systemctl restart middleway

# Verifica as saidas 
tail -f /var/log/middleway.log
