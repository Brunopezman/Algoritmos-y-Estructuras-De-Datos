#!/usr/bin/python3
# coding=utf-8

import sys
from red_de_delincuencia import RedDeDelincuentes
from funciones_de_grafos import *

sys.setrecursionlimit(10000)

# Constantes para los comandos del menú
MIN_SEGUIMIENTOS = "min_seguimientos"
DELINCUENTES_MAS_IMP = "mas_imp"
PERSECUCION = "persecucion"
COMUNIDADES = "comunidades"
DIVULGACION_DE_RUMOR = "divulgar"
DIVULGAR_CICLO = "divulgar_ciclo"
CFC = "cfc"

def crear_red_delincuencia(archivo):
    """
    Carga un grafo desde un archivo .tsv
    """
    grafo = Grafo()
    with open(archivo, 'r') as archivo_grafo:
        for linea in archivo_grafo:
            v, w = linea.strip().split('\t')
            grafo.agregarArista(v, w)
    return grafo

def procesar_comando(redDelincuencia, instruccion):
    instruccion = instruccion.strip()
    if not instruccion:
        return

    orden = instruccion.split()
    comando = orden[0]
    if len(orden) > 1:
        parametros = orden[1:]
    else:
        parametros = []

    if comando == MIN_SEGUIMIENTOS:
        origen = parametros[0]
        destino = parametros[1]
        redDelincuencia.minimoSeguimiento(origen, destino)

    elif comando == DELINCUENTES_MAS_IMP:
        cant = int(parametros[0])
        redDelincuencia.mas_imp(cant)

    elif comando == PERSECUCION:
        agentesEncubiertos = parametros[0].split(",")
        kMasImportantes = int(parametros[1])
        redDelincuencia.persecucion(agentesEncubiertos, kMasImportantes)

    elif comando == COMUNIDADES:
        min_integrantes = int(parametros[0])
        redDelincuencia.comunidades(min_integrantes)

    elif comando == DIVULGACION_DE_RUMOR:
        delincuente = parametros[0]
        distMax = int(parametros[1])
        redDelincuencia.divulgarRumor(delincuente, distMax)

    elif comando == DIVULGAR_CICLO:
        delincuenteCiclo = parametros[0]
        redDelincuencia.divulgarCiclo(delincuenteCiclo)

    elif comando == CFC:
        redDelincuencia.cfc()

    else:
        print(f"Comando desconocido: {comando}")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python comandos.py <archivo_grafo>")
        sys.exit(1)
    
    archivo_grafo = sys.argv[1]
    grafo = crear_red_delincuencia(archivo_grafo)
    redDelincuencia = RedDeDelincuentes(grafo)

    try:
        while True:
            instruccion = input()
            if instruccion == "":
                break
            procesar_comando(redDelincuencia, instruccion)
    except EOFError:
        pass
