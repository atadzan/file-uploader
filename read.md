# **minIO Multi-Cloud Object Storage**
## Mini project to learn basic functionality of minIO storage with Go Client SDK

### Endpoints
 #### POST ```/bucket``` - creates a new bucket
 #### GET ```/buckets``` - get list of buckets
 #### POST ```/upload``` - upload file to specific bucket
 #### POST ```/download``` - download file from specific bucket of storage
 #### GET ```/file``` - get download link of file  


### Creates a new bucket
**POST** `/bucket`
#### Request body:
    {
      "bucketName" string 
    }
#### Responses:
##### Success:
    {
      "code":    200
      "message": "Successfully created"
    }
##### Error:
    {
      "code"    int
      "message" error message
    }

### Get list of buckets
**GET** `/buckets`
#### Request body:
None
#### Responses:
##### Success:
    [
     {
       "name"          string  
       "creationDate"  time.Time 
     }
    ]   
##### Error:
    {
      "code"    int
      "message" error_message
    }

### Upload file to specific bucket
**POST** `/upload`
#### Request body:
    {
      "bucketName"  string 
	  "fileName"    string 
	  "filePath"    string
	  "contentType" string 
    }
#### Responses:
##### Success:
    {
      "code":    200
      "message": "Successfully uploaded <filename> of size <int> "
    }
##### Error:
    {
      "code"    int
      "message" error message
    }

### Download file from specific bucket of storage
**POST** `/download`
#### Request body:
    {
      "bucketName"       string 
	  "fileName"         string
	  "destinationPath"  string 
    }
#### Responses:
##### Success:
    {
      "code":    200
      "message": "Successfully downloaded in <destinationPath>"
    }
##### Error:
    {
      "code"    int
      "message" error message
    }

### Get download link of file
**GET** `/file`
#### Request body:
    {
      "bucketName"       string 
	  "fileName"         string
    }
#### Responses:
##### Success:
    {
      "code":    200
      "message": "<download link>"
    }
##### Error:
    {
      "code"    int
      "message" error message
    }