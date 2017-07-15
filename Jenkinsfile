//Jenkins Kubenetes Plugin 插件地址: https://github.com/jenkinsci/kubernetes-plugin
// 这个插件目前只支持 pipeline, 需要先在 Jenkins master 中添加一个新的 kubernetes 的 cloud, 配置完成就是做一些集群的配置
// Manage Jenkins -> Configure System -> add a new cloud

// 注意其中几项
// Jenkins URL: 需要是 master 的 url, 必须带 http
// Jenkins tunnel: 是 jnlp slave 访问 master 的 url, 默认端口是50000, 这个一定不能带http
// Container Cap: 这个值并不是字面上看起来的意思, 它表示这个 k8s 的 cloud 最多同时能提供多少个 slave/agent
// Kubernetes server certificate key: X509 PEM encoded, 不能有换行, 不能有头尾, 就是一个字符串

def image_tag = "caicloud/kubernetes-admin:${params.imageTag}"
def registry = "cargo.caicloudprivatetest.com"
//运行了一个叫podTemplate的step, 
//定义了这个step运行在`dev-cluster`的cloud上面, 名字叫`dev-cluster`, 
//label是`kubernetes-admin`, podTemplate step并不会去创建pod, 它只是定义了一个podTemplate, 注册到cloud中
podTemplate(
    cloud: 'dev-cluster',// The name of the cloud as defined in Jenkins settings. Defaults to kubernetes
    namespace: 'kube-system',//The namespace of the pod.
    name: 'kubernetes-admin',//The name of the pod.  这个名字会影响到slave的名字, slave实际名字是kube-system-${UUID}
    // 这个地方是一个trick, 一旦遇到always-或者always_开头的label
    // 则表示这个pod是一个长期运行的pod, retentionStrategy改为Always, 长期存在
    label: 'test',// The label of the pod. 这个最重要, 可以说是唯一标示
    instanceCap: 1,// 这个表示这个pod template在k8s集群中最多同时可以有几个实例
    //nodeSelector: "os=centos,lg=golang", // k8s node selector
    idleMinutes: 1440,
    // 下面这个Container是这个插件强制要求启动的, 是一个jnlp-slave, 用来跟master通讯
    containers: [//The container templates that are use to create the containers of the pod (see below). 用于创建pod容器的容器模板
        // jnlp with kubectl
        containerTemplate(
            name: 'jnlp',
            alwaysPullImage: true,
            image: 'cargo.caicloud.io/circle/jnlp:2.62',
            command: '',
            args: '${computer.jnlpmac} ${computer.name}',// 强制要求这么写
        ),
        // docker in docker
        containerTemplate(
            name: 'dind', 
            image: 'cargo.caicloud.io/caicloud/docker:17.03-dind', 
            ttyEnabled: true, 
            command: '', 
            args: '--host=unix:///home/jenkins/docker.sock',
            privileged: true,// docker in docker 要求privileged模式
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
    node('test') {// 这个地方表面使用demo-job-echo的标签的node
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
                        # ln link /home/jenkins/workspace/xxxx to /go/src/github.com/caicloud/baomengjiang at first time
                        # ln will link /home/jenkins/workspace/xxxx to /go/src/github.com/caicloud/baomengjiang/xxxx at second time
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
                    withCredentials ([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'bmj', usernameVariable: 'GIT_USERNAME', passwordVariable: 'GIT_PASSWORD']]){
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
