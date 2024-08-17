class Vertice:

    def __init__(self, clave):
        """
        Constructor de clase
        Pre: Recibe por parametro una clave utilizada con id para el vertice.
        Posee un diccionario con los vertices a los cuales se encuentra conectado por una arista
        """
        self.id = clave
        self.conectadoA = {}

    def __str__(self):
        return str(self.id)

    def agregarVecino(self, vecino, ponderacion=0):
        """
        Añade una arista dirigida hacia un vertice.
        Pre: Recibe por parametro el vertice vecino y una ponderacion para la arista, por default 0.
        """
        self.conectadoA[vecino] = ponderacion

    def obtenerAdyacentes(self):
        """
        Retorna una lista con los vertices adyacentes.
        
        """
        return list(self.conectadoA.keys())

    def obtenerId(self):
        """
        Retorna la Id del vertice en cuestion
        
        """
        return self.id

    # Retorna la ponderacion de la arista dirigida del vertice en cuestion hacia el vertice vecino.
    def obtenerPonderacion(self, vecino):
        return self.conectadoA[vecino]

class Grafo:

    def __init__(self):
        self.listaVertices = {}
        self.numVertices = 0

    def __contains__(self, vertice):
        return vertice in self.listaVertices

    def __iter__(self):
        return iter(self.listaVertices.values())

    def __len__(self):
        return self.numVertices

    # Agrega un vertice al grafo
    # Pre: Recibe por parametro un clave como id para el vertice nuevo
    # Post: Retorna el vertice creado
    def agregarVertice(self, clave):
        self.numVertices += 1
        nuevoVertice = Vertice(clave)
        self.listaVertices[clave] = nuevoVertice
        return nuevoVertice

    # Retorna el vertice asociado a idVertice
    # Pre: Recibe un idVertice
    # Post: Retorna el vertice, None en caso que la id no pertenezca al grafo.
    def obtenerVertice(self, idVertice):
        if idVertice in self.listaVertices:
            return self.listaVertices[idVertice]
        else:
            return None

    # Añade una arista dirigado y opcionalmente ponderada 
    # Pre: Recibe por parametro el id del vertice origen y destino. Opcionalmente
    # una ponderacion
    # Post: Añade un arista dirigida desde el origen hacia el destino. Si los vertices no
    # existen, son creados. 
    def agregarArista(self, idOrigen, idDestino, ponderacion=0):
        if idOrigen not in self.listaVertices:
            self.agregarVertice(idOrigen)
        if idDestino not in self.listaVertices:
            self.agregarVertice(idDestino)
        self.listaVertices[idOrigen].agregarVecino(self.listaVertices[idDestino], ponderacion)

    # Retorna una lista de todos los vertices del arbol
    def obtenerVertices(self):
        return list(self.listaVertices.values())