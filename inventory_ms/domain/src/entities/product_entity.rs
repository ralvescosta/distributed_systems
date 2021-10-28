pub struct ProductEntity {
    pub id: String,
    pub product_category: String,
    pub tag: String,
    pub title: String,
    pub subtitle: String,
    pub authors: Vec<String>,
    pub amount_in_stock: i32,
    pub created_at: String,
    pub updated_at: String,
    pub num_pages: i32,
    pub tags: Vec<String>,
}
