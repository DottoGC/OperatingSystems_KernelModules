Para compilar e inicializar archivos de modulo, ejecutando desde la ubicacion del Makefile
> make initmodules


Para insertar el modulo creado a los procesos del kernel
> sudo insmod memo_201020975_201212623.ko

Para ver el log de los modulos de kernel, y ver el mensaje que se imprime al insertar el modulo
> dmesg

Para listar todos los modulos que estan en ejecucion en el kernel
>lsmod

Para remover el modulo de los procesos del kernel
> sudo rmmod memo_201020975_201212623


Para remover y limpiar archivos compilados del modulo.c, ejecutando desde la ubicacion del Makefile
> make cleanmodules

