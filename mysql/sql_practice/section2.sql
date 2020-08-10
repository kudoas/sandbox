
-- q6
select 
  sum(price * item_quantity) sum_price 
from (
  select item_id, item_quantity 
  from (
    select b.id 
    from customers a inner join orders b on a.id = b.customer_id 
      where name = "B商会"
      ) a 
      inner join order_details b on a.id = b.order_id
      ) c 
      inner join items d on c.item_id = d.id;