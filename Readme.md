# Image Provider / Document Store

Simple api to save and retrieve images.


## Create Images

POST `/{storage}`

Creates a folder in ./storage and returns the document Id as a Guid.


POST `/{client}/{folder}`

Created a folder in ./static/{client}/{folder} that can be served as static content. Returns the image path


## Get images

With DocumentId an image can be retrieved by the call

GET `/{storage}/{id}`

You can also get images in the static folder by 

GET `/static/{path}`



## Dependencies

1) github.com/gorilla/mux
2) github.com/gofrs/uuid