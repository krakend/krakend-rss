Krakend RSS
====

Enable RSS encoding in the [KrakenD](github.com/devopsfaith/krakend) framework

**Notice: this package requires KrakenD >= 0.4**

## Documentation

For more details, check the auto-generated documentation: https://godoc.org/github.com/krakend/krakend-rss

## Using it

Just register the module before parsing your config and everything will work smoothly

	rss.Register()

## Build the example

Go 1.8 is a requirement

	$ make

## Run

Running it as a common executable, logs are send to the stdOut and some options are available at the CLI

	$ ./krakend_rss_example
	Usage of ./krakend_rss_example:
	  -c string
	    	Path to the configuration filename (default "/etc/krakend/configuration.json")
	  -d	Enable the debug
	  -p int
	    	Port of the service

This is an example of a suitable json configuration

	{
	    "version": 1,
	    "name": "TV shows gateway",
	    "port": 8080,
	    "cache_ttl": 3600,
	    "timeout": "3s",
	    "endpoints": [
	        {
	            "endpoint": "/showrss/{id}",
	            "backend": [
	                {
	                    "host": [
	                        "http://showrss.info/"
	                    ],
	                    "url_pattern": "/user/schedule/{id}.rss",
	                    "encoding": "rss",
	                    "group": "schedule",
	                    "allow": ["items", "title"]
	                },
	                {
	                    "host": [
	                        "http://showrss.info/"
	                    ],
	                    "url_pattern": "/user/{id}.rss",
	                    "encoding": "rss",
	                    "group": "available",
	                    "allow": ["items", "title"]
	                }
	            ]
	        }
	    ]
	}

You can see the resulting output with this simple curl command (replace `<YOUR_SHOWRSS_USER_ID>` with your own id)

	$ curl -i http://127.0.0.1:8080/showrss/<YOUR_SHOWRSS_USER_ID>