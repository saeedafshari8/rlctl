![flixinit](logo.png)
# Flixinit

Flixinit is a simple CLI tool to make your application a great tenant for cloud environments.

# Table of Contents

- [Overview](#overview)
- [Concepts](#concepts)
  * [Commands](#commands)
    * [spring](#spring)
    * [gitlab](#gitlab)
      * [namespaces](#namespaces)
- [Installing](#installing)
  * [Building binary](#bulding-binary)
    * [MAC OS](#mac-os)
    * [Linux](#Linux)
    * [Windows](#Windows)
- [Getting Started](#getting-started)
  * [Example1](#example1)
  * [Example2](#example2)
- [Contributing](#contributing)
- [License](#license)

# Overview
It automates everything from generating different types of projects to configuring CI pipeline and deployment on Kubernetes clusters.
To make the claim more tangible, usually in order to create a Spring Boot project and assign it to a development team the following steps needs to be taken:
- Create a project repository (ex: using [gitlab](https://about.gitlab.com/))
- Create a new Spring Boot project (ex: using [spring initializr](https://start.spring.io/))
- Configure .gitignore
- Create multiple environments configuration files (application-{environment name}.yml) 
- Containerise the project (ex: using Dockerfile)
- Create Kubernetes deployment and service configurations
- Create .gitlab-ci.yml with proper steps from build, test, integration test, ..., deploy on kubernetes
- Commit and push the generated project to the VCS

In addition to aforementioned steps, usually many projects needs to be integrated with a couple of tools like:
- Quality check tool (ex: [SonarQube](https://www.sonarqube.org))
- OAuth2 active directory and Oauth2 API
- Newrelic
- Kafka
- Schema Registry (ex: [confluence](https://docs.confluent.io/current/schema-registry/index.html))
- and so on so forth

So as you can see provisioning a new projects needs a great effort and time. However, you might think, as you need to create projects
once in a blue moon, it does not make any sense to think about automation. Nevertheless, the life is not as easy as you think
and specially with emerge of Microservices architecture, it has become even more complicated.
With the new architecture, teams needs to create more and more fine-grained projects and continuously involve in system integrations.
Therefore, one of the consequences of migrating from monolith to scalable Microservices which are aligned with business needs is acceptance
of the fact for creating and integrating new smaller granular systems. 
To do so, we started flixinit to help developers bring concentrate back on writing codes instead of repeating time consuming amd error prone 
tasks of project provisioning.     

# Concepts
## Commands

### spring

Use spring command to generate a spring boot application. This command uses [SpringInitializr](https://start.spring.io/) service
to create the project.

***Usage***

`flixinit spring [flags]`

***Flags***

| ***Flag*** | ***Description*** |
| ----------- | ----------- |
|       --azure-enabled                      |Enable Azure Active Directory |
|       --build-tool string                  |Spring project type [gradle-project | maven-project] (default "gradle-project") |
|       --container-image string             |Docker exposed port (default "openjdk:11.0.5-jdk-stretch") |
|       --container-port string              |Docker exposed port (default "8080") |
|       --container-registry string          |Docker Registry URL (default "dcr.flix.tech/charter/cust") |
|       --description string                 |Spring application description |
|       --git-repo-url string                |git remote repository url |
|       --gitlab-ci-enabled                  |Create .gitlab-ci config (default true) |
|       --gitlab-ci-except stringArray       |.gitlab-ci except (default [schedules]) |
|       --gitlab-ci-tags stringArray         |.gitlab-ci tags (default [docker,autoscaling]) |
|   -g, --group string                       |Spring application groupId |
|   -h, --help                               |help for spring |
|   -j, --java-source-compatibility string   |Java source compatibility version (default "11") |
|       --jpa-database string                |JPA Database Name (default "MYSQL") |
|       --jpa-enabled                        |Enable JPA-Hibernate (default true) |
|       --kafka-enabled                      |Enable Kafka integration |
|   -l, --language string                    |Spring project language [java , kotlin , groovy] (default "java") |
|       --liquibase-enabled                  |Enable Liquibase migration |
|       --name string                        |Spring application name |
|       --security-enabled                   |Enable Spring security |
|       --security-oauth2                    |Enable OAuth2 |
|       --server-host string                 |Spring application base url host (default "localhost") |
|       --server-port string                 |Spring boot application port (default "8080") |
|       --server-protocol string             |Spring application base url protocol (default "http") |
|       --spring-boot-version string         |Spring boot version (default "2.2.4.RELEASE") |
|   -v, --version string                     |Spring boot application version |


### gitlab

Use gitlab command to generate a new project in gitlab repository.

***Usage***

`flixinit gitlab [flags]`

`flixinit gitlab [command]`

***Available Commands***:

namespaces  gitlab command get list of existing namespaces.

***Flags***

| ***Flag*** | ***Description*** |
| ----------- | ----------- |
|      --approvals_before_merge int32                       | How many approvers should approve merge requests by default (default 1) |
|       --auto_cancel_pending_pipelines string              | Auto-cancel pending pipelines (Note: this is not a boolean, but enabled/disabled (default "enabled") |
|   -h, --help                                              | help for gitlab |
|   -r, --initialize_with_readme                            | Initialise by README.md (default true) |
|   -n, --name string                                       | The name of the new project. Equals path if not provided. |
|       --namespace_id int32                                | Namespace for the new project (defaults to the current user’s namespace) |
|       --only_allow_merge_if_all_discussions_are_resolved  | Set whether merge requests can only be merged when all the discussions are resolved (default true) |
|      --only_allow_merge_if_pipeline_succeeds              | Auto-cancel pending pipelines (Note: this is not a boolean, but enabled/disabled (default true) |
|   -p, --path string                                       | Repository name for new project. Generated based on name if not provided (generated lowercased with dashes). |
|       --token string                                      | Gitlab token. |
|   -v, --visibility string                                 | private|internal|public (default "private") |

### gitlab namespaces
To get list of existing namespaces.

***Usage***
`flixinit gitlab namespaces [flags]`

***Flags***

| ***Flag*** | ***Description*** |
| ----------- | ----------- |
|       -h, --help           | help for namespaces |
|      --token string        | Gitlab token. |

# Installing

The only thing you need to have is the executable file. Thanks packr (https://github.com/gobuffalo/packr/tree/master/v2).

## Building binary
To build the binary file for a specific operating system, you can execute the following commands in the root of project. 

### MAC OS
      GOOS=darwin GOARCH=amd64 packr build

### Linux
      GOOS=linux GOARCH=amd64 packr build

### Windows 
      GOOS=windows GOARCH=386 packr build

# Getting Started

## Example1: Create Spring Boot-Gradle application** 

`flixinit spring --group=charter.flixbus.com --name=testApp -v=1.0.0 --liquibase=true --jpa=true --language=kotlin --kafka=true --git-remote=https://git.com/...`
    
## Example2: Create Spring Boot-Gradle with all flags placed** 

`flixinit spring --group=com.example --name=sample --description="Sample application" --language=java --version=0.0.1 --java-source-compatibility=11 --build-tool=gradle-project --spring-boot-version=2.2.5.RELEASE --server-port=9090 --server-host=0.0.0.0 --server-protocol=http --jpa-enabled=true --jpa-database=MYSQL --liquibase-enabled=true --security-enabled=true --security-oauth2=true --kafka-enabled=true --azure-enabled=true --container-port=9999 --container-image=jdk-11.0.6_10-alpine-slim --container-registry=dcr.flix.tech/charter/cust --gitlab-ci-enabled=true --gitlab-ci-tags=docker --gitlab-ci-tags=autoscaling --gitlab-ci-except=schedules`

The above command create a spring boot/gradle application with the following config:
- groupId = charter.flixbus.com
- artifactId = tesApp
- Dockerfile
- application.yml|application-int.yml|application.local|application-prod.yml
- Enable JPA-Hibernate with MYSQL config
- Liquibase integration

# Contributing
# License

flixinit is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/saeedafshari8/flixinit/blob/master/LICENSE.txt)