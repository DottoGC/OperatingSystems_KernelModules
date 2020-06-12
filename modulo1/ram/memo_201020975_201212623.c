#include <linux/proc_fs.h>
#include <linux/seq_file.h> 
#include <asm/uaccess.h> 
#include <linux/hugetlb.h>
#include <linux/module.h>
#include <linux/init.h>
#include <linux/kernel.h>   
#include <linux/fs.h>

#define BUFSIZE  	150

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Escritura de un modulo con informacion de RAM.");
MODULE_AUTHOR("Ottoniel Guarchaj & Kenia Zepeda");
struct sysinfo estructuraInfoSist;

static int escribir_archivo(struct seq_file * archivo, void *v) {	
    si_meminfo(&estructuraInfoSist);
    long memoriaTotal 	= (estructuraInfoSist.totalram * 4);
    long memoriaLibre 	= (estructuraInfoSist.freeram * 4 );

    seq_printf(archivo, "***************************************************\n");
    seq_printf(archivo, "*  CARNETS:     201020975      &     201212623    *\n");
    seq_printf(archivo, "*  NOMBRES: Ottoniel Guarchaj  &   Kenia Zepeda   *\n");
    seq_printf(archivo, "***************************************************\n");
    seq_printf(archivo, " MEMORIA TOTAL:  %8lu KB - %8lu MB\n",memoriaTotal, memoriaTotal / 1024);
    seq_printf(archivo, " MEMORIA LIBRE: :  %8lu KB - %8lu MB \n", memoriaLibre, memoriaLibre / 1024);
    seq_printf(archivo, " PORCENTAJE MEMORIA UTILIZADA:  %i %%\n", (memoriaLibre * 100)/memoriaTotal) ;
    seq_printf(archivo, "***************************************************\n");
    seq_printf(archivo, "***************************************************\n\n");
    return 0;
}

static int al_abrir(struct inode *inode, struct  file *file) {
  return single_open(file, escribir_archivo, NULL);
}

static struct file_operations operaciones =
{    
    .open = al_abrir,
    .read = seq_read
};


static int inicializar(void)
{
    proc_create("memo_201020975_201212623", 0, NULL, &operaciones);
    printk(KERN_INFO "CARNETS: 201020975 & 201212623\n");

    return 0;
}
 
static void finalizar(void)
{
    remove_proc_entry("memo_201020975_201212623", NULL);
    printk(KERN_INFO "CURSO: Sistemas Operativos 1\n");
}
 


module_init(inicializar);

module_exit(finalizar); 
