## Remaining TODO Parts :-

1. Add the thrid stage to pipeline which will take the aggregation bits and find the attestation yes/no and save for the combination of validator and a slot (explained below in detaisl for db design)
2. Need to check and redirect the calls which are failed first with 429 due to rate-limit limitation in the beacon server
3. Web server for the participation rate, which will data from the levelDb file per slot and per validators for missed and all attesttations.

## DB Design and data save logic - 

LevelDB Reason -

I can see the slot value and the validator index are in integers and can be saved in the leveldb as bytes keys and I would need the composite keys as well so to traverse the composite keys and operate on, its efficient to take the operation between integers and byte arrays.

Since we need indexing operations the data is best to be saved in key value model and modularised in files for each epoch and all validator index in another file to keep the workload and data both minimal

### DB Design -

1. Would create files in server memory for level db using apis of lib - github.com/syndtr/goleveldb/leveldb
   Files will be each for each epoch and one file from all validators missed and total attestations which will help us in catering participation rate of specific/any validator

2. the Epoch file will save format as -
   SlotNumber [bytes] validatorIndex[bytes] : 0/1 [bytes].  (yes/no attested ?)
3. the Validator file will contain
   validatorIndex[bytes] : cumulativeMissedAttestation [bytes] + cumulatineTotalAttestation[]
4. Can also keep and global synchronised variable for the total global missed attestation which will be maintained every slot-come-in

## Morphing logic for realtime indexing

1. use see client lib - github.com/r3labs/sse to subscribe handle event with topic head
2. inside event handler :-
   a. get the slot data and it should have the attestation data
   b.if it doesnt have fetch the attestation data from attestation api and list of validators
   c. for each validator, add key for slot:validator and attestaion yes/no and add the cumulative validator performace in validator file
   d. keep the track of oldest slot/epoch in a variable and delete that slot:validaor(s) and adjust the values accordingly from the total cumulative
3. This application can be deployed in any cloud container.


### How to run-

1. Make sure docker is installed in system
2. to build the image - make build-dev
3. to run the container instance make run-dev


