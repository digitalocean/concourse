module github.com/concourse/concourse

require (
	code.cloudfoundry.org/clock v0.0.0-20180518195852-02e53af36e6c
	code.cloudfoundry.org/credhub-cli v0.0.0-20190415201820-e3951663d25c
	code.cloudfoundry.org/garden v0.0.0-20181108172608-62470dc86365
	code.cloudfoundry.org/lager v2.0.0+incompatible
	code.cloudfoundry.org/localip v0.0.0-20170223024724-b88ad0dea95c
	code.cloudfoundry.org/urljoiner v0.0.0-20170223060717-5cabba6c0a50
	github.com/Azure/go-autorest/autorest v0.11.4 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.2 // indirect
	github.com/DataDog/datadog-go v3.2.0+incompatible
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace v0.11.0
	github.com/Masterminds/squirrel v1.1.0
	github.com/NYTimes/gziphandler v1.1.1
	github.com/aryann/difflib v0.0.0-20170710044230-e206f873d14a
	github.com/aws/aws-sdk-go v1.34.9
	github.com/caarlos0/env v3.5.0+incompatible
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cloudfoundry/go-socks5 v0.0.0-20180221174514-54f73bdb8a8e // indirect
	github.com/cloudfoundry/socks5-proxy v0.0.0-20180530211953-3659db090cb2 // indirect
	github.com/concourse/baggageclaim v1.8.0
	github.com/concourse/dex v0.3.0
	github.com/concourse/flag v1.1.0
	github.com/concourse/go-archive v1.0.1
	github.com/concourse/retryhttp v1.0.2
	github.com/containerd/cgroups v0.0.0-20191220161829-06e718085901 // indirect
	github.com/containerd/containerd v1.3.4
	github.com/containerd/continuity v0.0.0-20191214063359-1097c8bae83b // indirect
	github.com/containerd/fifo v0.0.0-20191213151349-ff969a566b00 // indirect
	github.com/containerd/go-cni v0.0.0-20200107172653-c154a49e2c75
	github.com/containerd/ttrpc v0.0.0-20191028202541-4f1b8fe65a5c // indirect
	github.com/containerd/typeurl v0.0.0-20190911142611-5eb25027c9fd
	github.com/coreos/go-iptables v0.4.5
	github.com/cppforlife/go-semi-semantic v0.0.0-20160921010311-576b6af77ae4
	github.com/cyberark/conjur-api-go v0.6.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/go-events v0.0.0-20190806004212-e31b211e4f1c // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/fatih/color v1.10.0
	github.com/felixge/httpsnoop v1.0.1
	github.com/frankban/quicktest v1.10.0 // indirect
	github.com/gobuffalo/packr v1.13.7
	github.com/goccy/go-yaml v1.8.8
	github.com/gogo/googleapis v1.3.1 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e
	github.com/google/jsonapi v0.0.0-20180618021926-5d047c6bc66b
	github.com/googleapis/gnostic v0.4.0 // indirect
	github.com/gophercloud/gophercloud v0.12.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/go-hclog v0.14.0 // indirect
	github.com/hashicorp/go-multierror v1.1.0
	github.com/hashicorp/go-rootcerts v1.0.2
	github.com/hashicorp/vault/api v1.0.5-0.20191108163347-bdd38fca2cff
	github.com/hashicorp/vault/sdk v0.1.14-0.20191112033314-390e96e22eb2 // indirect
	github.com/imdario/mergo v0.3.6
	github.com/inconshreveable/go-update v0.0.0-20160112193335-8152e7eb6ccf
	github.com/influxdata/influxdb1-client v0.0.0-20191209144304-8bf82d3c094d
	github.com/jessevdk/go-flags v1.4.0
	github.com/klauspost/compress v1.10.10
	github.com/kr/pty v1.1.8
	github.com/kr/text v0.2.0 // indirect
	github.com/krishicks/yaml-patch v0.0.10
	github.com/lib/pq v1.3.0
	github.com/markbates/pkger v0.17.1
	github.com/mattn/go-colorable v0.1.8
	github.com/mattn/go-isatty v0.0.12
	github.com/maxbrunsfeld/counterfeiter/v6 v6.2.3
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/miekg/dns v1.1.31
	github.com/mitchellh/mapstructure v1.3.2
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/onsi/ginkgo v1.13.0
	github.com/onsi/gomega v1.10.1
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/runtime-spec v1.0.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/peterhellberg/link v1.0.0
	github.com/pierrec/lz4 v2.5.2+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/pkg/term v0.0.0-20190109203006-aa71e9d9e942
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.13.0 // indirect
	github.com/racksec/srslog v0.0.0-20180709174129-a4725f04ec91
	github.com/satori/go.uuid v1.2.1-0.20181028125025-b2ce2384e17b // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/skratchdot/open-golang v0.0.0-20160302144031-75fb7ed4208c
	github.com/square/certstrap v1.1.1
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/syndtr/gocapability v0.0.0-20180916011248-d98352740cb2 // indirect
	github.com/tedsuo/ifrit v0.0.0-20180802180643-bea94bb476cc
	github.com/tedsuo/rata v1.0.1-0.20170830210128-07d200713958
	github.com/vbauerster/mpb/v4 v4.6.1-0.20190319154207-3a6acfe12ac6
	github.com/vito/go-interact v0.0.0-20171111012221-fa338ed9e9ec
	github.com/vito/go-sse v0.0.0-20160212001227-fd69d275caac
	github.com/vito/houdini v1.1.1
	github.com/vito/twentythousandtonnesofcrudeoil v0.0.0-20180305154709-3b21ad808fcb
	go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.11.0
	go.opentelemetry.io/otel/sdk v0.11.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sync v0.0.0-20201020160332-67f06af15bc9
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/api v0.34.0 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/square/go-jose.v2 v2.4.1
	gopkg.in/yaml.v2 v2.3.0
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/client-go v0.18.8
	k8s.io/utils v0.0.0-20200414100711-2df71ebbae66 // indirect
	sigs.k8s.io/yaml v1.2.0
)

go 1.13

replace github.com/docker/distribution v2.7.1+incompatible => github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible

replace github.com/jessevdk/go-flags => github.com/vito/go-flags v1.4.1-0.20200428200343-c7161c3bd74d
