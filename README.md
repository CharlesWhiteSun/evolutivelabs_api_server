# Gin Api Server

## Design Requirements

  1. Metaphorpsum API <br>
    - Call third-party API. <br>
    - Implement the getSentence() method and return content. <br>
    - The method should not contain parameters. <br>
    - Need to add unit tests. <br>
  ```
    Example
      - Http Method : GET
      - Request URL : 127.0.0.1:8000/api/v1/sentences
      - Response :
        {
          "code": 200,
          "data": "Before shocks, dirts were only methanes. They were lost without the halftone list that composed their respect. Extending this logic, a bedimmed bull's lobster comes with it the thought that the toothless client is a Sunday.",
          "msg": "ok"
        }
  ```
  
  2. Itsthisforthat API <br>
    - Call third-party API. <br>
    - The method should different parameters. <br>
    - Need to add unit tests. <br>
  ```
    Example
      - Http Method : GET
      - Request URL : 127.0.0.1:8000/api/v1/sentences/{text}
      - Response :
        {
          "code": 200,
          "data": "So, Basically, It's Like A Airbnb for Blacking Out.",
          "msg": "ok"
        }
  ```
