Bulk DNS Lookup
===============

Author: **[Afaan Bilal](https://afaan.dev)**

## Introduction
**Bulk DNS Lookup** is a small program written in Go that takes a newline separated text file of domain names and returns
the IP addresses associated with them in a JSON file.  

To run:
````
go run src/bulk-dns-lookup.go src/input.txt src/output.json
````
  
Example input (`input.txt`):
````
google.com
github.com
````
  
Example output (`output.json`):
````
[
  {
    "domain": "google.com",
    "ips": [
      "74.125.24.102",
      "74.125.24.113",
      "74.125.24.138",
      "74.125.24.139",
      "74.125.24.100",
      "74.125.24.101"
    ]
  },
  {
    "domain": "github.com",
    "ips": [
      "52.74.223.119"
    ]
  }
]
````

## Contributing
All contributions are welcome. Please create an issue first for any feature request
or bug. Then fork the repository, create a branch and make any changes to fix the bug 
or add the feature and create a pull request. That's it!
Thanks!

## License
**Bulk DNS Lookup** is released under the MIT License.
Check out the full license [here](LICENSE).
