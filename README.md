Problems to solve:
    - Need to be told when my applications don't behave as expected, rather than hoping I find out somehow like a user reporting the issue.
    - Need to figure out what the problem is fast so I can start working on a solution.

Solution:
    - Ensure applications report their issues by asking each of the for their health status on an interval.
    - If applications fail to report their status after so long then we consider that a failure and notify on that issue.
    - Aggregate all health statuses to a single UI page.
    - Have a way for applications to report unexpected failures (e.i., assertion failures, errors that bubble up and we would normally exit(1) on them)

Purpose:


Ensure local ~/.ssh/config file contains an entry for "deploy.target"

Deploy:
```./deploy.sh```

See logs:
```journalctl -l --no-pager -u healthy.service | less```

Turn off deployed service:
```sudo systemctl stop healthy.service```

Config
    Upload with:
```aws s3 cp <local_file_path> s3://<bucket_name>/<object_key>```
    Download with:
```aws s3 cp s3://<bucket_name>/<object_key> <local_file_path>```

Config locations:
    Config is grabbed from s3 but make sure you emplace s3 files at:
```/root/.aws/config```
```/root/.aws/credentials```

Mass Deploy:
```./mass-deploy.sh```


Bulk config upload testing:

aws s3 cp ./api/config.json s3://config-bunker/testing/api/config.json 
aws s3 cp ./cron/config.json s3://config-bunker/testing/cron/config.json 
aws s3 cp ./message/config.json s3://config-bunker/testing/message/config.json 
aws s3 cp ./payment/config.json s3://config-bunker/testing/payment/config.json 
aws s3 cp ./proxy/config.json s3://config-bunker/testing/proxy/config.json 
aws s3 cp ./pubsub/config.json s3://config-bunker/testing/pubsub/config.json 
aws s3 cp ./share/config.json s3://config-bunker/testing/share/config.json 
aws s3 cp ./upload/config.json s3://config-bunker/testing/upload/config.json 
aws s3 cp ./user/config.json s3://config-bunker/testing/user/config.json 
aws s3 cp ./healthy/config.json s3://config-bunker/testing/healthy/config.json 
aws s3 cp ./jobby/config.json s3://config-bunker/testing/jobby/config.json 



Bulk config downloads testing:

aws s3 cp .s3://config-bunker/testing/api/config.json ./api/config.json 
aws s3 cp .s3://config-bunker/testing/cron/config.json ./cron/config.json 
aws s3 cp .s3://config-bunker/testing/message/config.json ./message/config.json 
aws s3 cp .s3://config-bunker/testing/payment/config.json ./payment/config.json 
aws s3 cp .s3://config-bunker/testing/proxy/config.json ./proxy/config.json 
aws s3 cp .s3://config-bunker/testing/pubsub/config.json ./pubsub/config.json 
aws s3 cp .s3://config-bunker/testing/share/config.json ./share/config.json 
aws s3 cp .s3://config-bunker/testing/upload/config.json ./upload/config.json 
aws s3 cp .s3://config-bunker/testing/user/config.json ./user/config.json 


Bulk config upload production:

aws s3 cp ./api/config.json s3://config-bunker/api/config.json 
aws s3 cp ./cron/config.json s3://config-bunker/cron/config.json 
aws s3 cp ./message/config.json s3://config-bunker/message/config.json 
aws s3 cp ./payment/config.json s3://config-bunker/payment/config.json 
aws s3 cp ./proxy/config.json s3://config-bunker/proxy/config.json 
aws s3 cp ./pubsub/config.json s3://config-bunker/pubsub/config.json 
aws s3 cp ./share/config.json s3://config-bunker/share/config.json 
aws s3 cp ./upload/config.json s3://config-bunker/upload/config.json 
aws s3 cp ./user/config.json s3://config-bunker/user/config.json 



Bulk config downloads production:

aws s3 cp .s3://config-bunker/api/config.json ./api/config.json 
aws s3 cp .s3://config-bunker/cron/config.json ./cron/config.json 
aws s3 cp .s3://config-bunker/message/config.json ./message/config.json 
aws s3 cp .s3://config-bunker/payment/config.json ./payment/config.json 
aws s3 cp .s3://config-bunker/proxy/config.json ./proxy/config.json 
aws s3 cp .s3://config-bunker/pubsub/config.json ./pubsub/config.json 
aws s3 cp .s3://config-bunker/share/config.json ./share/config.json 
aws s3 cp .s3://config-bunker/upload/config.json ./upload/config.json 
aws s3 cp .s3://config-bunker/user/config.json ./user/config.json 


