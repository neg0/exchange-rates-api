# Foreign Exchange Rates API
Provides latest rates for GBP and USD currencies and a prediction based on last seven days (a week) to 
sell the targeted currency.


## Available Endpoints
Base URL is localhost with port `8091` _(http://localhost:8091)_. There are three endpoints available:

 * __Health:__ `/v1/health`
 * __Rates:__ `/v1/rates?base=EUR`
 * __Predict:__ `/v1/predict?currency=USD`


### Health
This is for testing purpose to ensure the health of the application and also used 
for Health check of docker container inside the Dockerfile at the root of the project


### Rates
Gives a rate of two main currencies:
 * __United Kingdom:__ GBP
 * __United States:__ USD

by default base currency used for conversion is Euro (_EUR_) but it could be changed by 
passing string query `base`. E.g. `/v1/rates?base=GBP`


### Predict
It's shallow algorithm identifying if it's good to sell based on `currency` parameter E.g. 
`/v1/predict?currency=GBP` _if `currency` is not passed by default it will be `GBP`_

> It calculates days that currency is less than today's rate if number is higher than half of
period of time, it will suggest to sell otherwise it will respond sell by false value in JSON


# Architecture & Software Patterns
I have taken DDD approach for development of this application. Unit tests are also provided 
but due to time limit on this assessment unit tests are not as rigorous that why I can't say 
it's TDD :)


## Domain Driven Architecture
This application created using Domain Driven Model (Rich Domain Model). There are three layers 
defined and for ease of understanding separated to relevant package name.

 * __Application:__ contains endpoints (controllers) for two main functionality and determining health of application
 
 * __Domain:__ contains our business logic that infrastructure layer implements and provides main business logic such as: prediction to sell or not withing a week period window 
 
 * __Infrastructure:__ contains API calls to Exchange Rates API that are being utilised in application layer


## Design Patterns & SOLID Principles
Below I have given few example of patterns and SOLID Principles within this repository. I have provided comments in 
the code when I thought there might be a question.

 * __Future/Promise Concurrency Pattern:__ used for HTTP calls to external API inside 
 infrastructure layer
 
 * __Interface Segregation Principle & Polymorphism:__ Used to create rate object inside infrastructure, it creates 
 a bridge between infrastructure and domain layer

 * __Single Responsibility Principles:__ All packages created with this principle in mind, if it's not it's a mistake :P
 
 * __Open/Closed Principle & Factory Pattern:__ Used for creation of objects from payload, you could see examples such as: `fromJSON`
 
 

# Build & Deployment
I have created a _Dockerfile_ for build and deployment of this microservice and _Makefile_ 
for ease of use. ExchangeRatesAPI is defined inside `.env.dist` which should be renamed to 
`.env` to be used when making http calls.


## Docker and Make
Please use `make` command to build the Docker image for deployment by:

    ~$: make build

then you could run following to create and run a container from built image:

    ~$: make up
    
after running a container, you should be able to access all defined endpoints in this documentation 
on your local machine.

when you are finished with application you could run following to exit and remove the container:
    
    ~$: make down


### Credit
Built with :heart: & :coffee: at the heart of London