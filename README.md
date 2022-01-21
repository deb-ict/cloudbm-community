# CloudBM Community
Open Source Cloud Business Management suite

---

## Introduction
Cloud Business Management is an open source web application to handle daily business activities.

**Available modules**
- N/A

**Future modules**
- Employee management
- Customer management
- Product management
- Project management
- Invoice mangement
- Time registration
- Service desk

## Getting started

##### Running on [Go environment].

```
mkdir -p $GOPATH/src/deb-ict
cd $GOPATH/src/deb-ict
git clone https://github.com/deb-ict/cloudbm-community
cd cloudbm-community
go run ./cmd/webhost
```

##### Running on [Docker environment].

```
docker build -f build/container/Dockerfile -t cloudbm/community:dev .
docker run -d -p 5000:80 cloudbm/community:dev
```

##### Running on [Kubernetes environment].

*comming soon*

## Hosted Cloud Business Management (Enterprise edition)

*comming soon*

[DEB-ICT] will provide a hosted version of the application, known as enterprise edition.\
Enterprise edition is a multi-tenant microservice version designed for [Kubernetes environment].


More information: https://www.cloudbm.eu

[DEB-ICT]: https://www.deb-ict.com
[Go environment]: https://golang.org/doc/install
[Docker environment]: https://docs.docker.com/engine
[Kubernetes environment]: https://kubernetes.io