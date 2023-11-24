package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Feriados struct {
	Fecha   string `json:"fecha"`
	Detalle string `json:"detalle"`
}

func findString(arreglo []string, cadena string) int {
	for i := 0; i < len(arreglo); i++ {
		if arreglo[i] == cadena {
			return i
		}
	}
	return 0
}

func getFeriados(anio string) []Feriados {
	// Descarga la web mediante un GET request
	apiUrl := fmt.Sprintf("https://www.argentina.gob.ar/interior/feriados-nacionales-%s", anio)
	request, error := http.NewRequest("GET", apiUrl, nil)

	if error != nil {
		fmt.Println(error)
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)

	// Libera la memoria
	defer response.Body.Close()

	if error != nil {
		fmt.Println(error)
	}

	strResponseBody := string(responseBody)
	archivos(strResponseBody)

	splitted := strings.Split(strResponseBody, "{")

	var struct_feriados Feriados
	var array_of_struct []Feriados
	var Fecha string
	var Detalle string
	var data []Feriados

	for i := 0; i < len(splitted); i++ {
		boolDate := strings.Contains(splitted[i], "\"date\":")
		boolLabel := strings.Contains(splitted[i], "\"label\":")

		if boolDate == true && boolLabel == true {

			array_fecha := strings.Split(strings.TrimRight(splitted[i], " "), " ")
			Fecha = strings.TrimLeft(strings.TrimRight(array_fecha[2], "\","), "\"")

			type_pos := findString(array_fecha, "\"type\":")
			if type_pos >= 4 {
				Detalle = strings.TrimLeft(strings.TrimRight(strings.Join(array_fecha[4:type_pos], " "), "\","), "\"")

				if !((strings.Contains(Detalle, "(a)")) || (strings.Contains(Detalle, "(b)")) || (strings.Contains(Detalle, "(c)"))) { // Si no es un feriado festivo se descarta
					struct_feriados = Feriados{
						Fecha:   Fecha,
						Detalle: Detalle,
					}
					array_of_struct = append(array_of_struct, struct_feriados)
					data = append(data, struct_feriados)
				}
			}
			Detalle = ""
			Fecha = ""
		}
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = os.WriteFile("feriados.json", file, 0644)

	return (array_of_struct)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func archivos(val string) {

	data := []byte(val)
	err := os.WriteFile("filedump.html", data, 0644)
	check(err)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Obteniendo Feriados...")
	feriados := getFeriados(os.Args[1])
	for i := 0; i < len(feriados); i++ {
		fmt.Printf("Fecha: %v \n", feriados[i].Fecha)
		fmt.Printf("Detalle: %v \n", feriados[i].Detalle)
	}
}
