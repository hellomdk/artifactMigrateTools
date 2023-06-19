set -e

source="nexus"


# 1.启动pod服务

# 2.执行制品迁移脚本
kubectl -n test exec -it repo-mdk-deployment-7b55d87c59-nmlqm bash migrate-xx.sh

# 3.复制SQL文件到本到
kubectl -n test exec -it repo-mdk-deployment-7b55d87c59-nmlqm bash migrate-xx.sh

# 4.执行SQL脚本--插入binary表
kubectl -n test exec -it repo-mdk-deployment-7b55d87c59-nmlqm bash migrate-xx.sh

# 5.执行migrate工具
./migrate xxx

