import org.jetbrains.kotlin.gradle.plugin.KotlinSourceSet
import org.jetbrains.kotlin.gradle.tasks.KotlinCompile
import org.springframework.boot.gradle.tasks.run.BootRun

buildscript {
    repositories {
        jcenter()
        maven("https://plugins.gradle.org/m2/")
        {{if eq .EnableKafka true}}
        maven("http://packages.confluent.io/maven/")
        {{end}}
    }
    dependencies {
        {{if eq .EnableKafka true}}
        classpath("com.github.imflog:kafka-schema-registry-gradle-plugin:0.8.0")
        {{end}}
        classpath(kotlin("gradle-plugin:1.3.61"))
    }
}

plugins {
    id("org.springframework.boot") version "2.2.1.RELEASE"
    id("io.spring.dependency-management") version "1.0.8.RELEASE"
    id("org.sonarqube") version "2.7"
    id("jacoco")
    id("net.ltgt.apt") version "0.8"
    id("com.diffplug.gradle.spotless") version "3.26.1"
    id("org.jetbrains.kotlin.plugin.jpa") version "1.3.61"
    id("org.jetbrains.kotlin.plugin.noarg") version "1.3.61"
    {{if eq .EnableKafka true}}
    id("com.github.imflog.kafka-schema-registry-gradle-plugin") version "0.8.0"
    id("com.commercehub.gradle.plugin.avro") version "0.17.0"
    {{end}}
    kotlin("jvm") version "1.3.61"
    kotlin("kapt") version "1.3.61"
    kotlin("plugin.spring") version "1.3.61"
}

apply(plugin = "io.spring.dependency-management")
apply(plugin = "checkstyle")
apply(plugin = "kotlin")

group = "com.flixbus.charter"
java.sourceCompatibility = JavaVersion.VERSION_1_10

val newrelicVersion = "4.11.0";
val swaggerVersion = "2.9.2";
val mysqlConnectorVersion = "8.0.11"
val azureVersion = "2.1.6"
val prodProfile = project.hasProperty("prod")

extra["springCloudVersion"] = "Hoxton.RC2"

{{if eq .EnableKafka true}}
val confluentVersion = "5.3.0"
{{end}}

repositories {
    mavenCentral()
    maven { url = uri("https://repo.spring.io/milestone") }
    maven("http://packages.confluent.io/maven/")
}

sourceSets {
    create("integrationTest") {
        withConvention(KotlinSourceSet::class) {
            kotlin.srcDir("src/integrationTest/kotlin")
            resources.srcDir("src/integrationTest/resources")
            compileClasspath += sourceSets["main"].output + configurations["testRuntimeClasspath"]
            runtimeClasspath += output + compileClasspath + sourceSets["test"].runtimeClasspath
        }
    }
}

task<Test>("integrationTest") {
    description = "Runs the integration tests"
    group = "verification"
    testClassesDirs = sourceSets["integrationTest"].output.classesDirs
    classpath = sourceSets["integrationTest"].runtimeClasspath
    mustRunAfter(tasks["test"])
    useJUnitPlatform()
}

dependencies {
    implementation("org.springframework.boot:spring-boot-starter")
    implementation("org.springframework.boot:spring-boot-starter-web")
    implementation("org.springframework.boot:spring-boot-starter-actuator")
    implementation("org.springframework.boot:spring-boot-starter-tomcat")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin")
    implementation("org.springframework.boot:spring-boot-starter-data-jpa")
    implementation("org.jetbrains.kotlin:kotlin-reflect")
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    testImplementation("org.springframework.boot:spring-boot-starter-test")
    testCompile("com.h2database:h2:1.4.199")

    compile("org.springframework.cloud:spring-cloud-starter-openfeign")
    compile("org.springframework.cloud:spring-cloud-starter-netflix-ribbon")

    compile("org.jadira.usertype:usertype.core:7.0.0.CR1") {
        exclude(group = "org.hibernate", module = "hibernate-entitymanager")
    }

    compile("com.newrelic.agent.java:newrelic-api:$newrelicVersion")

    compile("org.joda:joda-money:1.0.1")
    compile("joda-time:joda-time:2.10.1")

    compile("mysql:mysql-connector-java:$mysqlConnectorVersion")

    compile("io.springfox:springfox-swagger2:$swaggerVersion")
    compile("io.springfox:springfox-swagger-ui:$swaggerVersion")

    compile("net.lingala.zip4j:zip4j:2.2.7")

    compile("org.apache.commons:commons-csv:1.7")
    compile("com.amazonaws:aws-java-sdk:1.11.681")

    {{if eq .EnableKafka true}}
    // Kafka
    compile("io.confluent:kafka-avro-serializer:$confluentVersion")
    compile("io.confluent:kafka-schema-registry-client:$confluentVersion")
    compile("io.confluent:kafka-streams-avro-serde:$confluentVersion")
    compile("io.confluent:monitoring-interceptors:$confluentVersion")
    compile("org.springframework.kafka:spring-kafka:2.3.3.RELEASE")

    compile("org.apache.avro:avro:1.9.1")
    {{end}}

    compile("org.springframework.boot:spring-boot-starter-test")
    compile("org.springframework.security:spring-security-test")

    compile("com.microsoft.azure:azure-active-directory-spring-boot-starter:$azureVersion")
}

dependencyManagement {
    imports {
        mavenBom("org.springframework.cloud:spring-cloud-dependencies:${property("springCloudVersion")}")
    }
}

tasks {
    withType<Test> {
        useJUnitPlatform()
    }

    withType<KotlinCompile> {
        dependsOn("copyNewRelic")

        kotlinOptions {
            freeCompilerArgs = listOf("-Xjsr305=strict")

            jvmTarget = "1.8"
        }
    }

    create("copyNewRelic", Copy::class) {

        from("newrelic") {
            include("**/*.jar")
            include("**/*.yml")
        }

        into("$buildDir/libs")
        includeEmptyDirs = false
    }

    getByName<BootRun>("bootRun") {
        val activeProfile = System.getProperty("spring.profiles.active")
        systemProperty("spring.profiles.active", activeProfile)

        dependsOn("copyNewRelic")

        var newRelicEnabled = false;
        if (System.getProperty("newrelic.enabled") != null) {
            newRelicEnabled = System.getProperty("newrelic.enabled").toBoolean()
        }
        if (newRelicEnabled) {
            jvmArgs = listOf("-javaagent:newrelic/newrelic.jar")
        }
    }
}

// Checkstyle
spotless {
    kotlin {
    }
}

tasks.withType<Checkstyle>().configureEach {
    source("src/main/kotlin")
    exclude("**")
    include("src/main/kotlin")
    reports {
        xml.isEnabled = false
        html.isEnabled = true
        html.stylesheet = resources.text.fromFile("config/checkstyle/checkstyle.xml")
    }
}
// end of Checkstyle

// sonarqube
val allTestCoverageFile = "$buildDir/jacoco/allTestCoverage.exec"
val allITCoverageFile = "$buildDir/jacoco/allITCoverage.exec"
sonarqube {
    properties {
        property("sonar.host.url", "https://sonar.mfb.io")
        property("sonar.login", "04ea568319221bcc3852d24cfefe887efce5a3d3")
        property("sonar.gitlab.user_token", "Ga7bxxmyoX3tTgLAkgsT")
        property("sonar.projectKey", "charter-finance-backend")
        property("sonar.projectName", "Charter-Finance-Backend")
        property("sonar.java.binaries", "${project.buildDir}")
        property("sonar.junit.reportPaths", "./build/test-results/test,./build/test-results/integrationTest")
        property("sonar.jacoco.reportPaths", allTestCoverageFile)
        property("sonar.jacoco.itReportPath", allITCoverageFile)
        property("sonar.tests", "src/test,src/integrationTest")
        property("sonar.sources", "src/main")
        property("sonar.exclusions", "**/*.js")
    }
}

task<JacocoMerge>("jacocoMergeTest") {
    destinationFile = file(allTestCoverageFile)
    executionData = project.fileTree(mapOf("dir" to ".", "include" to "**/build/jacoco/test.exec"))
}

task<JacocoMerge>("jacocoMergeIntegrationTest") {
    destinationFile = file(allITCoverageFile)
    executionData = project.fileTree(mapOf("dir" to ".", "includes" to listOf("**/build/jacoco/integrationTest.exec")))
}

task<JacocoMerge>("jacocoMerge") {
    executionData = project.fileTree(mapOf("dir" to ".", "includes" to listOf("**/build/jacoco/integrationTest.exec")))
    dependsOn("jacocoMergeTest", "jacocoMergeIntegrationTest")
}

tasks.withType<JacocoReport> {
    group = "verification"
    description = "Generate Jacoco coverage reports after running tests."

    executionData.setFrom(fileTree(mapOf("dir" to "build/jacoco", "include" to "**/*.exec")))

    reports {
        xml.apply {
            isEnabled = true
            //Following value is a file
            destination = file("${buildDir}/reports/jacoco/xml/jacoco.xml")
        }
        csv.isEnabled = false
        html.apply {
            isEnabled = true
            //Following value is a folder
            destination = file("${buildDir}/reports/jacoco/html")
        }
    }

    afterEvaluate {
        classDirectories.setFrom(files(classDirectories.files.map {
            fileTree(mapOf(
                    "dir" to it,
                    "exclude" to listOf(
                            "**/dto/**",
                            "**/controller/**"
                    )
            ))
        }))
    }
}

// end of sonarqube

// kafka
{{if eq .EnableKafka true}}
var schemaRegistryUrl = "https://schema-registry.buzz-dev.flixos.io"
if (prodProfile) {
    schemaRegistryUrl = "https://schema-registry.buzz.flixos.io"
}

configure<com.github.imflog.schema.registry.SchemaRegistryExtension> {
    url = schemaRegistryUrl
}

configure<com.github.imflog.schema.registry.download.DownloadSubjectExtension> {
    //subject("fxt-charter-cust-booking-invoiceneeded-value", "src/main/avro")
}

configure<com.github.imflog.schema.registry.compatibility.CompatibilitySubjectExtension> {
    //subject("fxt-charter-cust-booking-invoiceneeded-value", "src/main/avro/fxt-charter-cust-booking-invoiceneeded-value.avsc")
}
// end of kafka
{{end}}