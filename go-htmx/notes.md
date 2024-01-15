#### Go web app

#### Running the Go server
```Go
    go run 'server.go'
```
This command will build and compile the server.


#### Go web app Architecture

In Go you can create your controllers as follows:
_long hand_
```go

http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "Hello, World!")
    w.WriteHeader(http.StatusNotFound)
    templ := template.Must(template.ParseFiles("404.html"))
})
```
_short hand_
```go
h1 := func(w http.ResponseWriter, r *http.Request) {
    // code here
}

#### pass to url handler
http.HandleFunc("/", h1) 
```

#### Expose the server
`http.ListenAndServe(":8000", nil)`
Go app will run on localhost and serve at http://localhost:8000.

#### Loops in Go - Html template
Assuming we have a list of film that we want to display on our index.html, we can use range to loop thro and siplay them.
```go
h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Matrix", Year: "1999", Director: "Wachowski"},
				{Title: "The Matrix Reloaded", Year: "2003", Director: "Wachowski"},
				{Title: "The Matrix Revolutions", Year: "2003", Director: "Wachowski"},
			},
		}
		tmpl.Execute(w, films)
	}
```
`index.html`
```html
<div class="container">
    <div class="col-8" id="films">
            <h1 class="mb-4">Films List</h1>
            <ul class="list-group fs-5 me-5">
                <!-- {{ range .Films }}
                    <li class="list-group-item bg-secondary text-white">{{ .Title }} - {{ .Director }} - {{ .Year }}</li>
                
                {{ end }} -->
                <!-- template fragment -->
                {{ range .Films }}
                    {{ block "film-list-item" . }}
                        <li class="list-group-item bg-secondary text-white">{{ .Title }} - {{ .Director }} - {{ .Year }}</li>
                    {{ end }}
                {{ end}}
            </ul>
        </div>
</div>
```

#### Templating in Go 
You render html in go using the `template.Must`
```go
    h1 := func(w http.ResponseWriter, r *http.Request) {
	    tmpl := template.Must(template.ParseFiles("index.html"))
    }
``` 

You can also render a template frage using the `template.ExecuteTemplate(w, "fragment eg(film-list-item)", context )` method instead of a whole page.

>Example.
>This contollers handles fiml creation. Using htmx, we can return a fragement, just the new change on the page. This will return the new film list.

```go
h2 := func(w http.ResponseWriter, r *http.Request) {
    // simulate latency to test hx-indicator
    time.Sleep(5 * time.Second)
    title := r.PostFormValue("title")
    year := r.PostFormValue("year")
    director := r.PostFormValue("director")
    templ := template.Must(template.ParseFiles("index.html"))
    templ.ExecuteTemplate(w, "film-list-item", Film{Title: title, Year: year, Director: director})
    
```
```go
http.HandleFunc('/addfilm/', h2)
```

##### Resources
Go net/http package: https://pkg.go.dev/net/http
Go html/template package: https://pkg.go.dev/html/template
Go templates: https://golangforall.com/en/post/temp...
HTMX Install: https://htmx.org/docs/#installing
HTMX Swap Methods: https://htmx.org/docs/#swapping


#### Side Note
##### Storing Git Credentials with Git Credential Helpe
“gitcredentials” module is used to request these credentials from the user as well as stores these credentials to avoid inputting these credentials repeatedly.

_Avoid Repetition_
### Git Credentials Helper
By default git credentials are not cached at all.
Every connection will prompt you for your username and password.

*Available models for credential.helper config*
1. cache
2. store
3. osxkeychain
4. manager

##### Git Credentials Helper: cache
Cache credentials in memory for a short period of time.
None of the passwords are ever stored on disk, and they are purged from the cache after `15 minutes` (default cache timeout).

`git config --global credential.helper cache`

__Increase timeout for `git config --global credential.helper`__
`git config --global credential.helper "cache --timeout=3600"`


##### Git Credentials Helper: store
Store credentials indefinitely on disk.
`git config --global credential.helper store`
By default, the git credentials in the “store” mode will be stored in the
_“.git-credentials” file in the user’s home directory (~/.git-credentials)_

You can also specify the `store` file like
`git config --global credential.helper "store --file ~/.my-credentials"`

> stores password in plaintext in own url like
> https://<url_encoded_plain_text_username>:<url_encoded_plain_text_password>@github.com

##### Git Credentials Helper: osxkeychain
This method stores the credentials on disk, and they never expire, but they’re encrypted with the same system that stores HTTPS certificates
`git config --global credential.helper osxkeychain`

##### Git Credentials Helper: manager
For Windows using e Git Credential Manager for Windows (GCM).
`git config --global credential.helper manager`