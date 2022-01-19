package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) {
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

//go test -> executa todos os testes
//go test -cover -> mostrar % de coverage
//go test -coverprofile=result.out -> gera um arquivo result.out com o resultado
//go tool cover -func=result.out -> ferramenta para ler os resultados
//go tool cover -html=result.out -> ferramenta para ler os resultados via HTML
