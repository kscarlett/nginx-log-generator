# Nginx Log Generator

A tiny Go utility to generate a large amount realistic-looking Nginx logs quickly. It was written to aid in testing logging pipelines and other such tools, and demoing them in Kubernetes.

Most of the heavy lifting is done by the amazing [gofakeit](https://github.com/brianvoe/gofakeit) library, with some extra work to skew the results towards typical values.

## Usage

The most important step is to set the desired rate in the `RATE` environment variable. The simplest way to do this is the following:

```sh
$ # Will generate 10 entries per second
$ RATE=10 ./nginx-log-generator
```

The reason this is an environment variable is so it's easier to run via Docker as well:

```sh
$ docker pull kscarlett/nginx-log-generator
$ docker run -e "RATE=10" kscarlett/nginx-log-generator
```

### Configuration

The following environment variables can be set to modify the output.

| Name              | Default | Notes                                                                                                                                           |
| ----------------- | ------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| RATE              | 1       | Logs to output per second.                                                                                                                      |
| IPV4_PERCENT      | 100     | Percentage of IP addresses that will be IPv4. Change to 0 to only get IPv6.                                                                     |
| STATUS_OK_PERCENT | 80      | _Roughly_ the percentage of `200` status codes. The rest will be randomised and may contain `200` as well.                                      |
| PATH_MIN          | 1       | Minimum elements to put in the request path.                                                                                                    |
| PATH_MAX          | 5       | Maximum elements to put in the request path.                                                                                                    |
| GET_PERCENT       | 60      | Percentage of requests that will be `GET` requests. If the total adds up to less than 100%, the rest will be made up of random HTTP methods.    |
| POST_PERCENT      | 30      | Percentage of requests that will be `POST` requests. If the total adds up to less than 100%, the rest will be made up of random HTTP methods.   |
| PUT_PERCENT       | 0       | Percentage of requests that will be `PUT` requests. If the total adds up to less than 100%, the rest will be made up of random HTTP methods.    |
| PATCH_PERCENT     | 0       | Percentage of requests that will be `PATCH` requests. If the total adds up to less than 100%, the rest will be made up of random HTTP methods.  |
| DELETE_PERCENT    | 0       | Percentage of requests that will be `DELETE` requests. If the total adds up to less than 100%, the rest will be made up of random HTTP methods. |

## Note

This is a tool I made in no time at all, because I needed a tool that did exactly this right that second. The code quality isn't optimal and it can probably be optimised. I will be coming back to it at some other time.

## License

This tool is released under the [MIT License](LICENSE).