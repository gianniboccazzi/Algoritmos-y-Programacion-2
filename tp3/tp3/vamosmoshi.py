#!/usr/bin/python3
import sys
import csv
import libreria
from grafo import Grafo
sys.setrecursionlimit(4000)

def cargar_en_memoria(archivo):
    grafo = Grafo()
    coordenadas = {}
    with open(archivo) as archivo:
        archivo = csv.reader(archivo, delimiter=",")
        contador_lineas = 0
        for fila in archivo:
            if contador_lineas == 0:
                cant_ciudades = int(fila[0])
            if 0 < contador_lineas <= cant_ciudades:
                coordenadas[fila[0]] = (fila[1], fila[2])
                grafo.agregar_vertice(fila[0])
            if cant_ciudades + 1 < contador_lineas:
                grafo.agregar_arista(fila[0], fila[1], int(fila[2]))
            contador_lineas += 1
    return grafo, coordenadas


def crear_archivo_kml(ruta, recorrido, coordenadas):
    f = open(ruta, "w")
    f.write('<?xml version="1.0" encoding="UTF-8"?>\n')
    f.write('<kml xmlns="http://earth.google.com/kml/2.1">\n')
    f.write('\t<Document>\n')
    f.write(f'\t\t<name>Camino desde {recorrido[0]} hacia {recorrido[len(recorrido) - 1]}</name>\n\n')
    visitadas = set()
    for ciudad in recorrido:
        if ciudad not in visitadas:
            f.write('\t\t<Placemark>\n'
                    f'\t\t\t<name>{ciudad}</name>\n'
                    '\t\t\t<Point>\n'
                    f'\t\t\t\t<coordinates>{coordenadas[ciudad][0]}, {coordenadas[ciudad][1]}</coordinates>\n'
                    '\t\t\t</Point>\n'
                    '\t\t</Placemark>\n')
        visitadas.add(ciudad)
    for i in range(len(recorrido) - 1):
        f.write('\t\t<Placemark>\n'
                '\t\t\t<LineString>\n'
                f'\t\t\t\t<coordinates>{coordenadas[recorrido[i]][0]}, {coordenadas[recorrido[i]][1]} '
                f'{coordenadas[recorrido[i+1]][0]}, {coordenadas[recorrido[i+1]][1]}</coordinates>\n'
                '\t\t\t</LineString>\n'
                '\t\t</Placemark>\n')
    f.write('\t</Document>\n')
    f.write('</kml>')
    f.close()


def crear_archivo_pajek(grafo, ruta, aristas, coordenadas):
    f = open(ruta, "w")
    cant_vertices = len(grafo.obtener_vertices())
    cant_aristas = len(aristas)
    f.write(f"{cant_vertices}\n")
    for ciudad in grafo.obtener_vertices():
        f.write(f"{ciudad},{coordenadas[ciudad][0]},{coordenadas[ciudad][1]}\n")
    f.write(f"{cant_aristas}\n")
    for a in aristas:
        f.write(f"{a[0]},{a[1]},{a[2]}\n")
    f.close()


def ejecutar_ir(grafo, coordenadas, ciudades, ruta):
    if not libreria.existe_vertice(grafo, ciudades[0]) or not libreria.existe_vertice(grafo, ciudades[1]):
        print("No se encontro recorrido")
        return
    padres, distancia = libreria.camino_minimo(grafo, ciudades[0], ciudades[1])
    if distancia[ciudades[1]] == float("inf"):
        print("No se encontro recorrido")
        return
    recorrido = libreria.reconstruir_camino(padres, ciudades[1])
    crear_archivo_kml(ruta, recorrido, coordenadas)
    print(" -> ".join(recorrido))
    print("Tiempo total:", distancia[ciudades[1]])


def ejecutar_itinerario(grafo, archivo):
    nuevo_grafo = libreria.generar_grafo_recomendado(archivo)
    itinerario = libreria.orden_topologico(nuevo_grafo)
    if itinerario is None:
        print("No se encontro recorrido")
    restantes = libreria.vertices_restantes(grafo, nuevo_grafo)
    if restantes is not None:
        for elemento in restantes:
            itinerario.append(elemento)
    print(" -> ".join(itinerario))


def ejecutar_viaje(grafo, coordenadas, ciudad, ruta):
    recorrido, distancia = libreria.recorrido_hierholzer(grafo, ciudad)
    if recorrido is not None and distancia is not None:
        crear_archivo_kml(ruta, recorrido, coordenadas)
        print(" -> ".join(recorrido))
        print("Tiempo total:", distancia)
    else:
        print("No se encontro recorrido")


def ejecutar_reducir_caminos(grafo, coordenadas, ruta):
    arbol = libreria.mst_kruskal(grafo)
    aristas = libreria.obtener_aristas(arbol)
    peso_total = libreria.peso_total_grafo(aristas)
    crear_archivo_pajek(grafo, ruta, aristas, coordenadas)
    print("Peso total:", peso_total)


def ejecutar_programa(grafo, coordenadas):
    for linea in sys.stdin:
        if linea == "":
            break
        entrada = linea.strip("\n")
        comando = entrada[:entrada.find(" ")]
        ciudades = entrada[entrada.find(" ") + 1:entrada.rfind(",")].split(", ")
        archivo = entrada[entrada.rfind(",") + 1:].strip(" ").split(" ")
        if comando == "ir":
            ejecutar_ir(grafo, coordenadas, ciudades, archivo[0])
        elif comando == "itinerario":
            ejecutar_itinerario(grafo, archivo[1])
        elif comando == "viaje":
            ejecutar_viaje(grafo, coordenadas, ciudades[0], archivo[0])
        elif comando == "reducir_caminos":
            ejecutar_reducir_caminos(grafo, coordenadas, archivo[1])


def main():
    args = sys.argv
    
    grafo, coordenadas = cargar_en_memoria(args[1])
    ejecutar_programa(grafo, coordenadas)


if __name__ == "__main__":
    main()
