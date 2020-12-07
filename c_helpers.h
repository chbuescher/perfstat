#include <libperfstat.h>

extern perfstat_cpu_t *get_cpu_stat(perfstat_cpu_t *, int);
extern perfstat_diskadapter_t *get_diskadapter_stat(perfstat_diskadapter_t *, int);
extern double get_partition_mhz(perfstat_partition_config_t);
