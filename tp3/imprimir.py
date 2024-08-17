MAS_IMP = "mas_imp"
DIVULGAR = "divulgar"
PERSECUCION = "persecucion"

def imprimir_camino(padre, destino):
    if destino not in padre:
        print("Seguimiento imposible")
        return
    camino = []
    act = destino 
    while act is not None:
        camino.append(act)
        act = padre[act]
    camino.reverse()

    print(" -> ".join(camino))

def imprimir_mas_imp(lista, comando):
    if lista:
        if comando == PERSECUCION:
            print(" -> ".join(lista))
        elif comando == MAS_IMP:
            print(", ".join(map(str, lista[:-1]))+ ", " + str(lista[-1])) 
        elif comando == DIVULGAR:
            print(", ".join(map(str, lista)))

def imprimir_comunidades(lista):
    for i, c in enumerate(lista, start=1):
        print(f"CFC {i}:", end="")
        print(", ".join(map(str, c)))
              

def imprimir_ciclo(ciclo):
    if ciclo is None:
        print("No se encontro recorrido")
    else:
        for i, v in enumerate(ciclo):
            if i < len(ciclo) - 1:
                print(f"{v} -> ", end="")
            else:
                print(v)
