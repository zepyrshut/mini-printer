# Mini Printer Server

## Descripción
Un sencillo servidor en el que proporciona una interfaz para imprimir tickets numéricos. Orientado a las comandas de
los restaurantes. Estos tickets irán adjuntado al plato que corrasponda en el pase de cocina.

## Instalación
Clonar el repositorio en una carpeta y ejecutar el comando `go run main.go` en la carpeta raíz del proyecto.

## Uso
Conectarse al servidor, a través del puerto `8082` y acceder a la ruta `/` para imprimir un ticket. No tiene más.

## Impresoras compatibles
Cualquiera que soporten el formato ESC/POS en sistemas Windows. Se ha probado con una impresora de 80 mm.

No se han probado en otros sistemas como Linux o Mac ni en impresoras de 58 mm.

## Este proyecto no hubiese sido posible sin las siguientes aportaciones
- [alexbrainman/printer](https://github.com/alexbrainman/printer) - Librería de impresión Windows.
- [go-chi/chi](https://github.com/go-chi/chi) - Router HTTP.
- [simple-keyboard](https://github.com/simple-keyboard) - Teclado virtual en pantalla.
- [BigJk/snd](https://github.com/BigJk/snd) - Todo el código relacionado con la transformación de plantilla HTML
 a imagen para su impresión se debe gracia a este usuario y este fantástico proyecto.