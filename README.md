# ws17_pp
Praxisprojekt (BHTB MIB 13 W17)

## Dokumentation:
https://www.ws17-pp.mhertel.de

## Routing of the API

**Endpoints Jobs:**
```
URI: /jobs                                    METHOD: GET         DESC: Retrieve all Jobs
URI: /jobs                                    METHOD: POST        DESC: Create a new Job
URI: /jobs                                    METHOD: PUT         DESC: Update an existing Job
URI: /jobs                                    METHOD: DELETE      DESC: Delete an exisiting Job
URI: /jobs/{jobID}                            METHOD: GET         DESC: Get specific job
```

**Endpoints Templates:**
```
URI: /jobs/{jobID}/templates                  METHOD: GET         DESC: Retrieve all Templates for Job
URI: /jobs/{jobID}/templates                  METHOD: POST        DESC: Create a new Template for Job
URI: /jobs/{jobID}/templates                  METHOD: PUT         DESC: Update an existing Template
URI: /jobs/{jobID}/templates                  METHOD: DELETE      DESC: Delete an exisiting Template
URI: /jobs/{jobID}/templates/{templateID}     METHOD: GET         DESC: Get specific Template
```

## Models

```
type Job struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Param       string        `bson:"param" json:"param"`
	Templates   []Template    `bson:"templates" json:"templates"`
}
```

```
type Template struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	JobID string        `bson:"jobid" json:"jobid"`
	Name  string        `bson:"name" json:"name"`
	Path  string        `bson:"path" json:"path"`
}
```



## Development

1. required for creating tmux session (.tmuxinator.yml)

  `./scripts/install_tmuxinator.sh`

  (requires ruby <=2.2.1)

2. recommened to start the tmux-session

  `./scripts/up_tmux.sh`

  (alternativly execute the scripts in ./scripts folder to start the services )

3. foreach service is a tmux-window with useful panes

  ```
  windows:
    - working:
      - panes
        - git:
        - docker:
    - go-rest-api:
      - panes
        - server:
        - test-server:
    - kong:
      - panes
        - compose-up:
        - seeding:
    - minio:
      - panes
        - compose-up:
        - seeding:
    - nomad-consul:
      - panes:
        - start nomad
        - start consul
        - submit testjob fib
  ```

4. quit Development with stopping service-containers)

  `./scripts/down_tmux.sh`

## Architecture (High Level)

```

                            +                                                +
CLIENT_LAYER                |              API_LAYER                         |            SOLVER_LAYER
                            |                                                |
                            |                                                |
                            |                                                |
                            | Forward Requests                    Submit     |                           Submit Job
         +------------------+                   +--------------+             +----------------------+      Docker         +-------+
         | Kong Api Gateway +-----------------> | HPC-REST-API +-----------> | NOMAD Job Scheduling +-------------------> | Node1 |
         | :8000            |                   | :7777        |             | :4646                |                     +-------+
         | :8001            |                   +--------------+             +----------------------+
         | :8443            |                    |                           |                                            +-------+
         | :8444            |                    |                           |                                            | Node2 |
         +------------------+                    | FILES                     |                                            +-------+
                            |                    |                           |
                            |                   +v------------------+        |                                            +-------+
                            |                   | MINIO Objectstore |        |                                            | Node3 |
                            |                   | :9001             |        |                                            +-------+
                            |                   | :9002             |        |
                            |                   | :9003             |        |                                            +-------+
                            |                   | :9004             |        |                                            | Node4 |
                            |                   +-------------------+        |                                            +-------+
                            |                                                |
+---------------------------------------------------------------------------------------------------------------------------------+
                            |                                                |
PERSISTENCE                 |                                                |
         +------------------+                   +----------------------------+
         | POSTGRES         |                   | MONGODB                    |
         | :5432            |                   | :27017                     |
         +------------------+                   +----------------------------+

```

## Aufgaben

1. Exposee (Fälligkeitsdatum	Dienstag, 17. Oktober 2017, 00:00)

2. Status KW.41-42 (Fälligkeitsdatum	Dienstag, 24. Oktober 2017, 00:00)

3. Status KW.43-44 (Fälligkeitsdatum	Dienstag, 7. November 2017, 00:00)

4. Status KW.45-46 (Fälligkeitsdatum	Dienstag, 21. November 2017, 23:55)

5. Status KW.47-48 (Fälligkeitsdatum	Dienstag, 5. Dezember 2017, 23:55)

6. Status KW.49-02 (Fälligkeitsdatum	Dienstag, 16. Januar 2018, 23:55)

7. Abschlussdokument (Fälligkeitsdatum ...)

8. Folien / Präsentation (Fälligkeitsdatum	Samstag, 3. Februar 2018, 00:00)

## Ressourcen

https://prof.beuth-hochschule.de/edlich/praxisprojekt-online/
