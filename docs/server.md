### Работа со стендами (инициализация)

Установить Kubectl -cli  для подключение к кубам

    `curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl`
    `chmod +x ./kubectl`
    `sudo mv ./kubectl /usr/local/bin/kubectl`

Конфигурация для подлкючения находиться здесь:

    `config-project`

Выполняем команду (для перманентой работы можно прописать в систему):

    `export KUBECONFIG=config-project`
    
### Работа со стендами (команды)    
    
    kubectl get pods -n project-develop    - вывод запущенных подов на dev
    kubectl get pods -n project-staging    - вывод запущенных подов на stage
    kubectl get pods -n project-production    - вывод запущенных подов на prod

Подсключение и проброс портов для docom

    sudo kubectl port-forward docom-f9b6c86cf-7hjr6 90:80

Подсключение и проброс портов для просмотра логов на стендах

    kubectl logs -f [NAME|ID пода]   - просмотр логов на стендах

Подсключение и проброс портов до баз данных для back

(БД со стенда будет доступа по порту 5433 на localhost)
    
    kubectl port-forward postgresql-0 5433:5432 -n project-develop
    kubectl port-forward postgresql-0 5433:5432 -n project-staging
    kubectl port-forward postgresql-0 5433:5432 -n project-production