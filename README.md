# cipherbreaker-go
Simple criptoanalysis in Go

## How to start? :star:
Como começar? Para começar a desenvolver nesse projeto, você deve cloná-lo na sua máquina local. Nas instruções a seguir, considere que "$ ..." significa executar na linha de comando.

### Usando go get
A opção mais tranquila é:

$ go get github.com/titosilva/cipherbreaker-go/...

Isso clonará o repositório na pasta GOPATH/src/github.com/titosilva/cipherbreaker-go, onde GOPATH é dado por

$ go env

### Manualmente
Para isso, primeiro vocẽ deverá descobrir a sua pasta GOPATH (que corresponde ao local onde o go
busca pelo seu código). Isso pode ser feito pelo comando

$ go env

Que retorna as variáveis de ambiente associadas ao Go, e, entre elas, o GOPATH. Navegue até a pasta 
indicada pelo GOPATH. Procure pela pasta src, navegue até ela, confira se há uma pasta chamada
"github.com". Se houve, navegue até ela. Crie a pasta "titosilva", navegue até ela e use o comando

$ git clone github.com/titosilva/cipherbreaker-go.git

para clonar o repositório. Após isso, você terá uma pasta acessível pelo go build e outras ferramentas do go.

## Instalando as dependências
Você também precisará instalar as dependências:

$ go get golang.org/x/term


## How to know what I can do?
Como saber o que posso fazer?
Na aba de issues, as tarefas foram adicionadas. Eventualmente e conforme a necessidade, adicionamos mais tarefas aos poucos :happy:

## Obrigado!!!!
