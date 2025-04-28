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
```journalctl -u healthy.service```

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

