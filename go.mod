module github.com/tidb-community-bots/prow-github

go 1.15

replace (
	k8s.io/api => k8s.io/api v0.19.3
	k8s.io/client-go => k8s.io/client-go v0.19.3
)

require (
	github.com/google/go-cmp v0.5.2
	github.com/prometheus/client_golang v1.7.1
	github.com/shurcooL/githubv4 v0.0.0-20191102174205-af46314aec7b
	github.com/sirupsen/logrus v1.6.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	k8s.io/apimachinery v0.19.3
	k8s.io/test-infra v0.0.0-20201114015505-f09ff0e80535
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800
	sigs.k8s.io/yaml v1.2.0
)
