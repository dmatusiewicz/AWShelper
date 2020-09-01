# AWShelper 
# Problem statement.

Helper for exporting assumed role redentials. 



### Hierarchy
```
---
hierarchy:
  - name: Task specific configuration
    path: tasks/%{ENVIRONMENT}/%{TASK}.yaml
  - name: Env specific configuraiton
    path: environments/%{ENVIRONMENT}.yaml
  - name: Global configuration
    path: global.yaml

data: data/sample
```

### Data

*data/sample/tasks/production/task1.yaml*
```
key1: val1 
list1: 



