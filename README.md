# serverless-aurora


inspired by https://github.com/aws-samples/serverless-go-demo

### rest api

@rootUrl = https://41dtcaimwi.execute-api.us-east-1.amazonaws.com/dev

- get {{rootUrl}}/user
- get {{rootUrl}}/user/id
- post {{rootUrl}}/user
  
  {
    "firstname": "John",
    "lastname": "Wick"
  }

- put {{rootUrl}}/user/id
  
  {
    "firstname": "Sherlock",
    "lastname": "Holmes"
  }

- delete {{rootUrl}}/user/id
