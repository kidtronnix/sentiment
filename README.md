# sentiment

Example of google language service implementation in golang.

[![Build Status](https://travis-ci.org/smaxwellstewart/sentiment.svg?branch=master)](https://travis-ci.org/smaxwellstewart/sentiment)


## requirements

- golang (1.8+) - [follow install docs](https://golang.org/doc/install)
- gcloud - [follow install docs](https://cloud.google.com/sdk/gcloud/)
- a google user account with permissions to use Language API

## usage

### Setup! Authenticate with Google APIs

To run this code your local enviroment needs to have google's [default application credentials](https://developers.google.com/identity/protocols/application-default-credentials) setup.

```
$ gcloud init
$ gcloud auth application-default login
```

Follow the prompts and login with your google creds. For [more info](https://cloud.google.com/ml-engine/docs/command-line).

You should now have a file under `~/.config/gcloud/application_default_credentials.json` that looks something like this:

```
$ cat ~/.config/gcloud/application_default_credentials.json
```

### Run! Start the web service

There is no `golang` dependencies to install because the code is shipped pre-vendored. I have found this the best approach for portability with Golang. In my experience the size of the `vendor` folder has never become an issue for my use cases.

To start the service locally on port `8000`, in the root of the project...

```
$ go run main.go -addr :8000 -key YOUR_GOOGLE_LANGUAGE_API_KEY
```

### Test! To run test suite

To run all tests (except eveything in vendor folders). Wanring!
The tests in analyze require internet connection and will cause usage of Language API.

```
go test -v $(go list ./... | grep -v /vendor/)
```
## api

### **POST** `/api?order=<asc|desc>&limit=<positive integer>`

Default order is `desc`.
In case of tied sentiment score oredering shall proceed alphabetically for either `asc` or `desc` order. See [tests](https://github.com/smaxwellstewart/sentiment/blob/master/word/sort_test.go) for more details.

Default limit is 10.


## todo (improvements / clarifications)

- Improve / extend coverage of tests.
- Currently all punctuation is stripped, but certain punctuation can effect sentiment! Stakeholder clarification needed.
- Improvement, deploy to Google cloud using Travis CI. [Travis docs](https://docs.travis-ci.com/user/deployment/google-app-engine/) and [Google docs](https://cloud.google.com/solutions/continuous-delivery-with-travis-ci).
