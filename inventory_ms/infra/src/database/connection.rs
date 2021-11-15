use mongodb::{options::ClientOptions, Client, Collection, Database};
use std::{env, error::Error};

#[derive(Debug)]
pub struct DbConnection {
    pub client: Client,
    pub app_name: String,
    pub db_name: String,
    pub inventory_collection_name: String,
}

pub async fn connect_to_database() -> Result<DbConnection, Box<dyn Error>> {
    let mongo_connection_uri = env::var("MONGO_CONNECTION_URI")?;
    let app_name = env::var("APP_NAME")?;
    let db_name = env::var("MONGO_DB_NAME")?;
    let inventory_collection_name = env::var("MONGO_INVENTORY_COLLECTION")?;

    let mut client_options = ClientOptions::parse(mongo_connection_uri).await?;

    client_options.app_name = Some(app_name.clone());
    let client = Client::with_options(client_options)?;

    Ok(DbConnection {
        client,
        app_name,
        db_name,
        inventory_collection_name,
    })
}

impl DbConnection {
    pub fn get_database(&self) -> Database {
        self.client.database(&self.db_name)
    }

    pub fn get_collection<T>(&self, collection_name: &str) -> Collection<T> {
        self.client
            .database(&self.db_name)
            .collection::<T>(collection_name)
    }
}
