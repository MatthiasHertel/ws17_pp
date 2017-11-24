# ws17_pp
Praxisprojekt (BHTB MIB 13 W17)

https://www.ws17-pp.mhertel.de

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
         | POSTGRES         |                   | DOCUMENTSTORAGE ?          |
         | :5432            |                   |                            |
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
