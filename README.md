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

## Configurations

## GCP
*Nodes*:
- Instance: n1_standard_1
- 1 CPU
- 3.25 GB Memory
- K8s storage volume, but none on nodes

*Database*:
- 2 CPU
- 7.5GB memory
- 10 GB storage
- (Doesn't use predefined instances for PG)

### AWS
*Nodes*:
- Instance: t2 small
- 1 CPU
- 2 GB memory
- 20 GB storage

*Database*:
- Instance: t2 large
- 20 GB storage
- 2 CPU
- 8 GB memory 


### Azure

*Nodes*:
- Instance:
- 1 CPU
- 3.5 GB memory
- No storage

*Database*:
- 2 CPU
- 100 GB
- Memory? Who knows!

## WTFs Per Session

A record of everytime I went WTF.

### GCP
Time to Functional: 1.5 hours
Number of WTFS: 7

1. The default method to get access to my DB is using a Cloud Proxy
2. Activation time for every single API
3. Took me 5 tries to create a service key. The interface is pretty but can be clumsy.  
4. Tiny differences between the tutorial and the repo code. They did things slightly differently which confused me
5. Stack driver is strange. It offered to log things for my AWS. I oblidged, but then suddenly those credentials appeared everywhere
6. Confusion between the tut and repo meant I used a couple of commands incorrectly. 
7. Interface made chrome crash

### AWS
Time to Functional: 3 hours
Number of WTFS: 11

1. The tutorial tells you to go to CloudFormation, but doesn't tell you what it is or why it is necessary
2. There is a separate configuration setup cloudformation mission for master node and worker nodes
3. The default example to set this up has literally no explanation. Just lines of code. Hard to cross apply
3. There is no Kubernetes dashboard
4. The AWS authenticator for kubectl is 26 MB. For what?
5. You need to manually copy paths to make use of the AWS authenticator
6. In order to create clusters and nodes via the CLI, you need to save the subnets, iam groups, security groups. 
7. The AWS authenticator is a nightmare to actually get working
8. The documentation for the AWS authenticator, albiet detailed, did not help solve my problem
9. I needed to know that the RDS should be in same security group. No tutorial to explain or highlight this
10. The documentation and the cloudformation interfaces don't match - seems their are a few interface versions.
11. What is a cloudformation template for?


### Azure
Time to Functional: 4 hours
Number of WTFS: 15

1. Interface made chrome crash
2. Only certain machine sizes support accelerated networking. This will make connections to PG slower.
3. The instance type I'm using doesn't appear in the documentation
4. The interface is buggy when it came to specifying my database size. 
5. It doesn't tell me how much memory my database has and I can't select it
6. The Azure + PG tutorial exists but has no information in it
7. The container registry is impossible to find in the interface
8. Smaller size machines do not support virtual networks. What?
9. Azure has a cloud shell that doesn't support docker. So no building of images
10. There recommended way for installing the Azure CLI (if their standard releases didn't work) is via a docker, which once again, means I can't build docker images
11. I had to install the github bleeding new release version 
12. There are inconsistencies between the Azure CLI and the documentation
13. The Kubernetes autoscaling is in preview mode
14. 10-15 minute create of cluster
15. Pay for storage for using their cloud shell

## WOWs per Session

A record of everytime I went WOW.

### GCP

Number of WOWs: 9

1. Cloud bash is AMAZING. It works beautifully - 
2. Documentation is well supported and links well together. 
3. "What's Next" on each tutorial are really useful
4. Very complete tutorials with many combinations really considering how we can vary stuff. 
5. Permission management is AMAZING. If you're using does not have a permission, the error messages provide you a link that takes you straight to the dashboard to fix it 
6. Cluster took 2-3 minutes to be created.
7. Cost optimizations from the dashboard. In your face!
8. K8s dashboards are amazing
9. Stack driver logging works out the box

### AWS
Number of WOWs: 2

1. Autoscaling configurations mentioned up front. 

### Azure

Number of WOWs: 8

1. The database creation enforces password strength of SQL users
2. The interface is very detailed providing every type of metric you'd want
3. Has infrastructure to discuss connection pooling
4. Very beautiful dashboard for Kubernetes - even giving you easy links to get to the K8s dashboard
5. Price in Rands
6. The price is directly on the dashboard when you create the services
7. Amazing defaults but allow customization
