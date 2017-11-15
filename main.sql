CREATE audit (
  user_id int
  date date,
  id int,
  -- ....
  UNIQUE(user_id, date, id)
)

INSERT --- .....

INSERT --- ON CONFLICT
DO UPADTE 
  old.money
  old.tmoney
 === 
  new.money
  new.tmoney

ELSE 
  insert into audit_rejected (l....)
DO NOTHING

  -- log.Error("inst ... int")

CREATE audit_rejected (
  user_id int
  date date,
  id int,
  -- ....
)
