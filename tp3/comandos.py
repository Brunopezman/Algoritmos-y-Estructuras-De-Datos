import sys
from red_de_delincuencia import *
from funciones_de_grafos import *
from imprimir import *
from grafo import Grafo

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
    Carga un grafo
    Pre: Recibe por parámetro un fichero .tsv en formato:
    id_v_origen   id_vertice_destino
    Post: Retorna un grafo con los vértices y aristas contenidas en el archivo
    """
    grafo = Grafo()
    with open(archivo, 'r') as archivo:
        for linea in archivo:
            v, w = linea.strip().split('\t')
            grafo.agregar_arista(v, w)
    return grafo

def comandos(redDelincuencia, instruccion):
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
        largoCiclo = int(parametros[1])
        redDelincuencia.divulgar_ciclo(delincuenteCiclo, largoCiclo)

    elif comando == CFC:
        redDelincuencia.cfc()

    else:
        print(f"Comando desconocido: {comando}")
