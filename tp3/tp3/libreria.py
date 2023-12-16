import heapq
import csv
from grafo import Grafo
from union_find import UnionFind


def grados(grafo):
    g_entrada, g_salida = {}, {}
    for v in grafo.obtener_vertices():
        g_entrada[v] = 0
        g_salida[v] = 0
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            g_entrada[w] += 1
            g_salida[v] += 1
    return g_entrada, g_salida


def camino_minimo(grafo, origen, destino):
    distancia = {}
    padres = {}
    for v in grafo.obtener_vertices():
        distancia[v] = float("inf")
    distancia[origen] = 0
    padres[origen] = None
    q = []
    heapq.heappush(q, (0, origen))
    while len(q) != 0:
        _, v = heapq.heappop(q)
        if v == destino:
            return padres, distancia
        for w in grafo.adyacentes(v):
            distancia_por_aca = distancia[v] + grafo.peso(v, w)
            if distancia_por_aca < distancia[w]:
                distancia[w] = distancia_por_aca
                padres[w] = v
                heapq.heappush(q, (distancia[w], w))
    return None, None


def reconstruir_camino(padres, destino):
    recorrido = []
    while destino is not None:
        recorrido.append(destino)
        destino = padres[destino]
    return recorrido[::-1]


def orden_topologico(grafo):
    g_ent, _ = grados(grafo)
    cola = []
    for v in grafo.obtener_vertices():
        if g_ent[v] == 0:
            cola.append(v)
    res = []
    while len(cola) != 0:
        v = cola.pop(0)
        res.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                cola.append(w)
    if len(res) == len(grafo.obtener_vertices()):
        return res
    return None


def vertices_restantes(grafo1, grafo2):
    """
    Devuelve los elementos del grafo 1 que no se encuentran en el grafo 2
    """
    if len(grafo1.obtener_vertices()) == len(grafo2.obtener_vertices()):
        return None
    vertices_grafo1, vertices_grafo2 = grafo1.obtener_vertices(), grafo2.obtener_vertices()
    res = []
    for elemento in vertices_grafo1:
        if elemento not in vertices_grafo2:
            res.append(elemento)
    return res


def generar_grafo_recomendado(archivo):
    nuevo_grafo = Grafo(True)
    with open(archivo) as archivo:
        archivo = csv.reader(archivo, delimiter=",")
        for fila in archivo:
            nuevo_grafo.agregar_vertice(fila[0])
            nuevo_grafo.agregar_vertice(fila[1])
            nuevo_grafo.agregar_arista(fila[0], fila[1])
    return nuevo_grafo


def mst_kruskal(grafo):
    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo), key=lambda arista: arista[2])
    arbol = Grafo()
    for v in grafo.obtener_vertices():
        arbol.agregar_vertice(v)
    for a in aristas:
        v, w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol


def obtener_aristas(grafo):
    aristas = []
    visitados = set()
    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v, w, grafo.peso(v, w)))
        visitados.add(v)
    return aristas


def peso_total_grafo(aristas):
    peso_total = 0
    for a in aristas:
        peso_total += a[2]
    return peso_total


def tiene_ciclo_euleriano(grafo):
    g_ent, _ = grados(grafo)
    contador = 0
    for clave in g_ent:
        if g_ent[clave] % 2 != 0:
            contador += 1
    return contador == 0 or contador == 2


def recorrido_hierholzer(grafo, origen):
    if tiene_ciclo_euleriano(grafo) and origen in grafo.obtener_vertices() and es_conexo(grafo):
        visitados = set()
        aristas = obtener_aristas(grafo)
        camino = []
        distancia = _dfs_hierholzer(grafo, origen, visitados, origen, camino, aristas)
        return camino, distancia
    return None, None


def _dfs_hierholzer(grafo, v, visitados, origen, camino, aristas):
    suma = 0
    for w in grafo.adyacentes(v):
        if len(camino) == len(aristas):
            break
        if (v, w) in visitados:
            continue
        visitados.add((v, w))
        visitados.add((w, v))
        suma += grafo.peso(v, w)
        suma += _dfs_hierholzer(grafo, w, visitados, origen, camino, aristas)
    camino.append(v)
    return suma


def es_conexo(grafo):
    visitados = set()
    cant = 0
    for v in grafo.obtener_vertices():
        if v not in visitados:
            cant += 1
            _dfs_es_conexo(grafo, v, visitados)
    if cant > 1:
        return False
    return True


def _dfs_es_conexo(grafo, v, visitados):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            _dfs_es_conexo(grafo, w, visitados)


def existe_vertice(grafo, vertice):
    for v in grafo.obtener_vertices():
        if v == vertice:
            return True
    return False
