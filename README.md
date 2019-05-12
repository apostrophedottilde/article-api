# blog-article-api #

## Deployment instructions ##
This app is containerised with docker and I am using docker-compose to deploy the API and the database container together.
- Please use the command <code>docker-compose up --build</code> in the root directory of this project and the stack should start up correctly.

- Unit tests will be run during building the docker image, however if you want to run them individually please run the following command from the command line:
<code>go test ./...</code> from the terminal in the root of this project.

## Assumptions ##
- By 'free' schema I think you mean that the fields needed for an article will not be known in advance so I should be able to supply any fields.
- You would be interested to know when the article was created and updated.
- The ID of each article does not necessarily need to be an integer.
- You are able to run a dockerised application stack using docker-compose.

## Design justifications ##
- I have chosen XML mime type for this API because more metadata can be easily applied to XML compared with JSON. 
- XSLT can be used to translate the unstructured (schema-less) XML into a desired format. For example it could be easily translated to HTML for web browsing. 
- XML feels like a good choice because the text can be 'marked up' to describe the structure.
- Generic style can be specified in XML element attributes and easily rendered differently for different devices or even ignored where not necessary.
- I chose a mongoDB datastore because it provides an out-of-the-box support for text search which would require a more complex solution using something like Apache Lucene if I chose a relational store such as MYSQL.
- I am not certain that storing an embedded XML document as a string inside an attribute of a parent XML document is the best idea in 
production, however it was an inventive way to solve the problem as I did not have the available time to find a more elegant solution.

## Endpoints ##
All endpoints consume and/or expose mime type 'application/xml'
- <code>POST /articles</code> - A body of XML is passed in this POST request and is persisted to the datastore. A location header specifying the new resource will be added to the response.
- <code>GET /articles/{articleId}</code> - Returns the Article with this if it exists and a 404 response is returned otherwise.
- <code>GET /articles</code> - A collection of all stored Articles is returned.
- <code>PUT /articles/{articleId}</code> - An XML body is passed in this PUT request and the Article with the specified ID is updated.

*** NOTE *** For the POST and PUT requests: You <b>must</b> pass valid and well formed XML as the reqest body or a 400 error code will be returned with a message explaining this.

## Limitations ##
- Because the structure of the xml data is unknown up-front I have stored it as a string in the entity model. As a side effect,
 depending on which rest client I use to consume this service I may see escape characters in the content attribute in the response. 
 This may mean the content needs to be unescaped by the consumer but it may not be necessary.
- The error codes returned upon error are not all perfect. If I had more time I would pass more context about the nature
of the error from the repository layer up to the controller where I could use more specific error code. They should mostly be acceptable.
- <b>I have only had time to test the usecase package, but there is a full test suite available. The testing framework is in place and it 
would not be difficult to test the rest.</b>