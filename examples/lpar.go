package main

import (
        "fmt"

        "github.com/power-devops/perfstat"
)

func main() {
        pinfo, err := perfstat.Partition()
        if err != nil {
                fmt.Println(err)
                return
        }
        fmt.Printf("==========Configuration Information of Partition==========\n");
        fmt.Printf("Partition Name =       %v\n",pinfo.Name);
        fmt.Printf("Node Name =            %v\n",pinfo.Node)
        fmt.Printf("Partition Number =     %v\n",pinfo.Number);
        fmt.Printf("Group ID =             %v\n",pinfo.GroupID);
        fmt.Printf("\n\n========General Partition Properties(1=YES, 0=NO)=========\n");
        fmt.Printf("SMT Capable =          %v\n",pinfo.Conf.SmtCapable); 
        fmt.Printf("SMT Enabled =          %v\n",pinfo.Conf.SmtEnabled); 
        fmt.Printf("LPAR Capable =         %v\n",pinfo.Conf.LparCapable); 
        fmt.Printf("LPAR Enabled =         %v\n",pinfo.Conf.LparEnabled); 
        fmt.Printf("Shared Capable =       %v\n",pinfo.Conf.SharedCapable); 
        fmt.Printf("Shared Enabled =       %v\n",pinfo.Conf.SharedEnabled); 
        fmt.Printf("DLPAR Capable =        %v\n",pinfo.Conf.DLparCapable); 
        fmt.Printf("Capped =               %v\n",pinfo.Conf.Capped); 
        fmt.Printf("64-Bit Kernel =        %v\n",pinfo.Conf.Kernel64bit); 
        fmt.Printf("Pool Util Authority =  %v\n",pinfo.Conf.PoolUtilAuthority); 
        fmt.Printf("Donate Capable =       %v\n",pinfo.Conf.DonateCapable); 
        fmt.Printf("Donate Enabled =       %v\n",pinfo.Conf.DonateEnabled); 
        fmt.Printf("AMS Capable =          %v\n",pinfo.Conf.AmsCapable); 
        fmt.Printf("AMS Enabled =          %v\n",pinfo.Conf.AmsEnabled); 
        fmt.Printf("Power Saving Mode =    %v\n",pinfo.Conf.PowerSave); 
        fmt.Printf("AME Enabled =          %v\n",pinfo.Conf.AmeEnabled); 
        fmt.Printf("Shared Extended =      %v\n",pinfo.Conf.SharedExtended); 
        fmt.Printf("\n\n==================Hardware Configuration==================\n");
        fmt.Printf("Processor Type =               %v\n",pinfo.ProcessorFamily);
        fmt.Printf("Processor Model =              %v\n",pinfo.ProcessorModel);
        fmt.Printf("Machine ID =                   %v\n",pinfo.MachineID);
        fmt.Printf("Processor Clock Speed =        %v MHz\n",pinfo.ProcessorMhz);
	fmt.Printf("Online Configured Processors = %v\n",pinfo.NumProcessors.Online);
	fmt.Printf("Max Configured Processors =    %v\n",pinfo.NumProcessors.Max);
	fmt.Printf("\n\n==================Software Configuration==================\n");
	fmt.Printf("OS Name =                      %v\n",pinfo.OSName);
	fmt.Printf("OS Version =                   %v\n",pinfo.OSVersion);
	fmt.Printf("OS Build =                     %v\n",pinfo.OSBuild);
	fmt.Printf("\n\n====================LPAR Configuration====================\n");
	fmt.Printf("Number of Logical CPUs =       %v\n",pinfo.LCpus);
	fmt.Printf("Number of SMT Threads =        %v\n",pinfo.SmtThreads);
	fmt.Printf("Number of Drives =             %v\n",pinfo.Drives);
	fmt.Printf("Number of NW Adapters =        %v\n",pinfo.NetworkAdapters);
	fmt.Printf("\n\n===========Physical CPU Related Configuration=============\n");
	fmt.Printf("Minimum CPU Capacity =         %v\n",float32(pinfo.CpuCap.Min)/100.0);
	fmt.Printf("Maximum CPU Capacity =         %v\n",float32(pinfo.CpuCap.Max)/100.0);
	fmt.Printf("CPU Capacity Weightage =       %v\n",pinfo.Weightage);
	fmt.Printf("Entitled Proc Capacity =       %v\n",float32(pinfo.EntCapacity)/100.0);
	fmt.Printf("\n\n============Virtual CPU Related Configuration=============\n");
	fmt.Printf("Minimum Virtual CPUs =         %v\n",pinfo.VCpus.Min);
	fmt.Printf("Maximum Virtual CPUs =         %v\n",pinfo.VCpus.Max);
	fmt.Printf("Online Virtual CPUs =          %v\n",pinfo.VCpus.Online);
	fmt.Printf("\n\n==========Processor Pool Related Configuration============\n");
	fmt.Printf("Processor Pool Id =            %v\n",pinfo.PoolID);
	fmt.Printf("Active CPUs in pool =          %v\n",pinfo.ActiveCpusInPool);
	fmt.Printf("Pool Weightage =               %v\n",pinfo.PoolWeightage);
	fmt.Printf("Shared processors Count =      %v\n",pinfo.SharedPCpu);
	fmt.Printf("Max pool Capacity =            %v\n",pinfo.MaxPoolCap);
	fmt.Printf("Entitled pool Capacity =       %v\n",pinfo.EntPoolCap);
	fmt.Printf("\n\n==============Memory Related Configuration================\n");
	fmt.Printf("Minimum Memory =               %v\n",pinfo.Mem.Min);
	fmt.Printf("Maximum memory =               %v\n",pinfo.Mem.Max);
	fmt.Printf("Online memory =                %v\n",pinfo.Mem.Online);
	fmt.Printf("Memory capacity Weightage =    %v\n",pinfo.MemWeightage);
	fmt.Printf("\n\n===============AMS Related Configuration==================\n");
	fmt.Printf("I/O memory Entitlement =       %v\n",pinfo.TotalIOMemoryEntitlement);
	fmt.Printf("AMS Pool ID =                  %v\n",pinfo.MemPoolID);
	fmt.Printf("Hypervisor Page Size =         %v\n",pinfo.HyperPgSize);
	fmt.Printf("\n\n===============AME Related Configuration==================\n");
	fmt.Printf("Minimum Expanded memory =      %v\n",pinfo.ExpMem.Min);
	fmt.Printf("Maximum Expanded Memory =      %v\n",pinfo.ExpMem.Max);
	fmt.Printf("Online Expanded memory =       %v\n",pinfo.ExpMem.Online);
	fmt.Printf("Target memory Expansion factor = %v\n",pinfo.TargetMemExpFactor);
	fmt.Printf("Target Memory Expansion Size = %v\n",pinfo.TargetMemExpSize);
	fmt.Printf("\n==========================================================\n");
}
