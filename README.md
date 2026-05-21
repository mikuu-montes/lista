<h1 align="center">Lista</h1>

El contenido de este repositorio es la implementación de una Lista.

Implementado mediante el uso de una Lista Enlazada.

---
## Contenido:
- [Requisitos previos para la lista](#requisitos-previos-para-la-lista)
- [Pasos para levantar la lista](#pasos-para-levantar-la-lista)
- [Estructura del repo](#estructura-del-repo)
- [Enunciado del TP](#enunciado-del-tp)
- [Créditos](#créditos)
---

## Requisitos Previos para la lista:
- [Tener Git instalado](https://git-scm.com/book/es/v2/Inicio---Sobre-el-Control-de-Versiones-Instalaci%C3%B3n-de-Git)
- [El repo debe estar dentro de una carpeta llamada "tdas".](#estructura-del-repo)

## Pasos para levantar la lista:

### 1- Clona el repositorio:

```
git clone git@github.com:mikuu-montes/lista.git

``` 
o 

```
git clone https://github.com/mikuu-montes/lista.git
``` 
(te va a pedir que coloques tu contraseña)

### 2- Moverse a la carpeta del proyecto:
```
cd lista
```


## Estructura del repo:
Explicamos que contiene cada archivo:
- ` lista.go `: En este se declaran las interfaces de la lista y de su iterador externo.
- ` lista_enlazada.go `: En este se implementan las primitivas de la interfaz de la lista.
- ` iterador_externo.go `: En este se implementan las primitivas de la interfaz del iterador externo.
- ` lista_test.go `: En este se encuentras las pruebas que verifican que lo implementado, funcione efectivamente como una lista.
- ` iterador_externo_test.go `: En este se encuentran las pruebas que verifican que el iterador externo de la lista funcione correctamente.
- ` iterador_interno_test.go `: En este se encuentran las pruebas que verifican que el iterador interno de la lista funcione correctamente.

Para usarse deben tener esta estructura de archivos:

```
tdas/
└── lista/ ← (repositorio actual)
    ├── lista.go
    ├── lista_enlazada.go
    ├── iterador_externo.go
    ├── lista_test.go
    ├── iterador_interno_test.go
    └── iterador_externo_test.go

```

## Enunciado del TP:
- [Link al enunciado de la cátedra.](https://algoritmos-rw.github.io/algoritmos/tps/lista/)

## Créditos:
Este trabajo fue realizado como tarea para la materia de **Algoritmos y Estructuras de Datos**, cátedra **Buchwald**, en **FIUBA**, por la alumna:
- **Montes Brisa Micaela**