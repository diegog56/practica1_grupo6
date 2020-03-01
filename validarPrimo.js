//funcion para determinar si un numero es primo o no
//Brandon Pedroza 201314079
function esPrimo(numero) {
    let contador= 0; 
    for(let i = 1; i <= numero; i++)
    {
        if((numero % i) == 0)
        {
            contador++;
        }
    }
    if (contador<=2) alert('es primo');
    else alert('NO es primo');
}

