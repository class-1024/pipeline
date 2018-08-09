node {
       stage('Clone') {
                checkout scm
                echo "1.Clone Stage" 
                script {
                   build_tag = sh(returnStdout: true, script: 'git rev-parse --short HEAD').trim()
                   echo "build_tag:${build_tag}"
               }      
        }
        stage('Test') {
                echo "2.Test Stage"   
                 sh "sudo  go build  ."   
        }
        stage('Build') {            
                echo "3.Build Docker Image Stage"  
                 sh "sudo docker build -t registry.cn-shenzhen.aliyuncs.com/kinot-k8s/kinot-web-api:${build_tag} ."   
        }
        stage('Push') {             
              echo "4.Push Docker Image Stage"
              sh "sudo docker login -u=广州科玛技术有限公司 -p kinot-2018 registry.cn-shenzhen.aliyuncs.com"
              sh "sudo docker push registry.cn-shenzhen.aliyuncs.com/kinot-k8s/kinot-web-api:${build_tag}"    
        } 
        stage('YAML') {             
               echo "5. Change YAML File Stage"    
        }
        stage('Deploy') {
              echo "6. Deploy Stage"    
        }
}
