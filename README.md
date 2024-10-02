# ReadMe - Obtener Feriados de Argentina

Este script desarrollado en Go, obtiene los feriados festivos nacionales de Argentina de un año específico desde el sitio web oficial del gobierno (https://www.argentina.gob.ar/interior/). El código realiza una solicitud HTTP a la URL del sitio, luego extrae y almacena los datos en un archivo JSON.

## Uso

``` go run main.go 2023 ```


Asegúrate de tener Go instalado en tu sistema antes de ejecutar el código. Si no lo tienes instalado, puedes descargarlo [aquí](https://golang.org/dl/).

1. Descarga el código fuente o clónalo desde este repositorio.
2. Abre una terminal y navega hasta el directorio donde se encuentra el código.
3. Ejecuta el programa proporcionando el año del que deseas obtener los feriados como argumento. 
4. El programa realizará la solicitud HTTP, analizará los datos, generará un archivo JSON llamado "feriados.json" que contiene la información de los feriados y otro llamado "feriados.txt" el cual contiene solo las fechas. También guardará una copia de la página web en formato HTML como "filedump.html".


