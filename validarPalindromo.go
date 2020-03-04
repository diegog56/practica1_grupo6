//funcion para determinar si un cadena es palindromo o no
//Brandon Pedroza 201314079
package main

import ("fmt"

)

func validadPalindromo(cadena string) {
	inc := 0
	des := len(cadena)-1
	banderaError := false
	for inc<des && !banderaError{
			
	if (cadena[inc]==cadena[des]){				
		inc++
		des--
	}else if cadena [inc]==' '{
		inc++
	}else if cadena [des]==' '{
		des--
	} else {
		banderaError = true
	}
}
if (!banderaError){
	fmt.Println(cadena + " ES un PALINDROMO")
	}else{
	fmt.Println(cadena + " NO es un PALINDROMO")
	}
	
}

func main() {
    validadPalindromo("anita lava la tina")
    validadPalindromo("Mañana")
    validadPalindromo("amo la pacifica paloma")
    validadPalindromo("rompecabeza")
}