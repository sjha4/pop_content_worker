# echo_worker
Test repo with yggdrasil worker to track pulp replication via yggdrasil

Example ruby code to send messages to this worker:
```
require 'mqtt'
require 'json'
require 'time'
require 'securerandom'


def with_mqtt_client(&block)
        ::MQTT::Client.connect("dhcp-2-22.vms.sat.rdu2.redhat.com",
                               1883,
                               :ssl => false,
                               &block)
end

def mqtt_payload_base
      {
        'type': 'data',
        'message_id': SecureRandom.uuid,
        'version': 1,
        'sent': DateTime.now.iso8601,
        'directive': 'echo'
      }.to_json
end

def publish(topic, payload, retain: false, qos: 1)
        with_mqtt_client do |c|
          c.publish(topic,payload, retain, qos)
        end
      end
puts mqtt_payload_base
puts "Hello World!"
publish('yggdrasil/859e1e4b49f84928aab3006e38f104a3/data/in', mqtt_payload_base())
```
