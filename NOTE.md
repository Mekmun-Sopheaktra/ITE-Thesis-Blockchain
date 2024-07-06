# Docker :
1. Docker Swarm Init
   `docker swarm init --advertise-addr MANAGER_IP`
   2. Example : `docker swarm init --advertise-addr 192.168.33.10`
2. Docker Swarm Join : just copy token to paste in another machine
   `docker swarm join --token SWMTKN-1-3dhcnsyga5gb29v4woiqnt92tdxzout69okx71gv0c5auc1ote-enyjkpsvxhh0f96ho137wwyeq 192.168.33.10:2377`

3. Create Service
   `docker service create --replicas 3 -p 80:80 --name serviceName nginx`
   2. Example : `docker service create --replicas 3 -p 80:80 --name web nginx`

4. Scale Service
    `docker service scale serviceName=2`

5. Create Overlay Network
    `docker network create --driver=overlay --attachable network`

# Setup Network
    1. Add Label to Node
        `docker node update --label-add name=manager vfy1bwhj1t1gbu1m5u2xiphn4`
        `docker node update --label-add name=worker1 vfy1bwhj1t1gbu1m5u2xiphn4`
    2. Deploy CA
        `docker stack deploy -c docker/docker-compose-ca.yaml hlf`
    3. Generate Certificates
        `source ./organizations/fabric-ca/registerEnroll.sh` //it may have some error on dos2unix
    4. Call Create Org1
        `createOrg1`



export EXPLORER_CONFIG_FILE_PATH=./config.json
export EXPLORER_PROFILE_DIR_PATH=./connection-profile
export FABRIC_CRYPTO_PATH=./organizations