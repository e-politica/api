CREATE TABLE comment_project
(
    cp_user_id uuid NOT NULL
    cp_project_id VARCHAR NOT NULL
    cp_comment VARCHAR(200) NOT NULL 
    
    CONSTRAINT cp_fk_user    
        FOREIGN KEY(cp_user_id)
        REFERENCES user_account(user_id)
)