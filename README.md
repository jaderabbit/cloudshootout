# Cloud Shootout

The aim of this project is to provide a benchmark for testing kubernetes deployments on the 3 major clouds.
Still in early development. 

## Project Structure

- `aws/` contains the AWS EKS configuration
- `azure/` contains the Azure AKS configuration
- `gcp/` contains the Google Cloud GKS configuration
- `helloapp` contains the tiny golang application we're going to deploy and it's docker file


## helloapp

Helloapp is a small example web server that connects to a postgresql database. 

Required environment variables:
- DATABASE_HOST
- DATABASE_PORT
- DATABASE_USER
- DATABASE_PASS
- DATABASE_NAME

The server will create a single table `tests` if it does not exist. It's a dummy table that has an id, a text field called `random` and a timestamp `created_at`.
When the web server is run and you access it, the system will add a row to the database with a random value in the text field and return the generated random value and the id 


