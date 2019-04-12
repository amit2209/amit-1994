step1 ->Establishing connection with redis

step2->Checking if coming request is withing given limit

step3->If the difference between current timestamp and timestamp coming from redis is within the  limit  and counter reach the maximum then reject the request .

step4->If not upsert the new timestamp and reduce the counter with one .
