# GO Image Builder

Create and serve images according to parameters passed in URL

## Up and Running

    $ go run main.go utils.go imagebuilder.go

Type something like this in your browser:

    http://localhost:8080/?width=500&height=50&background=abc123

or even better, use like a placeholder for images in HTML

```html
<img src="http://localhost:8080/?width=500&height=50&background=abc123" alt="Image Generated on Server">
```

## More
It's a work in progress