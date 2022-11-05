CREATE TABLE user_project
(
    up_user_id uuid NOT NULL 
    up_project_id VARCHAR NOT NULL 

     CONSTRAINT up_fk_user    
        FOREIGN KEY(up_user_id)
        REFERENCES user_account(user_id)
)
