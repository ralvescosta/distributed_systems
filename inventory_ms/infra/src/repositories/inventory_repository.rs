use crate::database::{connection::MongoConnection, documents::product_documents::ProductDocument};

pub struct InventoryRepository {
    connection: MongoConnection,
}

impl InventoryRepository {
    pub fn new(connection: MongoConnection) -> InventoryRepository {
        InventoryRepository { connection }
    }
}

impl InventoryRepository {
    pub async fn get_product_by_id(&self) -> Result<(), Box<dyn std::error::Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        match collection.find_one(None, None).await? {
            None => Ok(()),
            _ => Ok(()),
        }
    }
}
