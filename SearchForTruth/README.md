# Concept idea for state machine 

There are Projects and there are Models. A Project owns a series of Models. For the project to proceed, it should initiate Models that represent computation. There should also be a mechanism for the computation to communicate with Project, this could take the form of a queue or REST endpoint. 

### More Detail 

The idea here is that there are "Models" that are computational workloads. Each computation requires inputs and ouputs. The required model should be updated and the Project be updated and begin downstream models. 

To achieve this; a "Project State Service" (PSS) could be makde to store and update state of a Project, it is also responsible for kicking off Model runs.

The Model strucutre is much of a Linked List, or nested object. 

### Mechanism
A PPS service is created, it is web enabled and is backed by a persistence layer. A call can initiate a project to kick off the Projects Model run. This should off load the computation off to another location like a autoscaling group or on prem server. This initial will only kick off what it can, because of run dependecies, likley to only progress one step initially. The idea here the action of starting a Projects Model is stateless

The Models will run, Inputs would be given to the Models runtime, the Models will do their compute and send back the results via REST or another mechanism to the PPS. The PPS will then update the Project and attempt to kick off the next model with the given result
