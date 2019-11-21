![flixinit](logo.png)
# flixinit

Flixinit is a simple CLI tool to make your application a great tenant for cloud environments.
Complete documentation is available at https://github.com/saeedafshari8/flixinit

# Table of Contents

- [Overview](#overview)
- [Concepts](#concepts)
  * [Commands](#commands)
- [Installing](#installing)
- [Getting Started](#getting-started)
  * [Using the flixinit Generator](#using-the-flixinit-generator)
  * [Example](#example)
  * [Help Command](#help-command)
- [Contributing](#contributing)
- [License](#license)

# Overview
# Concepts
# Commands
**java**

Use java command to generate a spring boot application. This command uses [SpringInitializr](https://start.spring.io/) service
to create the project.

**Usage**:

     flixinit java [flags]
**Flags:**

     --app-host string               Spring application base url host (default localhost) (default "localhost")
          --app-port string               Spring boot application port (default is 8080) (default "8080")
          --app-protocol string           Spring application base url protocol (default http (default "http")
      -v, --app-version string            Spring boot application version (default is empty and there will not be any version defined for the project)
          --azure-ad                      Enable Azure Active Directory (default false)
      -i, --container-image string        Docker exposed port (default is openjdk:11.0.5-jdk-stretch) (default "openjdk:11.0.5-jdk-stretch")
      -p, --container-port string         Docker exposed port (default is 8080) (default "8080")
          --database string               JPA Database Name (default is MYSQL) (default "MYSQL")
          --description string            Spring application description
          --docker-registry string        Docker Registry URL (default is https://index.docker.io/v1) (default "https://index.docker.io/v1")
          --git-remote string             git remote repository url
          --gitlabci                      Create .gitlab-ci config (default is true) (default true)
          --gitlabci-except stringArray   .gitlab-ci except (default is schedules) (default [schedules])
          --gitlabci-tags stringArray     .gitlab-ci tags (default is docker,autoscaling) (default [docker,autoscaling])
      -g, --group string                  Spring application groupId (default is empty)
      -h, --help                          help for java
      -j, --java-version string           Gradle (java)sourceCompatibility version (default is 11) (default "11")
          --jpa                           Enable JPA-Hibernate (default is true) (default true)
      -l, --language string               Spring project language [java | kotlin | groovy] (default is java) (default "java")
          --liquibase                     Enable Liquibase migration (default is false)
          --name string                   Spring application name
          --oauth2                        Enable OAuth2 (default false)
          --security                      Enable Spring security (default false)
          --spring-boot-version string    Spring boot version (default is 2.2.1.RELEASE) (default "2.2.1.RELEASE")
      -t, --type string                   Spring project type [gradle-project | maven-project] (default is gradle-project) (default "gradle-project")


# Installing
# Getting Started
**Use the flixinit Generator**

**Example**

Example1: Create Spring Boot/Gradle application 

    flixinit java --group=charter.flixbus.com --name=testApp -v=1.0.0 --liquibase=true --jpa=true

The above command create a spring boot/gradle application with the following config:
* groupId = charter.flixbus.com
* artifactId = tesApp
* Dockerfile
* application.yml|application-int.yml|application.local|application-prod.yml
* Enable JPA-Hibernate with MYSQL config
* Liquibase integration

# Contributing
# License

flixinit is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/saeedafshari8/flixinit/blob/master/LICENSE.txt)