# Ad Campaigns

Ad Campaigns is a web application that allows users to register, login and view a list of pre-loaded ad campaigns.
It's implemented using a Microservice architecture.
A lot of the implementation patterns and code snippets I've used can be found in my Instagram Bot application.

The reason I've chosen a microservice architecture is, in a case like this, modularity. Adding new features, views or 
functionality would be as simple as writing a new microservice. One microservice crashing would not bring the whole application 
down. 

The backend is written in Golang. The reason is because Golang is, among other things, performant, easy to read,
well suited for writing microservices and is cross-platform.

The frontend is written in React.js. The reason I chose it is because it's lightweight and makes it easy to generalize 
to various data structures and render HTML according to the data retrieved from the server.
The detailed view of ad campaign is a good example of rendering complex structures where some object attributes can be omitted.

I found that displaying the detailed view in a modal simplifies implementation, in a sense that we don't need additional views, 
routing and potentially additional requests to the server.

The application is deployed with Docker containers, as separate microservices, to a Kubernetes cluster.
Kubernetes provides an easy way to have service discovery by utilizing it's inner DNS, and also makes scaling convenient.

The database that stores the ad campaigns is MongoDB. Because of a relatively complex structure of the JSON found in
 data.json, I've decided to use a No-SQL database, which will make it easier to store data without having to deal 
 with references between nested objects.
 
The part where I had a difficulty deciding which path to take is the model. The model is supposed to represent the structure 
found in the JSON data file. The same structure needs to be implemented in the protocol buffer descriptor file.
In order to communicate this data between the services using gRPC protocol, at some point conversion between proto and model 
needs to be done. 
So since it's a fairly complex structure, converting adds a bit of overhead. So an alternative would be to use the generated 
same-structured model found in the proto descriptor for persisting to the database. This would remove the overhead of mapping these
 complex objects, but the problem I had with this approach is that the proto model adds some additional fields that would also 
 be persisted along with the ones found in the JSON, and also gives less control over the model definition, since it's generated.
 
## Installation

At this time, I'm keeping the deployment descriptors in the same repository with the source, but ideally they would 
be in a separate repository. 

#### Prerequisites:

- Docker
- Minikube

#### Steps
- Clone the repository

        git clone https://github.com/nanosapp/fullstack-dev-assesment.git

- Navigate to the nanosapp project directory 

        cd nanosapp

- Deploy the microservices to your local Minikube cluster

        make deploy
        
- When all the pods are running, navigate to
        
        http://192.168.99.100:30001

## Features

- [Register a new account](#register-a-new-account)
- [Login to your account](#login-to-your-account)
- [See an overview of all Ad Campaigns](#see-an-overview-of-all-ad-campaigns)
- [See a detailed view of a single Ad Campaign](#see-a-detailed-view-of-a-single-ad-campaign)

## Microservices

- **web** - Go
- **account** - Go
- **adcampaign** - Go

## Technologies/libraries used

- [Golang](https://golang.org)
- [Micro](https://micro.mu)
- [gRPC](https://grpc.io)
- [Kubernetes](https://kubernetes.io)
- [Docker](https://www.docker.com)
- [Labstack Echo](https://echo.labstack.com)
- [PostgreSQL](https://www.postgresql.org)
- [MongoDB](https://www.mongodb.com)
- [React.js](https://reactjs.org)
- [Gorm](http://gorm.io)
- [dep](https://golang.github.io/dep)
- [Bootstrap](https://getbootstrap.com)


## Register a new account

Register a new account on the Ad Campaigns web application, with provided username and password.

To achieve this,
- An HTTP request is sent to the **web** web server with provided username, password and confirmPassword
- Validation is performed
- The **web** microservice sends a request to **account** microservice to create a new account
- **account** saves the new account to a PostgreSQL database
- **web** receives a response with new user's credentials, creates a cookie and redirects to the
    home page

## Login to your account

Login to your Ad Campaigns account by passing your credentials.

To achieve this,
- An HTTP request is sent to **web**, to log in to the user account.
- **web** sends a request to the **account** service to fetch an account with provided username
- **web** will then verify that the account exists and that the password is correct, create 
    a cookie and a JWT token and send the token back to the client


## See an overview of all Ad Campaigns

The user will see an overview of all the available ad campaigns in a table.

To achieve this,
- An HTTP request is sent to **web** to fetch all the ad campaigns 
- **web** sends a request to **adcampaign** to get all the ad campaigns that were saved to a MongoDB database 
    during the **adcampaign** service startup
- **adcampaign** gets all the ad campaigns from the MongoDB database and returns it to **web**
- **web** marshals the ad campaign objects to JSON and sends the response back to the client

## See a detailed view of a single Ad Campaign

By clicking on an ad campaign with a specific ID within the overview, a detailed view will be shown to the user 
in a modal. The JSON response from the web server will be rendered by React.js into a combined view.

