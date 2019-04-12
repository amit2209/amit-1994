establishing connection with redis
Checking if coming request is withing given limit
If the diference between current timestamp and timestamp coming from redis is within the  limit  and counter reach the maximum then reject the request .
if not upsert the new timestamp and reduce the counter with one .
