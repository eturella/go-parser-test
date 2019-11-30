copy favoritemovies from 'dynamodb://Movies'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
readratio 50;

copy listing
from 's3://mybucket/data/listing/' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';

copy sales
from 'emr://j-SAMPLE2B500FC/myoutput/part-*' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '\t' lzop;

copy sales
from 'emr://j-SAMPLE2B500FC/myoutput/json/' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
JSON 's3://mybucket/jsonpaths.txt';

copy category
from 's3://mybucket/custdata' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';

copy customer
from 's3://mybucket/cust.manifest' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
manifest;

copy listing 
from 's3://mybucket/data/listings_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole';

copy listing 
from 's3://mybucket/data/listings/parquet/' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
format as parquet;

copy event
from 's3://mybucket/data/allevents_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
removequotes
emptyasnull
blanksasnull
maxerror 5
delimiter '|'
timeformat 'YYYY-MM-DD HH:MI:SS';

copy venue
from 's3://mybucket/data/venue_fw.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
fixedwidth 'venueid:3,venuename:25,venuecity:12,venuestate:2,venueseats:6';

copy category
from 's3://mybucket/data/category_csv.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
csv;

copy category
from 's3://mybucket/data/category_csv.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
csv quote as '%';

copy venue
from 's3://mybucket/data/venue.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
explicit_ids;

copy time
from 's3://mybucket/data/timerows.gz' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
gzip
delimiter '|';

copy timestamp1 
from 's3://mybucket/data/time.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
timeformat 'YYYY-MM-DD HH:MI:SS';

copy venue_new(venueid, venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_noseats.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|';

copy venue(venueid, venuecity, venuestate) 
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|';

copy venue(venueid, venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|' explicit_ids;

copy venue(venuename, venuecity, venuestate) 
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|' explicit_ids;

copy venue(venueid, venuename, venuecity, venuestate)
from 's3://mybucket/data/venue_pipe.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
delimiter '|';

copy redshiftinfo from 's3://mybucket/data/redshiftinfo.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
delimiter '|' escape;

copy category
from 's3://mybucket/category_object_auto.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 'auto';

copy category
from 's3://mybucket/category_object_paths.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 's3://mybucket/category_jsonpath.json';

copy category
from 's3://mybucket/category_array_data.json'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
json 's3://mybucket/category_array_jsonpath.json';

copy category
from 's3://mybucket/category_auto.avro'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'
format as avro 'auto';

copy category
from 's3://mybucket/category_object_paths.avro'
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole' 
format avro 's3://mybucket/category_path.avropath ';

copy t2 from 's3://mybucket/data/nlTest2.txt' 
iam_role 'arn:aws:iam::0123456789012:role/MyRedshiftRole'  
escape
delimiter as '|';