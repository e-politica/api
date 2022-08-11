CREATE TABLE IF NOT EXISTS default_account (
    df_account_id uuid PRIMARY KEY default uuid_generate_v4(),
    
    df_account_password CHARACTER varying(500) NOT NULL,

    df_account_update_timestamp TIMESTAMP default NOW()
)