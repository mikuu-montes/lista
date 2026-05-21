<h1 align="center">Diccionario</h1>

El contenido de este repositorio es la implementación de un Diccionario.

Implementado mediante el uso de un Hash Abierto.

---
## Contenido:
- [Requisitos previos para el diccionario](#requisitos-previos-para-el-diccionario)
- [Pasos para levantar el diccionario](#pasos-para-levantar-el-diccionario)
- [Estructura del repo](#estructura-del-repo)
- [Enunciado del TP](#enunciado-del-tp)
- [Créditos](#créditos)
---

## Requisitos Previos para el diccionario:
- [Tener Git instalado](https://git-scm.com/book/es/v2/Inicio---Sobre-el-Control-de-Versiones-Instalaci%C3%B3n-de-Git)
- [Tener una implementación de una lista al mismo nivel que el repo](https://github.com/mikuu-montes/lista)
- Tanto la lista como el diccionario deben estar dentro de una carpeta llamada "tdas".

## Pasos para levantar el diccionario:

### 1- Clona el repositorio:

```
git clone git@github.com:mikuu-montes/diccionario.git

``` 
o 

```
git clone https:m//github.com/mikuu-montes/diccionario.git
``` 
(te va a pedir que coloques tu contraseña)

### 2- Moverse a la carpeta del proyecto:
```
cd diccionario
```


## Estructura del repo:
Explicamos que contiene cada archivo:
- ` diccionario.go `: En este se declaran las interfaces del diccionario y de su iterador externo.
- ` hash.go `: En este se implementan las primitivas de las interfaces antes declaradas.
- ` diccionario_test.go `: En este se encuentras las pruebas que verifican que lo implementado, funcione efectivamente como un diccionario.

Para usarse deben tener esta estructura de archivos:

```
tdas/
├── lista/  ← (repositorio a descargar)
└── diccionario/ ← (repositorio actual)
    ├── diccionario.go
    ├── diccionario_test.go
    └── hash.go

```

## Enunciado del TP:
- [Link al enunciado de la cátedra.](https://algoritmos-rw.github.io/algoritmos/tps/hash/)

## Créditos:
Este trabajo fue realizado como tarea para la materia de **Algoritmos y Estructuras de Datos**, cátedra **Buchwald**, en **FIUBA**, por las alumnas:
- **Montes Brisa Micaela**
- **Tejeda Carrasco Angie Rubi**