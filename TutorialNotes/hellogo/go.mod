module github.com/nicholas-karimi/hellogo

go 1.21.5

replace github.com/nicholas-karimi/mystrings v0.0.0 => ../mystrings // tell go no to look for the package remotely but look locally for mystrings

require ( 
    github.com/nicholas-karimi/mystrings v0.0.0
)