**alle routen bekommen userid prefix und werden authorisiert durch kong**

---
# 1. Uploading your data

### Description:

_In order to upload your files to the platform, you need to send them as a multipart HTTP POST requests to the API. The form should contain a file field which is the file you want to upload._

_Once the file is uploaded, the server should return a JSON blob which contains amongst other things, the id of the File you just uploaded. This field is going to come in handy when referring to this particular file in all the other calls to the API._

### Endpoint:
`files/content/`

### Request:
```
curl -X POST -H 'Content-Type:multipart/form-data' \
-H 'Authorization: Token <api-token>' \
-F "file=@airfoil2d.zip" \
https://platform.rescale.com/api/v2/files/contents/
```

### Response:

**HTTP 201 CREATED**

```
{
  "typeId": 1,
  "name": "airfoil2d.zip",
  "dateUploaded": "2014-11-12T22:20:09.596996Z",
  "relativePath": "airfoil2d.zip",
  "encodedEncryptionKey": "3VgE0Ql/6hGwGNOF+TXMHNsthAdae9C+8tiIfuoDpPA=",
  "sharedWith": [],
  "decryptedSize": 663368,
  "owner": "demouser@example.com",
  "path": "user/user_OvdRk/airfoil2d-06feda72-6b3d-4f24-86e2-cd56722d1a42.zip",
  "isUploaded": true,
  "viewInBrowser": false,
  "id": "XkpTse",
  "md5": "0b66f3069732b02fe6c132f4cbd2f5b8"
}
```


---

# 2. Choosing an analysis

### Description

_A Rescale job is composed of one or more analyses (software packages). A GET call to the analysis endpoint lists out all the analyses available on the platform. A single analysis may have multiple versions. You can use the code and versionCode returned here to select the analysis while creating a Job. Additionally the meta data returned by this call is instructive for expected values in the “command” string for a given analysis. In this example we’ll use OpenFOAM._

### Endpoint:
`/analyses`

### Request:

```
curl -H "Authorization: Token <api-token>" https://platform.rescale.com/api/v2/analyses/
```

### Response:

**HTTP 200 OK**

```
{
  "industries": [
    {
      "name": "Aerospace",
      "icon": "https://d1n0dfo9fhokwf.cloudfront.net/thumbnails/tools-aero_1.png"
    },
    {
      "name": "Automotive",
      "icon": "https://d1n0dfo9fhokwf.cloudfront.net/thumbnails/tools-auto_1.png"
    }
  ],
  "code": "openfoam",
  "description": "<p><b>OpenFOAM</b> (<b>Open</b>-source <b>F</b>ield <b>O</b>peration <b>A</b>nd <b>M</b>anipulation is a C++ toolbox for the development of customized numerical solvers, and pre-/post-processing utilities for the solution of continuum mechanics problems, including computational fluid dynamics (CFD).</p>",
  "versions": [
    {
      "eula": null,
      "allowedCoreTypes": [
        "hi-io",
        "hi-io-plus",
        "hi-mem",
        "hi-mem-hpc",
        "hpc",
        "hpc-plus",
        "lo-mem",
        "standard",
        "standard-plus"
      ],
      "stdCommand": "foamExec <module> <input-file>",
      "version": "2.3.0",
      "mpiCommand": "mpirun -np <mpi-ranks> foamExec <module> <input-file> -parallel",
      "versionCode": "2.3.0-openmpi",
      "smpCommand": "mpirun -np <smp-ranks> foamExec <module> <input-file> -parallel"
    },
    {
      "eula": null,
      "allowedCoreTypes": [
        "hi-io",
        "hi-io-plus",
        "hi-mem",
        "hi-mem-hpc",
        "hpc",
        "hpc-plus",
        "lo-mem",
        "standard",
        "standard-plus"
      ],
      "stdCommand": "foamExec <module> <input-file>",
      "version": "2.2.2",
      "mpiCommand": "mpirun -np <mpi-ranks> foamExec <module> <input-file> -parallel",
      "versionCode": "2.2.2-openmpi",
      "smpCommand": "mpirun -np <smp-ranks> foamExec <module> <input-file> -parallel"
    },
    {
      "eula": null,
      "allowedCoreTypes": [
        "hi-io",
        "hi-io-plus",
        "hi-mem",
        "hi-mem-hpc",
        "hpc",
        "hpc-plus",
        "lo-mem",
        "standard",
        "standard-plus"
      ],
      "stdCommand": "foamExec <module> <input-file>",
      "version": "2.2.0",
      "mpiCommand": "mpirun -np <mpi-ranks> foamExec <module> <input-file> -parallel",
      "versionCode": "2.2.0-openmpi",
      "smpCommand": "mpirun -np <smp-ranks> foamExec <module> <input-file> -parallel"
    },
    {
      "eula": null,
      "allowedCoreTypes": [
        "hi-io",
        "hi-io-plus",
        "hi-mem",
        "hi-mem-hpc",
        "hpc",
        "hpc-plus",
        "lo-mem",
        "standard",
        "standard-plus"
      ],
      "stdCommand": "foamExec <module> <input-file>",
      "version": "2.1.1",
      "mpiCommand": "mpirun -np <mpi-ranks> foamExec <module> <input-file> -parallel",
      "versionCode": "2.1.1-openmpi",
      "smpCommand": "mpirun -np <smp-ranks> foamExec <module> <input-file> -parallel"
    }
  ],
  "supportDesks": [
    {
      "code": "rescale",
      "displayName": "Rescale Support",
      "email": "support@rescale.com"
    }
  ],
  "hasRescaleLicense": false,
  "vendorName": "",
  "pricing": "",
  "licenseSettings": [],
  "optimizerType": null,
  "thumbnail": "https://d1n0dfo9fhokwf.cloudfront.net/thumbnails/openfoam_m.png",
  "resources": [],
  "name": "OpenFOAM"
}
```

---

# 3. Setting up your Rescale Job

### Description

_A Rescale Job is composed of one or more software packages which we refer to as analyses. Thus the properties can be either at the Job level (i.e Global) or on a JobAnalysis level. See the job endpoint for information on the fields available._

**The id field in the job creation data at right is the id returned by the API from uploading the input file.**

### Endpoint:
`/jobs/`

### Request:

```
cat <<EOF > data.json
{
  "name": "Basic Job",
  "jobanalyses": [
    {
      "useMpi": false,
      "command": "./airFoil2D/Allrun",
      "analysis": {
        "code": "openfoam",
        "version": "2.3.0-openmpi"
      },
      "hardware": {
        "coresPerSlot": 1,
        "slots": 1,
        "coreType": "standard-plus"
      },
      "inputFiles": [
        {
          "id": "XkpTse"
        }
      ]
    }
  ]
}
EOF
curl -X POST --data @data.json \
-H "Authorization: Token <api-token>" \
-H "Content-Type: application/json" \
https://platform.rescale.com/api/v2/jobs/
```

### Response:

**HTTP 201 CREATED**

```
{
  "monteCarloIterations": null,
  "paramFile": null,
  "name": "Basic Job",
  "includeNominalRun": false,
  "jobanalyses": [
    {
      "envVars": null,
      "useMpi": false,
      "postProcessScriptCommand": "",
      "preProcessScriptCommand": "",
      "useRescaleLicense": false,
      "templateTasks": [],
      "analysis": {
        "code": "openfoam",
        "version": "2.3.0-openmpi"
      },
      "hardware": {
        "coreSummary": {
          "storagePerNode": 4000,
          "numberOfNodes": 1,
          "memoryPerNode": 3750
        },
        "coresPerSlot": 1,
        "slots": 1,
        "coreType": "standard-plus"
      },
      "command": "./airFoil2D/Allrun",
      "preProcessScript": null,
      "postProcessScript": null,
      "inputFiles": [
        {
          "typeId": 1,
          "name": "airfoil2d.zip",
          "dateUploaded": "2014-11-12T22:20:09.596996Z",
          "relativePath": "airfoil2d.zip",
          "encodedEncryptionKey": "3VgE0Ql/6hGwGNOF+TXMHNsthAdae9C+8tiIfuoDpPA=",
          "sharedWith": [],
          "decryptedSize": 663368,
          "owner": "demouser@rescale.com",
          "path": "user/user_OvdRk/airfoil2d-06feda72-6b3d-4f24-86e2-cd56722d1a42.zip",
          "isUploaded": true,
          "viewInBrowser": false,
          "id": "XkpTse",
          "md5": "0b66f3069732b02fe6c132f4cbd2f5b8"
        }
      ]
    }
  ],
  "resourceFilters": [
    {
      "caseSensitive": false,
      "include": true,
      "filterType": "Analysis generated files",
      "selector": "Analysis generated files"
    },
    {
      "caseSensitive": false,
      "include": true,
      "filterType": "Completed templates",
      "selector": "Completed templates"
    },
    {
      "caseSensitive": false,
      "include": true,
      "filterType": "Input files",
      "selector": "Input files"
    }
  ],
  "jobvariables": [],
  "isTemplateDryRun": false,
  "remoteVizConfig": null,
  "caseFile": null,
  "isLowPriority": false,
  "owner": "demouser@rescale.com",
  "optimizer": null,
  "expectedRuns": null,
  "id": "QTVia",
  "archiveFilters": []
}
```

---

# 4. Submitting your job

### Description:

_If the job creation was successful, server should respond back with the full schema of the job as stored on the backend. Contained within the JSON response is the job id. We will use this id to submit, monitor and get the results from the job._

**The id in the submit url at right is the job id returned by the API**

### Endpoint:
`/jobs/{JobID}/submit/`

### Request:

```
curl -X POST -H 'Authorization: Token <api-token>' \
https://platform.rescale.com/api/v2/jobs/QTVia/submit/
```

### Response:

TODO


---

# 5. Monitoring your job

### Description:

_A GET call to this endpoint returns all the statuses associated with the specified job in order of their recency._

### Endpoint:

`jobs/{JobID}/statuses`

### Request:

```
curl -H 'Authorization: Token <api-token>'  https://platform.rescale.com/api/v2/jobs/QTVia/statuses/
```

### Response:

**HTTP 200 OK**

```
{
  "count": 5,
  "previous": null,
  "results": [
    {
      "status": "Executing",
      "statusDate": "2014-11-12T22:55:45.804000Z",
      "statusReason": null
    },
    {
      "status": "Validated",
      "statusDate": "2014-11-12T22:52:12.065000Z",
      "statusReason": null
    },
    {
      "status": "Started",
      "statusDate": "2014-11-12T22:52:11.456000Z",
      "statusReason": null
    },
    {
      "status": "Queued",
      "statusDate": "2014-11-12T22:52:11.073170Z",
      "statusReason": null
    },
    {
      "status": "Pending",
      "statusDate": "2014-11-12T22:52:10.125011Z",
      "statusReason": null
    }
  ],
  "next": null
}
```

---

# 6. Getting results

### Description:

_A single execution of your workflow is called a run in Rescale parlance. When a job is in either “Completed” or “Executing” state, you can query the job runs endpoint for a list of finished runs._

_This is a basic job and therefore has 1 run. DOE jobs would have multiple runs, which could be queried individually by sending a request to the run endpoint with the run id in the path._

### Endpoint:

### Request:

```
curl -H 'Authorization: Token <api-token>' https://platform.rescale.com/api/v2/jobs/QTVia/runs/
```

### Response:

**HTTP 200 OK**

```
{
  "count": 1,
  "previous": null,
  "results": [
    {
      "dateCompleted": null,
      "parent": null,
      "variables": [],
      "outputFileSize": null,
      "outputFileCount": 0,
      "dateStarted": "2014-11-12T22:55:56.659000Z",
      "type": 3,
      "id": "1",
      "displayOrder": 1,
      "isOptimal": false
    }
  ],
  "next": null
}
```

---

# File download

### Description:

_We can list all the output files for the job by sending a request to the output files endpoint. With those ids, we can download files by issuing a GET call to the file download endpoint._


### Endpoint:

`/files/{FileID}/contents/`

### Request:

```
curl -H 'Authorization: Token <api-token>' \
https://platform.rescale.com/api/v2/files/iudjqe/contents/ > process_output.log
```
