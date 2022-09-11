#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <asm/uaccess.h>	
#include <linux/seq_file.h>
#include <linux/hugetlb.h>

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo RAM");
MODULE_AUTHOR("Diego Fernando Cotez Lopez, 201900955");

struct sysinfo ram_info;

static int writeFile(struct seq_file * file, void * v)
{   
    long ram_total;
    long ram_free;

    si_meminfo(&ram_info);
    ram_total = (ram_info.totalram * ram_info.mem_unit)/(1024*1024);
    ram_free = (ram_info.freeram * ram_info.mem_unit)/(1024*1024);

    seq_printf(file, "{");
    seq_printf(file, "\"total\": ");
    seq_printf(file, "%d", ram_total);
    seq_printf(file, ", \"free\": ");
    seq_printf(file, "%d", ram_free);
    seq_printf(file, "}");
    return 0;
}

static int openFile(struct inode * inode, struct file * file){
    return single_open(file, writeFile, NULL);
}

static struct proc_ops operaciones ={
    .proc_open = openFile,
    .proc_read = seq_read
};

static int insertMod(void){
    proc_create("ram_201900955", 0, NULL, &operaciones);
    printk(KERN_INFO "201900955\n");
    return 0;
}

static void removeMod(void){
    remove_proc_entry("ram_201900955", NULL);
    printk(KERN_INFO "Segundo Semestre 2022\n");
}

module_init(insertMod);
module_exit(removeMod);