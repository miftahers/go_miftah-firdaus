# Praktikum Compute Service
## tugas 2
1. Membuat VM di EC2, dan Implementasi Security Group EC2
    - Create Instances
    !['create instances'](./screenshots/1-instances-created.png)
    - Set security Group
    !['set security group'](./screenshots/2-instances-set_security_group.png)

2. Melakukan ssh remote ke VM di AWS EC2 menggunakan key
   - aws configure & connect
    !['aws configure & connect'](./screenshots/3-connected-ssh-to-console.png)
3. Deploy your program to EC2
   - Deploy program
    !['deploy using docer compose'](screenshots/4-Docker-container-initiated.png)
   - Test:
     - Create user
        !['create user'](screenshots/5-test-create-user.png)
     - User login
        !['user login'](screenshots/6-user-login.png)
     - Get Users
        !['get users'](screenshots/7-get-users.png)

## tugas 3
1. Membuat DB di RDS
   !['choosing_database'](screenshots/8-db-choosing_database.png)
   !['set_master_name_and_password'](screenshots/9-db-set_master_name_and_password.png)
   !['instances_config'](screenshots/10-db-instances_config.png)
   !['setting_storage'](screenshots/11-db-setting_storage.png)
   !['set_vpc'](screenshots/12-db-set_vpc.png)
   !['choose_existing_instance_security_group'](screenshots/13-db-choose_existing_instance_security_group.png)
   !['set_db_port'](screenshots/14-db-set_db_port.png)
   !['finish_create_db'](screenshots/15-db-finish_create_db.png)
   !['database_list'](screenshots/16-db-database_list.png)
   !['db-summary'](screenshots/17-db-summary.png)

2. Migrate your local Data to RDS
   
   skip

3. Connect app to RDS

    - work using mysql in ec2 instance
    !['work using mysql in ec2 instance'](screenshots/18-work-using-instance-mysql.png)
    - settingan docker-compose (bukti gak pake container mysql)
    !['settingan docker-compose'](screenshots/19-settingan-docker-compose.png)
    - test postman
    !['test postman'](screenshots/20-test-postman.png)
    - data exist in DB
    !['data exist in db'](screenshots/21-data-exist-in-db.png)