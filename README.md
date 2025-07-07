# kumparan-backend-position-interview

## How to Run the Project

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/kumparan-backend-position-interview.git
    ```

2. Navigate to the project directory:
    ```bash
    cd kumparan-backend-position-interview
    ```

3. Install the dependencies:
    ```bash
    go mod tidy
    ```

4. Start the application:
    ```bash
    go run main.go
    ```

5. Access the application in your browser at `http://localhost:8080`.


## Existing API
The following endpoints are available in the application:

1. **GET /v1/articles**
    - Description: Retrieves example Articles.
    - Response: JSON object containing article data.

2. **POST /v1/articles**
    - Description: Creates a new articles entry.
    - Request Body: JSON object with article data.
      ```json
      {
        "title": "Programming in HTML",
        "body": "Learning HTML",
        "author_id": "auth-003",
        "created_at": "2025-07-06T20:45:00Z"
      }
      ```
    - Response: Confirmation message with the created entry.

3. **Get /v1/articles/search**
    - Description: Updates an existing example entry.
    3. **GET /v1/articles/search**
        - Description: Updates an existing example entry.
        - Request Body: JSON object with updated example data.
          ```json
          {
              "title": "HTML",
              "body": "HTML"
          }
          ```
        - Response: Confirmation message with the updated entry.
    - Response: Confirmation message with the updated entry.
