package arquitetura

import (
	"runtime"
	"testing"
)

func TestDependente(t *testing.T) {
	t.Parallel()
	if runtime.GOARCH == "amd64" { //filtro o test para executar apenas se a arquitetura for amd64
		t.Skip("Não funciona em arquitetura amd64")
	}

	//..
	t.Fail()
}
