//funcion para determinar si un numero es primo o no
//Brandon Pedroza 201314079
package main

import ("fmt"
"strconv"
)

func validarPrimo(numero int) {
    contador:= 0 
    for i := 1; i <= numero; i++{
        if numero % i  == 0{
            contador++
        }
    }
    if contador<=2{ 
        fmt.Println(strconv.Itoa(numero)+ " Es primo")
    } else {
        fmt.Println(strconv.Itoa(numero) + " NO es primo")
    }
}

func main() {
    validarPrimo(65)
    validarPrimo(97)
}


