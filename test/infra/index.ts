import * as koyeb from "@koyeb/pulumi-koyeb";

const koyebApp = new koyeb.App("sample-app", {
  name: "sample-app",
});

const serviceResource = new koyeb.Service("serviceResource", {
    appName: koyebApp.name,
    definition: {
        instanceTypes: [{
            type: "micro",
            // scopes: ["string"],
        }],
        scalings: [{
            max: 0,
            min: 0,
            scopes: ["string"],
            targets: [{
                averageCpus: [{
                    value: 1,
                }],
                averageMems: [{
                    value: 0,
                }],
                concurrentRequests: [{
                    value: 10,
                }],
                requestResponseTimes: [{
                    value: 1,
                }],
                requestsPerSeconds: [{
                    value: 10,
                }],
            }],
        }],
        regions: ["string"],
        name: "string",
        healthChecks: [{
            gracePeriod: 0,
            http: {
                path: "/api/v1/health",
                port: 0,
                headers: [{
                    key: "string",
                    value: "string",
                }],
                method: "string",
            },
            interval: 0,
            restartLimit: 0,
            tcp: {
                port: 0,
            },
            timeout: 0,
        }],
        docker: {
            image: "string",
            args: ["string"],
            command: "string",
            entrypoints: ["string"],
            imageRegistrySecret: "string",
            privileged: false,
        },
        ports: [{
            port: 0,
            protocol: "string",
        }],
        git: {
            branch: "string",
            repository: "string",
            buildpack: {
                buildCommand: "string",
                privileged: false,
                runCommand: "string",
            },
            dockerfile: {
                args: ["string"],
                command: "string",
                dockerfile: "string",
                entrypoints: ["string"],
                privileged: false,
                target: "string",
            },
            noDeployOnPush: false,
            workdir: "string",
        },
        routes: [{
            path: "string",
            port: 0,
        }],
        envs: [{
            key: "string",
            scopes: ["string"],
            secret: "string",
            value: "string",
        }],
        skipCache: false,
        type: "string",
        volumes: [{
            id: "string",
            path: "string",
            replicaIndex: 0,
            scopes: ["string"],
        }],
    },
    messages: "string",
});