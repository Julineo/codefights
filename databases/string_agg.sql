#standardSQL
# recall the earlier query that showed multiple product_names for each SKU 
SELECT 
DISTINCT 
COUNT(DISTINCT v2ProductName) AS product_count,
STRING_AGG(DISTINCT v2ProductName LIMIT 5) AS product_name,
productSKU
FROM `data-to-insights.ecommerce.all_sessions_raw` 
WHERE v2ProductName IS NOT NULL
GROUP BY productSKU
HAVING product_count > 1
ORDER BY product_count DESC
