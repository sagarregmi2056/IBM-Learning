#!/usr/bin/env python3
from zoautil_py import datasets, jobs, zsystem
import sys, os

# Prompt for data set name:
dsname = input("Enter the Sequential Data Set name: ")
dsname = dsname.upper().replace("'", "").replace(os.environ.get('LOGNAME') + ".", "")
dsname = f"{os.environ.get('LOGNAME')}.{dsname}"
print(f"Using dataset: {dsname}")

# Check if dataset exists
if datasets.exists(dsname):
    print("Data set found! We will use it")
else:
    create_new = input("Data set not found. Create it? (y/n): ")
    if create_new.upper() == "Y":
        datasets.create(dsname, dataset_type="SEQ", primary_space="1k", secondary_space="1k")
        print("Data set created.")
    else:
        sys.exit("Without a data set name, we cannot continue. Quitting!")

# Get the system's linklist
linklist_output = zsystem.list_linklist()

# Format output in required style
python_list_repr = "[\n" + "\n".join(f"'{entry}'" for entry in linklist_output) + "\n]"

# Write to dataset (overwrite each run)
datasets.write(dsname, python_list_repr)
print(f"Linklist successfully written to {dsname}")
# Submit CHKAUTO job for validation
jcl_dataset = "ZXP.PUBLIC.JCL(CHKAUTO)"
try:
    job_id = jobs.submit(jcl_dataset)
    print(f"JCL submitted successfully. Job ID: {job_id}")
except Exception as e:
    print(f"Failed to submit JCL {jcl_dataset}: {e}")
    sys.exit(1)