<h1 align="center" id="title">Rede Social em GoLang</h1>

<p id="description">O objetivo é criar uma rede social para servir como projeto de estudo e melhorias durante o aprendizado e melhoria no conhecimento do GoLang</p>

  
<h2> Próximos passsos: </h2>

Here're some of the project's best features:

*  Criar um endpoint que lê um CSV e carrega dados para a base de usuários e publicações
*  Refatorar algumas classes para tentar aplicar clean code
*  Adicionar testes unitarios ao projeto
*  Melhorar o read-me explicando como funciona a arquitetura básica desse projeto

<h2>🛠️ Para executar o projeto:</h2>

<p><strong>Utilizando docker</strong></p>
<p>1. Na raiz do projeto executar o comando</p>

```
docker-compose up 
```

<p><strong>Sem docker</strong></p>
<p>1. Instalar uma instancia do mysql, com os seguintes dados</p>

```
Usuário: golang
Senha: golang
Database: devbook
```

<p>2. Na raiz /api, vai subir o backend na porta 5000 </p>

```
go run main.go
```

<p>3. Na raiz /webapp, vai subir o frontend na porta 3000 </p>

```
go run main.go
```

<h2>Endereço para acesso local:</h2>

```
http://localhost:3000/login
```
    
<h2>💻 Construido com</h2>

Technologies used in the project:

*   GoLang
*   Mysql
