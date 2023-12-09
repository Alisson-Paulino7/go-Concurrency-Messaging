# "As builder" eu vou rodar uma execução na imagem compilada e depois jogo numa menor

FROM golang:latest as builder

# Defino o local que vou rodar a aplicação
WORKDIR /app

# Copiar tudo daqui pra pasta app. Ele já entende só
COPY . .

# Detalhar como Linux porquê senão ele gera o executável para Windows direto
# Ele vai gerar o build do arquivo de api neste caso
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api ./cmd/api/main.go

FROM scratch
COPY --from=builder /app/api /

# Executando agora o arquivo que foi buildado antes

CMD ["./api"]