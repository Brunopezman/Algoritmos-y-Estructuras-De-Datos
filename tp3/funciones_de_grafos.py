from io import open
from random import choice
from math import inf
from grafo import *
from collections import deque
from collections import Counter

def bfs(grafo, origen, destino=None, ordenMax=None):
    """
    Recorrido tipo "Breadth First Search" 
    Pre: Recibe como parametro el grafo dirigido y no ponderado, y el vector origen desde el cual se 
    desea iniciar el recorrido. Opcionalmente puede recibir un vector destino y/o un orden maximo a fin de 
    finalizar el recorrido una vez alcanzada el destino y/o orden maximo.
    Post: Retorna un diccionario con el orden de cada vertice en el grafo y un diccionario con
    el padre correspondiente a cada vertice, de modo de poder reconstruir el recorrido.
    """
    visitados = set()
    padre = {}
    orden = {}
    cola = deque()
    visitados.add(origen)
    padre[origen] = None
    orden[origen] = 0
    cola.append(origen)
    while (cola):
        vertice = cola.popleft()

        if (destino != None and vertice == destino):
            return padre, orden
        if (ordenMax and ordenMax == orden[vertice]):
            return padre, orden

        for adyacente in vertice.obtenerAdyacentes():
            if adyacente not in visitados:
                visitados.add(adyacente)
                padre[adyacente] = vertice
                orden[adyacente] = orden[vertice] + 1
                cola.append(adyacente)

    return padre, orden

def recorridoMinimoBfs(grafo, origen, destino, ordenMax=None):

    """Recorrido minimo con BFS con destino vertice.
    Pre: Recibe por parametro un grafo dirigido y no ponderado, vertice origen y vertice destino.
    Post: Retorna el orden del vertice destino y una lista con el camino correspondiente
    de origen a destino."""
    padres, orden = bfs(grafo, origen, destino, ordenMax)
    if (destino not in padres): return None, None
    camino = []
    vertice = destino
    camino.append(vertice)
    while (padres[vertice] != None):
        camino.append(padres[vertice])
        vertice = padres[vertice]
    return orden[destino], camino[::-1]

def recorridoMinimoBfsMaximo(grafo, origen, distMax):
    """
    Recorrido radial con BFS de n distancia maxima.
    Pre: Recibe por parametro un grafo dirigido y no ponderado, vertice origen y una distancia maxima.
    Post: Retorna una lista con los vertices afectados distMax un radio n del origen
    """
    padres, orden = bfs(grafo, origen, None, distMax)
    verticesAfectados = []
    for key in padres.keys():
        verticesAfectados.append(key)
    return verticesAfectados


def label_propagation(grafo, n=3):
    """
    Comunidades en forma aproximada a traves de label_propagation
    Pre: Recibe por parametro un grafo dirigido y no ponderado, opcionamente un 'n' iteraciones, por default 3,
    esta ultima permite ajustar el numero de iteraciones, a mayor n mas compactas seran las comunidades
    Post: Retorna una lista de comunidades con sus respectivos vertices.
    """
    label = {}
    i = 1
    for vertice in grafo:
        label[vertice] = i
        i += 1
    for i in range(n):
        for vertice in grafo:
            if vertice.obtenerAdyacentes():
                label[vertice] = max_freq(label, vertice.obtenerAdyacentes())

    comunidades = {}
    for vertice, comunidad in label.items():
        if comunidad not in comunidades:
            comunidades[comunidad] = []
        comunidades[comunidad].append(vertice)

    return comunidades


def max_freq(label, adyacentes):
    """
    Retorna el elemento mas repetido de una lista de elementos. Funcion auxiliar label_propagation
    """
    labelAdyacentes = {}
    for adyacente in adyacentes:
        if label[adyacente] in labelAdyacentes:
            labelAdyacentes[label[adyacente]] += 1
        else:
            labelAdyacentes[label[adyacente]] = 1
    label_mas_repetido = None
    num_repeticiones_label = 0
    for label in labelAdyacentes:
        if labelAdyacentes[label] > num_repeticiones_label:
            label_mas_repetido = label
            num_repeticiones_label = labelAdyacentes[label]
    return label_mas_repetido

def centralidad_aprox(grafo, cantidad):
    """
    Centralidad aproximada, a traves de random_walks, de n "cantidad" de vertices
    Pre: Recibe por parametro un grafo dirigido y no ponderado, n 'cantidad' de vertices centrales
    que se buscan
    Post: Retorna una lista de los n "cantidad" de vertices mas centrales de mayor a menor
    """
    caminos = random_walks(grafo)
    recorridoTotal = []
    for camino in caminos:
        for vertice in camino:
            recorridoTotal.append(vertice)

    verticesCentrales = []
    candidatosCentrales = Counter(recorridoTotal).most_common()
    for i in range(cantidad):
        if i < len(grafo) - 1:
            verticesCentrales.append(candidatosCentrales[i][0])
    return verticesCentrales

def random_walks(grafo, longitudCamino=500, cantidadCaminos=100):
    """
    Randon walks - Realiza caminos aleatoreos a lo largo del grafo.
    Pre: Recibe por parametro un grafo dirigido y no ponderado, una longitudCamino y cantidadCaminos
    Post: Retorna un lista de n 'cantidadCaminos' de longitud 'longitudCaminos' con
    todos los vertices que componen al camino  
    """
    caminos = []
    vertices = grafo.obtenerVertices()
    for i in range(cantidadCaminos):
        camino = []
        verticeActual = choice(vertices)
        camino.append(verticeActual)
        for j in range(longitudCamino):
            adyacentes = verticeActual.obtenerAdyacentes()
            if adyacentes:
                verticeActual = choice(verticeActual.obtenerAdyacentes())
                camino.append(verticeActual)
        caminos.append(camino)

    return caminos

def recorrido_min_multi_origen_multi_destino(grafo, idVerticesOrigen, kMasImp):
    """
    Recorrido minimo de tipo BFS, con multiples origenes hacia 'kMasImp' mas centrales
    Pre: Recibe por parametro un grafo dirigido y no ponderado, una lista de vertices origen y
    un int 'kMasImp' vertices de mayor Centalidad. 
    Post: Retorna el camino de menor longitud, desde uno de los vertices origen hacia uno de los
    "kMasImp" mas centrales. Optando a estos, orgien y destino, del modo mas conveniente a la finalidad
    del camino mas corto
    """
    kVerticesMasImp = centralidad_aprox(grafo, kMasImp)
    verticesOrigen = []
    for idVerticeOrigen in idVerticesOrigen:
        vertice = grafo.obtenerVertice(idVerticeOrigen)
        if (vertice != None):
            verticesOrigen.append(vertice)
    caminoMin = []
    if len(verticesOrigen) == 0: return caminoMin
    ordenMin = inf
    for verticeOrigen in verticesOrigen:
        for kVerticeMasImp in kVerticesMasImp:
            orden, camino = recorridoMinimoBfs(grafo, verticeOrigen, kVerticeMasImp, ordenMin)
            if orden and orden < ordenMin:
                ordenMin = orden
                caminoMin = camino

    return caminoMin


def dfs_cfc(grafo, v, visitados, orden, p, s, cfcs, en_cfs):
    """
    Componente fuertemente conexo - Algoritmo de Tarjan  
    """
    visitados.add(v)
    s.append(v)
    p.append(v)
    for w in v.obtenerAdyacentes():
        if w not in visitados:
            orden[w] = orden[v] + 1
            dfs_cfc(grafo, w, visitados, orden, p, s, cfcs, en_cfs)
        elif w not in en_cfs:
            while p and orden[p[-1]] > orden[w]:
                p.pop()

    if p and p[-1] == v:
        p.pop()
        z = None
        nueva_cfc = []
        while z != v:
            z = s.pop()
            en_cfs.add(z)
            nueva_cfc.append(z)
        cfcs.append(nueva_cfc)

def cfc(grafo):
    """
    Componente fuertemente conexo - Algoritmo de Tarjan
    """
    visitados = set()
    orden = {}
    p = []
    s = []
    cfcs = []
    en_cfs = set()
    for v in grafo:
        if v not in visitados:
            orden[v] = 0
            dfs_cfc(grafo, v, visitados, orden, p, s, cfcs, en_cfs)
    return cfcs


def ciclo_mas_corto(grafo, origen):
    ciclo = bfs_ciclo(grafo, origen, origen)
    if ciclo is not None:
        return ciclo
    else:
        return None

def bfs_ciclo(grafo, origen, v):
    dist = {origen: 0}
    padre = {origen: None}
    cola = deque([origen])
    while cola:
        v = cola.popleft()
        for w in v.obtenerAdyacentes():  
            if w not in dist:
                dist[w] = dist[v] + 1
                padre[w] = v
                cola.append(w)
            elif w == origen:
                return reconstruir_ciclo(padre, v, origen)
    return None

def reconstruir_ciclo(padre, fin, origen):
    ciclo = []
    ciclo.append(origen)
    actual = fin
    while actual is not None:
        ciclo.append(actual)
        actual = padre[actual]
    return ciclo[::-1]