package appliance

import (
	"reflect"
	"testing"
)

const erroPadrao = "Valor Esperado %v, mas o resultado encontrado foi %v."

func TestExampleTesteFunc(t *testing.T) {
	testes := []struct {
		entrada  string
		esperado string
	}{
		//testes ok
		{"testeOK", "testeOK"},
		//Testes de entradas invalidas
		//{"teste", []byte("teste")},
		//{ 257,0},
	}

	for _, teste := range testes {
		t.Logf("Teste %v", teste)
		atual := ExampleTesteFunc(teste.entrada)
		if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}

}

func TestReadFileErr(t *testing.T) {
	testes := []struct {
		entrada  string
		esperado []byte
	}{
		//testes ok
		{"hard.conf", []byte("medias\n")},
		//Testes de entradas invalidas
		//{"teste", []byte("teste")},
		//{ 257,0},
	}

	for _, teste := range testes {
		t.Logf("Teste %v", teste)
		atual, _ := ReadFileErr(teste.entrada)
		if !reflect.DeepEqual(atual, teste.esperado) {
			//if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}

}

/*
func TestReadFile(t *testing.T) {
	testes := []struct {
		entrada  string
		esperado []byte
	}{
		//testes ok
		{"hard.conf", []byte("medias\n")},
		//Testes de entradas invalidas
		//{"teste", []byte("teste")},
		//{ 257,0},
	}

	for _, teste := range testes {
		t.Logf("Teste %v", teste)
		atual := ReadFile(teste.entrada)
		if !reflect.DeepEqual(atual, teste.esperado) {
			//if atual != teste.esperado {
			t.Errorf(erroPadrao, teste.esperado, atual)
		}
	}

}
*/