obj-m += cpu_201020975_201212623.o

initmodules:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules

cleanmodules:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean
