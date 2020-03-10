//funcion para determinar el determinante de una funcion de segundo grado
//jose carlos 201504231
package main

import ("fmt"
"strconv"
"math"
"math/rand"
"time"
)


func determinante2x2(coordenada00 int , coordenada01 int ,coordenada10 int ,coordenada11 int){
    determinante:= 0
    determinante = coordenada00 * coordenada11 - coordenada10 * coordenada11
    if determinante == 0{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales y diferentes")
    }else if determinante > 0{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales e iguales ")
    }else{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones imaginaria")
    }
}

func determinante3x2(coordenada000 int , coordenada010 int ,coordenada001 int, coordenada100 int, coordenada110 int, coordenada111 int ){
  determinante:= 0
  determinante = coordenada000 * coordenada010 * coordenada001 - coordenada100 * coordenada110 * coordenada111
  if determinante == 0{
      fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales y diferentes")
  }else if determinante > 0{
      fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales e iguales ")
  }else{
      fmt.Println(strconv.Itoa(determinante)+ "Soluciones imaginaria")
  }
}

func absolute(number float64) {
  number = math.Abs(number);
  fmt.Println("El valor absoluto del numero es: ", number)
}

func randomize(min int, max int) {
  rand.Seed(time.Now().UnixNano())
  fmt.Print("El numero random es: ")
  fmt.Println(rand.Intn(max - min + 1) + min)
}

func raizCubica(num float64) {
  val := math.Cbrt(num)
  s := fmt.Sprintf("%.2f", val)

	fmt.Println("La raiz cubica de", num, "es", s)
}

func potenciaCubo(num float64) {
	fmt.Println(num, "elevado a 3 es:", math.Pow(num, 3))
}

func mult(factor1 int,factor2 int) int{
	return factor1*factor2
}

func div(dividend int,divider int) string{
  if divider==0 { return "Indefinido" }
	return strconv.Itoa(dividend/divider)
}

func decimalABinario(dec int){
  val := int64(dec)

  fmt.Printf("El número %d en binario es '%s'\n", dec, strconv.FormatInt(val, 2))
}

func menu(){
  fmt.Println("Bienvenido al programa")
  fmt.Println("A continuacion ingrese una opcion:")
  fmt.Println("1. Calculo de determinante---------")
  fmt.Println("2. Absoluto de un numero-----------")
  fmt.Println("3. Cubo de un numero---------------")
  fmt.Println("4. Raiz cubica de un numero--------")
  fmt.Println("5. Multiplicacion de dos numeros---")
  fmt.Println("6. Division de entre numeros-------")
  fmt.Println("7. Salir del programa--------------")
  fmt.Println("8. Numero Random")
  fmt.Println("9. decimal a binario--------------")
}

func main() {
    fmt.Println("Antes de empezar este es un cambio, el bugfix 6.2")
    opcion := ""
    for {
      menu();
      fmt.Scan(&opcion)
      switch opcion {
      case "1":
        fmt.Println("Ingrese la coordenada 00")
        fmt.Println("Ingrese la coordenada 01")
        fmt.Println("Ingrese la coordenada 10")
        fmt.Println("Ingrese la coordenada 11")
      case "2":
        var number float64
        fmt.Println("Ingrese el numero")
        fmt.Scan(&number)
        absolute(number)
      case "3":
        fmt.Println("Ingrese el numero")
      case "4":
        fmt.Println("Ingrese el numero")
      case "5":
        fmt.Println("Ingrese el primer factor")
        var f1 int
        fmt.Scan(&f1)
        fmt.Println("Ingrese el segundo factor")
        var f2 int
        fmt.Scan(&f2)
        fmt.Print("El producto es: ")
        fmt.Println(mult(f1,f2))
      case "6":
        fmt.Println("Ingrese el dividendo")
        var dividend int
        fmt.Scan(&dividend)
        fmt.Println("Ingrese el divisor")
        var divider int
        fmt.Scan(&divider)
        fmt.Print("El cociente es: ")
        fmt.Println(div(dividend,divider))
      case "7":
        fmt.Println("Saliendo del programa")
      case "8":
        fmt.Println("Ingrese el minimo")
        var min int
        fmt.Scan(&min)
        fmt.Println("Ingrese el maximo")
        var max int
        fmt.Scan(&max)
        randomize(min, max)
      case "9":
        fmt.Println("Ingrese un numero: ")
        var num_ int
        fmt.Scan(&num_)
        decimalABinario(num_)
      default:
        fmt.Println("Opcion incorrecta")
      }
      if opcion=="7" { break }
    }
}
