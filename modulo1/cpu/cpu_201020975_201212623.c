#include <linux/proc_fs.h>
#include <linux/seq_file.h> 
#include <asm/uaccess.h> 
#include <linux/hugetlb.h>
#include <linux/module.h>       // Needed by all modules
#include <linux/init.h>
#include <linux/kernel.h>       // KERN_INFO
#include <linux/fs.h>
#include <linux/sched.h>        // for_each_process, pr_info
#include <linux/sched/signal.h>  

#define BUFSIZE         150

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Escritura de un modulo con informacion de RAM.");
MODULE_AUTHOR("Ottoniel Guarchaj & Kenia Zepeda");
struct task_struct* task_list;
struct task_struct *task_list_child;
struct list_head *list;

static int escribir_archivo(struct seq_file * archivo, void *v) {       

        seq_printf(archivo, "CARNETS:     201020975      &     201212623       \n");
        seq_printf(archivo, "NOMBRES: Ottoniel Guarchaj  &   Kenia Zepeda      \n");

        seq_printf(archivo, "PID \t NAME \t STATE \t CHILDS \n");
        for_each_process(task_list) {
                seq_printf(archivo,"[%d] \t [%s] \t [%ld] \n", task_list->pid,task_list->comm,task_list->state);
                
                list_for_each(list, &task_list->children){         
                    task_list_child = list_entry( list, struct task_struct, sibling );             
                    seq_printf(archivo,"      [%d] \t [%s] \t [%ld] \n",task_list_child->pid, task_list_child->comm, task_list_child->state);
                }
                seq_printf(archivo, "\n");
        }

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
    proc_create("cpu_201020975_201212623", 0, NULL, &operaciones);
    printk(KERN_INFO "CARNETS: 201020975 & 201212623\n");

    return 0;
}
 
static void finalizar(void)
{
    remove_proc_entry("cpu_201020975_201212623", NULL);    
    printk(KERN_INFO "CURSO: Sistemas Operativos 1\n");
}


module_init(inicializar);

module_exit(finalizar); 
