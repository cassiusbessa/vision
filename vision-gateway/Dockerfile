# Utiliza uma imagem base oficial do Golang
FROM golang:latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o arquivo go.mod e go.sum e instala as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte da aplicação para o contêiner
COPY . .

# Compila a aplicação
RUN go build -o app

# Expõe a porta em que sua aplicação estará escutando (ajuste conforme necessário)
EXPOSE 8888

# Define o comando que será executado ao iniciar o contêiner
CMD ["./app"]