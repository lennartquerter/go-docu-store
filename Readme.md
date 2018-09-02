# Image Provider / Document Store

Simple api to save and retrieve images.


## Create Images

POST `/api/v1/{storage}`

Creates a folder in ./storage and returns the document Id as a Guid.


POST `/api/v1/static/{client}/{folder}`

Created a folder in ./static/{client}/{folder} that can be served as static content. Returns the image path


## Get images

With DocumentId an image can be retrieved by the call

GET `/{storage}/{id}`

You can also get images in the static folder by 

GET `/static/{path}`


## Todo

1) Image compression
2) retrieve resized images
3) Save thumbnails


## Dependencies

1) github.com/gorilla/mux
2) github.com/gofrs/uuid

