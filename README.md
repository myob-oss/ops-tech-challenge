# MYOB DevOps Technical Test

The goal of this challenge is to enable you to demonstrate how you work and
solve problems. There are only 2 _required_ components, plus some optional ones.
You have the option to only focus on the operations and deployment, or dive
into more of the software. We just want to know _your reasoning_ for your
decisions.

Please consider the following in your solution, we may ask you about them in
the interview:

- What security implications does my solution have?
- How does my solution scale?
- What assumptions have I made with my solution?
- Does my development process work well with a team of developers?

Note: if you have existing code that you think will demonstrate the principals
we are looking for, _and you are able to share it with us_, then feel free to
send us a link to the repository so we can discuss it in the interview.

## The challenge

This repository contains the code for a simple HTTP service with a number of
endpoints. The primary endpoint returns an HMAC from a shared secret and the
incoming request. The service is written in Go and should compile without
issue, creating a binary for use in the challenge.

There are only 2 required parts to this challenge, the other parts are optional:

1. Define a Dockerfile and a pipeline to build an artefact, explaining your
   choices and your method of handling secrets. Please use any CI configuration
   you see fit, such as Buildkite, Github Actions, Travis etc.

   Note: you do not need to actually deploy anything for this part.

2. Provide a Git repository containing your solution for part 1.

These items are optional, and in _no specific order_:

- (optional) Create a script or Makefile that automates the building of the Go binary.

- (optional) Find and fix the many issues in the Go source code, explaining the
  reasons for your changes. A definition of the service is given below in
  "Service interface".

- (optional) Deploy the artefact to a cloud provider such as AWS, Heroku etc.

- (optional) Add some additional tests that ensure the code is more robust.


## Building the binary

You will need a Go development environment installed to build the binary for
your pipeline in part 1. To build the binary execute the following:

```shell

go build

```

## Service interface

- POST /token

  Return a token based on a shared secret. The shared secret is passed in the
  environment as the variable `SECRET`.

- GET /health

  Used to check the service is 'up'. It should return an HTTP code >= 200

- GET /metrics

  Return some basic metrics about the running service.

All endpoints should:

- Respond with valid JSON.
- Return appropriate HTTP codes.
- Enable concurrent access.
