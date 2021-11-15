#[derive(Debug)]
pub struct CreateProductDto {
    pub product_category: String,
    pub tag: String,
    pub title: String,
    pub subtitle: String,
    pub authors: Vec<String>,
    pub amount_in_stock: i64,
    pub num_pages: i64,
    pub tags: Vec<String>,
}
