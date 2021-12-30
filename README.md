# Awair Prometheus Exporter

This is a prometheus exporter for Awair products, I wanted to use my Awair element in my prometheus + grafana setup but ended up building my own. So I'm just throwing this out here in case anyone would like to use it (esp now that awair basically known for their crypto)

Basically just get a token for Awair's API from [here](https://docs.developer.getawair.com/) and run the app like so:

`AWAIR_TOKEN=token make run`

Then point your prometheus instance at ip:2112/metrics and you'll be good to go. 

---

Feel free to open issues if you're using this I would be happy to help.
