# SO1P1
Proyecto 1 - MODULOS KERNEL, PLANIFICACIÓN DE PROCESOS Y MONITOREO DE RECURSOS
Sistemas Operativos 1
Escuela de Vacaciones de Junio 2020
Facultad de Ingenieria, Universidad de San Carlos de Guatemala


# MODULOS KERNEL, PLANIFICACIÓN DE PROCESOS Y MONITOREO DE RECURSOS
Descripcion:
Se trata de desarrollar una aplicación que permita monitorear y gestionar los procesos de un servidor Linux, por medio de una interfaz web de fácil acceso desde el navegador de una computadora o de un dispositivo móvil como teléfono o Tablet. 
La información de los procesos, uso de memoria RAM y uso de CPU será extraída mediante la implementación de módulos kernel que estarán adquiriendo y exponiendo esta información utilizando archivos ubicados en /proc.

# Modulo1
## Módulo de Memoria(sysinfo)
El módulo deberá sobrescribir un archivo en el directorio /proc con la siguiente información:
- Carné Numero carnet del Estudiante 
- Nombre Nombre del estudiante 
- Memoria total Memoria total utilizada en el momento en MB Memoria libre 
- Memoria total libre en MB % de memoria utilizada 
- Porcentaje de utilización de memoria

Características a implementar: 
- Debe imprimir el número de Carné del estudiante al cargar el módulo (insmod). 
- Debe imprimir el nombre del curso al momento de descargar el módulo (rmmod). 
- La información que se mostrará en el módulo debe ser obtenida por medio de los struct de información del sistema operativo y no de la lectura de otro archivo. 
- El contenido del archivo se debe actualizar al abrir el archivo (evento open). 
- El nombre del módulo será: memo_<<carne>>


## Módulo CPU(task_struct)
El módulo deberá sobrescribir un archivo en el directorio /proc con la siguiente información de encabezado: 
- Carné Numero de carné del estudiante 
- Nombre Nombre del estudiante
Posterior a esto deberá listar todos los procesos, mostrando:
- PID Identificador del proceso 
- Nombre Nombre del proceso 
- Estado Estado en el que se encuentra el proceso 
- Hijos Lista de procesos hijos

Características a implementar: 
- Importar librerías: <linux/sched.h>, <linux/sched/signal.h> 
- Debe imprimir el nombre del estudiante al cargar el módulo (insmod). 
- Debe imprimir el nombre del curso al momento de descargar el módulo (rmmod). 
- La información a mostrar debe ser obtenida por medio de los struct de datos del sistema operativo y no de la lectura de archivos o comandos de consola. 
- El contenido del archivo se debe actualizar al abrir el archivo (evento open). - El nombre del módulo será: cpu_<<carne>
  


# Modulo2
## APPLICACION WEB
La aplicación web permite visualizar gráficas dinámicas que muestren el uso del CPU y de la memoria RAM del servidor. 
La aplicación web permite mostrará la información básica de los procesos que se ejecutan y permite terminar los procesos(kill) que se encuentran en ejecución.

Página Principal
Esta debe mostrar de manera tabulada todos los procesos que están siendo ejecutados en el servidor, así como un resumen general de los procesos

Monitor de CPU 
El monitor de CPU debe mostrar la información del consumo de CPU del servidor, en el cual se podrá visualizar la siguiente información
La gráfica debe ser similar a un polígono de frecuencia, el cual debe mostrar el consumo del CPU del servidor en tiempo real sin que el usuario necesite estar refrescando la página para monitorear el comportamiento de la utilización del CPU. Queda a discreción del estudiante la herramienta para realizar las gráficas.

Monitor de RAM 
El monitor de memoria RAM es similar al de CPU, debe mostrar la información del consumo de RAM del servidor.


# Requisitos/Tecnologias
- Ubuntu 18.04
- Golang (Go)
