# logs-monitor-api

### Ambientes
1) Ambiente Docker  
2) Instalação Manual.  

### 1- Ambiente Docker
a) Baixar e instalar a última versão do docker.  
b) Realizar o clone do repositório.  
c) Através do Power Shell ou CMD, acessar o diretório clonado.  
d) Executar: docker-compose up -d  
e) Executar comando para visualizar os logs: docker-compose logs -f -t   

Este processo gerará dois containers e através deles, poderemos inciar e parar as aplicações de maneira independente.  
Para listar os containers: docker container ps ou docker container ps -a  
Serão listados os containers da seguinte forma:  

|CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES |
|-------------|-------|---------|---------|--------|-------|-------|
|218c3fa8c2ed|mongo:3.6|"docker-entrypoint.s…"|About an hour ago|Up About an hour| 0.0.0.0:27017->27017/tcp|logs-monitor-api_db_1
|b0fcfa30606a|golang:1.11.5|"bash -c 'cd /usr/lo…"|About an hour ago|Up About an hour|0.0.0.0:3000->3000/tcp|logs-monitor-api_backend_1

a) Container de banco de dados: mongodb  
Para testar se o banco está online: http://localhost:27017  

b) Container de backend: golang  
Para testar se o backend está online: http://localhost:3000/logs  

### 2- Instalação Manual
a) Baixar e instalar o GOlang na versão 1.11.5  
b) Baixar e instalar o MongoDB na versão 3.6  
c) Realizar o clone do repositório.  
d) Apontar para localhost na chave "server" do arquivo config.toml  
e) Através do Power Shell ou CMD, acessar o diretório clonado.  
f) Iniciar o MongoDB. Comando: mongod  
g) Iniciar o backend. Comando: go run app.go  

Este processo deverá inciar a API, ficando disponível na porta 3000.  

### Testes
a) GET ALL  
http://localhost:3000/logs  

b) GET BY ID  
http://localhost:3000/logs/1  

c) POST  
http://localhost:3000/logs  
{
"codigo_filial": "1",
"nome_filial": "Teste",
"tipo_notificacao": "M",
"descricao_erro": "Erro ao integrar pedido",
"conteudo_mensagem_erro": "NPE",
"data_hora": "19/02/2019"
}



