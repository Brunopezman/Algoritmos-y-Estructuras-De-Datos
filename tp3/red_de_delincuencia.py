import grafo
from funciones_de_grafos import *

class RedDeDelincuentes:

    def __init__(self, grafo):
        """
        Constructor 
        Recibe por parametro un grafo que representa las comunicaciones de delincuentes
        """
        self.grafoDelincuentes = grafo

    def minimoSeguimiento(self, origen, destino):
        """
        Imprime una lista con los delincuentes (su código identificador) con los cuáles vamos del delincuente origen al delincuente destino de la forma más rápida. En caso de no poder hacer el seguimiento (i.e. no existe camino), imprimir Seguimiento imposible.
        """
        vOrigen = self.grafoDelincuentes.obtenerVertice(origen)
        vDestino = self.grafoDelincuentes.obtenerVertice(destino)
        orden, camino = recorridoMinimoBfs(self.grafoDelincuentes, vOrigen, vDestino)
        if ((orden and camino) is None):
            return print("Seguimiento imposible")
        self.__visualizar_resultado(camino, ' -> ')

    def divulgarRumor(self, origen, distMax):
        """
        Imprime una lista con todos los delincuentes a los cuales les termina llegando un rumor que comienza en el delincuente pasado por parámetro, y a lo sumo realiza n saltos (luego, se empieza a tergiversar el mensaje), teniendo en cuenta que todos los delincuentes transmitirán el rumor a sus allegados.
        """
        vOrigen = self.grafoDelincuentes.obtenerVertice(origen)
        if (vOrigen == None): return print("El delincuente no existe")
        afectados = recorridoMinimoBfsMaximo(grafo, vOrigen, distMax)
        self.__visualizar_resultado(afectados[1:], ', ')

    def comunidades(self, n):
        """
        Imprime un listado de comunidades de al menos n integrantes.
        
        """
        comunidades = label_propagation(self.grafoDelincuentes)
        num_comunidad = 1
        for comunidad, integrantes in comunidades.items():
            if len(integrantes) >= n:
                print("Comunidad " + str(num_comunidad) + ":", end=" ")
                self.__visualizar_resultado(integrantes, ', ')
                num_comunidad += 1

    def mas_imp(self, cant):
        """
        Imprime, de mayor a menor importancia, los cant delincuentes más importantes.
        
        """
        mas_importantes = centralidad_aprox(self.grafoDelincuentes, cant)
        for i in range(cant - 1):
            if i < len(self.grafoDelincuentes) - 1:
                print(mas_importantes[i], end=", ")
        print(mas_importantes[-1])

    def persecucion(self, agentesEncubiertos, kMasImp):
        """
        Dado cada uno de los delincuentes pasados (agentes encubiertos), obtener cuál es el camino más corto para llegar desde alguno de los delincuentes pasados por parámetro, a alguno de los K delincuentes más importantes. En caso de tener caminos de igual largo, priorizar los que vayan 
        a un delincuente más importante.
        """
        caminoMin = recorrido_min_multi_origen_multi_destino(self.grafoDelincuentes, agentesEncubiertos, kMasImp)
        if (len(caminoMin) == 0):
            return print("La persecución no es posible")
        self.__visualizar_resultado(caminoMin, ' -> ')

    def cfc(self):
        """
        Imprime cada conjunto de vértices entre los cuales todos están conectados con todos.
        
        """
        componente_fuertemente_conexas = cfc(self.grafoDelincuentes)
        comp_contador = 1
        for componente in componente_fuertemente_conexas:
            print("CFC " + str(comp_contador) + ": ", end="")
            self.__visualizar_resultado(componente, ', ')
            comp_contador += 1

    def divulgarCiclo(self, delincuente):
        """
        Imprime un camino simple que empiece y termine en el delincuente pasado por parámetro,
        de largo n. En caso de no encontrarse un ciclo, imprimir No se 
        encontro recorrido.
        """
        vDelincuente = self.grafoDelincuentes.obtenerVertice(delincuente)
        if vDelincuente is None:
            print("El delincuente no existe")
            return

        ciclo = ciclo_mas_corto(self.grafoDelincuentes, vDelincuente)
        if ciclo is None:
            print("No se encontro recorrido")
            return

        self.__visualizar_resultado(ciclo, ' -> ')

    def __visualizar_resultado(self, lista, separador):
        """
        Visuliza por consola los elementos de una lista separados por el parametro 'separador'
        
        """
        for i in range(0, len(lista) - 1):
            print(lista[i], end=separador)
        print(lista[-1])
