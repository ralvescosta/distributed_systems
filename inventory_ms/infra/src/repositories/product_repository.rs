use async_trait::async_trait;
use mongodb::bson::doc;

use application::interfaces::i_product_repository::IProductRepository;
use domain::entities::product_entity::ProductEntity;

use crate::database::{connection::DbConnection, documents::product_documents::ProductDocument};

pub struct ProductRepository {
    connection: DbConnection,
}

impl ProductRepository {
    pub fn new(connection: DbConnection) -> ProductRepository {
        ProductRepository { connection }
    }
}

#[async_trait]
impl IProductRepository for ProductRepository {
    async fn get_product_by_id(
        &self,
        id: String,
    ) -> Result<Option<ProductEntity>, Box<dyn std::error::Error>> {
        let collection = self
            .connection
            .get_collection::<ProductDocument>(&self.connection.inventory_collection_name);

        let filter = doc! { "id": id };
        match collection.find_one(filter, None).await? {
            None => Ok(None),
            Some(document) => Ok(Some(document.to_entity())),
        }
    }
}

#[cfg(test)]
mod test {
    use super::*;

    use mongodb::Client;

    #[tokio::test]
    async fn should_execute_method_with_connection_error() -> mongodb::error::Result<()> {
        let client = Client::with_uri_str("mongodb://example.com").await?;
        let db_connection = DbConnection {
            client,
            app_name: String::from(""),
            db_name: String::from(""),
            inventory_collection_name: String::from(""),
        };

        let sut = ProductRepository::new(db_connection);

        match sut.get_product_by_id(String::from("some")).await {
            Err(_err) => assert!(true),
            _ => assert!(false),
        }

        Ok(())
    }
}
