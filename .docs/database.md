# Database
Since Soluno relies on a PostgreSQL database in order to store the parsed data, one of the most important things is to create such database. To do this the first thing you need to do is install [PostgreSQL](https://www.postgresql.org/). 

Once installed you need to create a new database, and a new user that is going to read and write data inside it.  
Then, once that's one, you need to run the SQL queries that you can find inside the [`database/schema` folder](../database/schema).  

Once that's done, you are ready to [continue the setup](setup.md).

# Backup accounts status

To backup all the accounts status in order to migrate to the new database, execute the following command:

```bash
pg_dump soljuno --table=account_balance --table=multisig --table=nonce_account \
    --table=program_account --table=program_data_account --table=stake_account \
    --table=stake_delegation --table=stake_lockup --table=token --table=token_account \
    --table=token_account_balance --table=token_delegation --table=token_supply \
    --table=validator --table=validator_config \
    --section=data --column-inserts --on-conflict-do-nothing -f backup.dump
```