![flixinit](logo.png)
# flixinit

Flixinit is a simple CLI tool to make your application a great tenant for cloud environments.
Complete documentation is available at https://github.com/saeedafshari8/flixinit

Usage:<br />
flixinit [flags]<br />
flixinit [command]

Available Commands:<br />
help &nbsp;&nbsp; Help about any command<br />
java &nbsp;&nbsp; java command generates a new spring/java project

Flags:<br />
  -a, --author string &nbsp;&nbsp; author name for copyright attribution (default "Saeed Afshari")<br />
      --config string &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; config file (default is $HOME/.flixinit.yaml)<br />
  -h, --help &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; help for flixinit<br />
  --license string &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp; name of license for the project (default "Apache 2.0")<br />

Use "flixinit [command] --help" for more information about a command.

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
**java**<br/>
Use java command to generate a spring boot application. This command use [SpringInitializr](https://start.spring.io/) service
to create the project.<br/>

**Usage**:<br/>
*     flixinit java [flags]
**Flags:**<br/>
*     -v, --app-version string           Spring boot application version (default is empty and there will not be any version defined for the project)
*     --description string               Spring application description
*     -g, --group string                 Spring application groupId (default is empty)
*     -h, --help                         help for java
*     -j, --java-version string          Gradle (java)sourceCompatibility version (default is 11) (default "11")
*     -l, --language string              Spring project language [java | kotlin | groovy] (default is java) (default "java")
*     --name string                      Spring application name
*     --spring-boot-version string       Spring boot version (default is 2.2.1.RELEASE) (default "2.2.1.RELEASE")
*     -t, --type string                  Spring project type [gradle-project | maven-project] (default is gradle-project) (default "gradle-project")


# Installing
# Getting Started
**Use the flixinit Generator**
**Example**
**Help Command**
# Contributing
# License

flixinit is released under the Apache 2.0 license. See [LICENSE.txt](https://github.com/saeedafshari8/flixinit/blob/master/LICENSE.txt)