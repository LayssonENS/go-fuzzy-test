# go-fuzzy-test
Exemplo de teste utilizando fuzzy


# Comandos para executar benchmark
go test -bench . -cpu 8 -benchmem -benchtime 5s -count 5 -memprofile mem.prof ./words

# Comando gerar profile memória
go tool pprof -http :8081 mem.prof