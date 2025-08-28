
# Mainframe Virtualization & Memory Notes

## 1. Virtualization on the Mainframe
- A **mainframe** is physically huge: processors, memory, I/O adapters, cooling, power, networking.
- Virtualization allows carving this physical machine into multiple **logical systems**.
- These logical systems are called **LPARs (Logical Partitions)**.
  - Each LPAR behaves like its own computer:
    - Runs an operating system
    - Can be started/stopped independently
    - Has virtual display and I/O
  - Multiple LPARs can coexist on the same mainframe:
    - Some big (lots of processors + memory)
    - Some small (just enough to run basic programs)
    - Flexible allocation of resources.

**Analogy:** Like a box of LEGO blocks — you can build many different systems (boat, house, car) from the same set of resources.

---

## 2. Mainframe Processors
Mainframes contain many processors, each configured for different roles:

- **CP (Central Processor)**  
  General-purpose, runs operating systems and applications.

- **SAP (System Assist Processor)**  
  Handles I/O data movement, freeing up CP to focus on business apps.

- **IFL (Integrated Facility for Linux)**  
  Specialized processor for running Linux workloads (cheaper licensing).

- **ZIIP (z Integrated Information Processor)**  
  Offloads specific workloads (e.g., Java) to save cost and free CP capacity.

---

## 3. I/O Adapters
- Mainframes have many **I/O adapters**, connected to:
  - Disk storage
  - Network adapters
  - Other mainframes
- Identified using **CHPID (Channel Path ID)**:
  - **PCHID** → physical port location
  - **Subsystem** → logical connection to LPAR
- CHPID types:
  - **Dedicated** → only one LPAR
  - **Shared** → multiple LPARs
  - **Reconfigurable** → switchable between LPARs
  - **Spanned** → used by multiple subsystems.

---

## 4. Virtualization Software: z/VM
- **z/VM** = Type-1 hypervisor for mainframes.
- Runs directly on hardware (not on another OS).
- Can create **hundreds/thousands of virtual machines** inside it.
- To the mainframe, it looks like one OS, but internally it can host many.

---

## 5. Memory & Storage Concepts
- **Storage** on mainframes = combination of real memory + auxiliary (AUX) storage.
- **Real Memory** = physical RAM.
- **Auxiliary (AUX) Storage** = DASD (Direct Access Storage Device: disk/SSD).
- Applications always see it as one continuous **storage pool**.

### Virtual Memory
- Managed by **z/OS** (mainframe OS).
- Applications get their own **address space** (up to 16 exabytes).
- Memory is divided into **pages** (4 KB each).
  - In real memory → called a **frame**.
  - In AUX storage → called a **slot**.
- Moving data between them = **paging**.
  - Paging should be minimized (slows system).
  - Well-tuned system = page out once, rarely page back in.

---

## 6. Mainframe Architecture
- Current generation = **z/Architecture**.
- Includes:
  - Mainframe hardware
  - Operating systems (z/OS, Linux on Z, z/VM, etc.)
- Tight integration between hardware + OS makes virtualization and memory management efficient and scalable.

---

## ✅ Key Takeaways
- **Virtualization**: Splits one big machine into many logical computers (LPARs).  
- **Processors**: Different types (CP, SAP, IFL, ZIIP) optimize cost and workload.  
- **Memory**: Combination of RAM (real) + disk/SSD (AUX) managed by paging.  
- **z/VM**: Hypervisor enabling thousands of VMs.  
- **z/Architecture**: Modern mainframe design blending hardware + OS features.  





Resources Video url (free) : https://learn.ibm.com/mod/video