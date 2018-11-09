#standardSQL
# take the one name associated with a SKU
WITH product_query AS (
  SELECT 
  DISTINCT 
  v2ProductName,
  productSKU
  FROM `data-to-insights.ecommerce.all_sessions_raw` 
  WHERE v2ProductName IS NOT NULL 
)

SELECT k.* FROM (

  # aggregate the products into an array and 
  # only take 1 result
  SELECT ARRAY_AGG(x LIMIT 1)[OFFSET(0)] k 
  FROM product_query x 
  GROUP BY productSKU # this is the field we want deduplicated
);
