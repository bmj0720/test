def image_tag = "caicloud/test:${params.imageTag}"
def registry = "cargo.caicloudprivatetest.com"

podTemplate(
    cloud: 'dev-cluster',
    namespace: 'kube-system',
    name: 'kubernetes-admin',
    label: 'kubernetes-admin',
    instanceCap: 1,
    idleMinutes: 1440,
    containers: [
        // jnlp with kubectl
        containerTemplate(
            name: 'jnlp',
            alwaysPullImage: true,
            image: 'cargo.caicloud.io/circle/jnlp:2.62',
            command: '',
            args: '${computer.jnlpmac} ${computer.name}',
        ),
        // docker in docker
        containerTemplate(
            name: 'dind', 
            image: 'cargo.caicloud.io/caicloud/docker:17.03-dind', 
            ttyEnabled: true, 
            command: '', 
            args: '--host=unix:///home/jenkins/docker.sock',
            privileged: true,
        ),
        // golang with docker client
        containerTemplate(
            name: 'golang',
            image: 'cargo.caicloud.io/caicloud/golang-docker:1.8-17.03',
            ttyEnabled: true,
            command: '',
            args: '',
            envVars: [
                containerEnvVar(key: 'DEBUG', value: 'true'),
                containerEnvVar(key: 'NOT_LOCAL', value: 'true'),
                containerEnvVar(key: 'IMAGE', value: "${registry}/${image_tag}"),
                containerEnvVar(key: 'DOCKER_HOST', value: 'unix:///home/jenkins/docker.sock'),
                containerEnvVar(key: 'DOCKER_API_VERSION', value: '1.26'),
                containerEnvVar(key: 'WORKDIR', value: '/go/src/github.com/caicloud/test')
            ],
        ),
    ]
) {
    node('test') {
        stage('Checkout') {
            checkout scm
        }
        container('golang') {
            ansiColor('xterm') {

                stage("Complie") {
                    sh('''
                        set -e 
                        mkdir -p $(dirname ${WORKDIR}) 

                        echo "clean previous build garbage"

                        # if you do not remove target dir manually
                        # ln will not work according to what you want
                        # ln link /home/jenkins/workspace/xxxx to /go/src/github.com/caicloud/cyclone at first time
                        # ln will link /home/jenkins/workspace/xxxx to /go/src/github.com/caicloud/cyclone/xxxx at second time
                        # so remove the target workdir before you link
                        rm -rf ${WORKDIR}
                        ln -sfv $(pwd) ${WORKDIR}

                        cd ${WORKDIR}

                        echo "buiding test"
                        GOOS=linux GOARCH=amd64 go build -o test
                    ''')
                }

                stage('Run e2e test') {
                    if (!params.integration) {
                        echo "skip integration"
                        return
                    }
                    sh('''
                        set -e
                        cd ${WORKDIR}
                        # get host ip
                        HOST_IP=$(ifconfig eth0 | grep 'inet addr:'| grep -v '127.0.0.1' | cut -d: -f2 | awk '{ print $1}')
                        
                        export CDS_SERVER="http://cds-server.default:9000"

                        echo "run E2E script"
                        /bin/bash tests/run-e2e.sh
                    ''')
                }
            }

            stage("Build image and publish") {
                if (!params.publish) {
                    echo "skip publish"
                    return
                }
                sh "docker build -t ${image_tag} -f Dockerfile ."
                echo "skip push"
                if (params.autoGitTag) {
                    echo "auto git tag: " + params.imageTag
                    withCredentials ([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'baomengjiang', usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD']]){
                        sh("git config --global user.email \"info@caicloud.io\"")
                        sh("git tag -a $imageTag -m \"$tagDescribe\"")
                        sh("git push https://${GIT_USERNAME}:${GIT_PASSWORD}@github.com/baomengjiang/test $imageTag")
                   }
                } 
            }
        }

        stage("deploy") {
            echo "skip deploy"
        }
    }
}
