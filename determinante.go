//funcion para determinar el determinante de una funcion de segundo grado
//jose carlos 201504231
package main

import ("fmt"
"strconv"
)


func determinante2x2(coordenada00 int , coordenada01 int ,coordenada10 int ,coordenada11 int){
    determinante:= 0
    determinante:= coordenada00 * coordenada11 - coordenada10 * coordenada11;
    if determinante == 0{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales y diferentes")
    }else if determinante > 0{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones Reales e iguales ")
    }else{
        fmt.Println(strconv.Itoa(determinante)+ "Soluciones imaginaria")
    }
}

func main() {
    determinante2x2(4,2,2,2)
}