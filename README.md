# pubsub-event-maker

Simple utility to stream mocked events into Google Cloud PubSub. Good for demo and sample apps.

## Run

To start the `eventmaker` and have it mock Cloud PubSub events run the following command on the [released version](https://github.com/mchmarny/pubsub-event-maker/releases) of this binary

```shell
./eventmaker --project YOUR_PROJECT_ID --topic TARGET_TOPIC_NAME
```

> The above command uses few defaults, you may want to consider the [Additional Configuration](#additional-configuration) section below

The output from the above command will look something like that

```shell
[EVENT-MAKER] Publishing: {"source_id":"device-0","event_id":"eid-b6569857-232c-4e6f-bd51-cda4e81f3e1f","event_ts":"2019-06-05T11:39:50.403778Z","label":"utilization","mem_used":34.47265625,"cpu_used":6.5,"load_1":1.55,"load_5":2.25,"load_15":2.49,"random_metric":94.05090880450125}
```

The JSON payload that will be posted to PubSub topic will have following format:

```json
{
    "source_id": "device-0",
    "event_id": "eid-b6569857-232c-4e6f-bd51-cda4e81f3e1f",
    "event_ts": "2019-06-05T11:39:50.403778Z",
    "label": "utilization",
    "mem_used": 34.47265625,
    "cpu_used": 6.5,
    "load_1": 1.55,
    "load_5": 2.25,
    "load_15": 2.49,
    "random_metric": 94.05090880450125
}
```

## Cost

There is a pretty generous free tier (first 10GB) on PubSub which is priced based on data volume transmitted in a calendar month. For more information see [PubSub Pricing](https://cloud.google.com/pubsub/pricing)

## Additional Configuration

The `eventmaker` also supports the following runtime configuration parameters:

* **project** - (string) GCP Project ID
* **topic** (string) Name of the PubSub topic (will be created if does not exist)
* **freq** (optional, string) Frequency in which the mocked events will be posted to PubSub topic [default "5s"]
* **sources** (optional, int) Number of event sources to mock [default 1]
* **metric** (optional, string) Name of the metric label for each event [default "utilization"]
* **range** (optional, string) Range of the random metric value [default "0-100"]
* **maxErrors** (optional, int) Maximum number of PubSub push errors [default 10]


## Disclaimer

This is my personal project and it does not represent my employer. I take no responsibility for issues caused by this code. I do my best to ensure that everything works, but if something goes wrong, my apologies is all you will get.