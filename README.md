# pop_content_worker

Test repo with yggdrasil worker to track pulp replication via yggdrasil

Example ruby code to send messages to this worker:
```
require 'mqtt'
require 'json'
require 'time'
require 'securerandom'


def with_mqtt_client(&block)
        ::MQTT::Client.connect("hostname",
                               1883,
                               :ssl => false,
                               &block)
end

def mqtt_payload_base_create
      {
        'type': 'data',
        'message_id': SecureRandom.uuid,
        'version': 1,
        'sent': DateTime.now.iso8601,
        'directive': 'echo',
        'status':'ON',
        "metadata": {
        "action": "create",  
        "name": "UpstreamPulp",
        "base_url": "upstream_url",
        "api_root": "/pulp/",
    },
      }.to_json
end

def mqtt_payload_base_list
      {
        'type': 'data',
        'message_id': SecureRandom.uuid,
        'version': 1,
        'sent': DateTime.now.iso8601,
        'directive': 'echo',
        'status':'ON',
        "metadata": {
        "action": "list"
         },
      }.to_json
end

def mqtt_payload_base_replicate
      {
        'type': 'data',
        'message_id': SecureRandom.uuid,
        'version': 1,
        'sent': DateTime.now.iso8601,
        'directive': 'echo',
        'status':'ON',
        "metadata": {
        "action": "replicate",
        "href": "",
         },
      }.to_json
end



def publish(topic, payload, retain: false, qos: 1)
        with_mqtt_client do |c|
          c.publish(topic,payload, retain, qos)
        end
      end
puts "Hello World!"
publish('yggdrasil/859e1e4b49f84928aab3006e38f104a3/data/in', mqtt_payload_base_replicate())
```
