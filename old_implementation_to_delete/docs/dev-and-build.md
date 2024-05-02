# Develope

## UI

The use makes use of Gulp



    cd web/app/
    gulp watch
    
this will complile, unify and minify the js / css / html files present in this folders, and put them in the static folder

## About webfonts
The fontello webfont is a special case, when the font is updated, the uncompressed needs to go in web/app but you will 
need to copy the fonts into static/font  

# Build

## UI
    cd web/app/
    gulp build
    
## packr2 static files
see https://github.com/gobuffalo/packr/tree/master/v2

    cd internal/fe26/
adapt the import of packrd to something like:

    import _ "github.com/AndresBott/Fe26/internal/fe26/packrd"

and then:
  
    go build


## go binary

    go get -u -v github.com/go-task/task/cmd/task
    task build
    
## circleCI

circleCI will create a new release for every pushed code for witch the build does not fail
