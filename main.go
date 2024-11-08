package main

// Importa pacotes necessários para o funcionamento do código.
import (
	"bufio"   // Pacote para ler entradas do usuário de forma mais controlada
	"fmt"     // Pacote para imprimir mensagens no console
	"os"      // Pacote para interagir com o sistema operacional, como ler entradas do teclado
	"strconv" // Pacote para converter valores entre diferentes tipos, como de string para inteiro
)

var memoria = [8]string{"00001", "00010", "00000", "00000", "00000", "00000", "00000", "00000"}
var acumulador = "00000"

// Converte uma string binária para 5 bits e verifica se excede o limite
func converte5Bits(valor string) (string, bool) {
	for len(valor) < 5 {
		valor = "0" + valor
	}
	if len(valor) > 5 {
		return "00000", true
	}
	return valor, false
}

// Função para validar se o valor é binário (apenas 0 e 1)
func validaBinario(valor string, tamanho int) bool {
	if len(valor) != tamanho {
		return false
	}
	for _, c := range valor {
		if c != '0' && c != '1' {
			return false
		}
	}
	return true
}

// Função para validar se o endereço é binário (apenas 0 e 1) e com 3 bits
func validarEndereco(valor string) bool {
	return validaBinario(valor, 3)
}

// Mostra o estado atual da memória
func mostraMemoria() {
	fmt.Print("Memória: ")
	for _, valor := range memoria {
		fmt.Print(valor, " ")
	}
	fmt.Println()
}

// Mostra o valor do acumulador
func mostraAcumulador() {
	fmt.Println("AC:", acumulador)
}

// Comandos
func ler(posicao int) {
	acumulador = memoria[posicao]
	mostraAcumulador()
}

func escrever(posicao int, dado string) {
	// Verifica se o dado é válido (5 bits binários)
	if !validaBinario(dado, 5) {
		fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
		return
	}
	memoria[posicao] = dado
	mostraMemoria()
}

func somar(posicao int, dado string) {
	// Verifica se o dado é válido (5 bits binários)
	if !validaBinario(dado, 5) {
		fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
		return
	}
	valorMem, _ := strconv.ParseInt(memoria[posicao], 2, 64)
	valorDado, _ := strconv.ParseInt(dado, 2, 64)
	resultado := valorMem + valorDado
	resultadoBin := strconv.FormatInt(resultado, 2)
	res, overflow := converte5Bits(resultadoBin)
	if overflow {
		fmt.Println("Overflow detectado!")
	} else {
		memoria[posicao] = res
	}
	mostraMemoria()
}

func subtrair(posicao int, dado string) {
	// Verifica se o dado é válido (5 bits binários)
	if !validaBinario(dado, 5) {
		fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
		return
	}
	valorMem, _ := strconv.ParseInt(memoria[posicao], 2, 64)
	valorDado, _ := strconv.ParseInt(dado, 2, 64)
	resultado := valorMem - valorDado
	resultadoBin := strconv.FormatInt(resultado, 2)
	res, overflow := converte5Bits(resultadoBin)
	if overflow || resultado < 0 {
		fmt.Println("Overflow detectado!")
		memoria[posicao] = "00000"
	} else {
		memoria[posicao] = res
	}
	mostraMemoria()
}

func multiplicar(posicao int, dado string) {
	// Verifica se o dado é válido (5 bits binários)
	if !validaBinario(dado, 5) {
		fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
		return
	}
	valorMem, _ := strconv.ParseInt(memoria[posicao], 2, 64)
	valorDado, _ := strconv.ParseInt(dado, 2, 64)
	resultado := valorMem * valorDado
	resultadoBin := strconv.FormatInt(resultado, 2)
	res, overflow := converte5Bits(resultadoBin)
	if overflow {
		fmt.Println("Overflow detectado!")
	} else {
		memoria[posicao] = res
	}
	mostraMemoria()
}

func dividir(posicao int, dado string) {
	// Verifica se o dado é válido (5 bits binários)
	if !validaBinario(dado, 5) {
		fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
		return
	}
	valorMem, _ := strconv.ParseInt(memoria[posicao], 2, 64)
	valorDado, _ := strconv.ParseInt(dado, 2, 64)
	if valorDado == 0 {
		fmt.Println("Erro: divisão por zero!")
	} else {
		resultado := valorMem / valorDado
		resultadoBin := strconv.FormatInt(resultado, 2)
		res, overflow := converte5Bits(resultadoBin)
		if overflow {
			fmt.Println("Overflow detectado!")
		} else {
			memoria[posicao] = res
		}
	}
	mostraMemoria()
}

func comparar(posicao1, posicao2 int) {
	valor1, _ := strconv.ParseInt(memoria[posicao1], 2, 64)
	valor2, _ := strconv.ParseInt(memoria[posicao2], 2, 64)
	if valor1 > valor2 {
		fmt.Printf("Memória[%d] é maior que Memória[%d]\n", posicao1, posicao2)
	} else if valor1 < valor2 {
		fmt.Printf("Memória[%d] é menor que Memória[%d]\n", posicao1, posicao2)
	} else {
		fmt.Printf("Memória[%d] é igual a Memória[%d]\n", posicao1, posicao2)
	}
}

func compararArray() {
	var maior, menor int64 = -1, 32
	for _, valor := range memoria {
		valorInt, _ := strconv.ParseInt(valor, 2, 64)
		if valorInt > maior {
			maior = valorInt
		}
		if valorInt < menor {
			menor = valorInt
		}
	}
	fmt.Printf("Maior valor na memória: %d, Menor valor na memória: %d\n", maior, menor)
}

func limparArray() {
	for i := range memoria {
		memoria[i] = "00000"
	}
	fmt.Println("Memória limpa!")
	mostraMemoria()
}

func limparPosicao(posicao int) {
	memoria[posicao] = "00000"
	fmt.Printf("Memória[%d] limpa!\n", posicao)
	mostraMemoria()
}

func executaInstrucao(instrucao string) bool {
	comando := instrucao[:4]
	endereco := instrucao[4:7]
	dado := instrucao[7:]

	posicao1, _ := strconv.ParseInt(endereco, 2, 64)

	switch comando {
	case "0000":
		fmt.Println("Encerrando o programa...")
		mostraMemoria()
		return false
	case "0001":
		ler(int(posicao1))
	case "0010":
		escrever(int(posicao1), dado)
	case "0011":
		somar(int(posicao1), dado)
	case "0100":
		subtrair(int(posicao1), dado)
	case "0101":
		multiplicar(int(posicao1), dado)
	case "0111":
		dividir(int(posicao1), dado)
	case "1000":
		// Comando 1000: Comparar dois endereços
		// Solicitar o "outro endereço" para comparação
		fmt.Print("Digite o outro endereço (3 bits): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		endereco2 := scanner.Text()

		// Valida se o endereço inserido tem 3 bits e é válido
		if !validarEndereco(endereco2) {
			fmt.Println("Erro: endereço inválido. Apenas 0 e 1 são permitidos, com 3 bits.")
			return true
		}

		// Convertendo o segundo endereço para posição
		posicao2, _ := strconv.ParseInt(endereco2, 2, 64)

		// Exibe os endereços que estão sendo comparados
		fmt.Printf("Memória[%d] = %s e Memória[%d] = %s\n", posicao1, memoria[posicao1], posicao2, memoria[posicao2])

		// Verifica se o segundo endereço está dentro do intervalo de memória
		if posicao2 >= 0 && posicao2 < 8 {
			comparar(int(posicao1), int(posicao2))
		} else {
			fmt.Println("Erro: o endereço está fora do intervalo de memória.")
		}
	case "1001":
		// Comando 1001: Comparar os valores na memória (maior e menor valor)
		// Este comando não precisa de endereço ou dado
		compararArray()
	case "1010":
		// Comando 1010: Limpar uma posição específica
		// Solicitar o endereço a ser limpo
		fmt.Print("Digite o endereço (3 bits) para limpar: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		endereco := scanner.Text()

		// Valida o endereço (deve ser 3 bits binários)
		if !validarEndereco(endereco) {
			fmt.Println("Erro: endereço inválido. Apenas 0 e 1 são permitidos, com 3 bits.")
			return true
		}

		// Convertendo o endereço para posição
		posicao, _ := strconv.ParseInt(endereco, 2, 64)

		// Limpa o dado na posição especificada
		memoria[posicao] = "00000"
		fmt.Printf("Memória[%d] foi limpa!\n", posicao)
		mostraMemoria()
	case "1011":
		// Comando 1011: Limpar toda a memória
		limparArray()
	default:
		fmt.Println("Comando inválido!")
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Solicita cada parte da instrução separadamente
		fmt.Print("Digite o comando (4 bits): ")
		scanner.Scan()
		comando := scanner.Text()

		// Valida o comando (deve ser 4 bits binários)
		if !validaBinario(comando, 4) {
			fmt.Println("Erro: comando inválido. Apenas 0 e 1 são permitidos, com 4 bits.")
			continue
		}

		// O comando 0000 não precisa de mais nada (nem endereço nem dado)
		if comando == "0000" {
			instrucao := comando + "000" + "00000" // Preenche com valores fictícios para completar a instrução
			if len(instrucao) != 12 {
				fmt.Println("Instrução inválida! Deve ter 12 bits.")
				continue
			}
			// Executa o comando de encerramento
			if !executaInstrucao(instrucao) {
				break // Encerra o programa
			}
			continue
		}

		if comando == "1011" {
			instrucao := comando + "000" + "00000"
			if len(instrucao) != 12 {
				fmt.Println("Instrução inválida! Deve ter 12 bits.")
				continue
			}
			executaInstrucao(instrucao)
			continue
		}

		// O comando 1001 não exige o endereço
		if comando == "1001" {
			// Para o comando 1001, podemos chamar diretamente a função compararArray sem precisar de endereço ou dado
			instrucao := comando + "000" + "00000" // Endereço fictício e dado fictício
			if len(instrucao) != 12 {
				fmt.Println("Instrução inválida! Deve ter 12 bits.")
				continue
			}
			executaInstrucao(instrucao)
			continue
		}

		// O comando 1010 exige um endereço
		if comando == "1010" {
			// Para o comando 1010, chamamos diretamente a execução de limparPosicao
			instrucao := comando + "000" + "00000" // Endereço fictício e dado fictício
			if len(instrucao) != 12 {
				fmt.Println("Instrução inválida! Deve ter 12 bits.")
				continue
			}
			executaInstrucao(instrucao)
			continue
		}

		// Solicitar o endereço para outros comandos
		fmt.Print("Digite o endereço (3 bits): ")
		scanner.Scan()
		endereco := scanner.Text()

		// Valida o endereço (deve ser 3 bits binários)
		if !validaBinario(endereco, 3) {
			fmt.Println("Erro: endereço inválido. Apenas 0 e 1 são permitidos, com 3 bits.")
			continue
		}

		// Para o comando 0001 (ler), não precisamos de dado
		dado := ""
		if comando != "0001" && comando != "1000" && comando != "1010" && comando != "1011" { // Comando de comparação 1000 e 1001 não necessitam de dado
			fmt.Print("Digite os dados (5 bits): ")
			scanner.Scan()
			dado = scanner.Text()

			// Valida o dado (deve ser 5 bits binários)
			if !validaBinario(dado, 5) {
				fmt.Println("Erro: dado inválido. Apenas 0 e 1 são permitidos, com 5 bits.")
				continue
			}
		}

		// Monta a instrução completa (12 bits)
		var instrucao string
		if comando == "1000" || comando == "1001" || comando == "1010" || comando == "1011" {
			// Comando 1000, 1001 e 1010 não precisam de dado
			instrucao = comando + endereco + "00000" // "00000" é um preenchimento para o dado
		} else {
			instrucao = comando + endereco + dado
		}

		// Valida a instrução (deve ter 12 bits)
		if len(instrucao) != 12 {
			fmt.Println("Instrução inválida! Deve ter 12 bits.")
			continue
		}

		if !executaInstrucao(instrucao) {
			break
		}
	}
}
